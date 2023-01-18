// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode.proto

package peer

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	common "github.com/hyperledger/fabric-protos-go/common"
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

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}

var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}

func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{2, 0}
}

//ChaincodeID contains the path as specified by the deploy transaction
//that created it as well as the hashCode that is generated by the
//system for the path. From the user level (ie, CLI, REST API and so on)
//deploy transaction is expected to provide the path and other requests
//are expected to provide the hashCode. The other value will be ignored.
//Internally, the structure could contain both values. For instance, the
//hashCode will be set when first generated using the path
type ChaincodeID struct {
	//deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	//all other requests will use the name (really a hashcode) generated by
	//the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//user friendly version name for the chaincode
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeID) Reset()         { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()    {}
func (*ChaincodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{0}
}

func (m *ChaincodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeID.Unmarshal(m, b)
}
func (m *ChaincodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeID.Marshal(b, m, deterministic)
}
func (m *ChaincodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeID.Merge(m, src)
}
func (m *ChaincodeID) XXX_Size() int {
	return xxx_messageInfo_ChaincodeID.Size(m)
}
func (m *ChaincodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeID.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeID proto.InternalMessageInfo

func (m *ChaincodeID) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args        [][]byte          `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	Decorations map[string][]byte `protobuf:"bytes,2,rep,name=decorations,proto3" json:"decorations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// is_init is used for the application to signal that an invocation is to be routed
	// to the legacy 'Init' function for compatibility with chaincodes which handled
	// Init in the old way.  New applications should manage their initialized state
	// themselves.
	IsInit               bool     `protobuf:"varint,3,opt,name=is_init,json=isInit,proto3" json:"is_init,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeInput) Reset()         { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()    {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{1}
}

func (m *ChaincodeInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInput.Unmarshal(m, b)
}
func (m *ChaincodeInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInput.Marshal(b, m, deterministic)
}
func (m *ChaincodeInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInput.Merge(m, src)
}
func (m *ChaincodeInput) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInput.Size(m)
}
func (m *ChaincodeInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInput.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInput proto.InternalMessageInfo

func (m *ChaincodeInput) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ChaincodeInput) GetDecorations() map[string][]byte {
	if m != nil {
		return m.Decorations
	}
	return nil
}

func (m *ChaincodeInput) GetIsInit() bool {
	if m != nil {
		return m.IsInit
	}
	return false
}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type                 ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId          *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	Input                *ChaincodeInput    `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Timeout              int32              `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChaincodeSpec) Reset()         { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()    {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{2}
}

func (m *ChaincodeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeSpec.Unmarshal(m, b)
}
func (m *ChaincodeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeSpec.Merge(m, src)
}
func (m *ChaincodeSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeSpec.Size(m)
}
func (m *ChaincodeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeSpec proto.InternalMessageInfo

func (m *ChaincodeSpec) GetType() ChaincodeSpec_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeSpec_UNDEFINED
}

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeSpec) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec        *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
	CodePackage          []byte         `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChaincodeDeploymentSpec) Reset()         { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()    {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{3}
}

func (m *ChaincodeDeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Unmarshal(m, b)
}
func (m *ChaincodeDeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeDeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeDeploymentSpec.Merge(m, src)
}
func (m *ChaincodeDeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Size(m)
}
func (m *ChaincodeDeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeDeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeDeploymentSpec proto.InternalMessageInfo

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetCodePackage() []byte {
	if m != nil {
		return m.CodePackage
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec        *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChaincodeInvocationSpec) Reset()         { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()    {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{4}
}

func (m *ChaincodeInvocationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInvocationSpec.Unmarshal(m, b)
}
func (m *ChaincodeInvocationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInvocationSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeInvocationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInvocationSpec.Merge(m, src)
}
func (m *ChaincodeInvocationSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInvocationSpec.Size(m)
}
func (m *ChaincodeInvocationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInvocationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInvocationSpec proto.InternalMessageInfo

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// LifecycleEvent is used as the payload of the chaincode event emitted by LSCC
type LifecycleEvent struct {
	ChaincodeName        string   `protobuf:"bytes,1,opt,name=chaincode_name,json=chaincodeName,proto3" json:"chaincode_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleEvent) Reset()         { *m = LifecycleEvent{} }
func (m *LifecycleEvent) String() string { return proto.CompactTextString(m) }
func (*LifecycleEvent) ProtoMessage()    {}
func (*LifecycleEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{5}
}

func (m *LifecycleEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleEvent.Unmarshal(m, b)
}
func (m *LifecycleEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleEvent.Marshal(b, m, deterministic)
}
func (m *LifecycleEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleEvent.Merge(m, src)
}
func (m *LifecycleEvent) XXX_Size() int {
	return xxx_messageInfo_LifecycleEvent.Size(m)
}
func (m *LifecycleEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleEvent proto.InternalMessageInfo

func (m *LifecycleEvent) GetChaincodeName() string {
	if m != nil {
		return m.ChaincodeName
	}
	return ""
}

// CDSData is data stored in the LSCC on instantiation of a CC
// for CDSPackage.  This needs to be serialized for ChaincodeData
// hence the protobuf format
type CDSData struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Metadatahash         []byte   `protobuf:"bytes,2,opt,name=metadatahash,proto3" json:"metadatahash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDSData) Reset()         { *m = CDSData{} }
func (m *CDSData) String() string { return proto.CompactTextString(m) }
func (*CDSData) ProtoMessage()    {}
func (*CDSData) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{6}
}

func (m *CDSData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDSData.Unmarshal(m, b)
}
func (m *CDSData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDSData.Marshal(b, m, deterministic)
}
func (m *CDSData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDSData.Merge(m, src)
}
func (m *CDSData) XXX_Size() int {
	return xxx_messageInfo_CDSData.Size(m)
}
func (m *CDSData) XXX_DiscardUnknown() {
	xxx_messageInfo_CDSData.DiscardUnknown(m)
}

var xxx_messageInfo_CDSData proto.InternalMessageInfo

func (m *CDSData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *CDSData) GetMetadatahash() []byte {
	if m != nil {
		return m.Metadatahash
	}
	return nil
}

// ChaincodeData defines the datastructure for chaincodes to be serialized by proto
// Type provides an additional check by directing to use a specific package after instantiation
// Data is Type specific (see CDSPackage and SignedCDSPackage)
type ChaincodeData struct {
	// Name of the chaincode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the chaincode
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Escc for the chaincode instance
	Escc string `protobuf:"bytes,3,opt,name=escc,proto3" json:"escc,omitempty"`
	// Vscc for the chaincode instance
	Vscc string `protobuf:"bytes,4,opt,name=vscc,proto3" json:"vscc,omitempty"`
	// Policy endorsement policy for the chaincode instance
	Policy *common.SignaturePolicyEnvelope `protobuf:"bytes,5,opt,name=policy,proto3" json:"policy,omitempty"`
	// Data data specific to the package
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// Id of the chaincode that's the unique fingerprint for the CC This is not
	// currently used anywhere but serves as a good eyecatcher
	Id []byte `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	// InstantiationPolicy for the chaincode
	InstantiationPolicy  *common.SignaturePolicyEnvelope `protobuf:"bytes,8,opt,name=instantiation_policy,json=instantiationPolicy,proto3" json:"instantiation_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ChaincodeData) Reset()         { *m = ChaincodeData{} }
func (m *ChaincodeData) String() string { return proto.CompactTextString(m) }
func (*ChaincodeData) ProtoMessage()    {}
func (*ChaincodeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{7}
}

func (m *ChaincodeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeData.Unmarshal(m, b)
}
func (m *ChaincodeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeData.Marshal(b, m, deterministic)
}
func (m *ChaincodeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeData.Merge(m, src)
}
func (m *ChaincodeData) XXX_Size() int {
	return xxx_messageInfo_ChaincodeData.Size(m)
}
func (m *ChaincodeData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeData.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeData proto.InternalMessageInfo

func (m *ChaincodeData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChaincodeData) GetEscc() string {
	if m != nil {
		return m.Escc
	}
	return ""
}

func (m *ChaincodeData) GetVscc() string {
	if m != nil {
		return m.Vscc
	}
	return ""
}

func (m *ChaincodeData) GetPolicy() *common.SignaturePolicyEnvelope {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *ChaincodeData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChaincodeData) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ChaincodeData) GetInstantiationPolicy() *common.SignaturePolicyEnvelope {
	if m != nil {
		return m.InstantiationPolicy
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterMapType((map[string][]byte)(nil), "protos.ChaincodeInput.DecorationsEntry")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterType((*LifecycleEvent)(nil), "protos.LifecycleEvent")
	proto.RegisterType((*CDSData)(nil), "protos.CDSData")
	proto.RegisterType((*ChaincodeData)(nil), "protos.ChaincodeData")
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor_202814c635ff5fee) }

var fileDescriptor_202814c635ff5fee = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6e, 0xeb, 0x36,
	0x10, 0xad, 0x64, 0xf9, 0x11, 0xca, 0x31, 0x54, 0x26, 0x69, 0x84, 0x6c, 0xea, 0x0a, 0x28, 0xea,
	0x45, 0x22, 0x03, 0x2e, 0xd0, 0x14, 0x45, 0x11, 0xc0, 0x8d, 0xdc, 0xc0, 0x41, 0xe0, 0x04, 0x4c,
	0xdb, 0x45, 0x37, 0x06, 0x43, 0x8d, 0x65, 0x22, 0x36, 0x25, 0x48, 0xb4, 0x50, 0xfd, 0x48, 0xd7,
	0xfd, 0x9c, 0xfe, 0xd5, 0xbd, 0x20, 0xe9, 0xe7, 0x4d, 0x16, 0x17, 0xb8, 0x2b, 0x8d, 0x0e, 0xcf,
	0x3c, 0xce, 0x90, 0x33, 0xe8, 0x34, 0x03, 0xc8, 0xfb, 0x6c, 0x4e, 0xb9, 0x60, 0x69, 0x0c, 0x61,
	0x96, 0xa7, 0x32, 0xc5, 0x0d, 0xfd, 0x29, 0x2e, 0xce, 0x58, 0xba, 0x5c, 0xa6, 0xa2, 0x9f, 0xa5,
	0x0b, 0xce, 0x38, 0x14, 0xe6, 0x38, 0x78, 0x44, 0xee, 0xed, 0xc6, 0x63, 0x1c, 0x61, 0x8c, 0x9c,
	0x8c, 0xca, 0xb9, 0x6f, 0x75, 0xad, 0xde, 0x11, 0xd1, 0xb6, 0xc2, 0x04, 0x5d, 0x82, 0x6f, 0x1b,
	0x4c, 0xd9, 0xd8, 0x47, 0xcd, 0x12, 0xf2, 0x82, 0xa7, 0xc2, 0xaf, 0x69, 0x78, 0xf3, 0x1b, 0xfc,
	0x6f, 0xa1, 0xce, 0x2e, 0xa2, 0xc8, 0x56, 0x52, 0x05, 0xa0, 0x79, 0x52, 0xf8, 0x56, 0xb7, 0xd6,
	0x6b, 0x13, 0x6d, 0xe3, 0x31, 0x72, 0x63, 0x60, 0x69, 0x4e, 0x25, 0x4f, 0x45, 0xe1, 0xdb, 0xdd,
	0x5a, 0xcf, 0x1d, 0xfc, 0x60, 0x8a, 0x2a, 0xc2, 0xc3, 0x00, 0x61, 0xb4, 0x63, 0x8e, 0x84, 0xcc,
	0x2b, 0xb2, 0xef, 0x8b, 0xcf, 0x51, 0x93, 0x17, 0x53, 0x2e, 0xb8, 0xd4, 0xb5, 0xb4, 0x48, 0x83,
	0x17, 0x63, 0xc1, 0xe5, 0xc5, 0x0d, 0xf2, 0x3e, 0xf5, 0xc4, 0x1e, 0xaa, 0xbd, 0x42, 0xb5, 0xd6,
	0xa7, 0x4c, 0x7c, 0x8a, 0xea, 0x25, 0x5d, 0xac, 0x8c, 0xbe, 0x36, 0x31, 0x3f, 0xbf, 0xd8, 0x3f,
	0x5b, 0xc1, 0x07, 0x0b, 0x1d, 0x6f, 0x2b, 0x79, 0xce, 0x80, 0xe1, 0x10, 0x39, 0xb2, 0xca, 0x40,
	0xbb, 0x77, 0x06, 0x17, 0x6f, 0xca, 0x55, 0xa4, 0xf0, 0x8f, 0x2a, 0x03, 0xa2, 0x79, 0xf8, 0x27,
	0xd4, 0xde, 0xde, 0xc7, 0x94, 0xc7, 0x3a, 0x85, 0x3b, 0x38, 0x79, 0x2b, 0x33, 0x22, 0xee, 0x96,
	0x38, 0x8e, 0xf1, 0x25, 0xaa, 0x73, 0xa5, 0x5c, 0x0b, 0x72, 0x07, 0xdf, 0xbc, 0xdf, 0x17, 0x62,
	0x48, 0xea, 0x32, 0x24, 0x5f, 0x42, 0xba, 0x92, 0xbe, 0xd3, 0xb5, 0x7a, 0x75, 0xb2, 0xf9, 0x0d,
	0x6e, 0x90, 0xa3, 0xaa, 0xc1, 0xc7, 0xe8, 0xe8, 0xcf, 0x49, 0x34, 0xfa, 0x7d, 0x3c, 0x19, 0x45,
	0xde, 0x57, 0x18, 0xa1, 0xc6, 0xdd, 0xe3, 0xc3, 0x70, 0x72, 0xe7, 0x59, 0xb8, 0x85, 0x9c, 0xc9,
	0x63, 0x34, 0xf2, 0x6c, 0xdc, 0x44, 0xb5, 0xdb, 0x21, 0xf1, 0x6a, 0x0a, 0xba, 0x1f, 0xfe, 0x35,
	0xf4, 0x9c, 0xe0, 0x3f, 0x0b, 0x9d, 0x6f, 0x73, 0x46, 0x90, 0x2d, 0xd2, 0x6a, 0x09, 0x42, 0xea,
	0x5e, 0xfc, 0x8a, 0x3a, 0x3b, 0x6d, 0x45, 0x06, 0x4c, 0x77, 0xc5, 0x1d, 0x9c, 0xbd, 0xdb, 0x15,
	0x72, 0xcc, 0x0e, 0x3a, 0xf9, 0x1d, 0x6a, 0x6b, 0xc7, 0x8c, 0xb2, 0x57, 0x9a, 0x80, 0x16, 0xda,
	0x26, 0xae, 0xc2, 0x9e, 0x0c, 0x74, 0xef, 0xb4, 0x6c, 0xaf, 0x76, 0xef, 0xb4, 0x1c, 0xaf, 0x4e,
	0x3a, 0x30, 0x9b, 0x01, 0x93, 0xbc, 0x84, 0x69, 0x4c, 0x25, 0x90, 0x16, 0xfc, 0x03, 0x6c, 0x0a,
	0xa2, 0x0c, 0xb2, 0xbd, 0x0a, 0xc7, 0xa2, 0x4c, 0x99, 0xbe, 0xed, 0x2f, 0xaf, 0xd0, 0xa4, 0x27,
	0x5f, 0xf3, 0x78, 0x9a, 0x80, 0x00, 0xf3, 0x88, 0xa6, 0x74, 0x91, 0x04, 0xd7, 0xa8, 0xf3, 0xc0,
	0x67, 0xc0, 0x2a, 0xb6, 0x80, 0x51, 0x09, 0x42, 0xe2, 0xef, 0xf7, 0x13, 0xe9, 0x59, 0x31, 0xef,
	0x6b, 0x17, 0x71, 0x42, 0x97, 0x10, 0x0c, 0x51, 0xf3, 0x36, 0x7a, 0x8e, 0xa8, 0xa4, 0x6a, 0x24,
	0xe6, 0xb4, 0x30, 0x73, 0xd6, 0x26, 0xda, 0xc6, 0x01, 0x6a, 0x2f, 0x41, 0xd2, 0x98, 0x4a, 0xaa,
	0xcf, 0xcc, 0x7b, 0x3c, 0xc0, 0x82, 0x7f, 0xed, 0xbd, 0x27, 0xb9, 0x89, 0xb4, 0x97, 0xf1, 0xcd,
	0x74, 0xda, 0x07, 0xd3, 0xa9, 0xd8, 0x50, 0x30, 0xb6, 0x1e, 0x5a, 0x6d, 0x2b, 0xac, 0x54, 0x98,
	0x63, 0x30, 0x65, 0xe3, 0x6b, 0xd4, 0xd0, 0x8b, 0xa2, 0xf2, 0xeb, 0xba, 0x65, 0xdf, 0x86, 0x66,
	0x7d, 0x84, 0xcf, 0x3c, 0x11, 0x54, 0xae, 0x72, 0x78, 0xd2, 0xc7, 0x23, 0x51, 0xc2, 0x22, 0xcd,
	0x80, 0xac, 0xe9, 0x2a, 0x98, 0x2a, 0xd6, 0x6f, 0x18, 0x61, 0xca, 0xc6, 0x1d, 0x64, 0xf3, 0xd8,
	0x6f, 0x6a, 0xc4, 0xe6, 0x31, 0x26, 0xe8, 0x94, 0x8b, 0x42, 0x52, 0x21, 0xb9, 0xe9, 0xea, 0x3a,
	0x55, 0xeb, 0xf3, 0x52, 0x9d, 0x1c, 0x38, 0x9b, 0xc3, 0xdf, 0x08, 0x0a, 0xd2, 0x3c, 0x09, 0xe7,
	0x55, 0x06, 0xf9, 0x02, 0xe2, 0x04, 0xf2, 0x70, 0x46, 0x5f, 0x72, 0xce, 0x36, 0x77, 0xad, 0x96,
	0xe3, 0xdf, 0x97, 0x09, 0x97, 0xf3, 0xd5, 0x8b, 0xca, 0xd0, 0xdf, 0xa3, 0xf6, 0x0d, 0xf5, 0xca,
	0x50, 0xaf, 0x92, 0xb4, 0xaf, 0xd8, 0x2f, 0x66, 0x75, 0xfe, 0xf8, 0x31, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0xcb, 0x76, 0xde, 0x59, 0x05, 0x00, 0x00,
}
