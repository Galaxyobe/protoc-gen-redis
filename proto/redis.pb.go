// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: redis.proto

package redis

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_Enabled = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65000,
	Name:          "redis.enabled",
	Tag:           "varint,65000,opt,name=enabled",
	Filename:      "redis.proto",
}

var E_Ttl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "redis.ttl",
	Tag:           "varint,65001,opt,name=ttl",
	Filename:      "redis.proto",
}

var E_StorageType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65002,
	Name:          "redis.storage_type",
	Tag:           "bytes,65002,opt,name=storage_type,json=storageType",
	Filename:      "redis.proto",
}

var E_StorageCodec = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "redis.storage_codec",
	Tag:           "bytes,65003,opt,name=storage_codec,json=storageCodec",
	Filename:      "redis.proto",
}

var E_StorageFieldType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65100,
	Name:          "redis.storage_field_type",
	Tag:           "bytes,65100,opt,name=storage_field_type,json=storageFieldType",
	Filename:      "redis.proto",
}

var E_StorageFieldCodec = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65101,
	Name:          "redis.storage_field_codec",
	Tag:           "bytes,65101,opt,name=storage_field_codec,json=storageFieldCodec",
	Filename:      "redis.proto",
}

func init() {
	proto.RegisterExtension(E_Enabled)
	proto.RegisterExtension(E_Ttl)
	proto.RegisterExtension(E_StorageType)
	proto.RegisterExtension(E_StorageCodec)
	proto.RegisterExtension(E_StorageFieldType)
	proto.RegisterExtension(E_StorageFieldCodec)
}

func init() { proto.RegisterFile("redis.proto", fileDescriptor_redis_bdfe8c07f3a1d489) }

var fileDescriptor_redis_bdfe8c07f3a1d489 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4a, 0x4d, 0xc9,
	0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x14, 0xd2, 0xf3, 0xf3,
	0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x82, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99,
	0x05, 0x25, 0xf9, 0x45, 0x10, 0x85, 0x56, 0xd6, 0x5c, 0xec, 0xa9, 0x79, 0x89, 0x49, 0x39, 0xa9,
	0x29, 0x42, 0xf2, 0x7a, 0x10, 0xd5, 0x7a, 0x30, 0xd5, 0x7a, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9,
	0xa9, 0xfe, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x12, 0x2f, 0x7e, 0x33, 0x2b, 0x30, 0x6a, 0x70,
	0x04, 0xc1, 0x74, 0x58, 0x19, 0x73, 0x31, 0x97, 0x94, 0xe4, 0x10, 0xd6, 0xf8, 0x12, 0xaa, 0x11,
	0xa4, 0xda, 0xca, 0x85, 0x8b, 0xa7, 0xb8, 0x24, 0xbf, 0x28, 0x31, 0x3d, 0x35, 0xbe, 0xa4, 0xb2,
	0x20, 0x95, 0xb0, 0xee, 0x57, 0x60, 0xdd, 0x9c, 0x41, 0xdc, 0x50, 0x6d, 0x21, 0x95, 0x05, 0xa9,
	0x56, 0x6e, 0x5c, 0xbc, 0x30, 0x53, 0x92, 0xf3, 0x53, 0x52, 0x93, 0x09, 0x1b, 0xf3, 0x1a, 0x6a,
	0x0c, 0xcc, 0x76, 0x67, 0x90, 0x36, 0x2b, 0x5f, 0x2e, 0x21, 0x98, 0x39, 0x69, 0x99, 0xa9, 0x39,
	0x29, 0x10, 0x37, 0xc9, 0x62, 0x18, 0xe6, 0x06, 0x92, 0x84, 0x19, 0x75, 0xe6, 0x0f, 0xc4, 0x28,
	0x01, 0xa8, 0x56, 0xb0, 0x24, 0xd8, 0x59, 0xfe, 0x5c, 0xc2, 0xa8, 0xc6, 0x41, 0x1c, 0x47, 0xc0,
	0xbc, 0xb3, 0x50, 0xf3, 0x04, 0x91, 0xcd, 0x03, 0xbb, 0xcf, 0x49, 0x99, 0x4b, 0x38, 0x39, 0x3f,
	0x17, 0x5d, 0xbb, 0x13, 0x6b, 0x10, 0x28, 0x7e, 0xa3, 0x20, 0xd1, 0x9c, 0xc4, 0x06, 0x16, 0x37,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x72, 0x81, 0x52, 0xb2, 0x03, 0x02, 0x00, 0x00,
}
