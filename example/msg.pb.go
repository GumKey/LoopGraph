// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package example

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

type Msg struct {
	MsgType              int32    `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	MsgInfo              string   `protobuf:"bytes,2,opt,name=MsgInfo,proto3" json:"MsgInfo,omitempty"`
	MsgFrom              string   `protobuf:"bytes,3,opt,name=MsgFrom,proto3" json:"MsgFrom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *Msg) GetMsgInfo() string {
	if m != nil {
		return m.MsgInfo
	}
	return ""
}

func (m *Msg) GetMsgFrom() string {
	if m != nil {
		return m.MsgFrom
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "example.Msg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x55, 0x0a,
	0xe6, 0x62, 0xf6, 0x2d, 0x4e, 0x17, 0x92, 0xe0, 0x62, 0xcf, 0x2d, 0x4e, 0x0f, 0xa9, 0x2c, 0x48,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71, 0x41, 0x32, 0xbe, 0xc5, 0xe9, 0x9e, 0x79,
	0x69, 0xf9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0x2e, 0x54, 0xc6, 0xad, 0x28, 0x3f,
	0x57, 0x82, 0x19, 0x2e, 0x03, 0xe2, 0x26, 0xb1, 0x81, 0x2d, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x6e, 0x2c, 0xa0, 0x71, 0x00, 0x00, 0x00,
}
