// Code generated by protoc-gen-go. DO NOT EDIT.
// source: json_codec.proto

package test

import (
	fmt "fmt"
	_ "github.com/galaxyobe/protoc-gen-redis/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type StringJsonCodec struct {
	SomeString           string           `protobuf:"bytes,1,opt,name=some_string,json=someString,proto3" json:"some_string,omitempty"`
	SomeBool             bool             `protobuf:"varint,2,opt,name=some_bool,json=someBool,proto3" json:"some_bool,omitempty"`
	SomeInt32            int32            `protobuf:"varint,3,opt,name=some_int32,json=someInt32,proto3" json:"some_int32,omitempty"`
	SomeUint32           uint32           `protobuf:"varint,4,opt,name=some_uint32,json=someUint32,proto3" json:"some_uint32,omitempty"`
	SomeInt64            int64            `protobuf:"varint,5,opt,name=some_int64,json=someInt64,proto3" json:"some_int64,omitempty"`
	SomeUint64           uint64           `protobuf:"varint,6,opt,name=some_uint64,json=someUint64,proto3" json:"some_uint64,omitempty"`
	SomeFloat            float32          `protobuf:"fixed32,7,opt,name=some_float,json=someFloat,proto3" json:"some_float,omitempty"`
	StringJsonCodec      *StringJsonCodec `protobuf:"bytes,8,opt,name=StringJsonCodec,proto3" json:"StringJsonCodec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StringJsonCodec) Reset()         { *m = StringJsonCodec{} }
func (m *StringJsonCodec) String() string { return proto.CompactTextString(m) }
func (*StringJsonCodec) ProtoMessage()    {}
func (*StringJsonCodec) Descriptor() ([]byte, []int) {
	return fileDescriptor_34fdd4f84d729f10, []int{0}
}

func (m *StringJsonCodec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringJsonCodec.Unmarshal(m, b)
}
func (m *StringJsonCodec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringJsonCodec.Marshal(b, m, deterministic)
}
func (m *StringJsonCodec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringJsonCodec.Merge(m, src)
}
func (m *StringJsonCodec) XXX_Size() int {
	return xxx_messageInfo_StringJsonCodec.Size(m)
}
func (m *StringJsonCodec) XXX_DiscardUnknown() {
	xxx_messageInfo_StringJsonCodec.DiscardUnknown(m)
}

var xxx_messageInfo_StringJsonCodec proto.InternalMessageInfo

func (m *StringJsonCodec) GetSomeString() string {
	if m != nil {
		return m.SomeString
	}
	return ""
}

func (m *StringJsonCodec) GetSomeBool() bool {
	if m != nil {
		return m.SomeBool
	}
	return false
}

func (m *StringJsonCodec) GetSomeInt32() int32 {
	if m != nil {
		return m.SomeInt32
	}
	return 0
}

func (m *StringJsonCodec) GetSomeUint32() uint32 {
	if m != nil {
		return m.SomeUint32
	}
	return 0
}

func (m *StringJsonCodec) GetSomeInt64() int64 {
	if m != nil {
		return m.SomeInt64
	}
	return 0
}

func (m *StringJsonCodec) GetSomeUint64() uint64 {
	if m != nil {
		return m.SomeUint64
	}
	return 0
}

func (m *StringJsonCodec) GetSomeFloat() float32 {
	if m != nil {
		return m.SomeFloat
	}
	return 0
}

func (m *StringJsonCodec) GetStringJsonCodec() *StringJsonCodec {
	if m != nil {
		return m.StringJsonCodec
	}
	return nil
}

type HashJsonCodec struct {
	SomeString           string         `protobuf:"bytes,1,opt,name=some_string,json=someString,proto3" json:"some_string,omitempty"`
	SomeBool             bool           `protobuf:"varint,2,opt,name=some_bool,json=someBool,proto3" json:"some_bool,omitempty"`
	SomeInt32            int32          `protobuf:"varint,3,opt,name=some_int32,json=someInt32,proto3" json:"some_int32,omitempty"`
	SomeUint32           uint32         `protobuf:"varint,4,opt,name=some_uint32,json=someUint32,proto3" json:"some_uint32,omitempty"`
	SomeInt64            int64          `protobuf:"varint,5,opt,name=some_int64,json=someInt64,proto3" json:"some_int64,omitempty"`
	SomeUint64           uint64         `protobuf:"varint,6,opt,name=some_uint64,json=someUint64,proto3" json:"some_uint64,omitempty"`
	SomeFloat            float32        `protobuf:"fixed32,7,opt,name=some_float,json=someFloat,proto3" json:"some_float,omitempty"`
	HashJsonCodec        *HashJsonCodec `protobuf:"bytes,8,opt,name=HashJsonCodec,proto3" json:"HashJsonCodec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HashJsonCodec) Reset()         { *m = HashJsonCodec{} }
func (m *HashJsonCodec) String() string { return proto.CompactTextString(m) }
func (*HashJsonCodec) ProtoMessage()    {}
func (*HashJsonCodec) Descriptor() ([]byte, []int) {
	return fileDescriptor_34fdd4f84d729f10, []int{1}
}

func (m *HashJsonCodec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashJsonCodec.Unmarshal(m, b)
}
func (m *HashJsonCodec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashJsonCodec.Marshal(b, m, deterministic)
}
func (m *HashJsonCodec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashJsonCodec.Merge(m, src)
}
func (m *HashJsonCodec) XXX_Size() int {
	return xxx_messageInfo_HashJsonCodec.Size(m)
}
func (m *HashJsonCodec) XXX_DiscardUnknown() {
	xxx_messageInfo_HashJsonCodec.DiscardUnknown(m)
}

var xxx_messageInfo_HashJsonCodec proto.InternalMessageInfo

func (m *HashJsonCodec) GetSomeString() string {
	if m != nil {
		return m.SomeString
	}
	return ""
}

func (m *HashJsonCodec) GetSomeBool() bool {
	if m != nil {
		return m.SomeBool
	}
	return false
}

func (m *HashJsonCodec) GetSomeInt32() int32 {
	if m != nil {
		return m.SomeInt32
	}
	return 0
}

func (m *HashJsonCodec) GetSomeUint32() uint32 {
	if m != nil {
		return m.SomeUint32
	}
	return 0
}

func (m *HashJsonCodec) GetSomeInt64() int64 {
	if m != nil {
		return m.SomeInt64
	}
	return 0
}

func (m *HashJsonCodec) GetSomeUint64() uint64 {
	if m != nil {
		return m.SomeUint64
	}
	return 0
}

func (m *HashJsonCodec) GetSomeFloat() float32 {
	if m != nil {
		return m.SomeFloat
	}
	return 0
}

func (m *HashJsonCodec) GetHashJsonCodec() *HashJsonCodec {
	if m != nil {
		return m.HashJsonCodec
	}
	return nil
}

func init() {
	proto.RegisterType((*StringJsonCodec)(nil), "test.StringJsonCodec")
	proto.RegisterType((*HashJsonCodec)(nil), "test.HashJsonCodec")
}

func init() { proto.RegisterFile("json_codec.proto", fileDescriptor_34fdd4f84d729f10) }

var fileDescriptor_34fdd4f84d729f10 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x92, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xc9, 0xd6, 0xed, 0xdd, 0xb2, 0x77, 0x28, 0x55, 0x21, 0x4c, 0xa4, 0xc1, 0x53, 0x2e,
	0xeb, 0x60, 0x2b, 0x15, 0x77, 0x11, 0x14, 0x44, 0x3d, 0x56, 0x3c, 0x8f, 0xb6, 0xcb, 0xda, 0x4a,
	0xd7, 0x47, 0x96, 0x0c, 0xf4, 0xc3, 0x09, 0x7e, 0x06, 0x8f, 0x1e, 0xfa, 0x59, 0x24, 0x89, 0xd6,
	0xb5, 0x7e, 0x04, 0x6f, 0xe9, 0xff, 0xf9, 0xfd, 0x42, 0x9f, 0x7f, 0x8b, 0xf7, 0x1f, 0x05, 0x14,
	0x8b, 0x18, 0x96, 0x3c, 0x76, 0x9f, 0x36, 0x20, 0xc1, 0xb6, 0x24, 0x17, 0x72, 0x74, 0x96, 0x64,
	0x32, 0xdd, 0x46, 0x6e, 0x0c, 0xeb, 0x49, 0x12, 0xe6, 0xe1, 0xf3, 0x0b, 0x44, 0x7c, 0xa2, 0x89,
	0x78, 0x9c, 0xf0, 0x62, 0xbc, 0xe1, 0xcb, 0x4c, 0x98, 0x60, 0xa2, 0xcf, 0x46, 0x1f, 0x8d, 0x77,
	0x45, 0x48, 0xc0, 0x20, 0xd1, 0x76, 0xa5, 0x9f, 0x0c, 0xaf, 0x4e, 0x06, 0x3f, 0x7d, 0x6d, 0xe1,
	0xbd, 0x7b, 0xb9, 0xc9, 0x8a, 0xe4, 0x4e, 0x40, 0x71, 0xa5, 0xde, 0xc3, 0x76, 0xf0, 0x40, 0xc0,
	0x9a, 0x2f, 0x84, 0xce, 0x09, 0xa2, 0x88, 0xf5, 0x03, 0xac, 0x22, 0x43, 0xda, 0xc7, 0xb8, 0xaf,
	0x81, 0x08, 0x20, 0x27, 0x2d, 0x8a, 0x58, 0x2f, 0xe8, 0xa9, 0xe0, 0x12, 0x20, 0xb7, 0x4f, 0xb0,
	0x46, 0x17, 0x59, 0x21, 0x67, 0x53, 0xd2, 0xa6, 0x88, 0x75, 0x02, 0x8d, 0xdf, 0xaa, 0xa0, 0xba,
	0x7c, 0x6b, 0xe6, 0x16, 0x45, 0x6c, 0x68, 0x2e, 0x7f, 0xd0, 0xc9, 0xae, 0xef, 0x7b, 0xa4, 0x43,
	0x11, 0x6b, 0x57, 0xbe, 0xef, 0xd5, 0x7c, 0xdf, 0x23, 0x5d, 0x8a, 0x98, 0xf5, 0xe3, 0xfb, 0x5e,
	0xe5, 0xaf, 0x72, 0x08, 0x25, 0xf9, 0x47, 0x11, 0x6b, 0x19, 0xff, 0x5a, 0x05, 0xf6, 0xc5, 0xaf,
	0x7d, 0x49, 0x8f, 0x22, 0x36, 0x98, 0x1e, 0xb9, 0xaa, 0x78, 0xb7, 0x31, 0x0c, 0x9a, 0xf4, 0xfc,
	0xff, 0x5b, 0xe9, 0xa0, 0x8f, 0xd2, 0xb1, 0xd4, 0x97, 0x53, 0xfd, 0x0d, 0x6f, 0x42, 0x91, 0xfe,
	0x91, 0xf6, 0xce, 0x1b, 0xdb, 0x7e, 0x75, 0x77, 0x60, 0xba, 0xab, 0x8d, 0x82, 0x3a, 0x39, 0x3f,
	0x54, 0xbd, 0xbd, 0x97, 0x8e, 0x95, 0x86, 0x22, 0xfd, 0xee, 0x2f, 0xea, 0xea, 0xdf, 0x70, 0xf6,
	0x19, 0x00, 0x00, 0xff, 0xff, 0x25, 0xf4, 0xe1, 0x38, 0x08, 0x03, 0x00, 0x00,
}
