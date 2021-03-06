// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	msg.proto
	pref.proto

It has these top-level messages:
	InnerMessage
	OuterMessage
	Prefecture
	GetPrefecturesResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type InnerMessage struct {
	// some_integer can only be in range (1, 100).
	SomeInteger int32 `protobuf:"varint,1,opt,name=some_integer,json=someInteger" json:"some_integer,omitempty"`
	// some_float can only be in range (0;1).
	SomeFloat float64 `protobuf:"fixed64,2,opt,name=some_float,json=someFloat" json:"some_float,omitempty"`
}

func (m *InnerMessage) Reset()                    { *m = InnerMessage{} }
func (m *InnerMessage) String() string            { return proto1.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()               {}
func (*InnerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InnerMessage) GetSomeInteger() int32 {
	if m != nil {
		return m.SomeInteger
	}
	return 0
}

func (m *InnerMessage) GetSomeFloat() float64 {
	if m != nil {
		return m.SomeFloat
	}
	return 0
}

type OuterMessage struct {
	// important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
	ImportantString string `protobuf:"bytes,1,opt,name=important_string,json=importantString" json:"important_string,omitempty"`
	// proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
	Inner *InnerMessage `protobuf:"bytes,2,opt,name=inner" json:"inner,omitempty"`
}

func (m *OuterMessage) Reset()                    { *m = OuterMessage{} }
func (m *OuterMessage) String() string            { return proto1.CompactTextString(m) }
func (*OuterMessage) ProtoMessage()               {}
func (*OuterMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OuterMessage) GetImportantString() string {
	if m != nil {
		return m.ImportantString
	}
	return ""
}

func (m *OuterMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

func init() {
	proto1.RegisterType((*InnerMessage)(nil), "validator.examples.InnerMessage")
	proto1.RegisterType((*OuterMessage)(nil), "validator.examples.OuterMessage")
}

func init() { proto1.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f,
	0xd2, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0x96, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf,
	0xd7, 0x05, 0x6b, 0xd0, 0x85, 0xab, 0x2f, 0xd6, 0x47, 0x68, 0x05, 0x4b, 0x29, 0x15, 0x71, 0xf1,
	0x78, 0xe6, 0xe5, 0xa5, 0x16, 0xf9, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x69, 0x73, 0xf1,
	0x14, 0xe7, 0xe7, 0xa6, 0xc6, 0x67, 0xe6, 0x95, 0xa4, 0xa6, 0xa7, 0x16, 0x49, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x3a, 0x71, 0x3c, 0xba, 0x2f, 0xcf, 0x22, 0xc0, 0x20, 0x91, 0x12, 0xc4, 0x0d, 0x92,
	0xf5, 0x84, 0x48, 0x0a, 0x99, 0x72, 0x71, 0x81, 0x15, 0xa7, 0xe5, 0xe4, 0x27, 0x96, 0x48, 0x30,
	0x29, 0x30, 0x6a, 0x30, 0x3a, 0x89, 0x3d, 0xba, 0x2f, 0x2f, 0xe4, 0xc9, 0x00, 0x05, 0x81, 0x10,
	0xea, 0x83, 0x7d, 0x10, 0x27, 0x48, 0xa5, 0x1b, 0x48, 0xa1, 0x52, 0x2f, 0x23, 0x17, 0x8f, 0x7f,
	0x69, 0x09, 0xc2, 0x52, 0x5b, 0x2e, 0x81, 0xcc, 0xdc, 0x82, 0xfc, 0xa2, 0x92, 0xc4, 0xbc, 0x92,
	0xf8, 0xe2, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0xb0, 0xc5, 0x9c, 0x4e, 0x42, 0x8f, 0xee, 0xcb, 0xf3,
	0x71, 0xf1, 0xc4, 0x45, 0x27, 0xea, 0x56, 0xc5, 0x56, 0x1b, 0xe9, 0x98, 0xd6, 0xaa, 0x04, 0xf1,
	0xc3, 0xd5, 0x06, 0x83, 0x95, 0x0a, 0xd9, 0x71, 0xb1, 0x66, 0x82, 0xfc, 0x00, 0x76, 0x01, 0xb7,
	0x91, 0x82, 0x1e, 0x66, 0xf8, 0xe8, 0x21, 0x7b, 0xd2, 0x89, 0xed, 0xd1, 0x7d, 0x79, 0x26, 0x05,
	0xc6, 0x20, 0x88, 0xb6, 0x24, 0x36, 0x70, 0x50, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0x84, 0x23, 0x47, 0x63, 0x01, 0x00, 0x00,
}
