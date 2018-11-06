// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/certs.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Proto representation of certificate details. Admin endpoint uses this wrapper for `/certs` to
// display certificate information. See :ref:`/certs <operations_admin_interface_certs>` for more
// information.
type Certificates struct {
	// List of certificates known to an Envoy.
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{0}
}

func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificates.Unmarshal(m, b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return xxx_messageInfo_Certificates.Size(m)
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	// Details of CA certificate.
	CaCert []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// Details of Certificate Chain
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{1}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

type CertificateDetails struct {
	// Path of the certificate.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Certificate Serial Number.
	SerialNumber string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// List of Subject Alternate names.
	SubjectAltNames []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	// Minimum of days until expiration of certificate and it's chain.
	DaysUntilExpiration uint64 `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	// Indicates the time from which the certificate is valid.
	ValidFrom *timestamp.Timestamp `protobuf:"bytes,5,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	// Indicates the time at which the certificate expires.
	ExpirationTime       *timestamp.Timestamp `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{2}
}

func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateDetails.Unmarshal(m, b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
}
func (m *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(m, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return xxx_messageInfo_CertificateDetails.Size(m)
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

func (m *CertificateDetails) GetValidFrom() *timestamp.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *CertificateDetails) GetExpirationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

type SubjectAlternateName struct {
	// Subject Alternate Name.
	//
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{3}
}

func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectAlternateName.Unmarshal(m, b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
}
func (m *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(m, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return xxx_messageInfo_SubjectAlternateName.Size(m)
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof"`
}

type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SubjectAlternateName_OneofMarshaler, _SubjectAlternateName_OneofUnmarshaler, _SubjectAlternateName_OneofSizer, []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
	}
}

func _SubjectAlternateName_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SubjectAlternateName)
	// name
	switch x := m.Name.(type) {
	case *SubjectAlternateName_Dns:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Dns)
	case *SubjectAlternateName_Uri:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Uri)
	case nil:
	default:
		return fmt.Errorf("SubjectAlternateName.Name has unexpected type %T", x)
	}
	return nil
}

func _SubjectAlternateName_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SubjectAlternateName)
	switch tag {
	case 1: // name.dns
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Name = &SubjectAlternateName_Dns{x}
		return true, err
	case 2: // name.uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Name = &SubjectAlternateName_Uri{x}
		return true, err
	default:
		return false, nil
	}
}

func _SubjectAlternateName_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SubjectAlternateName)
	// name
	switch x := m.Name.(type) {
	case *SubjectAlternateName_Dns:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Dns)))
		n += len(x.Dns)
	case *SubjectAlternateName_Uri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Uri)))
		n += len(x.Uri)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v2alpha.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v2alpha.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v2alpha.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v2alpha.SubjectAlternateName")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/certs.proto", fileDescriptor_c7cd1796e05ff7fa) }

var fileDescriptor_c7cd1796e05ff7fa = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xd9, 0x24, 0x04, 0x65, 0x12, 0xa8, 0x70, 0x41, 0xb2, 0x7a, 0xe9, 0x2a, 0x1c, 0x08,
	0x17, 0xaf, 0x14, 0x4e, 0xdc, 0x80, 0x94, 0x88, 0x53, 0x0f, 0x4b, 0x7b, 0xb6, 0x66, 0x77, 0x9d,
	0xc6, 0xc8, 0x7f, 0x56, 0xb6, 0x37, 0xa2, 0x4f, 0xc2, 0xab, 0xf1, 0x38, 0xc8, 0xde, 0xfe, 0x09,
	0x22, 0x52, 0xd4, 0xdb, 0xce, 0xf7, 0x7d, 0xbf, 0xf1, 0xec, 0x68, 0xe0, 0x5c, 0x98, 0x9d, 0xbd,
	0x2d, 0xb0, 0xd1, 0xd2, 0x14, 0xbb, 0x25, 0xaa, 0x76, 0x8b, 0x45, 0x2d, 0x5c, 0xf0, 0xac, 0x75,
	0x36, 0x58, 0x72, 0x9a, 0x02, 0x2c, 0x05, 0xd8, 0x5d, 0xe0, 0xec, 0xfc, 0xc6, 0xda, 0x1b, 0x25,
	0x8a, 0x14, 0xa9, 0xba, 0x4d, 0x11, 0xa4, 0x16, 0x3e, 0xa0, 0x6e, 0x7b, 0x6a, 0x7e, 0x05, 0xb3,
	0x95, 0x70, 0x41, 0x6e, 0x64, 0x8d, 0x41, 0x78, 0x72, 0x01, 0xb3, 0x7a, 0xaf, 0xa6, 0x59, 0x3e,
	0x5c, 0x4c, 0x97, 0x39, 0x3b, 0xd0, 0x9c, 0xed, 0x81, 0xe5, 0x3f, 0xd4, 0xfc, 0x77, 0x06, 0xd3,
	0x3d, 0x97, 0x7c, 0x86, 0x17, 0x35, 0xf2, 0x18, 0xb9, 0x6b, 0xf8, 0xfe, 0x58, 0xc3, 0x0b, 0x11,
	0x50, 0x2a, 0x5f, 0x8e, 0x6b, 0x8c, 0x2a, 0x59, 0x03, 0x44, 0x9c, 0xd7, 0x5b, 0x94, 0x86, 0x0e,
	0x9e, 0xd6, 0x64, 0x12, 0xd1, 0x55, 0x24, 0xe7, 0x7f, 0x06, 0x40, 0xfe, 0x4f, 0x10, 0x02, 0xa3,
	0x16, 0xc3, 0x96, 0x66, 0x79, 0xb6, 0x98, 0x94, 0xe9, 0x9b, 0xbc, 0x83, 0x97, 0x5e, 0x38, 0x89,
	0x8a, 0x9b, 0x4e, 0x57, 0xc2, 0xd1, 0x41, 0x32, 0x67, 0xbd, 0x78, 0x99, 0x34, 0x72, 0x0d, 0xaf,
	0x7d, 0x57, 0xfd, 0x14, 0x75, 0xe0, 0xa8, 0x02, 0x37, 0xa8, 0x85, 0xa7, 0xc3, 0x34, 0xde, 0x87,
	0x83, 0xe3, 0xfd, 0xe8, 0xd3, 0x5f, 0x54, 0x10, 0xce, 0x60, 0x10, 0x97, 0xa8, 0x45, 0x79, 0xe2,
	0x1f, 0xd4, 0x58, 0x7b, 0xb2, 0x84, 0xb7, 0x0d, 0xde, 0x7a, 0xde, 0x99, 0x20, 0x15, 0x17, 0xbf,
	0x5a, 0xe9, 0x30, 0x48, 0x6b, 0xe8, 0x28, 0xcf, 0x16, 0xa3, 0xf2, 0x34, 0x9a, 0xd7, 0xd1, 0xfb,
	0xf6, 0x60, 0x91, 0x4f, 0x00, 0x3b, 0x54, 0xb2, 0xe1, 0x1b, 0x67, 0x35, 0x7d, 0x9e, 0x67, 0x8b,
	0xe9, 0xf2, 0x8c, 0xf5, 0x07, 0xc0, 0xee, 0x0f, 0x80, 0x5d, 0xdd, 0x1f, 0x40, 0x39, 0x49, 0xe9,
	0xb5, 0xb3, 0x9a, 0xac, 0xe0, 0xe4, 0xf1, 0x0d, 0x1e, 0x6f, 0x84, 0x8e, 0x8f, 0xf2, 0xaf, 0x1e,
	0x91, 0x28, 0xce, 0xd7, 0xf0, 0xe6, 0xd0, 0xcf, 0x11, 0x02, 0xc3, 0xc6, 0xf8, 0x7e, 0xb5, 0xdf,
	0x9f, 0x95, 0xb1, 0x88, 0x5a, 0xe7, 0x64, 0xbf, 0xd1, 0xa8, 0x75, 0x4e, 0x7e, 0x1d, 0xc3, 0x28,
	0xae, 0xaf, 0x1a, 0xa7, 0xb7, 0x3e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xcf, 0xc9, 0x19,
	0xf2, 0x02, 0x00, 0x00,
}
