// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kservice.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

//string key
type MsgStringKey struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgStringKey) Reset()         { *m = MsgStringKey{} }
func (m *MsgStringKey) String() string { return proto.CompactTextString(m) }
func (*MsgStringKey) ProtoMessage()    {}
func (*MsgStringKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{0}
}

func (m *MsgStringKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgStringKey.Unmarshal(m, b)
}
func (m *MsgStringKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgStringKey.Marshal(b, m, deterministic)
}
func (m *MsgStringKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStringKey.Merge(m, src)
}
func (m *MsgStringKey) XXX_Size() int {
	return xxx_messageInfo_MsgStringKey.Size(m)
}
func (m *MsgStringKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStringKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStringKey proto.InternalMessageInfo

func (m *MsgStringKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

//int64 key
type MsgInt64Key struct {
	Key                  int64    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInt64Key) Reset()         { *m = MsgInt64Key{} }
func (m *MsgInt64Key) String() string { return proto.CompactTextString(m) }
func (*MsgInt64Key) ProtoMessage()    {}
func (*MsgInt64Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{1}
}

func (m *MsgInt64Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInt64Key.Unmarshal(m, b)
}
func (m *MsgInt64Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInt64Key.Marshal(b, m, deterministic)
}
func (m *MsgInt64Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInt64Key.Merge(m, src)
}
func (m *MsgInt64Key) XXX_Size() int {
	return xxx_messageInfo_MsgInt64Key.Size(m)
}
func (m *MsgInt64Key) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInt64Key.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInt64Key proto.InternalMessageInfo

func (m *MsgInt64Key) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

//value
type MsgValue struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgValue) Reset()         { *m = MsgValue{} }
func (m *MsgValue) String() string { return proto.CompactTextString(m) }
func (*MsgValue) ProtoMessage()    {}
func (*MsgValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{2}
}

func (m *MsgValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgValue.Unmarshal(m, b)
}
func (m *MsgValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgValue.Marshal(b, m, deterministic)
}
func (m *MsgValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgValue.Merge(m, src)
}
func (m *MsgValue) XXX_Size() int {
	return xxx_messageInfo_MsgValue.Size(m)
}
func (m *MsgValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgValue.DiscardUnknown(m)
}

var xxx_messageInfo_MsgValue proto.InternalMessageInfo

func (m *MsgValue) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *MsgValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

//string key - value
type MsgStringKeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgStringKeyValue) Reset()         { *m = MsgStringKeyValue{} }
func (m *MsgStringKeyValue) String() string { return proto.CompactTextString(m) }
func (*MsgStringKeyValue) ProtoMessage()    {}
func (*MsgStringKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{3}
}

func (m *MsgStringKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgStringKeyValue.Unmarshal(m, b)
}
func (m *MsgStringKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgStringKeyValue.Marshal(b, m, deterministic)
}
func (m *MsgStringKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStringKeyValue.Merge(m, src)
}
func (m *MsgStringKeyValue) XXX_Size() int {
	return xxx_messageInfo_MsgStringKeyValue.Size(m)
}
func (m *MsgStringKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStringKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStringKeyValue proto.InternalMessageInfo

func (m *MsgStringKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MsgStringKeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

//int64 key - value
type MsgInt64KeyValue struct {
	Key                  int64    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInt64KeyValue) Reset()         { *m = MsgInt64KeyValue{} }
func (m *MsgInt64KeyValue) String() string { return proto.CompactTextString(m) }
func (*MsgInt64KeyValue) ProtoMessage()    {}
func (*MsgInt64KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{4}
}

func (m *MsgInt64KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInt64KeyValue.Unmarshal(m, b)
}
func (m *MsgInt64KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInt64KeyValue.Marshal(b, m, deterministic)
}
func (m *MsgInt64KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInt64KeyValue.Merge(m, src)
}
func (m *MsgInt64KeyValue) XXX_Size() int {
	return xxx_messageInfo_MsgInt64KeyValue.Size(m)
}
func (m *MsgInt64KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInt64KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInt64KeyValue proto.InternalMessageInfo

func (m *MsgInt64KeyValue) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *MsgInt64KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

//string key - value - timeout
type MsgStringKeyValueTimeout struct {
	Timeout              int64    `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgStringKeyValueTimeout) Reset()         { *m = MsgStringKeyValueTimeout{} }
func (m *MsgStringKeyValueTimeout) String() string { return proto.CompactTextString(m) }
func (*MsgStringKeyValueTimeout) ProtoMessage()    {}
func (*MsgStringKeyValueTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{5}
}

func (m *MsgStringKeyValueTimeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgStringKeyValueTimeout.Unmarshal(m, b)
}
func (m *MsgStringKeyValueTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgStringKeyValueTimeout.Marshal(b, m, deterministic)
}
func (m *MsgStringKeyValueTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStringKeyValueTimeout.Merge(m, src)
}
func (m *MsgStringKeyValueTimeout) XXX_Size() int {
	return xxx_messageInfo_MsgStringKeyValueTimeout.Size(m)
}
func (m *MsgStringKeyValueTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStringKeyValueTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStringKeyValueTimeout proto.InternalMessageInfo

func (m *MsgStringKeyValueTimeout) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *MsgStringKeyValueTimeout) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MsgStringKeyValueTimeout) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

//int64 key - value - timeout
type MsgInt64KeyValueTimeout struct {
	Timeout              int64    `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Key                  int64    `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInt64KeyValueTimeout) Reset()         { *m = MsgInt64KeyValueTimeout{} }
func (m *MsgInt64KeyValueTimeout) String() string { return proto.CompactTextString(m) }
func (*MsgInt64KeyValueTimeout) ProtoMessage()    {}
func (*MsgInt64KeyValueTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{6}
}

func (m *MsgInt64KeyValueTimeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInt64KeyValueTimeout.Unmarshal(m, b)
}
func (m *MsgInt64KeyValueTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInt64KeyValueTimeout.Marshal(b, m, deterministic)
}
func (m *MsgInt64KeyValueTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInt64KeyValueTimeout.Merge(m, src)
}
func (m *MsgInt64KeyValueTimeout) XXX_Size() int {
	return xxx_messageInfo_MsgInt64KeyValueTimeout.Size(m)
}
func (m *MsgInt64KeyValueTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInt64KeyValueTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInt64KeyValueTimeout proto.InternalMessageInfo

func (m *MsgInt64KeyValueTimeout) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *MsgInt64KeyValueTimeout) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *MsgInt64KeyValueTimeout) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

//string key - timeout
type MsgStringKeyTimeout struct {
	Timeout              int64    `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgStringKeyTimeout) Reset()         { *m = MsgStringKeyTimeout{} }
func (m *MsgStringKeyTimeout) String() string { return proto.CompactTextString(m) }
func (*MsgStringKeyTimeout) ProtoMessage()    {}
func (*MsgStringKeyTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{7}
}

func (m *MsgStringKeyTimeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgStringKeyTimeout.Unmarshal(m, b)
}
func (m *MsgStringKeyTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgStringKeyTimeout.Marshal(b, m, deterministic)
}
func (m *MsgStringKeyTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStringKeyTimeout.Merge(m, src)
}
func (m *MsgStringKeyTimeout) XXX_Size() int {
	return xxx_messageInfo_MsgStringKeyTimeout.Size(m)
}
func (m *MsgStringKeyTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStringKeyTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStringKeyTimeout proto.InternalMessageInfo

func (m *MsgStringKeyTimeout) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *MsgStringKeyTimeout) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

//int64 key - timeout
type MsgInt64KeyTimeout struct {
	Timeout              int64    `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Key                  int64    `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInt64KeyTimeout) Reset()         { *m = MsgInt64KeyTimeout{} }
func (m *MsgInt64KeyTimeout) String() string { return proto.CompactTextString(m) }
func (*MsgInt64KeyTimeout) ProtoMessage()    {}
func (*MsgInt64KeyTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{8}
}

func (m *MsgInt64KeyTimeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInt64KeyTimeout.Unmarshal(m, b)
}
func (m *MsgInt64KeyTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInt64KeyTimeout.Marshal(b, m, deterministic)
}
func (m *MsgInt64KeyTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInt64KeyTimeout.Merge(m, src)
}
func (m *MsgInt64KeyTimeout) XXX_Size() int {
	return xxx_messageInfo_MsgInt64KeyTimeout.Size(m)
}
func (m *MsgInt64KeyTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInt64KeyTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInt64KeyTimeout proto.InternalMessageInfo

func (m *MsgInt64KeyTimeout) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *MsgInt64KeyTimeout) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

//长度
type MsgLen struct {
	HashLen              int32    `protobuf:"varint,1,opt,name=hashLen,proto3" json:"hashLen,omitempty"`
	ListLen              int32    `protobuf:"varint,2,opt,name=listLen,proto3" json:"listLen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgLen) Reset()         { *m = MsgLen{} }
func (m *MsgLen) String() string { return proto.CompactTextString(m) }
func (*MsgLen) ProtoMessage()    {}
func (*MsgLen) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee8e12a3d2812f4, []int{9}
}

func (m *MsgLen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgLen.Unmarshal(m, b)
}
func (m *MsgLen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgLen.Marshal(b, m, deterministic)
}
func (m *MsgLen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLen.Merge(m, src)
}
func (m *MsgLen) XXX_Size() int {
	return xxx_messageInfo_MsgLen.Size(m)
}
func (m *MsgLen) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLen.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLen proto.InternalMessageInfo

func (m *MsgLen) GetHashLen() int32 {
	if m != nil {
		return m.HashLen
	}
	return 0
}

func (m *MsgLen) GetListLen() int32 {
	if m != nil {
		return m.ListLen
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgStringKey)(nil), "pb.MsgStringKey")
	proto.RegisterType((*MsgInt64Key)(nil), "pb.MsgInt64Key")
	proto.RegisterType((*MsgValue)(nil), "pb.MsgValue")
	proto.RegisterType((*MsgStringKeyValue)(nil), "pb.MsgStringKeyValue")
	proto.RegisterType((*MsgInt64KeyValue)(nil), "pb.MsgInt64KeyValue")
	proto.RegisterType((*MsgStringKeyValueTimeout)(nil), "pb.MsgStringKeyValueTimeout")
	proto.RegisterType((*MsgInt64KeyValueTimeout)(nil), "pb.MsgInt64KeyValueTimeout")
	proto.RegisterType((*MsgStringKeyTimeout)(nil), "pb.MsgStringKeyTimeout")
	proto.RegisterType((*MsgInt64KeyTimeout)(nil), "pb.MsgInt64KeyTimeout")
	proto.RegisterType((*MsgLen)(nil), "pb.MsgLen")
}

func init() { proto.RegisterFile("kservice.proto", fileDescriptor_4ee8e12a3d2812f4) }

var fileDescriptor_4ee8e12a3d2812f4 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x52, 0x4d, 0x8b, 0xd3, 0x50,
	0x14, 0xa5, 0x2f, 0xb4, 0x8e, 0xd7, 0x12, 0xeb, 0xb3, 0x33, 0x13, 0x32, 0x82, 0x25, 0x28, 0xcc,
	0x42, 0x32, 0x32, 0x33, 0xcc, 0x80, 0xba, 0x50, 0xa4, 0x16, 0xb1, 0xd9, 0x24, 0x45, 0x17, 0xba,
	0x69, 0xe4, 0x36, 0x09, 0x49, 0x93, 0x90, 0xf7, 0x5a, 0xe8, 0x2f, 0x75, 0xe3, 0x8f, 0x91, 0x97,
	0x97, 0x40, 0x3e, 0x9a, 0x8a, 0xdd, 0xb8, 0x7b, 0xf7, 0xde, 0x73, 0xce, 0x3d, 0x39, 0xb9, 0xa0,
	0x86, 0x0c, 0xb3, 0x6d, 0xf0, 0x13, 0xcd, 0x34, 0x4b, 0x78, 0x42, 0x49, 0xea, 0xea, 0x17, 0x5e,
	0x92, 0x78, 0x11, 0x5e, 0xe5, 0x1d, 0x77, 0xb3, 0xba, 0xc2, 0x75, 0xca, 0x77, 0x12, 0x60, 0x4c,
	0x60, 0x68, 0x31, 0xcf, 0xe1, 0x59, 0x10, 0x7b, 0x5f, 0x70, 0x47, 0x47, 0xa0, 0x84, 0xb8, 0xd3,
	0x7a, 0x93, 0xde, 0xe5, 0x43, 0x5b, 0x3c, 0x8d, 0xe7, 0xf0, 0xc8, 0x62, 0xde, 0xe7, 0x98, 0xdf,
	0xdd, 0x36, 0x00, 0x8a, 0x04, 0xbc, 0x86, 0x13, 0x8b, 0x79, 0x5f, 0x97, 0xd1, 0x06, 0xa9, 0x0a,
	0x24, 0x09, 0xf3, 0xe1, 0x89, 0x4d, 0x92, 0x90, 0x8e, 0xa1, 0xbf, 0x15, 0x03, 0x8d, 0x4c, 0x7a,
	0x97, 0x43, 0x5b, 0x16, 0xc6, 0x5b, 0x78, 0x52, 0x5d, 0x2a, 0xa9, 0xad, 0xcd, 0x1d, 0xe4, 0x37,
	0x30, 0xaa, 0xf8, 0x69, 0x71, 0x95, 0x43, 0xdc, 0x1f, 0xa0, 0xb5, 0x16, 0x2f, 0x82, 0x35, 0x26,
	0x1b, 0x4e, 0x35, 0x78, 0xc0, 0xe5, 0xb3, 0xd0, 0x29, 0xcb, 0x52, 0x9d, 0xec, 0x71, 0xa6, 0x54,
	0xd5, 0xbf, 0xc3, 0x79, 0xd3, 0xd9, 0x3f, 0x89, 0x2b, 0x87, 0xc4, 0x3f, 0xc0, 0xd3, 0xaa, 0xf5,
	0x23, 0x5c, 0x1b, 0xef, 0x81, 0x56, 0xfc, 0x1d, 0x61, 0xcd, 0x78, 0x07, 0x03, 0x8b, 0x79, 0x73,
	0x8c, 0x05, 0xcb, 0x5f, 0x32, 0x7f, 0x8e, 0x71, 0xce, 0xea, 0xdb, 0x65, 0x29, 0x26, 0x51, 0xc0,
	0xb8, 0x98, 0x10, 0x39, 0x29, 0xca, 0xeb, 0xdf, 0x04, 0x4e, 0xeb, 0xd9, 0x3b, 0xf2, 0x58, 0xe9,
	0x2d, 0x28, 0x0e, 0x72, 0x7a, 0x6a, 0xa6, 0xae, 0xd9, 0xfa, 0x41, 0xfa, 0x99, 0x29, 0x2f, 0xd8,
	0x2c, 0x2f, 0xd8, 0x9c, 0x8a, 0x0b, 0xa6, 0x9f, 0x40, 0x75, 0x90, 0x7f, 0x0b, 0xb8, 0x5f, 0x7e,
	0xcb, 0xb3, 0xbd, 0x02, 0xc5, 0xb4, 0x53, 0xe7, 0x25, 0x28, 0x33, 0xe4, 0x74, 0xd4, 0x24, 0xeb,
	0xc3, 0xa2, 0x23, 0x8f, 0xec, 0x1e, 0xd4, 0x19, 0xf2, 0x85, 0x8f, 0xb1, 0x8d, 0xab, 0x0c, 0x99,
	0x4f, 0xcf, 0x9b, 0x8c, 0x72, 0x53, 0x9d, 0xf8, 0x0a, 0x06, 0x33, 0x14, 0x09, 0xd0, 0x0e, 0x07,
	0x3a, 0x14, 0x78, 0x81, 0xb9, 0x87, 0xfe, 0xc7, 0x08, 0x97, 0x59, 0x27, 0xb8, 0xa3, 0x7f, 0xfd,
	0x8b, 0xc0, 0xb8, 0x76, 0x7c, 0x65, 0xba, 0x37, 0x32, 0xdd, 0x71, 0xb1, 0xa4, 0x86, 0xe9, 0x0c,
	0x65, 0xda, 0x0a, 0xf7, 0x62, 0x1f, 0xff, 0x6f, 0xd9, 0xbe, 0x90, 0xd9, 0x3e, 0x6e, 0x70, 0x1b,
	0x09, 0xdd, 0xb5, 0xa2, 0x3d, 0x6b, 0x10, 0xfe, 0x67, 0xb2, 0xee, 0x20, 0xaf, 0x6f, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x38, 0x5f, 0x2f, 0xd7, 0x5e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StringKeyValueServiceClient is the client API for StringKeyValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringKeyValueServiceClient interface {
	//设置缓存
	Set(ctx context.Context, in *MsgStringKeyValue, opts ...grpc.CallOption) (*empty.Empty, error)
	//设置缓存-带超时
	SetWithTimeout(ctx context.Context, in *MsgStringKeyValueTimeout, opts ...grpc.CallOption) (*empty.Empty, error)
	//获取缓存
	Get(ctx context.Context, in *MsgStringKey, opts ...grpc.CallOption) (*MsgValue, error)
	//获取缓存后刷新
	GetThenRefresh(ctx context.Context, in *MsgStringKeyTimeout, opts ...grpc.CallOption) (*MsgValue, error)
	//获取缓存长度
	GetLen(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MsgLen, error)
	//清空缓存
	Clear(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type stringKeyValueServiceClient struct {
	cc *grpc.ClientConn
}

func NewStringKeyValueServiceClient(cc *grpc.ClientConn) StringKeyValueServiceClient {
	return &stringKeyValueServiceClient{cc}
}

func (c *stringKeyValueServiceClient) Set(ctx context.Context, in *MsgStringKeyValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.StringKeyValueService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringKeyValueServiceClient) SetWithTimeout(ctx context.Context, in *MsgStringKeyValueTimeout, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.StringKeyValueService/SetWithTimeout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringKeyValueServiceClient) Get(ctx context.Context, in *MsgStringKey, opts ...grpc.CallOption) (*MsgValue, error) {
	out := new(MsgValue)
	err := c.cc.Invoke(ctx, "/pb.StringKeyValueService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringKeyValueServiceClient) GetThenRefresh(ctx context.Context, in *MsgStringKeyTimeout, opts ...grpc.CallOption) (*MsgValue, error) {
	out := new(MsgValue)
	err := c.cc.Invoke(ctx, "/pb.StringKeyValueService/GetThenRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringKeyValueServiceClient) GetLen(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MsgLen, error) {
	out := new(MsgLen)
	err := c.cc.Invoke(ctx, "/pb.StringKeyValueService/GetLen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringKeyValueServiceClient) Clear(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.StringKeyValueService/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringKeyValueServiceServer is the server API for StringKeyValueService service.
type StringKeyValueServiceServer interface {
	//设置缓存
	Set(context.Context, *MsgStringKeyValue) (*empty.Empty, error)
	//设置缓存-带超时
	SetWithTimeout(context.Context, *MsgStringKeyValueTimeout) (*empty.Empty, error)
	//获取缓存
	Get(context.Context, *MsgStringKey) (*MsgValue, error)
	//获取缓存后刷新
	GetThenRefresh(context.Context, *MsgStringKeyTimeout) (*MsgValue, error)
	//获取缓存长度
	GetLen(context.Context, *empty.Empty) (*MsgLen, error)
	//清空缓存
	Clear(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterStringKeyValueServiceServer(s *grpc.Server, srv StringKeyValueServiceServer) {
	s.RegisterService(&_StringKeyValueService_serviceDesc, srv)
}

func _StringKeyValueService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStringKeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringKeyValueServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringKeyValueService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringKeyValueServiceServer).Set(ctx, req.(*MsgStringKeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringKeyValueService_SetWithTimeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStringKeyValueTimeout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringKeyValueServiceServer).SetWithTimeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringKeyValueService/SetWithTimeout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringKeyValueServiceServer).SetWithTimeout(ctx, req.(*MsgStringKeyValueTimeout))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringKeyValueService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStringKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringKeyValueServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringKeyValueService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringKeyValueServiceServer).Get(ctx, req.(*MsgStringKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringKeyValueService_GetThenRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStringKeyTimeout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringKeyValueServiceServer).GetThenRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringKeyValueService/GetThenRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringKeyValueServiceServer).GetThenRefresh(ctx, req.(*MsgStringKeyTimeout))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringKeyValueService_GetLen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringKeyValueServiceServer).GetLen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringKeyValueService/GetLen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringKeyValueServiceServer).GetLen(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringKeyValueService_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringKeyValueServiceServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringKeyValueService/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringKeyValueServiceServer).Clear(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringKeyValueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StringKeyValueService",
	HandlerType: (*StringKeyValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _StringKeyValueService_Set_Handler,
		},
		{
			MethodName: "SetWithTimeout",
			Handler:    _StringKeyValueService_SetWithTimeout_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StringKeyValueService_Get_Handler,
		},
		{
			MethodName: "GetThenRefresh",
			Handler:    _StringKeyValueService_GetThenRefresh_Handler,
		},
		{
			MethodName: "GetLen",
			Handler:    _StringKeyValueService_GetLen_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _StringKeyValueService_Clear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kservice.proto",
}

// Int64KeyValueServiceClient is the client API for Int64KeyValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Int64KeyValueServiceClient interface {
	//设置缓存
	Set(ctx context.Context, in *MsgInt64KeyValue, opts ...grpc.CallOption) (*empty.Empty, error)
	//设置缓存-带超时
	SetWithTimeout(ctx context.Context, in *MsgInt64KeyValueTimeout, opts ...grpc.CallOption) (*empty.Empty, error)
	//获取缓存
	Get(ctx context.Context, in *MsgInt64Key, opts ...grpc.CallOption) (*MsgValue, error)
	//获取缓存后刷新
	GetThenRefresh(ctx context.Context, in *MsgInt64KeyTimeout, opts ...grpc.CallOption) (*MsgValue, error)
	//获取缓存长度
	GetLen(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MsgLen, error)
	//清空缓存
	Clear(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type int64KeyValueServiceClient struct {
	cc *grpc.ClientConn
}

func NewInt64KeyValueServiceClient(cc *grpc.ClientConn) Int64KeyValueServiceClient {
	return &int64KeyValueServiceClient{cc}
}

func (c *int64KeyValueServiceClient) Set(ctx context.Context, in *MsgInt64KeyValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Int64KeyValueService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *int64KeyValueServiceClient) SetWithTimeout(ctx context.Context, in *MsgInt64KeyValueTimeout, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Int64KeyValueService/SetWithTimeout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *int64KeyValueServiceClient) Get(ctx context.Context, in *MsgInt64Key, opts ...grpc.CallOption) (*MsgValue, error) {
	out := new(MsgValue)
	err := c.cc.Invoke(ctx, "/pb.Int64KeyValueService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *int64KeyValueServiceClient) GetThenRefresh(ctx context.Context, in *MsgInt64KeyTimeout, opts ...grpc.CallOption) (*MsgValue, error) {
	out := new(MsgValue)
	err := c.cc.Invoke(ctx, "/pb.Int64KeyValueService/GetThenRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *int64KeyValueServiceClient) GetLen(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MsgLen, error) {
	out := new(MsgLen)
	err := c.cc.Invoke(ctx, "/pb.Int64KeyValueService/GetLen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *int64KeyValueServiceClient) Clear(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Int64KeyValueService/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Int64KeyValueServiceServer is the server API for Int64KeyValueService service.
type Int64KeyValueServiceServer interface {
	//设置缓存
	Set(context.Context, *MsgInt64KeyValue) (*empty.Empty, error)
	//设置缓存-带超时
	SetWithTimeout(context.Context, *MsgInt64KeyValueTimeout) (*empty.Empty, error)
	//获取缓存
	Get(context.Context, *MsgInt64Key) (*MsgValue, error)
	//获取缓存后刷新
	GetThenRefresh(context.Context, *MsgInt64KeyTimeout) (*MsgValue, error)
	//获取缓存长度
	GetLen(context.Context, *empty.Empty) (*MsgLen, error)
	//清空缓存
	Clear(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterInt64KeyValueServiceServer(s *grpc.Server, srv Int64KeyValueServiceServer) {
	s.RegisterService(&_Int64KeyValueService_serviceDesc, srv)
}

func _Int64KeyValueService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInt64KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Int64KeyValueServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Int64KeyValueService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Int64KeyValueServiceServer).Set(ctx, req.(*MsgInt64KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Int64KeyValueService_SetWithTimeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInt64KeyValueTimeout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Int64KeyValueServiceServer).SetWithTimeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Int64KeyValueService/SetWithTimeout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Int64KeyValueServiceServer).SetWithTimeout(ctx, req.(*MsgInt64KeyValueTimeout))
	}
	return interceptor(ctx, in, info, handler)
}

func _Int64KeyValueService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInt64Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Int64KeyValueServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Int64KeyValueService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Int64KeyValueServiceServer).Get(ctx, req.(*MsgInt64Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Int64KeyValueService_GetThenRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInt64KeyTimeout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Int64KeyValueServiceServer).GetThenRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Int64KeyValueService/GetThenRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Int64KeyValueServiceServer).GetThenRefresh(ctx, req.(*MsgInt64KeyTimeout))
	}
	return interceptor(ctx, in, info, handler)
}

func _Int64KeyValueService_GetLen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Int64KeyValueServiceServer).GetLen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Int64KeyValueService/GetLen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Int64KeyValueServiceServer).GetLen(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Int64KeyValueService_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Int64KeyValueServiceServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Int64KeyValueService/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Int64KeyValueServiceServer).Clear(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Int64KeyValueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Int64KeyValueService",
	HandlerType: (*Int64KeyValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Int64KeyValueService_Set_Handler,
		},
		{
			MethodName: "SetWithTimeout",
			Handler:    _Int64KeyValueService_SetWithTimeout_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Int64KeyValueService_Get_Handler,
		},
		{
			MethodName: "GetThenRefresh",
			Handler:    _Int64KeyValueService_GetThenRefresh_Handler,
		},
		{
			MethodName: "GetLen",
			Handler:    _Int64KeyValueService_GetLen_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _Int64KeyValueService_Clear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kservice.proto",
}
