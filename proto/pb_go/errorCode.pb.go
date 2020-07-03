// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errorCode.proto

package pb

import (
	fmt "fmt"
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

type ErrorCode int32

const (
	ErrorCode_OK                          ErrorCode = 0
	ErrorCode_HelloError                  ErrorCode = 1
	ErrorCode_LoginAccountOrPasswordError ErrorCode = 2
	ErrorCode_RegisterAccountExit         ErrorCode = 3
)

var ErrorCode_name = map[int32]string{
	0: "OK",
	1: "HelloError",
	2: "LoginAccountOrPasswordError",
	3: "RegisterAccountExit",
}

var ErrorCode_value = map[string]int32{
	"OK":                          0,
	"HelloError":                  1,
	"LoginAccountOrPasswordError": 2,
	"RegisterAccountExit":         3,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a55d56b9bad55ae, []int{0}
}

func init() {
	proto.RegisterEnum("pb.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("errorCode.proto", fileDescriptor_7a55d56b9bad55ae) }

var fileDescriptor_7a55d56b9bad55ae = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2d, 0x2a, 0xca,
	0x2f, 0x72, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0xd2,
	0x8a, 0xe5, 0xe2, 0x74, 0x85, 0x09, 0x0b, 0xb1, 0x71, 0x31, 0xf9, 0x7b, 0x0b, 0x30, 0x08, 0xf1,
	0x71, 0x71, 0x79, 0xa4, 0xe6, 0xe4, 0xe4, 0x83, 0x65, 0x04, 0x18, 0x85, 0xe4, 0xb9, 0xa4, 0x7d,
	0xf2, 0xd3, 0x33, 0xf3, 0x1c, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0xfc, 0x8b, 0x02, 0x12, 0x8b,
	0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x20, 0x0a, 0x98, 0x84, 0xc4, 0xb9, 0x84, 0x83, 0x52, 0xd3, 0x33,
	0x8b, 0x4b, 0x52, 0x8b, 0xa0, 0x6a, 0x5c, 0x2b, 0x32, 0x4b, 0x04, 0x98, 0x9d, 0x04, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92,
	0xd8, 0xc0, 0x76, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x25, 0x1c, 0x7a, 0x56, 0x8e, 0x00,
	0x00, 0x00,
}