// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/gzip/v2/gzip.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}
var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}
func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gzip_e2d8255c3cc4e261, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}
var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}
func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gzip_e2d8255c3cc4e261, []int{0, 0, 0}
}

type Gzip struct {
	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. The default value is 5.
	MemoryLevel *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	// Minimum response length, in bytes, which will trigger compression. The default value is 30.
	ContentLength *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	// A value used for selecting the zlib compression level. This setting will affect speed and
	// amount of compression applied to the content. "BEST" provides higher compression at the cost of
	// higher latency, "SPEED" provides lower compression with minimum impact on response time.
	// "DEFAULT" provides an optimal result between speed and compression. This field will be set to
	// "DEFAULT" if not specified.
	CompressionLevel Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	// A value used for selecting the zlib compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT" will be the best choice, though
	// there are situations which changing this parameter might produce better results. For example,
	// run-length encoding (RLE) is typically used when the content is known for having sequences
	// which same data occurs many consecutive times. For more information about each strategy, please
	// refer to zlib manual.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Set of strings that allows specifying which mime-types yield compression; e.g.,
	// application/json, text/html, etc. When this field is not defined, compression will be applied
	// to the following mime-types: "application/javascript", "application/json",
	// "application/xhtml+xml", "image/svg+xml", "text/css", "text/html", "text/plain", "text/xml".
	ContentType []string `protobuf:"bytes,6,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, disables compression when the response contains an etag header. When it is false, the
	// filter will preserve weak etags and remove the ones that require strong validation.
	DisableOnEtagHeader bool `protobuf:"varint,7,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	// If true, removes accept-encoding from the request headers before dispatching it to the upstream
	// so that responses do not get compressed before reaching the filter.
	RemoveAcceptEncodingHeader bool `protobuf:"varint,8,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits           *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_gzip_e2d8255c3cc4e261, []int{0}
}
func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (dst *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(dst, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetMemoryLevel() *wrappers.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetContentLength() *wrappers.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Gzip) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Gzip) GetRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.RemoveAcceptEncodingHeader
	}
	return false
}

func (m *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

type Gzip_CompressionLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gzip_CompressionLevel) Reset()         { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()    {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_gzip_e2d8255c3cc4e261, []int{0, 0}
}
func (m *Gzip_CompressionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip_CompressionLevel.Unmarshal(m, b)
}
func (m *Gzip_CompressionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip_CompressionLevel.Marshal(b, m, deterministic)
}
func (dst *Gzip_CompressionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip_CompressionLevel.Merge(dst, src)
}
func (m *Gzip_CompressionLevel) XXX_Size() int {
	return xxx_messageInfo_Gzip_CompressionLevel.Size(m)
}
func (m *Gzip_CompressionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip_CompressionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip_CompressionLevel proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Gzip)(nil), "envoy.config.filter.http.gzip.v2.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.config.filter.http.gzip.v2.Gzip.CompressionLevel")
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/gzip/v2/gzip.proto", fileDescriptor_gzip_e2d8255c3cc4e261)
}

var fileDescriptor_gzip_e2d8255c3cc4e261 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x8e, 0x93, 0x4c,
	0x14, 0xc7, 0x3f, 0x28, 0xdb, 0x2d, 0xc3, 0x7e, 0x2b, 0xce, 0x6e, 0x94, 0x34, 0xba, 0x69, 0x7a,
	0x45, 0xd6, 0x38, 0x24, 0xf4, 0xce, 0xec, 0x4d, 0x6b, 0xa9, 0xbb, 0x06, 0x57, 0x43, 0x5b, 0x2f,
	0xbc, 0x21, 0x94, 0x9e, 0x52, 0x12, 0x3a, 0x43, 0x60, 0x4a, 0xc3, 0x5e, 0xfa, 0x02, 0x26, 0x3e,
	0x8e, 0x57, 0xbe, 0x8e, 0x6f, 0x61, 0x18, 0xa8, 0xba, 0xab, 0x89, 0xee, 0x15, 0x13, 0xce, 0xef,
	0x77, 0xfe, 0x87, 0x33, 0xa0, 0x67, 0x40, 0x0b, 0x56, 0x5a, 0x21, 0xa3, 0xab, 0x38, 0xb2, 0x56,
	0x71, 0xc2, 0x21, 0xb3, 0xd6, 0x9c, 0xa7, 0x56, 0x74, 0x13, 0xa7, 0x56, 0x61, 0x8b, 0x27, 0x49,
	0x33, 0xc6, 0x19, 0xee, 0x09, 0x98, 0xd4, 0x30, 0xa9, 0x61, 0x52, 0xc1, 0x44, 0x40, 0x85, 0xdd,
	0x3d, 0x8b, 0x18, 0x8b, 0x12, 0xb0, 0x04, 0xbf, 0xd8, 0xae, 0xac, 0x5d, 0x16, 0xa4, 0x29, 0x64,
	0x79, 0xdd, 0xa1, 0xfb, 0xb8, 0x08, 0x92, 0x78, 0x19, 0x70, 0xb0, 0xf6, 0x87, 0xa6, 0x70, 0x1a,
	0xb1, 0x88, 0x89, 0xa3, 0x55, 0x9d, 0xea, 0xb7, 0xfd, 0x4f, 0x6d, 0xa4, 0xbc, 0xba, 0x89, 0x53,
	0xec, 0xa2, 0xa3, 0x0d, 0x6c, 0x58, 0x56, 0xfa, 0x09, 0x14, 0x90, 0x18, 0x52, 0x4f, 0x32, 0x35,
	0xfb, 0x09, 0xa9, 0xe3, 0xc8, 0x3e, 0x8e, 0xcc, 0xaf, 0x28, 0x1f, 0xd8, 0xef, 0x83, 0x64, 0x0b,
	0x23, 0xed, 0xcb, 0xb7, 0xaf, 0xad, 0xf6, 0xb9, 0x62, 0xa8, 0xa6, 0xe4, 0x69, 0xb5, 0xee, 0x56,
	0x36, 0xbe, 0x46, 0xc7, 0x21, 0xa3, 0x1c, 0x28, 0xf7, 0x13, 0xa0, 0x11, 0x5f, 0x1b, 0xf2, 0x3f,
	0xf4, 0x53, 0xab, 0x7e, 0xca, 0xb9, 0x6c, 0x9e, 0x79, 0xff, 0x37, 0xba, 0x2b, 0x6c, 0xbc, 0x45,
	0x0f, 0x43, 0xb6, 0x49, 0x33, 0xc8, 0xf3, 0x98, 0xd1, 0x66, 0xc4, 0x56, 0x4f, 0x32, 0x8f, 0xed,
	0x0b, 0xf2, 0xb7, 0x9d, 0x91, 0xea, 0x03, 0xc9, 0xcb, 0x9f, 0xbe, 0x98, 0x91, 0x38, 0x74, 0xbb,
	0x19, 0xa1, 0x2a, 0xf2, 0xe0, 0xa3, 0x24, 0xeb, 0x92, 0xa7, 0x87, 0x77, 0x10, 0x5c, 0xa2, 0xd3,
	0x5f, 0x63, 0x73, 0x9e, 0x05, 0x1c, 0xa2, 0xd2, 0x50, 0x44, 0xf2, 0x8b, 0xfb, 0x27, 0x4f, 0x9b,
	0x0e, 0xb7, 0x72, 0x4f, 0xc2, 0xdf, 0x01, 0xfc, 0x1c, 0x1d, 0xed, 0x37, 0xc8, 0xcb, 0x14, 0x8c,
	0x76, 0xaf, 0x65, 0xaa, 0x8d, 0xf6, 0x59, 0x92, 0x75, 0xdb, 0xd3, 0x9a, 0xfa, 0xac, 0x4c, 0x01,
	0x0f, 0xd0, 0xa3, 0x65, 0x9c, 0x07, 0x8b, 0x04, 0x7c, 0x46, 0x7d, 0xe0, 0x41, 0xe4, 0xaf, 0x21,
	0x58, 0x42, 0x66, 0x1c, 0xf6, 0x24, 0xb3, 0xe3, 0x9d, 0x34, 0xd5, 0xb7, 0xd4, 0xe1, 0x41, 0x74,
	0x29, 0x4a, 0x78, 0x88, 0x9e, 0x66, 0xb0, 0x61, 0x05, 0xf8, 0x41, 0x18, 0x42, 0xca, 0x7d, 0xa0,
	0x21, 0x5b, 0xc6, 0xf4, 0x87, 0xdb, 0x11, 0x6e, 0xb7, 0x86, 0x86, 0x82, 0x71, 0x1a, 0xa4, 0x69,
	0xf1, 0x1a, 0x69, 0xbb, 0x98, 0x2e, 0xd9, 0xce, 0x5f, 0xc4, 0x3c, 0x37, 0xd4, 0xfb, 0xfc, 0x35,
	0x0f, 0x4c, 0xd5, 0x43, 0xb5, 0x3d, 0x8a, 0x79, 0xde, 0xbd, 0x40, 0xfa, 0xdd, 0x4b, 0xea, 0x9b,
	0x48, 0xa9, 0xee, 0x09, 0x6b, 0xe8, 0x70, 0xec, 0x4c, 0x86, 0x73, 0x77, 0xa6, 0xff, 0x87, 0x3b,
	0x48, 0x19, 0x39, 0xd3, 0x99, 0x2e, 0x61, 0x15, 0x1d, 0x4c, 0xdf, 0x39, 0xce, 0x58, 0x97, 0xfb,
	0x13, 0x74, 0xf2, 0x87, 0x45, 0xdf, 0x16, 0x8f, 0x50, 0x67, 0x72, 0xe5, 0xce, 0x1c, 0xcf, 0x19,
	0xeb, 0x52, 0x55, 0xba, 0x9c, 0x4f, 0x26, 0x6f, 0x86, 0xd7, 0xba, 0x8c, 0x0f, 0x51, 0xcb, 0x73,
	0x1d, 0xbd, 0x35, 0x52, 0x3e, 0xc8, 0x85, 0xbd, 0x68, 0x8b, 0xd1, 0x07, 0xdf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x51, 0xd9, 0x19, 0xa1, 0xbe, 0x03, 0x00, 0x00,
}
