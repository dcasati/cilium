// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/hds.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import endpoint "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/endpoint"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Different Envoy instances may have different capabilities (e.g. Redis)
// and/or have ports enabled for different protocols.
type Capability_Protocol int32

const (
	Capability_HTTP  Capability_Protocol = 0
	Capability_TCP   Capability_Protocol = 1
	Capability_REDIS Capability_Protocol = 2
)

var Capability_Protocol_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "REDIS",
}
var Capability_Protocol_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"REDIS": 2,
}

func (x Capability_Protocol) String() string {
	return proto.EnumName(Capability_Protocol_name, int32(x))
}
func (Capability_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{0, 0}
}

// Defines supported protocols etc, so the management server can assign proper
// endpoints to healthcheck.
type Capability struct {
	HealthCheckProtocol  []Capability_Protocol `protobuf:"varint,1,rep,packed,name=health_check_protocol,json=healthCheckProtocol,proto3,enum=envoy.service.discovery.v2.Capability_Protocol" json:"health_check_protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Capability) Reset()         { *m = Capability{} }
func (m *Capability) String() string { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()    {}
func (*Capability) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{0}
}
func (m *Capability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capability.Unmarshal(m, b)
}
func (m *Capability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capability.Marshal(b, m, deterministic)
}
func (dst *Capability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capability.Merge(dst, src)
}
func (m *Capability) XXX_Size() int {
	return xxx_messageInfo_Capability.Size(m)
}
func (m *Capability) XXX_DiscardUnknown() {
	xxx_messageInfo_Capability.DiscardUnknown(m)
}

var xxx_messageInfo_Capability proto.InternalMessageInfo

func (m *Capability) GetHealthCheckProtocol() []Capability_Protocol {
	if m != nil {
		return m.HealthCheckProtocol
	}
	return nil
}

type HealthCheckRequest struct {
	Node                 *core.Node  `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Capability           *Capability `protobuf:"bytes,2,opt,name=capability,proto3" json:"capability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{1}
}
func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(dst, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *HealthCheckRequest) GetCapability() *Capability {
	if m != nil {
		return m.Capability
	}
	return nil
}

type EndpointHealth struct {
	Endpoint             *endpoint.Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	HealthStatus         core.HealthStatus  `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EndpointHealth) Reset()         { *m = EndpointHealth{} }
func (m *EndpointHealth) String() string { return proto.CompactTextString(m) }
func (*EndpointHealth) ProtoMessage()    {}
func (*EndpointHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{2}
}
func (m *EndpointHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealth.Unmarshal(m, b)
}
func (m *EndpointHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealth.Marshal(b, m, deterministic)
}
func (dst *EndpointHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealth.Merge(dst, src)
}
func (m *EndpointHealth) XXX_Size() int {
	return xxx_messageInfo_EndpointHealth.Size(m)
}
func (m *EndpointHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealth.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealth proto.InternalMessageInfo

func (m *EndpointHealth) GetEndpoint() *endpoint.Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *EndpointHealth) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

type EndpointHealthResponse struct {
	EndpointsHealth      []*EndpointHealth `protobuf:"bytes,1,rep,name=endpoints_health,json=endpointsHealth,proto3" json:"endpoints_health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EndpointHealthResponse) Reset()         { *m = EndpointHealthResponse{} }
func (m *EndpointHealthResponse) String() string { return proto.CompactTextString(m) }
func (*EndpointHealthResponse) ProtoMessage()    {}
func (*EndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{3}
}
func (m *EndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealthResponse.Unmarshal(m, b)
}
func (m *EndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealthResponse.Marshal(b, m, deterministic)
}
func (dst *EndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealthResponse.Merge(dst, src)
}
func (m *EndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_EndpointHealthResponse.Size(m)
}
func (m *EndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealthResponse proto.InternalMessageInfo

func (m *EndpointHealthResponse) GetEndpointsHealth() []*EndpointHealth {
	if m != nil {
		return m.EndpointsHealth
	}
	return nil
}

type HealthCheckRequestOrEndpointHealthResponse struct {
	// Types that are valid to be assigned to RequestType:
	//	*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest
	//	*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse
	RequestType          isHealthCheckRequestOrEndpointHealthResponse_RequestType `protobuf_oneof:"request_type"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *HealthCheckRequestOrEndpointHealthResponse) Reset() {
	*m = HealthCheckRequestOrEndpointHealthResponse{}
}
func (m *HealthCheckRequestOrEndpointHealthResponse) String() string {
	return proto.CompactTextString(m)
}
func (*HealthCheckRequestOrEndpointHealthResponse) ProtoMessage() {}
func (*HealthCheckRequestOrEndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{4}
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Unmarshal(m, b)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRequestOrEndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Merge(dst, src)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Size(m)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse proto.InternalMessageInfo

type isHealthCheckRequestOrEndpointHealthResponse_RequestType interface {
	isHealthCheckRequestOrEndpointHealthResponse_RequestType()
}

type HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest struct {
	HealthCheckRequest *HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof"`
}
type HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse struct {
	EndpointHealthResponse *EndpointHealthResponse `protobuf:"bytes,2,opt,name=endpoint_health_response,json=endpointHealthResponse,proto3,oneof"`
}

func (*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}
func (*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetRequestType() isHealthCheckRequestOrEndpointHealthResponse_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetHealthCheckRequest() *HealthCheckRequest {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest); ok {
		return x.HealthCheckRequest
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetEndpointHealthResponse() *EndpointHealthResponse {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse); ok {
		return x.EndpointHealthResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HealthCheckRequestOrEndpointHealthResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HealthCheckRequestOrEndpointHealthResponse_OneofMarshaler, _HealthCheckRequestOrEndpointHealthResponse_OneofUnmarshaler, _HealthCheckRequestOrEndpointHealthResponse_OneofSizer, []interface{}{
		(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest)(nil),
		(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse)(nil),
	}
}

func _HealthCheckRequestOrEndpointHealthResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HealthCheckRequestOrEndpointHealthResponse)
	// request_type
	switch x := m.RequestType.(type) {
	case *HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HealthCheckRequest); err != nil {
			return err
		}
	case *HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EndpointHealthResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HealthCheckRequestOrEndpointHealthResponse.RequestType has unexpected type %T", x)
	}
	return nil
}

func _HealthCheckRequestOrEndpointHealthResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HealthCheckRequestOrEndpointHealthResponse)
	switch tag {
	case 1: // request_type.health_check_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheckRequest)
		err := b.DecodeMessage(msg)
		m.RequestType = &HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest{msg}
		return true, err
	case 2: // request_type.endpoint_health_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EndpointHealthResponse)
		err := b.DecodeMessage(msg)
		m.RequestType = &HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HealthCheckRequestOrEndpointHealthResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HealthCheckRequestOrEndpointHealthResponse)
	// request_type
	switch x := m.RequestType.(type) {
	case *HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest:
		s := proto.Size(x.HealthCheckRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse:
		s := proto.Size(x.EndpointHealthResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LocalityEndpoints struct {
	Locality             *core.Locality       `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	Endpoints            []*endpoint.Endpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LocalityEndpoints) Reset()         { *m = LocalityEndpoints{} }
func (m *LocalityEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityEndpoints) ProtoMessage()    {}
func (*LocalityEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{5}
}
func (m *LocalityEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityEndpoints.Unmarshal(m, b)
}
func (m *LocalityEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityEndpoints.Marshal(b, m, deterministic)
}
func (dst *LocalityEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityEndpoints.Merge(dst, src)
}
func (m *LocalityEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityEndpoints.Size(m)
}
func (m *LocalityEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityEndpoints proto.InternalMessageInfo

func (m *LocalityEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityEndpoints) GetEndpoints() []*endpoint.Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// The cluster name and locality is provided to Envoy for the endpoints that it
// health checks to support statistics reporting, logging and debugging by the
// Envoy instance (outside of HDS). For maximum usefulness, it should match the
// same cluster structure as that provided by EDS.
type ClusterHealthCheck struct {
	ClusterName          string               `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	HealthChecks         []*core.HealthCheck  `protobuf:"bytes,2,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	Endpoints            []*LocalityEndpoints `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterHealthCheck) Reset()         { *m = ClusterHealthCheck{} }
func (m *ClusterHealthCheck) String() string { return proto.CompactTextString(m) }
func (*ClusterHealthCheck) ProtoMessage()    {}
func (*ClusterHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{6}
}
func (m *ClusterHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterHealthCheck.Unmarshal(m, b)
}
func (m *ClusterHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterHealthCheck.Marshal(b, m, deterministic)
}
func (dst *ClusterHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHealthCheck.Merge(dst, src)
}
func (m *ClusterHealthCheck) XXX_Size() int {
	return xxx_messageInfo_ClusterHealthCheck.Size(m)
}
func (m *ClusterHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHealthCheck proto.InternalMessageInfo

func (m *ClusterHealthCheck) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterHealthCheck) GetHealthChecks() []*core.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *ClusterHealthCheck) GetEndpoints() []*LocalityEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type HealthCheckSpecifier struct {
	HealthCheck []*ClusterHealthCheck `protobuf:"bytes,1,rep,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty"`
	// The default is 1 second.
	Interval             *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HealthCheckSpecifier) Reset()         { *m = HealthCheckSpecifier{} }
func (m *HealthCheckSpecifier) String() string { return proto.CompactTextString(m) }
func (*HealthCheckSpecifier) ProtoMessage()    {}
func (*HealthCheckSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_hds_41519b306ca672c7, []int{7}
}
func (m *HealthCheckSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckSpecifier.Unmarshal(m, b)
}
func (m *HealthCheckSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckSpecifier.Marshal(b, m, deterministic)
}
func (dst *HealthCheckSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckSpecifier.Merge(dst, src)
}
func (m *HealthCheckSpecifier) XXX_Size() int {
	return xxx_messageInfo_HealthCheckSpecifier.Size(m)
}
func (m *HealthCheckSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckSpecifier proto.InternalMessageInfo

func (m *HealthCheckSpecifier) GetHealthCheck() []*ClusterHealthCheck {
	if m != nil {
		return m.HealthCheck
	}
	return nil
}

func (m *HealthCheckSpecifier) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func init() {
	proto.RegisterType((*Capability)(nil), "envoy.service.discovery.v2.Capability")
	proto.RegisterType((*HealthCheckRequest)(nil), "envoy.service.discovery.v2.HealthCheckRequest")
	proto.RegisterType((*EndpointHealth)(nil), "envoy.service.discovery.v2.EndpointHealth")
	proto.RegisterType((*EndpointHealthResponse)(nil), "envoy.service.discovery.v2.EndpointHealthResponse")
	proto.RegisterType((*HealthCheckRequestOrEndpointHealthResponse)(nil), "envoy.service.discovery.v2.HealthCheckRequestOrEndpointHealthResponse")
	proto.RegisterType((*LocalityEndpoints)(nil), "envoy.service.discovery.v2.LocalityEndpoints")
	proto.RegisterType((*ClusterHealthCheck)(nil), "envoy.service.discovery.v2.ClusterHealthCheck")
	proto.RegisterType((*HealthCheckSpecifier)(nil), "envoy.service.discovery.v2.HealthCheckSpecifier")
	proto.RegisterEnum("envoy.service.discovery.v2.Capability_Protocol", Capability_Protocol_name, Capability_Protocol_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthDiscoveryServiceClient is the client API for HealthDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthDiscoveryServiceClient interface {
	// 1. Envoy starts up and if its can_healthcheck option in the static
	//    bootstrap config is enabled, sends HealthCheckRequest to the management
	//    server. It supplies its capabilities (which protocol it can health check
	//    with, what zone it resides in, etc.).
	// 2. In response to (1), the management server designates this Envoy as a
	//    healthchecker to health check a subset of all upstream hosts for a given
	//    cluster (for example upstream Host 1 and Host 2). It streams
	//    HealthCheckSpecifier messages with cluster related configuration for all
	//    clusters this Envoy is designated to health check. Subsequent
	//    HealthCheckSpecifier message will be sent on changes to:
	//    a. Endpoints to health checks
	//    b. Per cluster configuration change
	// 3. Envoy creates a health probe based on the HealthCheck config and sends
	//    it to endpoint(ip:port) of Host 1 and 2. Based on the HealthCheck
	//    configuration Envoy waits upon the arrival of the probe response and
	//    looks at the content of the response to decide whether the endpoint is
	//    healthy or not. If a response hasn’t been received within the timeout
	//    interval, the endpoint health status is considered TIMEOUT.
	// 4. Envoy reports results back in an EndpointHealthResponse message.
	//    Envoy streams responses as often as the interval configured by the
	//    management server in HealthCheckSpecifier.
	// 5. The management Server collects health statuses for all endpoints in the
	//    cluster (for all clusters) and uses this information to construct
	//    EndpointDiscoveryResponse messages.
	// 6. Once Envoy has a list of upstream endpoints to send traffic to, it load
	//    balances traffic to them without additional health checking. It may
	//    use inline healthcheck (i.e. consider endpoint UNHEALTHY if connection
	//    failed to a particular endpoint to account for health status propagation
	//    delay between HDS and EDS).
	// By default, can_healthcheck is true. If can_healthcheck is false, Cluster
	// configuration may not contain HealthCheck message.
	// TODO(htuch): How is can_healthcheck communicated to CDS to ensure the above
	// invariant?
	// TODO(htuch): Add @amb67's diagram.
	StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error)
	// TODO(htuch): Unlike the gRPC version, there is no stream-based binding of
	// request/response. Should we add an identifier to the HealthCheckSpecifier
	// to bind with the response?
	FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error)
}

type healthDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewHealthDiscoveryServiceClient(cc *grpc.ClientConn) HealthDiscoveryServiceClient {
	return &healthDiscoveryServiceClient{cc}
}

func (c *healthDiscoveryServiceClient) StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HealthDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.HealthDiscoveryService/StreamHealthCheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthDiscoveryServiceStreamHealthCheckClient{stream}
	return x, nil
}

type HealthDiscoveryService_StreamHealthCheckClient interface {
	Send(*HealthCheckRequestOrEndpointHealthResponse) error
	Recv() (*HealthCheckSpecifier, error)
	grpc.ClientStream
}

type healthDiscoveryServiceStreamHealthCheckClient struct {
	grpc.ClientStream
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Send(m *HealthCheckRequestOrEndpointHealthResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Recv() (*HealthCheckSpecifier, error) {
	m := new(HealthCheckSpecifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthDiscoveryServiceClient) FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error) {
	out := new(HealthCheckSpecifier)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v2.HealthDiscoveryService/FetchHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthDiscoveryServiceServer is the server API for HealthDiscoveryService service.
type HealthDiscoveryServiceServer interface {
	// 1. Envoy starts up and if its can_healthcheck option in the static
	//    bootstrap config is enabled, sends HealthCheckRequest to the management
	//    server. It supplies its capabilities (which protocol it can health check
	//    with, what zone it resides in, etc.).
	// 2. In response to (1), the management server designates this Envoy as a
	//    healthchecker to health check a subset of all upstream hosts for a given
	//    cluster (for example upstream Host 1 and Host 2). It streams
	//    HealthCheckSpecifier messages with cluster related configuration for all
	//    clusters this Envoy is designated to health check. Subsequent
	//    HealthCheckSpecifier message will be sent on changes to:
	//    a. Endpoints to health checks
	//    b. Per cluster configuration change
	// 3. Envoy creates a health probe based on the HealthCheck config and sends
	//    it to endpoint(ip:port) of Host 1 and 2. Based on the HealthCheck
	//    configuration Envoy waits upon the arrival of the probe response and
	//    looks at the content of the response to decide whether the endpoint is
	//    healthy or not. If a response hasn’t been received within the timeout
	//    interval, the endpoint health status is considered TIMEOUT.
	// 4. Envoy reports results back in an EndpointHealthResponse message.
	//    Envoy streams responses as often as the interval configured by the
	//    management server in HealthCheckSpecifier.
	// 5. The management Server collects health statuses for all endpoints in the
	//    cluster (for all clusters) and uses this information to construct
	//    EndpointDiscoveryResponse messages.
	// 6. Once Envoy has a list of upstream endpoints to send traffic to, it load
	//    balances traffic to them without additional health checking. It may
	//    use inline healthcheck (i.e. consider endpoint UNHEALTHY if connection
	//    failed to a particular endpoint to account for health status propagation
	//    delay between HDS and EDS).
	// By default, can_healthcheck is true. If can_healthcheck is false, Cluster
	// configuration may not contain HealthCheck message.
	// TODO(htuch): How is can_healthcheck communicated to CDS to ensure the above
	// invariant?
	// TODO(htuch): Add @amb67's diagram.
	StreamHealthCheck(HealthDiscoveryService_StreamHealthCheckServer) error
	// TODO(htuch): Unlike the gRPC version, there is no stream-based binding of
	// request/response. Should we add an identifier to the HealthCheckSpecifier
	// to bind with the response?
	FetchHealthCheck(context.Context, *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error)
}

func RegisterHealthDiscoveryServiceServer(s *grpc.Server, srv HealthDiscoveryServiceServer) {
	s.RegisterService(&_HealthDiscoveryService_serviceDesc, srv)
}

func _HealthDiscoveryService_StreamHealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthDiscoveryServiceServer).StreamHealthCheck(&healthDiscoveryServiceStreamHealthCheckServer{stream})
}

type HealthDiscoveryService_StreamHealthCheckServer interface {
	Send(*HealthCheckSpecifier) error
	Recv() (*HealthCheckRequestOrEndpointHealthResponse, error)
	grpc.ServerStream
}

type healthDiscoveryServiceStreamHealthCheckServer struct {
	grpc.ServerStream
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Send(m *HealthCheckSpecifier) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Recv() (*HealthCheckRequestOrEndpointHealthResponse, error) {
	m := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HealthDiscoveryService_FetchHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.HealthDiscoveryService/FetchHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, req.(*HealthCheckRequestOrEndpointHealthResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.HealthDiscoveryService",
	HandlerType: (*HealthDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHealthCheck",
			Handler:    _HealthDiscoveryService_FetchHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHealthCheck",
			Handler:       _HealthDiscoveryService_StreamHealthCheck_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/hds.proto",
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/hds.proto", fileDescriptor_hds_41519b306ca672c7)
}

var fileDescriptor_hds_41519b306ca672c7 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0xcd, 0x24, 0x7d, 0x8f, 0xf4, 0x26, 0x84, 0x74, 0x28, 0x25, 0x84, 0xaa, 0x2d, 0x56, 0x41,
	0x51, 0x11, 0x76, 0x65, 0x84, 0x90, 0x8a, 0xd8, 0x34, 0x69, 0x15, 0x04, 0x2a, 0x65, 0x52, 0xd6,
	0xd1, 0xc4, 0x99, 0xd6, 0x16, 0xae, 0xc7, 0xd8, 0x93, 0x48, 0xd9, 0xb2, 0x02, 0xb1, 0x41, 0xea,
	0x92, 0x0d, 0xff, 0xc0, 0x57, 0xb0, 0xe6, 0x17, 0xfa, 0x21, 0x28, 0x33, 0x63, 0xd7, 0x6d, 0x9a,
	0xd0, 0xee, 0xd8, 0xc5, 0x33, 0xe7, 0x9c, 0x7b, 0xee, 0x3d, 0xd7, 0x0e, 0xac, 0xb3, 0x60, 0xc8,
	0x47, 0x56, 0xcc, 0xa2, 0xa1, 0xe7, 0x30, 0xab, 0xef, 0xc5, 0x0e, 0x1f, 0xb2, 0x68, 0x64, 0x0d,
	0x6d, 0xcb, 0xed, 0xc7, 0x66, 0x18, 0x71, 0xc1, 0x71, 0x5d, 0xa2, 0x4c, 0x8d, 0x32, 0x53, 0x94,
	0x39, 0xb4, 0xeb, 0xcb, 0x4a, 0x81, 0x86, 0xde, 0x98, 0xe3, 0xf0, 0x88, 0x59, 0x3d, 0x1a, 0x33,
	0xc5, 0xac, 0xaf, 0x4f, 0xde, 0xba, 0x8c, 0xfa, 0xc2, 0xed, 0x3a, 0x2e, 0x73, 0x3e, 0x5c, 0x8a,
	0x62, 0x41, 0x3f, 0xe4, 0x5e, 0x20, 0xd2, 0x1f, 0x1a, 0xb5, 0x7c, 0xc4, 0xf9, 0x91, 0xcf, 0x24,
	0x8c, 0x06, 0x01, 0x17, 0x54, 0x78, 0x3c, 0xd0, 0x1e, 0xeb, 0x2b, 0xfa, 0x56, 0x3e, 0xf5, 0x06,
	0x87, 0x56, 0x7f, 0x10, 0x49, 0x80, 0xba, 0x37, 0xbe, 0x23, 0x80, 0x26, 0x0d, 0x69, 0xcf, 0xf3,
	0x3d, 0x31, 0xc2, 0x0e, 0xdc, 0xc9, 0x1a, 0xe9, 0x4a, 0x90, 0xc3, 0xfd, 0x1a, 0x5a, 0x2b, 0x34,
	0x2a, 0xb6, 0x65, 0x4e, 0x6f, 0xd9, 0x3c, 0x93, 0x31, 0xf7, 0x35, 0x8d, 0xdc, 0x56, 0x6a, 0xcd,
	0xb1, 0x58, 0x72, 0x68, 0x34, 0xa0, 0x98, 0xfc, 0xc6, 0x45, 0x98, 0x6b, 0x1f, 0x1c, 0xec, 0x57,
	0x73, 0xf8, 0x06, 0x14, 0x0e, 0x9a, 0xfb, 0x55, 0x84, 0xe7, 0xe1, 0x3f, 0xb2, 0xd3, 0x7a, 0xd5,
	0xa9, 0xe6, 0x8d, 0x2f, 0x08, 0x70, 0xfb, 0x4c, 0x81, 0xb0, 0x8f, 0x03, 0x16, 0x0b, 0xfc, 0x18,
	0xe6, 0x02, 0xde, 0x67, 0x35, 0xb4, 0x86, 0x1a, 0x25, 0xfb, 0xae, 0x36, 0x45, 0x43, 0x6f, 0x6c,
	0x63, 0x3c, 0x4d, 0x73, 0x8f, 0xf7, 0x19, 0x91, 0x20, 0xbc, 0x0b, 0xe0, 0xa4, 0xce, 0x6a, 0x79,
	0x49, 0x79, 0x74, 0xb5, 0x3e, 0x48, 0x86, 0x69, 0x9c, 0x20, 0xa8, 0xec, 0xe8, 0xd1, 0x2b, 0x4f,
	0xf8, 0x05, 0x14, 0x93, 0x30, 0xb4, 0x97, 0xd5, 0xf3, 0x5e, 0xd2, 0xa8, 0x12, 0x22, 0x49, 0x09,
	0xb8, 0x05, 0x37, 0xf5, 0xa8, 0x63, 0x41, 0xc5, 0x20, 0x96, 0xd6, 0x2a, 0x17, 0x15, 0x64, 0x37,
	0xaa, 0x5c, 0x47, 0xc2, 0x48, 0xd9, 0xcd, 0x3c, 0x19, 0x1c, 0x96, 0xce, 0x9b, 0x22, 0x2c, 0x0e,
	0x79, 0x10, 0x33, 0xfc, 0x1e, 0xaa, 0x49, 0xad, 0xb8, 0xab, 0x38, 0x32, 0xc5, 0x92, 0xbd, 0x31,
	0xab, 0xfb, 0x0b, 0x6a, 0xb7, 0x52, 0x0d, 0x75, 0x60, 0x7c, 0xcb, 0xc3, 0xc6, 0x64, 0x24, 0x6f,
	0xa3, 0x29, 0x2e, 0x7a, 0xb0, 0x78, 0x6e, 0xa1, 0x22, 0x85, 0xd7, 0xe3, 0x32, 0x67, 0x39, 0x99,
	0xac, 0xd2, 0xce, 0x11, 0xec, 0x4e, 0xae, 0x43, 0x00, 0xb5, 0xc4, 0xa5, 0x6e, 0xb4, 0x1b, 0xe9,
	0xfa, 0x3a, 0x6f, 0xfb, 0x1a, 0x1d, 0x6b, 0x66, 0x3b, 0x47, 0x96, 0xd8, 0xa5, 0x37, 0xdb, 0x15,
	0x28, 0xeb, 0x36, 0xba, 0x62, 0x14, 0x32, 0xe3, 0x2b, 0x82, 0x85, 0x37, 0xdc, 0xa1, 0xe3, 0x35,
	0x49, 0xc4, 0x62, 0xfc, 0x1c, 0x8a, 0xbe, 0x3e, 0xd4, 0xdd, 0xde, 0xbf, 0x24, 0xda, 0x84, 0x47,
	0x52, 0x30, 0x7e, 0x09, 0xf3, 0xe9, 0xd0, 0x6b, 0x79, 0x99, 0xd8, 0x5f, 0xd7, 0xea, 0x8c, 0x61,
	0xfc, 0x42, 0x80, 0x9b, 0xfe, 0x20, 0x16, 0x2c, 0xca, 0x4c, 0x10, 0x3f, 0x80, 0xb2, 0xa3, 0x4e,
	0xbb, 0x01, 0x3d, 0x56, 0xef, 0xce, 0x3c, 0x29, 0xe9, 0xb3, 0x3d, 0x7a, 0xcc, 0x70, 0x33, 0xdd,
	0x48, 0x99, 0x55, 0x52, 0x7c, 0x65, 0xea, 0x46, 0xaa, 0x14, 0xca, 0x99, 0x48, 0x62, 0xfc, 0x3a,
	0xeb, 0xbe, 0x20, 0x05, 0x9e, 0xcc, 0x9a, 0xfe, 0xc4, 0xe0, 0xb2, 0xbd, 0xfc, 0x40, 0xb0, 0x98,
	0x29, 0xd5, 0x09, 0x99, 0xe3, 0x1d, 0x7a, 0x2c, 0xc2, 0xef, 0xa0, 0x9c, 0xb5, 0xaa, 0x17, 0x7b,
	0xe6, 0x3a, 0x4d, 0xce, 0x84, 0x94, 0x32, 0xce, 0xf1, 0x33, 0x28, 0x7a, 0x81, 0x60, 0xd1, 0x90,
	0xfa, 0x7a, 0x6b, 0xee, 0x99, 0xea, 0xe3, 0x69, 0x26, 0x1f, 0x4f, 0xb3, 0xa5, 0x3f, 0x9e, 0x24,
	0x85, 0xda, 0xa7, 0x79, 0x58, 0x52, 0x9a, 0xad, 0xa4, 0x5a, 0x47, 0x95, 0xc7, 0x27, 0x08, 0x16,
	0x3a, 0x22, 0x62, 0xf4, 0x38, 0x1b, 0xc4, 0xee, 0xf5, 0x76, 0x7e, 0xda, 0x9b, 0x55, 0xdf, 0xbc,
	0xa2, 0x4e, 0x3a, 0x34, 0x23, 0xd7, 0x40, 0x9b, 0x08, 0xff, 0x44, 0x50, 0xdd, 0x65, 0xc2, 0x71,
	0xff, 0x0d, 0x53, 0x0f, 0x3f, 0xfd, 0x3e, 0x3d, 0xc9, 0xaf, 0x1a, 0xf5, 0xf1, 0x7f, 0x5c, 0x0a,
	0xdf, 0xca, 0xa6, 0xbb, 0x85, 0x36, 0xb6, 0x0b, 0x9f, 0x11, 0xea, 0xfd, 0x2f, 0x83, 0x78, 0xfa,
	0x27, 0x00, 0x00, 0xff, 0xff, 0xed, 0x53, 0x7e, 0x52, 0x9f, 0x07, 0x00, 0x00,
}
