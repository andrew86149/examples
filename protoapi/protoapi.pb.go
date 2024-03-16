// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoapi.proto

package protoapi

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

// For random number
type RandomParams struct {
	Seed                 int64    `protobuf:"varint,1,opt,name=Seed,proto3" json:"Seed,omitempty"`
	Place                int64    `protobuf:"varint,2,opt,name=Place,proto3" json:"Place,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomParams) Reset()         { *m = RandomParams{} }
func (m *RandomParams) String() string { return proto.CompactTextString(m) }
func (*RandomParams) ProtoMessage()    {}
func (*RandomParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a688fa86c7e81ef2, []int{0}
}

func (m *RandomParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomParams.Unmarshal(m, b)
}
func (m *RandomParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomParams.Marshal(b, m, deterministic)
}
func (m *RandomParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomParams.Merge(m, src)
}
func (m *RandomParams) XXX_Size() int {
	return xxx_messageInfo_RandomParams.Size(m)
}
func (m *RandomParams) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomParams.DiscardUnknown(m)
}

var xxx_messageInfo_RandomParams proto.InternalMessageInfo

func (m *RandomParams) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *RandomParams) GetPlace() int64 {
	if m != nil {
		return m.Place
	}
	return 0
}

type RandomInt struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomInt) Reset()         { *m = RandomInt{} }
func (m *RandomInt) String() string { return proto.CompactTextString(m) }
func (*RandomInt) ProtoMessage()    {}
func (*RandomInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_a688fa86c7e81ef2, []int{1}
}

func (m *RandomInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomInt.Unmarshal(m, b)
}
func (m *RandomInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomInt.Marshal(b, m, deterministic)
}
func (m *RandomInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomInt.Merge(m, src)
}
func (m *RandomInt) XXX_Size() int {
	return xxx_messageInfo_RandomInt.Size(m)
}
func (m *RandomInt) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomInt.DiscardUnknown(m)
}

var xxx_messageInfo_RandomInt proto.InternalMessageInfo

func (m *RandomInt) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// For date time
type DateTime struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateTime) Reset()         { *m = DateTime{} }
func (m *DateTime) String() string { return proto.CompactTextString(m) }
func (*DateTime) ProtoMessage()    {}
func (*DateTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_a688fa86c7e81ef2, []int{2}
}

func (m *DateTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateTime.Unmarshal(m, b)
}
func (m *DateTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateTime.Marshal(b, m, deterministic)
}
func (m *DateTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateTime.Merge(m, src)
}
func (m *DateTime) XXX_Size() int {
	return xxx_messageInfo_DateTime.Size(m)
}
func (m *DateTime) XXX_DiscardUnknown() {
	xxx_messageInfo_DateTime.DiscardUnknown(m)
}

var xxx_messageInfo_DateTime proto.InternalMessageInfo

func (m *DateTime) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestDateTime struct {
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestDateTime) Reset()         { *m = RequestDateTime{} }
func (m *RequestDateTime) String() string { return proto.CompactTextString(m) }
func (*RequestDateTime) ProtoMessage()    {}
func (*RequestDateTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_a688fa86c7e81ef2, []int{3}
}

func (m *RequestDateTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestDateTime.Unmarshal(m, b)
}
func (m *RequestDateTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestDateTime.Marshal(b, m, deterministic)
}
func (m *RequestDateTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestDateTime.Merge(m, src)
}
func (m *RequestDateTime) XXX_Size() int {
	return xxx_messageInfo_RequestDateTime.Size(m)
}
func (m *RequestDateTime) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestDateTime.DiscardUnknown(m)
}

var xxx_messageInfo_RequestDateTime proto.InternalMessageInfo

func (m *RequestDateTime) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// For random password
type RequestPass struct {
	Seed                 int64    `protobuf:"varint,1,opt,name=Seed,proto3" json:"Seed,omitempty"`
	Length               int64    `protobuf:"varint,8,opt,name=Length,proto3" json:"Length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPass) Reset()         { *m = RequestPass{} }
func (m *RequestPass) String() string { return proto.CompactTextString(m) }
func (*RequestPass) ProtoMessage()    {}
func (*RequestPass) Descriptor() ([]byte, []int) {
	return fileDescriptor_a688fa86c7e81ef2, []int{4}
}

func (m *RequestPass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPass.Unmarshal(m, b)
}
func (m *RequestPass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPass.Marshal(b, m, deterministic)
}
func (m *RequestPass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPass.Merge(m, src)
}
func (m *RequestPass) XXX_Size() int {
	return xxx_messageInfo_RequestPass.Size(m)
}
func (m *RequestPass) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPass.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPass proto.InternalMessageInfo

func (m *RequestPass) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *RequestPass) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type RandomPass struct {
	Password             string   `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomPass) Reset()         { *m = RandomPass{} }
func (m *RandomPass) String() string { return proto.CompactTextString(m) }
func (*RandomPass) ProtoMessage()    {}
func (*RandomPass) Descriptor() ([]byte, []int) {
	return fileDescriptor_a688fa86c7e81ef2, []int{5}
}

func (m *RandomPass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomPass.Unmarshal(m, b)
}
func (m *RandomPass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomPass.Marshal(b, m, deterministic)
}
func (m *RandomPass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomPass.Merge(m, src)
}
func (m *RandomPass) XXX_Size() int {
	return xxx_messageInfo_RandomPass.Size(m)
}
func (m *RandomPass) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomPass.DiscardUnknown(m)
}

var xxx_messageInfo_RandomPass proto.InternalMessageInfo

func (m *RandomPass) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*RandomParams)(nil), "RandomParams")
	proto.RegisterType((*RandomInt)(nil), "RandomInt")
	proto.RegisterType((*DateTime)(nil), "DateTime")
	proto.RegisterType((*RequestDateTime)(nil), "RequestDateTime")
	proto.RegisterType((*RequestPass)(nil), "RequestPass")
	proto.RegisterType((*RandomPass)(nil), "RandomPass")
}

func init() {
	proto.RegisterFile("protoapi.proto", fileDescriptor_a688fa86c7e81ef2)
}

var fileDescriptor_a688fa86c7e81ef2 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2c, 0xc8, 0xd4, 0x03, 0x33, 0x94, 0x2c, 0xb8, 0x78, 0x82, 0x12, 0xf3, 0x52, 0xf2,
	0x73, 0x03, 0x12, 0x8b, 0x12, 0x73, 0x8b, 0x85, 0x84, 0xb8, 0x58, 0x82, 0x53, 0x53, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0x80, 0x9c, 0xc4, 0xe4,
	0x54, 0x09, 0x26, 0xb0, 0x20, 0x84, 0xa3, 0xa4, 0xc8, 0xc5, 0x09, 0xd1, 0xe9, 0x99, 0x57, 0x02,
	0x52, 0x12, 0x96, 0x98, 0x53, 0x9a, 0x0a, 0xd5, 0x07, 0xe1, 0x28, 0x29, 0x70, 0x71, 0xb8, 0x24,
	0x96, 0xa4, 0x86, 0x64, 0xe6, 0xa6, 0xa2, 0xaa, 0xe0, 0x84, 0xa9, 0x50, 0xe7, 0xe2, 0x0f, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xc1, 0x54, 0xc8, 0x84, 0xac, 0xd0, 0x92, 0x8b, 0x1b, 0xaa, 0x30,
	0x20, 0xb1, 0x18, 0xbb, 0x33, 0xc5, 0xb8, 0xd8, 0x7c, 0x52, 0xf3, 0xd2, 0x4b, 0x32, 0x24, 0x38,
	0xc0, 0xa2, 0x50, 0x9e, 0x92, 0x06, 0x17, 0x17, 0xcc, 0x8b, 0xc5, 0xc5, 0x42, 0x52, 0x5c, 0x1c,
	0x20, 0xba, 0x3c, 0xbf, 0x28, 0x05, 0xea, 0x14, 0x38, 0xdf, 0xa8, 0x85, 0x91, 0x8b, 0x0d, 0xa2,
	0x54, 0x48, 0x8d, 0x8b, 0xdd, 0x3d, 0x15, 0xec, 0x28, 0x21, 0x01, 0x3d, 0x34, 0x27, 0x4a, 0x71,
	0xea, 0xc1, 0x5d, 0xab, 0xc6, 0xc5, 0xe9, 0x9e, 0x5a, 0x02, 0xd5, 0xc4, 0xab, 0x87, 0x1c, 0x96,
	0x52, 0x5c, 0x7a, 0x88, 0x00, 0xd2, 0xe2, 0xe2, 0x85, 0xab, 0x03, 0xbb, 0x83, 0x47, 0x0f, 0xc9,
	0x3f, 0x52, 0xdc, 0x7a, 0x08, 0x29, 0x27, 0xde, 0x28, 0x6e, 0x3d, 0x7d, 0x6b, 0x58, 0x44, 0x25,
	0xb1, 0x81, 0x59, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xfc, 0xb8, 0xf6, 0xbb, 0x01,
	0x00, 0x00,
}
