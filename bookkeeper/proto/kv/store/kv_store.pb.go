// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv_store.proto

package store

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	rpc "github.com/phemmer/go-pulsar-state/bookkeeper/proto/kv/rpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ValueType int32

const (
	ValueType_BYTES  ValueType = 0
	ValueType_NUMBER ValueType = 1
)

var ValueType_name = map[int32]string{
	0: "BYTES",
	1: "NUMBER",
}

var ValueType_value = map[string]int32{
	"BYTES":  0,
	"NUMBER": 1,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{0}
}

// KeyRecord holds mvcc metadata for a given key
type KeyMeta struct {
	// create_revision is the revision of the last creation on the key
	CreateRevision int64 `protobuf:"varint,1,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// mod_revision is the revision of the last modification on the key
	ModRevision int64 `protobuf:"varint,2,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
	// version is the version of the most recent value
	Version int64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// value type
	ValueType            ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=bookkeeper.proto.kv.store.ValueType" json:"value_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *KeyMeta) Reset()         { *m = KeyMeta{} }
func (m *KeyMeta) String() string { return proto.CompactTextString(m) }
func (*KeyMeta) ProtoMessage()    {}
func (*KeyMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{0}
}

func (m *KeyMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyMeta.Unmarshal(m, b)
}
func (m *KeyMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyMeta.Marshal(b, m, deterministic)
}
func (m *KeyMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyMeta.Merge(m, src)
}
func (m *KeyMeta) XXX_Size() int {
	return xxx_messageInfo_KeyMeta.Size(m)
}
func (m *KeyMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyMeta.DiscardUnknown(m)
}

var xxx_messageInfo_KeyMeta proto.InternalMessageInfo

func (m *KeyMeta) GetCreateRevision() int64 {
	if m != nil {
		return m.CreateRevision
	}
	return 0
}

func (m *KeyMeta) GetModRevision() int64 {
	if m != nil {
		return m.ModRevision
	}
	return 0
}

func (m *KeyMeta) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *KeyMeta) GetValueType() ValueType {
	if m != nil {
		return m.ValueType
	}
	return ValueType_BYTES
}

type NopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NopRequest) Reset()         { *m = NopRequest{} }
func (m *NopRequest) String() string { return proto.CompactTextString(m) }
func (*NopRequest) ProtoMessage()    {}
func (*NopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{1}
}

func (m *NopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NopRequest.Unmarshal(m, b)
}
func (m *NopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NopRequest.Marshal(b, m, deterministic)
}
func (m *NopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NopRequest.Merge(m, src)
}
func (m *NopRequest) XXX_Size() int {
	return xxx_messageInfo_NopRequest.Size(m)
}
func (m *NopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NopRequest proto.InternalMessageInfo

type Command struct {
	// Types that are valid to be assigned to Req:
	//	*Command_NopReq
	//	*Command_PutReq
	//	*Command_DeleteReq
	//	*Command_TxnReq
	//	*Command_IncrReq
	Req                  isCommand_Req `protobuf_oneof:"req"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{2}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

type isCommand_Req interface {
	isCommand_Req()
}

type Command_NopReq struct {
	NopReq *NopRequest `protobuf:"bytes,1,opt,name=nop_req,json=nopReq,proto3,oneof"`
}

type Command_PutReq struct {
	PutReq *rpc.PutRequest `protobuf:"bytes,2,opt,name=put_req,json=putReq,proto3,oneof"`
}

type Command_DeleteReq struct {
	DeleteReq *rpc.DeleteRangeRequest `protobuf:"bytes,3,opt,name=delete_req,json=deleteReq,proto3,oneof"`
}

type Command_TxnReq struct {
	TxnReq *rpc.TxnRequest `protobuf:"bytes,4,opt,name=txn_req,json=txnReq,proto3,oneof"`
}

type Command_IncrReq struct {
	IncrReq *rpc.IncrementRequest `protobuf:"bytes,5,opt,name=incr_req,json=incrReq,proto3,oneof"`
}

func (*Command_NopReq) isCommand_Req() {}

func (*Command_PutReq) isCommand_Req() {}

func (*Command_DeleteReq) isCommand_Req() {}

func (*Command_TxnReq) isCommand_Req() {}

func (*Command_IncrReq) isCommand_Req() {}

func (m *Command) GetReq() isCommand_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *Command) GetNopReq() *NopRequest {
	if x, ok := m.GetReq().(*Command_NopReq); ok {
		return x.NopReq
	}
	return nil
}

func (m *Command) GetPutReq() *rpc.PutRequest {
	if x, ok := m.GetReq().(*Command_PutReq); ok {
		return x.PutReq
	}
	return nil
}

func (m *Command) GetDeleteReq() *rpc.DeleteRangeRequest {
	if x, ok := m.GetReq().(*Command_DeleteReq); ok {
		return x.DeleteReq
	}
	return nil
}

func (m *Command) GetTxnReq() *rpc.TxnRequest {
	if x, ok := m.GetReq().(*Command_TxnReq); ok {
		return x.TxnReq
	}
	return nil
}

func (m *Command) GetIncrReq() *rpc.IncrementRequest {
	if x, ok := m.GetReq().(*Command_IncrReq); ok {
		return x.IncrReq
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Command) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Command_NopReq)(nil),
		(*Command_PutReq)(nil),
		(*Command_DeleteReq)(nil),
		(*Command_TxnReq)(nil),
		(*Command_IncrReq)(nil),
	}
}

type CheckpointMetadata struct {
	Files                []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Txid                 []byte   `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"`
	CreatedAt            uint64   `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointMetadata) Reset()         { *m = CheckpointMetadata{} }
func (m *CheckpointMetadata) String() string { return proto.CompactTextString(m) }
func (*CheckpointMetadata) ProtoMessage()    {}
func (*CheckpointMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3673e574d3b8b3, []int{3}
}

func (m *CheckpointMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointMetadata.Unmarshal(m, b)
}
func (m *CheckpointMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointMetadata.Marshal(b, m, deterministic)
}
func (m *CheckpointMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointMetadata.Merge(m, src)
}
func (m *CheckpointMetadata) XXX_Size() int {
	return xxx_messageInfo_CheckpointMetadata.Size(m)
}
func (m *CheckpointMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointMetadata proto.InternalMessageInfo

func (m *CheckpointMetadata) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CheckpointMetadata) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

func (m *CheckpointMetadata) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("bookkeeper.proto.kv.store.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*KeyMeta)(nil), "bookkeeper.proto.kv.store.KeyMeta")
	proto.RegisterType((*NopRequest)(nil), "bookkeeper.proto.kv.store.NopRequest")
	proto.RegisterType((*Command)(nil), "bookkeeper.proto.kv.store.Command")
	proto.RegisterType((*CheckpointMetadata)(nil), "bookkeeper.proto.kv.store.CheckpointMetadata")
}

func init() { proto.RegisterFile("kv_store.proto", fileDescriptor_cd3673e574d3b8b3) }

var fileDescriptor_cd3673e574d3b8b3 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x6f, 0x8b, 0xd3, 0x40,
	0x10, 0xc6, 0x2f, 0xd7, 0x3f, 0x31, 0xd3, 0x52, 0x8f, 0xc5, 0x17, 0x55, 0x10, 0x6a, 0x55, 0xac,
	0x1e, 0x97, 0x40, 0x7d, 0x2b, 0xa2, 0xad, 0x27, 0x8a, 0xde, 0x71, 0xac, 0x55, 0x50, 0x90, 0xb0,
	0x4d, 0xc6, 0x36, 0xa4, 0xd9, 0xdd, 0x6e, 0x36, 0xa1, 0xfd, 0x5e, 0x7e, 0x28, 0x3f, 0x86, 0x74,
	0xd2, 0xbb, 0x16, 0xb1, 0xfa, 0x6e, 0x9f, 0xd9, 0x67, 0x7e, 0x3c, 0x33, 0xbb, 0xd0, 0x49, 0xcb,
	0x30, 0xb7, 0xca, 0xa0, 0xaf, 0x8d, 0xb2, 0x8a, 0xdd, 0x9d, 0x2a, 0x95, 0xa6, 0x88, 0x1a, 0x4d,
	0x55, 0xf1, 0xd3, 0xd2, 0x27, 0xc3, 0xbd, 0x76, 0x5a, 0x86, 0x46, 0x47, 0x55, 0xb9, 0xff, 0xd3,
	0x01, 0xf7, 0x03, 0xae, 0x2f, 0xd0, 0x0a, 0xf6, 0x04, 0x6e, 0x47, 0x06, 0x85, 0xc5, 0xd0, 0x60,
	0x99, 0xe4, 0x89, 0x92, 0x5d, 0xa7, 0xe7, 0x0c, 0x6a, 0xbc, 0x53, 0x95, 0xf9, 0xb6, 0xca, 0x1e,
	0x40, 0x3b, 0x53, 0xf1, 0xce, 0x75, 0x4c, 0xae, 0x56, 0xa6, 0xe2, 0x1b, 0x4b, 0x17, 0xdc, 0x12,
	0x0d, 0xdd, 0xd6, 0xe8, 0xf6, 0x5a, 0xb2, 0x31, 0x40, 0x29, 0x16, 0x05, 0x86, 0x76, 0xad, 0xb1,
	0x5b, 0xef, 0x39, 0x83, 0xce, 0xf0, 0x91, 0x7f, 0x30, 0xaf, 0xff, 0x65, 0x63, 0x9e, 0xac, 0x35,
	0x72, 0xaf, 0xbc, 0x3e, 0xf6, 0xdb, 0x00, 0x97, 0x4a, 0x73, 0x5c, 0x16, 0x98, 0xdb, 0xfe, 0xaf,
	0x63, 0x70, 0xc7, 0x2a, 0xcb, 0x84, 0x8c, 0xd9, 0x2b, 0x70, 0xa5, 0xd2, 0xa1, 0xc1, 0x25, 0x85,
	0x6f, 0x0d, 0x1f, 0xff, 0x83, 0xbd, 0x63, 0xbc, 0x3b, 0xe2, 0x4d, 0x49, 0x8a, 0xbd, 0x04, 0x57,
	0x17, 0x96, 0x08, 0xc7, 0x44, 0x78, 0xf8, 0x57, 0xc2, 0x66, 0x87, 0x57, 0x85, 0xdd, 0xeb, 0xd7,
	0xa4, 0xd8, 0x47, 0x80, 0x18, 0x17, 0x48, 0x6b, 0x5c, 0xd2, 0xf4, 0xad, 0xe1, 0xe9, 0x41, 0xc4,
	0x1b, 0xb2, 0x72, 0x21, 0x67, 0xb8, 0x43, 0x79, 0x15, 0x60, 0x9b, 0xc6, 0xae, 0x24, 0xa1, 0xea,
	0xff, 0x49, 0x33, 0x59, 0xc9, 0xbd, 0x34, 0x96, 0x14, 0x7b, 0x0b, 0xb7, 0x12, 0x19, 0x19, 0x02,
	0x34, 0x08, 0xf0, 0xf4, 0x20, 0xe0, 0xbd, 0x8c, 0x0c, 0x66, 0x28, 0xf7, 0x86, 0x72, 0x37, 0xcd,
	0x1c, 0x97, 0xa3, 0x06, 0xd4, 0x0c, 0x2e, 0xfb, 0xdf, 0x81, 0x8d, 0xe7, 0x18, 0xa5, 0x5a, 0x25,
	0xd2, 0x6e, 0x7e, 0x4d, 0x2c, 0xac, 0x60, 0x77, 0xa0, 0xf1, 0x23, 0x59, 0x60, 0xde, 0x75, 0x7a,
	0xb5, 0x81, 0xc7, 0x2b, 0xc1, 0x18, 0xd4, 0xed, 0x2a, 0x89, 0x69, 0x8b, 0x6d, 0x4e, 0x67, 0x76,
	0x1f, 0xa0, 0xfa, 0x4c, 0x71, 0x28, 0x2c, 0x2d, 0xa7, 0xce, 0xbd, 0x6d, 0xe5, 0xb5, 0x7d, 0xd6,
	0x07, 0xef, 0xe6, 0xbd, 0x99, 0x07, 0x8d, 0xd1, 0xd7, 0xc9, 0xf9, 0xa7, 0x93, 0x23, 0x06, 0xd0,
	0xbc, 0xfc, 0x7c, 0x31, 0x3a, 0xe7, 0x27, 0xce, 0x28, 0x83, 0x53, 0x65, 0x66, 0xbe, 0xd0, 0x22,
	0x9a, 0xe3, 0xfe, 0x2c, 0xb9, 0x35, 0x28, 0xb2, 0x3f, 0xde, 0xf8, 0xca, 0xf9, 0xf6, 0x62, 0x96,
	0xd8, 0x79, 0x31, 0xf5, 0x23, 0x95, 0x05, 0x7a, 0x8e, 0x59, 0x86, 0x26, 0x98, 0xa9, 0x33, 0x5d,
	0x2c, 0x72, 0x61, 0xce, 0x72, 0x2b, 0x2c, 0x06, 0x3b, 0x4c, 0x40, 0xfd, 0x41, 0x5a, 0x06, 0xd4,
	0x3f, 0x6d, 0x92, 0x7e, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x24, 0x5f, 0x26, 0x76, 0x63, 0x03,
	0x00, 0x00,
}
