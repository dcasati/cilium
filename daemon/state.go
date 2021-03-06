// Copyright 2016-2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/cilium/cilium/common"
	"github.com/cilium/cilium/pkg/endpoint"
	"github.com/cilium/cilium/pkg/endpointmanager"
	identityPkg "github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/cilium/pkg/ipam"
	"github.com/cilium/cilium/pkg/labels"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/policy"
	"github.com/cilium/cilium/pkg/workloads"

	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

type endpointRestoreState struct {
	restored []*endpoint.Endpoint
	toClean  []*endpoint.Endpoint
}

// restoreOldEndpoints reads the list of existing endpoints previously managed
// Cilium when it was last run and associated it with container workloads. This
// function performs the first step in restoring the endpoint structure,
// allocating their existing IP out of the CIDR block and then inserting the
// endpoints into the endpoints list. It needs to be followed by a call to
// regenerateRestoredEndpoints() once the endpoint builder is ready.
//
// If clean is true, endpoints which cannot be associated with a container
// workloads are deleted.
func (d *Daemon) restoreOldEndpoints(dir string, clean bool) (*endpointRestoreState, error) {
	state := &endpointRestoreState{
		restored: []*endpoint.Endpoint{},
		toClean:  []*endpoint.Endpoint{},
	}

	if !option.Config.RestoreState {
		log.Info("Endpoint restore is disabled, skipping restore step")
		return state, nil
	}

	log.Info("Restoring endpoints from former life...")

	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return state, err
	}
	eptsID := endpoint.FilterEPDir(dirFiles)

	possibleEPs := readEPsFromDirNames(dir, eptsID)

	if len(possibleEPs) == 0 {
		log.Info("No old endpoints found.")
		return state, nil
	}

	for _, ep := range possibleEPs {
		scopedLog := log.WithField(logfields.EndpointID, ep.ID)
		skipRestore := false
		if _, err := netlink.LinkByName(ep.IfName); err != nil {
			scopedLog.Infof("Interface %s could not be found for endpoint being restored, ignoring", ep.IfName)
			skipRestore = true
		} else if !workloads.IsRunning(ep) {
			scopedLog.Info("No workload could be associated with endpoint being restored, ignoring")
			skipRestore = true
		}

		if clean && skipRestore {
			state.toClean = append(state.toClean, ep)
			continue
		}

		ep.Mutex.Lock()
		scopedLog.Debug("Restoring endpoint")
		ep.LogStatusOKLocked(endpoint.Other, "Restoring endpoint from previous cilium instance")

		if err := d.allocateIPsLocked(ep); err != nil {
			ep.Mutex.Unlock()
			scopedLog.WithError(err).Error("Failed to re-allocate IP of endpoint. Not restoring endpoint.")
			state.toClean = append(state.toClean, ep)
			continue
		}

		if option.Config.KeepConfig {
			ep.SetDefaultOpts(nil)
		} else {
			ep.SetDefaultOpts(option.Config.Opts)
			alwaysEnforce := policy.GetPolicyEnabled() == option.AlwaysEnforce
			ep.Opts.Set(option.IngressPolicy, alwaysEnforce)
			ep.Opts.Set(option.EgressPolicy, alwaysEnforce)
		}

		ep.Mutex.Unlock()

		state.restored = append(state.restored, ep)
	}

	log.WithFields(logrus.Fields{
		"count.restored": len(state.restored),
		"count.total":    len(possibleEPs),
	}).Info("Endpoints restored")

	return state, nil
}

func (d *Daemon) regenerateRestoredEndpoints(state *endpointRestoreState) {
	log.Infof("Regenerating %d restored endpoints", len(state.restored))

	// we need to signalize when the endpoints are regenerated, i.e., when
	// they have finished to rebuild after being restored.
	epRegenerated := make(chan bool, len(state.restored))

	for _, ep := range state.restored {

		// Insert into endpoint manager so it can be regenerated when calls to
		// TriggerPolicyUpdates() are made. This must be done synchronously (i.e.,
		// not in a goroutine) because regenerateRestoredEndpoints must guarantee
		// upon returning that endpoints are exposed to other subsystems via
		// endpointmanager.
		ep.Mutex.RLock()
		endpointmanager.Insert(ep)
		ep.Mutex.RUnlock()

		go func(ep *endpoint.Endpoint, epRegenerated chan<- bool) {
			ep.Mutex.Lock()

			scopedLog := log.WithField(logfields.EndpointID, ep.ID)

			state := ep.GetStateLocked()
			if state == endpoint.StateDisconnected || state == endpoint.StateDisconnecting {
				scopedLog.Warn("Endpoint to restore has been deleted")
				ep.Mutex.Unlock()
				return
			}

			ep.LogStatusOKLocked(endpoint.Other, "Synchronizing endpoint labels with KVStore")
			if err := d.syncLabels(ep); err != nil {
				scopedLog.WithError(err).Warn("Unable to restore endpoint")
				ep.Mutex.Unlock()
				epRegenerated <- false
				return
			}
			ready := ep.SetStateLocked(endpoint.StateWaitingToRegenerate, "Triggering synchronous endpoint regeneration while syncing state to host")
			ep.Mutex.Unlock()

			if !ready {
				scopedLog.WithField(logfields.EndpointState, ep.GetState()).Warn("Endpoint in inconsistent state")
				epRegenerated <- false
				return
			}
			if buildSuccess := <-ep.Regenerate(d, "syncing state to host"); !buildSuccess {
				scopedLog.Warn("Failed while regenerating endpoint")
				epRegenerated <- false
				return
			}

			ep.Mutex.RLock()
			scopedLog.WithField(logfields.IPAddr, []string{ep.IPv4.String(), ep.IPv6.String()}).Info("Restored endpoint")
			ep.Mutex.RUnlock()
			epRegenerated <- true
		}(ep, epRegenerated)
	}

	for _, ep := range state.toClean {
		go d.deleteEndpointQuiet(ep, true)
	}

	go func() {
		regenerated, total := 0, 0
		if len(state.restored) > 0 {
			for buildSuccess := range epRegenerated {
				if buildSuccess {
					regenerated++
				}
				total++
				if total >= len(state.restored) {
					break
				}
			}
		}
		close(epRegenerated)

		log.WithFields(logrus.Fields{
			"regenerated": regenerated,
			"total":       total,
		}).Info("Finished regenerating restored endpoints")
	}()
}

func (d *Daemon) allocateIPsLocked(ep *endpoint.Endpoint) error {
	err := ipam.AllocateIP(ep.IPv6.IP())
	if err != nil {
		// TODO if allocation failed reallocate a new IP address and setup veth
		// pair accordingly
		return fmt.Errorf("unable to reallocate IPv6 address: %s", err)
	}

	defer func(ep *endpoint.Endpoint) {
		if err != nil {
			ipam.ReleaseIP(ep.IPv6.IP())
		}
	}(ep)

	if !option.Config.IPv4Disabled {
		if ep.IPv4 != nil {
			if err = ipam.AllocateIP(ep.IPv4.IP()); err != nil {
				return fmt.Errorf("unable to reallocate IPv4 address: %s", err)
			}
		}
	}
	return nil
}

// readEPsFromDirNames returns a list of endpoints from a list of directory
// names that can possible contain an endpoint.
func readEPsFromDirNames(basePath string, eptsDirNames []string) []*endpoint.Endpoint {
	possibleEPs := []*endpoint.Endpoint{}
	for _, epID := range eptsDirNames {
		epDir := filepath.Join(basePath, epID)
		readDir := func() string {
			scopedLog := log.WithFields(logrus.Fields{
				logfields.EndpointID: epID,
				logfields.Path:       filepath.Join(epDir, common.CHeaderFileName),
			})
			scopedLog.Debug("Reading directory")
			epFiles, err := ioutil.ReadDir(epDir)
			if err != nil {
				scopedLog.WithError(err).Warn("Error while reading directory. Ignoring it...")
				return ""
			}
			cHeaderFile := common.FindEPConfigCHeader(epDir, epFiles)
			if cHeaderFile == "" {
				return ""
			}
			return cHeaderFile
		}
		// There's an odd issue where the first read dir doesn't work.
		cHeaderFile := readDir()
		if cHeaderFile == "" {
			cHeaderFile = readDir()
		}

		scopedLog := log.WithFields(logrus.Fields{
			logfields.EndpointID: epID,
			logfields.Path:       cHeaderFile,
		})

		if cHeaderFile == "" {
			scopedLog.Info("C header file not found. Ignoring endpoint")
			continue
		}

		scopedLog.Debug("Found endpoint C header file")

		strEp, err := common.GetCiliumVersionString(cHeaderFile)
		if err != nil {
			scopedLog.WithError(err).Warn("Unable to read the C header file")
			continue
		}
		ep, err := endpoint.ParseEndpoint(strEp)
		if err != nil {
			scopedLog.WithError(err).Warn("Unable to parse the C header file")
			continue
		}
		possibleEPs = append(possibleEPs, ep)
	}
	return possibleEPs
}

// syncLabels syncs the labels from the labels' database for the given endpoint.
// To be used with endpoint.Mutex locked.
func (d *Daemon) syncLabels(ep *endpoint.Endpoint) error {
	// Filter the restored labels with the new daemon's filter
	l, _ := labels.FilterLabels(ep.OpLabels.IdentityLabels())
	identity, _, err := identityPkg.AllocateIdentity(l)
	if err != nil {
		return err
	}

	if ep.SecurityIdentity != nil {
		if oldSecID := ep.SecurityIdentity.ID; identity.ID != oldSecID {
			log.WithFields(logrus.Fields{
				logfields.EndpointID:              ep.ID,
				logfields.IdentityLabels + ".old": oldSecID,
				logfields.IdentityLabels + ".new": identity.ID,
			}).Info("Security label ID for endpoint is different that the one stored, updating")
		}
	}

	ep.SetIdentity(identity)

	return nil
}
