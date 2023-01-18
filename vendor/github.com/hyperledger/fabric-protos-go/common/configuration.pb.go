// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/configuration.proto

package common

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

// HashingAlgorithm is encoded into the configuration transaction as a
// configuration item of type Chain with a Key of "HashingAlgorithm" and a
// Value of HashingAlgorithm as marshaled protobuf bytes
type HashingAlgorithm struct {
	// SHA256 is currently the only supported and tested algorithm.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashingAlgorithm) Reset()         { *m = HashingAlgorithm{} }
func (m *HashingAlgorithm) String() string { return proto.CompactTextString(m) }
func (*HashingAlgorithm) ProtoMessage()    {}
func (*HashingAlgorithm) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{0}
}

func (m *HashingAlgorithm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashingAlgorithm.Unmarshal(m, b)
}
func (m *HashingAlgorithm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashingAlgorithm.Marshal(b, m, deterministic)
}
func (m *HashingAlgorithm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashingAlgorithm.Merge(m, src)
}
func (m *HashingAlgorithm) XXX_Size() int {
	return xxx_messageInfo_HashingAlgorithm.Size(m)
}
func (m *HashingAlgorithm) XXX_DiscardUnknown() {
	xxx_messageInfo_HashingAlgorithm.DiscardUnknown(m)
}

var xxx_messageInfo_HashingAlgorithm proto.InternalMessageInfo

func (m *HashingAlgorithm) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// BlockDataHashingStructure is encoded into the configuration transaction as a configuration item of
// type Chain with a Key of "BlockDataHashingStructure" and a Value of HashingAlgorithm as marshaled protobuf bytes
type BlockDataHashingStructure struct {
	// width specifies the width of the Merkle tree to use when computing the BlockDataHash
	// in order to replicate flat hashing, set this width to MAX_UINT32
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockDataHashingStructure) Reset()         { *m = BlockDataHashingStructure{} }
func (m *BlockDataHashingStructure) String() string { return proto.CompactTextString(m) }
func (*BlockDataHashingStructure) ProtoMessage()    {}
func (*BlockDataHashingStructure) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{1}
}

func (m *BlockDataHashingStructure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockDataHashingStructure.Unmarshal(m, b)
}
func (m *BlockDataHashingStructure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockDataHashingStructure.Marshal(b, m, deterministic)
}
func (m *BlockDataHashingStructure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDataHashingStructure.Merge(m, src)
}
func (m *BlockDataHashingStructure) XXX_Size() int {
	return xxx_messageInfo_BlockDataHashingStructure.Size(m)
}
func (m *BlockDataHashingStructure) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDataHashingStructure.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDataHashingStructure proto.InternalMessageInfo

func (m *BlockDataHashingStructure) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

// OrdererAddresses is encoded into the configuration transaction as a configuration item of type Chain
// with a Key of "OrdererAddresses" and a Value of OrdererAddresses as marshaled protobuf bytes
type OrdererAddresses struct {
	Addresses            []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrdererAddresses) Reset()         { *m = OrdererAddresses{} }
func (m *OrdererAddresses) String() string { return proto.CompactTextString(m) }
func (*OrdererAddresses) ProtoMessage()    {}
func (*OrdererAddresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{2}
}

func (m *OrdererAddresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdererAddresses.Unmarshal(m, b)
}
func (m *OrdererAddresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdererAddresses.Marshal(b, m, deterministic)
}
func (m *OrdererAddresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdererAddresses.Merge(m, src)
}
func (m *OrdererAddresses) XXX_Size() int {
	return xxx_messageInfo_OrdererAddresses.Size(m)
}
func (m *OrdererAddresses) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdererAddresses.DiscardUnknown(m)
}

var xxx_messageInfo_OrdererAddresses proto.InternalMessageInfo

func (m *OrdererAddresses) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	MspId                string   `protobuf:"bytes,4,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	Identity             []byte   `protobuf:"bytes,5,opt,name=identity,proto3" json:"identity,omitempty"`
	ClientTlsCert        []byte   `protobuf:"bytes,6,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert        []byte   `protobuf:"bytes,7,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consenter) Reset()         { *m = Consenter{} }
func (m *Consenter) String() string { return proto.CompactTextString(m) }
func (*Consenter) ProtoMessage()    {}
func (*Consenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{3}
}

func (m *Consenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consenter.Unmarshal(m, b)
}
func (m *Consenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consenter.Marshal(b, m, deterministic)
}
func (m *Consenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consenter.Merge(m, src)
}
func (m *Consenter) XXX_Size() int {
	return xxx_messageInfo_Consenter.Size(m)
}
func (m *Consenter) XXX_DiscardUnknown() {
	xxx_messageInfo_Consenter.DiscardUnknown(m)
}

var xxx_messageInfo_Consenter proto.InternalMessageInfo

func (m *Consenter) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Consenter) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Consenter) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Consenter) GetMspId() string {
	if m != nil {
		return m.MspId
	}
	return ""
}

func (m *Consenter) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Consenter) GetClientTlsCert() []byte {
	if m != nil {
		return m.ClientTlsCert
	}
	return nil
}

func (m *Consenter) GetServerTlsCert() []byte {
	if m != nil {
		return m.ServerTlsCert
	}
	return nil
}

// Orderers is encoded into the configuration transaction as a configuration item of type Chain
// with a Key of "Orderers" and a Value of Orderers as marshaled protobuf bytes
type Orderers struct {
	ConsenterMapping     []*Consenter `protobuf:"bytes,1,rep,name=consenter_mapping,json=consenterMapping,proto3" json:"consenter_mapping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Orderers) Reset()         { *m = Orderers{} }
func (m *Orderers) String() string { return proto.CompactTextString(m) }
func (*Orderers) ProtoMessage()    {}
func (*Orderers) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{4}
}

func (m *Orderers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Orderers.Unmarshal(m, b)
}
func (m *Orderers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Orderers.Marshal(b, m, deterministic)
}
func (m *Orderers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orderers.Merge(m, src)
}
func (m *Orderers) XXX_Size() int {
	return xxx_messageInfo_Orderers.Size(m)
}
func (m *Orderers) XXX_DiscardUnknown() {
	xxx_messageInfo_Orderers.DiscardUnknown(m)
}

var xxx_messageInfo_Orderers proto.InternalMessageInfo

func (m *Orderers) GetConsenterMapping() []*Consenter {
	if m != nil {
		return m.ConsenterMapping
	}
	return nil
}

// Consortium represents the consortium context in which the channel was created
type Consortium struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consortium) Reset()         { *m = Consortium{} }
func (m *Consortium) String() string { return proto.CompactTextString(m) }
func (*Consortium) ProtoMessage()    {}
func (*Consortium) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{5}
}

func (m *Consortium) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consortium.Unmarshal(m, b)
}
func (m *Consortium) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consortium.Marshal(b, m, deterministic)
}
func (m *Consortium) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consortium.Merge(m, src)
}
func (m *Consortium) XXX_Size() int {
	return xxx_messageInfo_Consortium.Size(m)
}
func (m *Consortium) XXX_DiscardUnknown() {
	xxx_messageInfo_Consortium.DiscardUnknown(m)
}

var xxx_messageInfo_Consortium proto.InternalMessageInfo

func (m *Consortium) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Capabilities message defines the capabilities a particular binary must implement
// for that binary to be able to safely participate in the channel.  The capabilities
// message is defined at the /Channel level, the /Channel/Application level, and the
// /Channel/Orderer level.
//
// The /Channel level capabilties define capabilities which both the orderer and peer
// binaries must satisfy.  These capabilties might be things like a new MSP type,
// or a new policy type.
//
// The /Channel/Orderer level capabilties define capabilities which must be supported
// by the orderer, but which have no bearing on the behavior of the peer.  For instance
// if the orderer changes the logic for how it constructs new channels, only all orderers
// must agree on the new logic.  The peers do not need to be aware of this change as
// they only interact with the channel after it has been constructed.
//
// Finally, the /Channel/Application level capabilities define capabilities which the peer
// binary must satisfy, but which have no bearing on the orderer.  For instance, if the
// peer adds a new UTXO transaction type, or changes the chaincode lifecycle requirements,
// all peers must agree on the new logic.  However, orderers never inspect transactions
// this deeply, and therefore have no need to be aware of the change.
//
// The capabilities strings defined in these messages typically correspond to release
// binary versions (e.g. "V1.1"), and are used primarilly as a mechanism for a fully
// upgraded network to switch from one set of logic to a new one.
//
// Although for V1.1, the orderers must be upgraded to V1.1 prior to the rest of the
// network, going forward, because of the split between the /Channel, /Channel/Orderer
// and /Channel/Application capabilities.  It should be possible for the orderer and
// application networks to upgrade themselves independently (with the exception of any
// new capabilities defined at the /Channel level).
type Capabilities struct {
	Capabilities         map[string]*Capability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Capabilities) Reset()         { *m = Capabilities{} }
func (m *Capabilities) String() string { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()    {}
func (*Capabilities) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{6}
}

func (m *Capabilities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capabilities.Unmarshal(m, b)
}
func (m *Capabilities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capabilities.Marshal(b, m, deterministic)
}
func (m *Capabilities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capabilities.Merge(m, src)
}
func (m *Capabilities) XXX_Size() int {
	return xxx_messageInfo_Capabilities.Size(m)
}
func (m *Capabilities) XXX_DiscardUnknown() {
	xxx_messageInfo_Capabilities.DiscardUnknown(m)
}

var xxx_messageInfo_Capabilities proto.InternalMessageInfo

func (m *Capabilities) GetCapabilities() map[string]*Capability {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

// Capability is an empty message for the time being.  It is defined as a protobuf
// message rather than a constant, so that we may extend capabilities with other fields
// if the need arises in the future.  For the time being, a capability being in the
// capabilities map requires that that capability be supported.
type Capability struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Capability) Reset()         { *m = Capability{} }
func (m *Capability) String() string { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()    {}
func (*Capability) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba1ec2883858369, []int{7}
}

func (m *Capability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capability.Unmarshal(m, b)
}
func (m *Capability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capability.Marshal(b, m, deterministic)
}
func (m *Capability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capability.Merge(m, src)
}
func (m *Capability) XXX_Size() int {
	return xxx_messageInfo_Capability.Size(m)
}
func (m *Capability) XXX_DiscardUnknown() {
	xxx_messageInfo_Capability.DiscardUnknown(m)
}

var xxx_messageInfo_Capability proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HashingAlgorithm)(nil), "common.HashingAlgorithm")
	proto.RegisterType((*BlockDataHashingStructure)(nil), "common.BlockDataHashingStructure")
	proto.RegisterType((*OrdererAddresses)(nil), "common.OrdererAddresses")
	proto.RegisterType((*Consenter)(nil), "common.Consenter")
	proto.RegisterType((*Orderers)(nil), "common.Orderers")
	proto.RegisterType((*Consortium)(nil), "common.Consortium")
	proto.RegisterType((*Capabilities)(nil), "common.Capabilities")
	proto.RegisterMapType((map[string]*Capability)(nil), "common.Capabilities.CapabilitiesEntry")
	proto.RegisterType((*Capability)(nil), "common.Capability")
}

func init() { proto.RegisterFile("common/configuration.proto", fileDescriptor_cba1ec2883858369) }

var fileDescriptor_cba1ec2883858369 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0xa4, 0x09, 0xcd, 0x34, 0x85, 0x64, 0x05, 0x92, 0x89, 0x38, 0x44, 0x16, 0x8a,
	0x72, 0xa9, 0x03, 0xe5, 0x82, 0x38, 0x20, 0xb5, 0x01, 0x09, 0x2a, 0x21, 0x24, 0x17, 0x71, 0xe0,
	0x12, 0x6d, 0xec, 0xa9, 0xbd, 0xaa, 0xbd, 0x6b, 0xcd, 0x8e, 0x8b, 0xfc, 0x54, 0xbc, 0x05, 0xcf,
	0x85, 0xec, 0x75, 0x93, 0x54, 0xe5, 0x36, 0xff, 0x3f, 0xdf, 0xec, 0xcc, 0x68, 0x16, 0x66, 0xb1,
	0x29, 0x0a, 0xa3, 0x57, 0xb1, 0xd1, 0x37, 0x2a, 0xad, 0x48, 0xb2, 0x32, 0x3a, 0x2c, 0xc9, 0xb0,
	0x11, 0x43, 0x97, 0x0b, 0x16, 0x30, 0xf9, 0x22, 0x6d, 0xa6, 0x74, 0x7a, 0x91, 0xa7, 0x86, 0x14,
	0x67, 0x85, 0x10, 0x70, 0xa4, 0x65, 0x81, 0xbe, 0x37, 0xf7, 0x96, 0xa3, 0xa8, 0x8d, 0x83, 0xb7,
	0xf0, 0xf2, 0x32, 0x37, 0xf1, 0xed, 0x27, 0xc9, 0xb2, 0x2b, 0xb8, 0x66, 0xaa, 0x62, 0xae, 0x08,
	0xc5, 0x73, 0x18, 0xfc, 0x56, 0x09, 0x67, 0x6d, 0xc5, 0x69, 0xe4, 0x44, 0xf0, 0x06, 0x26, 0xdf,
	0x29, 0x41, 0x42, 0xba, 0x48, 0x12, 0x42, 0x6b, 0xd1, 0x8a, 0x57, 0x30, 0x92, 0xf7, 0xc2, 0xf7,
	0xe6, 0xfd, 0xe5, 0x28, 0xda, 0x1b, 0xc1, 0x5f, 0x0f, 0x46, 0x6b, 0xa3, 0x2d, 0x6a, 0x46, 0x12,
	0x4f, 0xa1, 0xa7, 0x92, 0xee, 0xc9, 0x9e, 0x4a, 0x9a, 0xb1, 0x32, 0x63, 0xd9, 0xef, 0xb9, 0xb1,
	0x9a, 0xb8, 0xf1, 0x4a, 0x43, 0xec, 0xf7, 0x5b, 0xaa, 0x8d, 0xc5, 0x0b, 0x18, 0x16, 0xb6, 0xdc,
	0xa8, 0xc4, 0x3f, 0x6a, 0xc9, 0x41, 0x61, 0xcb, 0xaf, 0x89, 0x98, 0xc1, 0xb1, 0x4a, 0x50, 0xb3,
	0xe2, 0xda, 0x1f, 0xcc, 0xbd, 0xe5, 0x38, 0xda, 0x69, 0xb1, 0x80, 0x67, 0x71, 0xae, 0x50, 0xf3,
	0x86, 0x73, 0xbb, 0x89, 0x91, 0xd8, 0x1f, 0xb6, 0xc8, 0xa9, 0xb3, 0x7f, 0xe4, 0x76, 0x8d, 0xc4,
	0x0d, 0x67, 0x91, 0xee, 0x90, 0xf6, 0xdc, 0x13, 0xc7, 0x39, 0xbb, 0xe3, 0x82, 0x2b, 0x38, 0xee,
	0x56, 0xb7, 0xe2, 0x23, 0x4c, 0xe3, 0xfb, 0x9d, 0x36, 0x85, 0x2c, 0x4b, 0xa5, 0xd3, 0x76, 0xf5,
	0x93, 0xf3, 0x69, 0xe8, 0xae, 0x10, 0xee, 0x96, 0x8e, 0x26, 0x3b, 0xf6, 0x9b, 0x43, 0x83, 0x39,
	0x40, 0x93, 0x36, 0xc4, 0xaa, 0xfa, 0xff, 0x6d, 0xfe, 0x78, 0x30, 0x5e, 0xcb, 0x52, 0x6e, 0x55,
	0xae, 0x58, 0xa1, 0x15, 0x57, 0x30, 0x8e, 0x0f, 0x74, 0xd7, 0x6d, 0xb1, 0xeb, 0x76, 0x90, 0x7b,
	0x20, 0x3e, 0x6b, 0xa6, 0x3a, 0x7a, 0x50, 0x3b, 0xbb, 0x86, 0xe9, 0x23, 0x44, 0x4c, 0xa0, 0x7f,
	0x8b, 0x75, 0x37, 0x44, 0x13, 0x8a, 0x25, 0x0c, 0xee, 0x64, 0x5e, 0x61, 0x7b, 0x9d, 0x93, 0x73,
	0xf1, 0xa8, 0x57, 0x1d, 0x39, 0xe0, 0x43, 0xef, 0xbd, 0x17, 0x8c, 0x01, 0xf6, 0x89, 0xcb, 0x9f,
	0xf0, 0xda, 0x50, 0x1a, 0x66, 0x75, 0x89, 0x94, 0x63, 0x92, 0x22, 0x85, 0x37, 0x72, 0x4b, 0x2a,
	0x76, 0x7f, 0xd5, 0x76, 0x6f, 0xfd, 0x0a, 0x53, 0xc5, 0x59, 0xb5, 0x6d, 0xe4, 0xea, 0x00, 0x5e,
	0x39, 0xf8, 0xcc, 0xc1, 0x67, 0xa9, 0x59, 0x39, 0x7e, 0x3b, 0x6c, 0x9d, 0x77, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x49, 0xe4, 0x92, 0xcc, 0x08, 0x03, 0x00, 0x00,
}
