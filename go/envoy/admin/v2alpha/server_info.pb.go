// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/server_info.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
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

type ServerInfo_State int32

const (
	// Server is live and serving traffic.
	ServerInfo_LIVE ServerInfo_State = 0
	// Server is draining listeners in response to external health checks failing.
	ServerInfo_DRAINING ServerInfo_State = 1
)

var ServerInfo_State_name = map[int32]string{
	0: "LIVE",
	1: "DRAINING",
}

var ServerInfo_State_value = map[string]int32{
	"LIVE":     0,
	"DRAINING": 1,
}

func (x ServerInfo_State) String() string {
	return proto.EnumName(ServerInfo_State_name, int32(x))
}

func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{0, 0}
}

// Proto representation of the value returned by /server_info, containing
// server version/server status information.
type ServerInfo struct {
	// Server version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// State of the server.
	State ServerInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v2alpha.ServerInfo_State" json:"state,omitempty"`
	// Uptime since current epoch was started.
	UptimeCurrentEpoch *duration.Duration `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	// Uptime since the start of the first epoch.
	UptimeAllEpochs *duration.Duration `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	// Which restart epoch the server is currently in.
	Epoch                uint32   `protobuf:"varint,5,opt,name=epoch,proto3" json:"epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetState() ServerInfo_State {
	if m != nil {
		return m.State
	}
	return ServerInfo_LIVE
}

func (m *ServerInfo) GetUptimeCurrentEpoch() *duration.Duration {
	if m != nil {
		return m.UptimeCurrentEpoch
	}
	return nil
}

func (m *ServerInfo) GetUptimeAllEpochs() *duration.Duration {
	if m != nil {
		return m.UptimeAllEpochs
	}
	return nil
}

func (m *ServerInfo) GetEpoch() uint32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.admin.v2alpha.ServerInfo_State", ServerInfo_State_name, ServerInfo_State_value)
	proto.RegisterType((*ServerInfo)(nil), "envoy.admin.v2alpha.ServerInfo")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/server_info.proto", fileDescriptor_ed0f406f9d75bf97)
}

var fileDescriptor_ed0f406f9d75bf97 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0x87, 0xcd, 0x5c, 0x75, 0xbe, 0xfe, 0x9b, 0x71, 0x87, 0xea, 0x41, 0xcb, 0x60, 0xd0, 0x53,
	0x0a, 0xf5, 0xe8, 0x69, 0xb8, 0x22, 0x45, 0xd9, 0xa1, 0x03, 0xaf, 0x25, 0xdb, 0xd2, 0xad, 0x90,
	0xe5, 0x2d, 0x69, 0x5a, 0xf0, 0x4b, 0xf9, 0x19, 0x65, 0x49, 0x87, 0x17, 0x61, 0xc7, 0x17, 0x7e,
	0xcf, 0xf3, 0x3e, 0x30, 0x11, 0xaa, 0xc5, 0xef, 0x88, 0xaf, 0x77, 0xa5, 0x8a, 0xda, 0x98, 0xcb,
	0x6a, 0xcb, 0xa3, 0x5a, 0xe8, 0x56, 0xe8, 0xbc, 0x54, 0x05, 0xb2, 0x4a, 0xa3, 0x41, 0x7a, 0x6f,
	0x67, 0xcc, 0xce, 0x58, 0x37, 0x7b, 0x7c, 0xda, 0x20, 0x6e, 0xa4, 0x88, 0xec, 0x64, 0xd9, 0x14,
	0xd1, 0xba, 0xd1, 0xdc, 0x94, 0xa8, 0x1c, 0x34, 0xfe, 0xe9, 0x01, 0x2c, 0xac, 0x2a, 0x55, 0x05,
	0x52, 0x1f, 0xce, 0x5b, 0xa1, 0xeb, 0x12, 0x95, 0x4f, 0x02, 0x12, 0x5e, 0x64, 0x87, 0x93, 0xbe,
	0x82, 0x57, 0x1b, 0x6e, 0x84, 0xdf, 0x0b, 0x48, 0x78, 0x13, 0x4f, 0xd8, 0x3f, 0xdf, 0xd8, 0x9f,
	0x89, 0x2d, 0xf6, 0xe3, 0xcc, 0x31, 0xf4, 0x03, 0x46, 0x4d, 0x65, 0xca, 0x9d, 0xc8, 0x57, 0x8d,
	0xd6, 0x42, 0x99, 0x5c, 0x54, 0xb8, 0xda, 0xfa, 0xa7, 0x01, 0x09, 0x2f, 0xe3, 0x07, 0xe6, 0x22,
	0xd9, 0x21, 0x92, 0xcd, 0xba, 0xc8, 0x8c, 0x3a, 0xec, 0xcd, 0x51, 0xc9, 0x1e, 0xa2, 0x09, 0xdc,
	0x75, 0x32, 0x2e, 0xa5, 0x13, 0xd5, 0x7e, 0xff, 0x98, 0xe9, 0xd6, 0x31, 0x53, 0x29, 0xad, 0xa5,
	0xa6, 0x23, 0xf0, 0x5c, 0x84, 0x17, 0x90, 0xf0, 0x3a, 0x73, 0xc7, 0xf8, 0x19, 0x3c, 0x5b, 0x4e,
	0x07, 0xd0, 0xff, 0x4c, 0xbf, 0x92, 0xe1, 0x09, 0xbd, 0x82, 0xc1, 0x2c, 0x9b, 0xa6, 0xf3, 0x74,
	0xfe, 0x3e, 0x24, 0xcb, 0x33, 0xab, 0x7e, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x62, 0xdb, 0x72,
	0x22, 0x95, 0x01, 0x00, 0x00,
}
