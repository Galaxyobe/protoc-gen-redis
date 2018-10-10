// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hash_type.proto

package test

import (
	fmt "fmt"
	_ "github.com/galaxyobe/protoc-gen-redis/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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

type HashStorageType_Enum int32

const (
	HashStorageType_E1 HashStorageType_Enum = 0
	HashStorageType_E2 HashStorageType_Enum = 1
)

var HashStorageType_Enum_name = map[int32]string{
	0: "E1",
	1: "E2",
}

var HashStorageType_Enum_value = map[string]int32{
	"E1": 0,
	"E2": 1,
}

func (x HashStorageType_Enum) String() string {
	return proto.EnumName(HashStorageType_Enum_name, int32(x))
}

func (HashStorageType_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f52cd8e24a8ed3f5, []int{0, 0}
}

type HashStorageType struct {
	SomeString           string               `protobuf:"bytes,1,opt,name=some_string,json=someString,proto3" json:"some_string,omitempty"`
	SomeBool             bool                 `protobuf:"varint,2,opt,name=some_bool,json=someBool,proto3" json:"some_bool,omitempty"`
	SomeInt32            int32                `protobuf:"varint,3,opt,name=some_int32,json=someInt32,proto3" json:"some_int32,omitempty"`
	SomeUint32           uint32               `protobuf:"varint,4,opt,name=some_uint32,json=someUint32,proto3" json:"some_uint32,omitempty"`
	SomeInt64            int64                `protobuf:"varint,5,opt,name=some_int64,json=someInt64,proto3" json:"some_int64,omitempty"`
	SomeUint64           uint64               `protobuf:"varint,6,opt,name=some_uint64,json=someUint64,proto3" json:"some_uint64,omitempty"`
	SomeFloat            float32              `protobuf:"fixed32,7,opt,name=some_float,json=someFloat,proto3" json:"some_float,omitempty"`
	SomeEnum             HashStorageType_Enum `protobuf:"varint,8,opt,name=some_enum,json=someEnum,proto3,enum=test.HashStorageType_Enum" json:"some_enum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HashStorageType) Reset()         { *m = HashStorageType{} }
func (m *HashStorageType) String() string { return proto.CompactTextString(m) }
func (*HashStorageType) ProtoMessage()    {}
func (*HashStorageType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52cd8e24a8ed3f5, []int{0}
}

func (m *HashStorageType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashStorageType.Unmarshal(m, b)
}
func (m *HashStorageType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashStorageType.Marshal(b, m, deterministic)
}
func (m *HashStorageType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashStorageType.Merge(m, src)
}
func (m *HashStorageType) XXX_Size() int {
	return xxx_messageInfo_HashStorageType.Size(m)
}
func (m *HashStorageType) XXX_DiscardUnknown() {
	xxx_messageInfo_HashStorageType.DiscardUnknown(m)
}

var xxx_messageInfo_HashStorageType proto.InternalMessageInfo

func (m *HashStorageType) GetSomeString() string {
	if m != nil {
		return m.SomeString
	}
	return ""
}

func (m *HashStorageType) GetSomeBool() bool {
	if m != nil {
		return m.SomeBool
	}
	return false
}

func (m *HashStorageType) GetSomeInt32() int32 {
	if m != nil {
		return m.SomeInt32
	}
	return 0
}

func (m *HashStorageType) GetSomeUint32() uint32 {
	if m != nil {
		return m.SomeUint32
	}
	return 0
}

func (m *HashStorageType) GetSomeInt64() int64 {
	if m != nil {
		return m.SomeInt64
	}
	return 0
}

func (m *HashStorageType) GetSomeUint64() uint64 {
	if m != nil {
		return m.SomeUint64
	}
	return 0
}

func (m *HashStorageType) GetSomeFloat() float32 {
	if m != nil {
		return m.SomeFloat
	}
	return 0
}

func (m *HashStorageType) GetSomeEnum() HashStorageType_Enum {
	if m != nil {
		return m.SomeEnum
	}
	return HashStorageType_E1
}

type HashStorageType2 struct {
	SomeString           string               `protobuf:"bytes,1,opt,name=some_string,json=someString,proto3" json:"some_string,omitempty"`
	SomeBool             bool                 `protobuf:"varint,2,opt,name=some_bool,json=someBool,proto3" json:"some_bool,omitempty"`
	SomeInt32            int32                `protobuf:"varint,3,opt,name=some_int32,json=someInt32,proto3" json:"some_int32,omitempty"`
	SomeUint32           uint32               `protobuf:"varint,4,opt,name=some_uint32,json=someUint32,proto3" json:"some_uint32,omitempty"`
	SomeInt64            int64                `protobuf:"varint,5,opt,name=some_int64,json=someInt64,proto3" json:"some_int64,omitempty"`
	SomeUint64           uint64               `protobuf:"varint,6,opt,name=some_uint64,json=someUint64,proto3" json:"some_uint64,omitempty"`
	SomeFloat            float32              `protobuf:"fixed32,7,opt,name=some_float,json=someFloat,proto3" json:"some_float,omitempty"`
	SomeMessage          *HashStorageType     `protobuf:"bytes,8,opt,name=some_message,json=someMessage,proto3" json:"some_message,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HashStorageType2) Reset()         { *m = HashStorageType2{} }
func (m *HashStorageType2) String() string { return proto.CompactTextString(m) }
func (*HashStorageType2) ProtoMessage()    {}
func (*HashStorageType2) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52cd8e24a8ed3f5, []int{1}
}

func (m *HashStorageType2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashStorageType2.Unmarshal(m, b)
}
func (m *HashStorageType2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashStorageType2.Marshal(b, m, deterministic)
}
func (m *HashStorageType2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashStorageType2.Merge(m, src)
}
func (m *HashStorageType2) XXX_Size() int {
	return xxx_messageInfo_HashStorageType2.Size(m)
}
func (m *HashStorageType2) XXX_DiscardUnknown() {
	xxx_messageInfo_HashStorageType2.DiscardUnknown(m)
}

var xxx_messageInfo_HashStorageType2 proto.InternalMessageInfo

func (m *HashStorageType2) GetSomeString() string {
	if m != nil {
		return m.SomeString
	}
	return ""
}

func (m *HashStorageType2) GetSomeBool() bool {
	if m != nil {
		return m.SomeBool
	}
	return false
}

func (m *HashStorageType2) GetSomeInt32() int32 {
	if m != nil {
		return m.SomeInt32
	}
	return 0
}

func (m *HashStorageType2) GetSomeUint32() uint32 {
	if m != nil {
		return m.SomeUint32
	}
	return 0
}

func (m *HashStorageType2) GetSomeInt64() int64 {
	if m != nil {
		return m.SomeInt64
	}
	return 0
}

func (m *HashStorageType2) GetSomeUint64() uint64 {
	if m != nil {
		return m.SomeUint64
	}
	return 0
}

func (m *HashStorageType2) GetSomeFloat() float32 {
	if m != nil {
		return m.SomeFloat
	}
	return 0
}

func (m *HashStorageType2) GetSomeMessage() *HashStorageType {
	if m != nil {
		return m.SomeMessage
	}
	return nil
}

func (m *HashStorageType2) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterEnum("test.HashStorageType_Enum", HashStorageType_Enum_name, HashStorageType_Enum_value)
	proto.RegisterType((*HashStorageType)(nil), "test.HashStorageType")
	proto.RegisterType((*HashStorageType2)(nil), "test.HashStorageType2")
}

func init() { proto.RegisterFile("hash_type.proto", fileDescriptor_f52cd8e24a8ed3f5) }

var fileDescriptor_f52cd8e24a8ed3f5 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x92, 0x4f, 0x6e, 0xd4, 0x30,
	0x14, 0xc6, 0x71, 0x9a, 0x0e, 0x13, 0xb7, 0xd0, 0x51, 0x24, 0x90, 0x15, 0x84, 0x62, 0xcd, 0xca,
	0x9b, 0xf1, 0x88, 0x34, 0x4a, 0x2b, 0x96, 0x48, 0x45, 0xb0, 0x60, 0x93, 0x96, 0xf5, 0xc8, 0x29,
	0xae, 0x13, 0x29, 0x89, 0xa3, 0xd8, 0x91, 0x98, 0xdb, 0x71, 0x01, 0x36, 0x1c, 0x60, 0xee, 0xc0,
	0x0d, 0x90, 0xed, 0x99, 0xcc, 0x1f, 0x71, 0x82, 0xae, 0xfc, 0xf2, 0xf9, 0xf7, 0xbd, 0x24, 0xdf,
	0x7b, 0xf0, 0xaa, 0x64, 0xaa, 0x5c, 0xe9, 0x75, 0xc7, 0x69, 0xd7, 0x4b, 0x2d, 0x43, 0x5f, 0x73,
	0xa5, 0xa3, 0x1b, 0x51, 0xe9, 0x72, 0x28, 0xe8, 0xa3, 0x6c, 0x96, 0x82, 0xd5, 0xec, 0xe7, 0x5a,
	0x16, 0x7c, 0x69, 0x89, 0xc7, 0x85, 0xe0, 0xed, 0xa2, 0xe7, 0x3f, 0x2a, 0xe5, 0x84, 0xa5, 0xad,
	0x9d, 0x3d, 0x5a, 0x1c, 0x1a, 0xa5, 0x90, 0x0e, 0x29, 0x86, 0x27, 0xfb, 0xe4, 0x78, 0x53, 0x6d,
	0xf1, 0x58, 0x48, 0x29, 0x6a, 0xbe, 0xa7, 0x74, 0xd5, 0x70, 0xa5, 0x59, 0xd3, 0x39, 0x60, 0xfe,
	0xdb, 0x83, 0x57, 0x5f, 0x98, 0x2a, 0xef, 0xb5, 0xec, 0x99, 0xe0, 0x0f, 0xeb, 0x8e, 0x87, 0x31,
	0xbc, 0x50, 0xb2, 0xe1, 0x2b, 0xa5, 0xfb, 0xaa, 0x15, 0x08, 0x60, 0x40, 0x82, 0x1c, 0x1a, 0xe9,
	0xde, 0x2a, 0xe1, 0x3b, 0x18, 0x58, 0xa0, 0x90, 0xb2, 0x46, 0x1e, 0x06, 0x64, 0x9a, 0x4f, 0x8d,
	0xf0, 0x49, 0xca, 0x3a, 0x7c, 0x0f, 0x2d, 0xba, 0xaa, 0x5a, 0x7d, 0x9d, 0xa0, 0x33, 0x0c, 0xc8,
	0x79, 0x6e, 0xf1, 0xaf, 0x46, 0x18, 0x9b, 0x0f, 0xee, 0xde, 0xc7, 0x80, 0xbc, 0x72, 0xcd, 0xbf,
	0x5b, 0xe5, 0xd0, 0x9f, 0xa5, 0xe8, 0x1c, 0x03, 0x72, 0x36, 0xfa, 0xb3, 0xf4, 0xc8, 0x9f, 0xa5,
	0x68, 0x82, 0x01, 0xf1, 0xf7, 0xfe, 0x2c, 0x1d, 0xfd, 0x4f, 0xb5, 0x64, 0x1a, 0xbd, 0xc4, 0x80,
	0x78, 0xce, 0xff, 0xd9, 0x08, 0xe1, 0xcd, 0xf6, 0xdb, 0x79, 0x3b, 0x34, 0x68, 0x8a, 0x01, 0x79,
	0x9d, 0x44, 0xd4, 0xcc, 0x84, 0x9e, 0xc4, 0x40, 0xef, 0xda, 0xa1, 0x71, 0xff, 0x65, 0xaa, 0xf9,
	0x5b, 0xe8, 0x9b, 0x33, 0x9c, 0x40, 0xef, 0xee, 0xc3, 0xec, 0x85, 0x3d, 0x93, 0x19, 0xf8, 0x78,
	0xf9, 0x6b, 0x13, 0x83, 0x3f, 0x9b, 0xd8, 0x37, 0x93, 0x9e, 0xff, 0xf5, 0xe0, 0xec, 0xa4, 0x51,
	0xf2, 0xbc, 0x03, 0xbd, 0x85, 0x97, 0xf6, 0xba, 0xe1, 0x4a, 0x31, 0xc1, 0x6d, 0xa6, 0x17, 0xc9,
	0x9b, 0xff, 0x66, 0x9a, 0xdb, 0x57, 0x7d, 0x73, 0x64, 0x78, 0x0b, 0x83, 0x71, 0x1d, 0x51, 0x60,
	0x6d, 0x11, 0x75, 0x0b, 0x4b, 0x77, 0x0b, 0x4b, 0x1f, 0x76, 0x44, 0xbe, 0x87, 0x8f, 0x33, 0x2f,
	0x26, 0x16, 0xbe, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x80, 0xef, 0x77, 0xb7, 0x6c, 0x03, 0x00,
	0x00,
}
