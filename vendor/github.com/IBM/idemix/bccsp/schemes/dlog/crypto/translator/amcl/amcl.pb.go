// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bccsp/schemes/dlog/crypto/translator/amcl/amcl.proto

package amcl

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

// ECP is an elliptic curve point specified by its coordinates
// ECP corresponds to an element of the first group (G1)
type ECP struct {
	X                    []byte   `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    []byte   `protobuf:"bytes,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ECP) Reset()         { *m = ECP{} }
func (m *ECP) String() string { return proto.CompactTextString(m) }
func (*ECP) ProtoMessage()    {}
func (*ECP) Descriptor() ([]byte, []int) {
	return fileDescriptor_250ddfa5c5f8dbbb, []int{0}
}

func (m *ECP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ECP.Unmarshal(m, b)
}
func (m *ECP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ECP.Marshal(b, m, deterministic)
}
func (m *ECP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ECP.Merge(m, src)
}
func (m *ECP) XXX_Size() int {
	return xxx_messageInfo_ECP.Size(m)
}
func (m *ECP) XXX_DiscardUnknown() {
	xxx_messageInfo_ECP.DiscardUnknown(m)
}

var xxx_messageInfo_ECP proto.InternalMessageInfo

func (m *ECP) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *ECP) GetY() []byte {
	if m != nil {
		return m.Y
	}
	return nil
}

// ECP2 is an elliptic curve point specified by its coordinates
// ECP2 corresponds to an element of the second group (G2)
type ECP2 struct {
	Xa                   []byte   `protobuf:"bytes,1,opt,name=xa,proto3" json:"xa,omitempty"`
	Xb                   []byte   `protobuf:"bytes,2,opt,name=xb,proto3" json:"xb,omitempty"`
	Ya                   []byte   `protobuf:"bytes,3,opt,name=ya,proto3" json:"ya,omitempty"`
	Yb                   []byte   `protobuf:"bytes,4,opt,name=yb,proto3" json:"yb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ECP2) Reset()         { *m = ECP2{} }
func (m *ECP2) String() string { return proto.CompactTextString(m) }
func (*ECP2) ProtoMessage()    {}
func (*ECP2) Descriptor() ([]byte, []int) {
	return fileDescriptor_250ddfa5c5f8dbbb, []int{1}
}

func (m *ECP2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ECP2.Unmarshal(m, b)
}
func (m *ECP2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ECP2.Marshal(b, m, deterministic)
}
func (m *ECP2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ECP2.Merge(m, src)
}
func (m *ECP2) XXX_Size() int {
	return xxx_messageInfo_ECP2.Size(m)
}
func (m *ECP2) XXX_DiscardUnknown() {
	xxx_messageInfo_ECP2.DiscardUnknown(m)
}

var xxx_messageInfo_ECP2 proto.InternalMessageInfo

func (m *ECP2) GetXa() []byte {
	if m != nil {
		return m.Xa
	}
	return nil
}

func (m *ECP2) GetXb() []byte {
	if m != nil {
		return m.Xb
	}
	return nil
}

func (m *ECP2) GetYa() []byte {
	if m != nil {
		return m.Ya
	}
	return nil
}

func (m *ECP2) GetYb() []byte {
	if m != nil {
		return m.Yb
	}
	return nil
}

func init() {
	proto.RegisterType((*ECP)(nil), "amcl.ECP")
	proto.RegisterType((*ECP2)(nil), "amcl.ECP2")
}

func init() {
	proto.RegisterFile("bccsp/schemes/dlog/crypto/translator/amcl/amcl.proto", fileDescriptor_250ddfa5c5f8dbbb)
}

var fileDescriptor_250ddfa5c5f8dbbb = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0x4a, 0x4e, 0x2e,
	0x2e, 0xd0, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2d, 0xd6, 0x4f, 0xc9, 0xc9, 0x4f, 0xd7, 0x4f,
	0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xd7, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0xce, 0x49, 0x2c, 0xc9, 0x2f,
	0xd2, 0x4f, 0xcc, 0x4d, 0xce, 0x01, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20,
	0xb6, 0x92, 0x22, 0x17, 0xb3, 0xab, 0x73, 0x80, 0x10, 0x0f, 0x17, 0x63, 0x85, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x4f, 0x10, 0x63, 0x05, 0x88, 0x57, 0x29, 0xc1, 0x04, 0xe1, 0x55, 0x2a, 0xb9, 0x71,
	0xb1, 0xb8, 0x3a, 0x07, 0x18, 0x09, 0xf1, 0x71, 0x31, 0x55, 0x24, 0x42, 0x15, 0x31, 0x55, 0x24,
	0x82, 0xf9, 0x49, 0x50, 0x65, 0x4c, 0x15, 0x49, 0x20, 0x7e, 0x65, 0xa2, 0x04, 0x33, 0x84, 0x5f,
	0x09, 0x96, 0xaf, 0x4c, 0x92, 0x60, 0x81, 0xf2, 0x93, 0x9c, 0x1c, 0xa3, 0xec, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x3d, 0x9d, 0x7c, 0xf5, 0x33, 0x53, 0x52, 0x73,
	0x33, 0x2b, 0xf4, 0x89, 0x76, 0x7e, 0x12, 0x1b, 0xd8, 0xe9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xae, 0x4d, 0xeb, 0xe4, 0xf2, 0x00, 0x00, 0x00,
}
