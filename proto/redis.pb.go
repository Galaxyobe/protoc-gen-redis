// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: redis.proto

package redis

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_Enabled = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65000,
	Name:          "redis.enabled",
	Tag:           "varint,65000,opt,name=enabled",
	Filename:      "redis.proto",
}

var E_StorageType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65002,
	Name:          "redis.storage_type",
	Tag:           "bytes,65002,opt,name=storage_type",
	Filename:      "redis.proto",
}

var E_StorageCodec = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "redis.storage_codec",
	Tag:           "bytes,65003,opt,name=storage_codec",
	Filename:      "redis.proto",
}

var E_HashGetter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65004,
	Name:          "redis.hash_getter",
	Tag:           "varint,65004,opt,name=hash_getter",
	Filename:      "redis.proto",
}

var E_HashSetter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65005,
	Name:          "redis.hash_setter",
	Tag:           "varint,65005,opt,name=hash_setter",
	Filename:      "redis.proto",
}

var E_FieldPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65006,
	Name:          "redis.field_prefix",
	Tag:           "varint,65006,opt,name=field_prefix",
	Filename:      "redis.proto",
}

var E_FunctionType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65007,
	Name:          "redis.function_type",
	Tag:           "bytes,65007,opt,name=function_type",
	Filename:      "redis.proto",
}

var E_StorageFieldType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65100,
	Name:          "redis.storage_field_type",
	Tag:           "bytes,65100,opt,name=storage_field_type",
	Filename:      "redis.proto",
}

var E_StorageFieldCodec = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65101,
	Name:          "redis.storage_field_codec",
	Tag:           "bytes,65101,opt,name=storage_field_codec",
	Filename:      "redis.proto",
}

var E_HashFieldGetter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65102,
	Name:          "redis.hash_field_getter",
	Tag:           "varint,65102,opt,name=hash_field_getter",
	Filename:      "redis.proto",
}

var E_HashFieldSetter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65103,
	Name:          "redis.hash_field_setter",
	Tag:           "varint,65103,opt,name=hash_field_setter",
	Filename:      "redis.proto",
}

func init() {
	proto.RegisterExtension(E_Enabled)
	proto.RegisterExtension(E_StorageType)
	proto.RegisterExtension(E_StorageCodec)
	proto.RegisterExtension(E_HashGetter)
	proto.RegisterExtension(E_HashSetter)
	proto.RegisterExtension(E_FieldPrefix)
	proto.RegisterExtension(E_FunctionType)
	proto.RegisterExtension(E_StorageFieldType)
	proto.RegisterExtension(E_StorageFieldCodec)
	proto.RegisterExtension(E_HashFieldGetter)
	proto.RegisterExtension(E_HashFieldSetter)
}

func init() { proto.RegisterFile("redis.proto", fileDescriptor_d954120a2319ff8f) }

var fileDescriptor_d954120a2319ff8f = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd3, 0xcb, 0x4e, 0x32, 0x31,
	0x14, 0x07, 0xf0, 0x7c, 0xf9, 0x82, 0x97, 0x0e, 0x46, 0x81, 0x8d, 0x31, 0x31, 0x92, 0xb8, 0x71,
	0x35, 0x2c, 0xdc, 0x8d, 0x3b, 0x24, 0xb8, 0x30, 0x04, 0x83, 0xae, 0xdc, 0x4c, 0xe6, 0x72, 0x66,
	0x98, 0x64, 0x9c, 0x36, 0x6d, 0x49, 0xe4, 0x1d, 0xbd, 0xbc, 0x86, 0xf1, 0xfa, 0x00, 0xb2, 0x30,
	0x3d, 0x6d, 0x0d, 0xc8, 0xa2, 0xb3, 0xa4, 0x3d, 0xff, 0x1f, 0xf0, 0x3f, 0x29, 0xf1, 0x38, 0xa4,
	0x85, 0xf0, 0x19, 0xa7, 0x92, 0xb6, 0x1b, 0xf8, 0xe1, 0xa0, 0x9b, 0x53, 0x9a, 0x97, 0xd0, 0xc3,
	0xc3, 0x78, 0x96, 0xf5, 0x52, 0x10, 0x09, 0x2f, 0x98, 0xa4, 0x5c, 0x0f, 0x06, 0x67, 0x64, 0x13,
	0xaa, 0x28, 0x2e, 0x21, 0x6d, 0x1f, 0xf9, 0x7a, 0xda, 0xb7, 0xd3, 0xfe, 0x08, 0x84, 0x88, 0x72,
	0x18, 0x33, 0x59, 0xd0, 0x4a, 0xec, 0xbf, 0x7c, 0xff, 0xef, 0xfe, 0x3b, 0xd9, 0x9a, 0xd8, 0x44,
	0x30, 0x20, 0x4d, 0x21, 0x29, 0x8f, 0x72, 0x08, 0xe5, 0x9c, 0x81, 0x5b, 0x78, 0x45, 0x61, 0x7b,
	0xe2, 0x99, 0xd8, 0xcd, 0x9c, 0x41, 0x30, 0x24, 0x3b, 0x56, 0x49, 0x68, 0x0a, 0x89, 0x9b, 0x79,
	0x33, 0x8c, 0xfd, 0xf6, 0x73, 0x15, 0x0b, 0xfa, 0xc4, 0x9b, 0x46, 0x62, 0x1a, 0xe6, 0x20, 0x25,
	0x70, 0xb7, 0xf2, 0x6e, 0xfe, 0x0e, 0x51, 0xa9, 0x0b, 0x0c, 0xfd, 0x1a, 0xa2, 0xa6, 0xf1, 0xb1,
	0x6c, 0x5c, 0x6b, 0x63, 0x40, 0x9a, 0x59, 0x01, 0x65, 0x1a, 0x32, 0x0e, 0x59, 0x71, 0xef, 0x46,
	0x3e, 0x0d, 0xe2, 0x61, 0xec, 0x0a, 0x53, 0xaa, 0x95, 0x6c, 0x56, 0x25, 0x6a, 0xa2, 0x66, 0xb9,
	0x5f, 0xb6, 0x15, 0x9b, 0xc3, 0x76, 0x47, 0xa4, 0x6d, 0xdb, 0xd5, 0xbf, 0x0a, 0xb1, 0xc3, 0x35,
	0x6c, 0xa8, 0x2e, 0x2d, 0xf5, 0xb0, 0xd0, 0xd4, 0x9e, 0x89, 0xe2, 0x25, 0x72, 0x63, 0xd2, 0x59,
	0xe5, 0xf4, 0xca, 0x1c, 0xde, 0xa3, 0xf1, 0x5a, 0xcb, 0x9e, 0xde, 0xda, 0x25, 0x69, 0x61, 0xe3,
	0x5a, 0x33, 0xbb, 0x73, 0x70, 0x4f, 0x0b, 0x5d, 0xd8, 0xae, 0x4a, 0xe2, 0x8d, 0x59, 0xdf, 0x2a,
	0x26, 0x6a, 0x61, 0xcf, 0x6b, 0x98, 0xde, 0x63, 0xff, 0x98, 0x74, 0x12, 0x7a, 0xf7, 0x37, 0xdc,
	0x6f, 0x4c, 0xd4, 0xd3, 0xba, 0xd5, 0x2f, 0x2c, 0xde, 0xc0, 0xf3, 0xd3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x60, 0xcf, 0x4f, 0x37, 0x7e, 0x03, 0x00, 0x00,
}
