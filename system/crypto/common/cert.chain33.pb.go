// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cert.chain.proto

package cert_chain33

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TODO 这个结构需要迁移到crypto/types.go中
type CertSignature struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Cert                 []byte   `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
	Uid                  []byte   `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertSignature) Reset()         { *m = CertSignature{} }
func (m *CertSignature) String() string { return proto.CompactTextString(m) }
func (*CertSignature) ProtoMessage()    {}
func (*CertSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_543d0dd13c546879, []int{0}
}

func (m *CertSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertSignature.Unmarshal(m, b)
}
func (m *CertSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertSignature.Marshal(b, m, deterministic)
}
func (m *CertSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertSignature.Merge(m, src)
}
func (m *CertSignature) XXX_Size() int {
	return xxx_messageInfo_CertSignature.Size(m)
}
func (m *CertSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_CertSignature.DiscardUnknown(m)
}

var xxx_messageInfo_CertSignature proto.InternalMessageInfo

func (m *CertSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CertSignature) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *CertSignature) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func init() {
	proto.RegisterType((*CertSignature)(nil), "CertSignature")
}

func init() { proto.RegisterFile("cert.chain.proto", fileDescriptor_543d0dd13c546879) }

var fileDescriptor_543d0dd13c546879 = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4e, 0x2d, 0x2a,
	0xd1, 0x4b, 0xce, 0x48, 0xcc, 0xcc, 0x33, 0x36, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x0a,
	0xe6, 0xe2, 0x75, 0x4e, 0x2d, 0x2a, 0x09, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0x15,
	0x92, 0xe1, 0xe2, 0x2c, 0x86, 0x71, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x10, 0x02, 0x42,
	0x42, 0x5c, 0x2c, 0x20, 0x43, 0x24, 0x98, 0xc0, 0x12, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x69,
	0x66, 0x8a, 0x04, 0x33, 0x58, 0x08, 0xc4, 0x4c, 0x62, 0x03, 0x9b, 0x6d, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x87, 0x45, 0x0a, 0x71, 0x00, 0x00, 0x00,
}
