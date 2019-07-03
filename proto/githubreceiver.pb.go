// Code generated by protoc-gen-go. DO NOT EDIT.
// source: githubreceiver.proto

package githubreceiver

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	TimeBetweenQueueProcess int64    `protobuf:"varint,1,opt,name=time_between_queue_process,json=timeBetweenQueueProcess,proto3" json:"time_between_queue_process,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7162ad0e80472c, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetTimeBetweenQueueProcess() int64 {
	if m != nil {
		return m.TimeBetweenQueueProcess
	}
	return 0
}

type Ping struct {
	Action               string      `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Issue                *Issue      `protobuf:"bytes,2,opt,name=issue,proto3" json:"issue,omitempty"`
	Ref                  string      `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	Repository           *Repository `protobuf:"bytes,4,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7162ad0e80472c, []int{1}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Ping) GetIssue() *Issue {
	if m != nil {
		return m.Issue
	}
	return nil
}

func (m *Ping) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *Ping) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type Repository struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7162ad0e80472c, []int{2}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repository) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

type Issue struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Issue) Reset()         { *m = Issue{} }
func (m *Issue) String() string { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()    {}
func (*Issue) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7162ad0e80472c, []int{3}
}

func (m *Issue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Issue.Unmarshal(m, b)
}
func (m *Issue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Issue.Marshal(b, m, deterministic)
}
func (m *Issue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Issue.Merge(m, src)
}
func (m *Issue) XXX_Size() int {
	return xxx_messageInfo_Issue.Size(m)
}
func (m *Issue) XXX_DiscardUnknown() {
	xxx_messageInfo_Issue.DiscardUnknown(m)
}

var xxx_messageInfo_Issue proto.InternalMessageInfo

func (m *Issue) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "githubreceiver.Config")
	proto.RegisterType((*Ping)(nil), "githubreceiver.Ping")
	proto.RegisterType((*Repository)(nil), "githubreceiver.Repository")
	proto.RegisterType((*Issue)(nil), "githubreceiver.Issue")
}

func init() { proto.RegisterFile("githubreceiver.proto", fileDescriptor_cc7162ad0e80472c) }

var fileDescriptor_cc7162ad0e80472c = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x93, 0x06, 0x3b, 0x82, 0xc8, 0xe0, 0x47, 0xac, 0x97, 0x92, 0x53, 0x41, 0xe8,
	0x41, 0x6f, 0x8a, 0x17, 0xc5, 0x43, 0x2f, 0x52, 0xf7, 0x0f, 0x2c, 0x49, 0x98, 0xc4, 0x85, 0x64,
	0x37, 0xee, 0x87, 0xe2, 0x4f, 0xf1, 0xdf, 0xca, 0x6e, 0xe2, 0x47, 0x73, 0x9b, 0x9d, 0xe7, 0x99,
	0x97, 0x97, 0x85, 0x93, 0x46, 0xd8, 0x57, 0x57, 0x6a, 0xaa, 0x48, 0xbc, 0x93, 0xde, 0xf4, 0x5a,
	0x59, 0x85, 0x47, 0xfb, 0xdb, 0xfc, 0x09, 0xd2, 0x47, 0x25, 0x6b, 0xd1, 0xe0, 0x1d, 0x2c, 0xad,
	0xe8, 0x88, 0x97, 0x64, 0x3f, 0x88, 0x24, 0x7f, 0x73, 0xe4, 0x88, 0xf7, 0x5a, 0x55, 0x64, 0x4c,
	0x16, 0xad, 0xa2, 0x75, 0xcc, 0xce, 0xbd, 0xf1, 0x30, 0x08, 0x2f, 0x9e, 0xef, 0x06, 0x9c, 0x7f,
	0x45, 0x90, 0xec, 0x84, 0x6c, 0xf0, 0x0c, 0xd2, 0xa2, 0xb2, 0x42, 0xc9, 0x70, 0xb1, 0x60, 0xe3,
	0x0b, 0xaf, 0x60, 0x2e, 0x8c, 0x71, 0x94, 0xcd, 0x56, 0xd1, 0xfa, 0xf0, 0xfa, 0x74, 0x33, 0x69,
	0xb7, 0xf5, 0x90, 0x0d, 0x0e, 0x1e, 0x43, 0xac, 0xa9, 0xce, 0xe2, 0x90, 0xe0, 0x47, 0xbc, 0x05,
	0xd0, 0xd4, 0x2b, 0x23, 0xac, 0xd2, 0x9f, 0x59, 0x12, 0x32, 0x96, 0xd3, 0x0c, 0xf6, 0x6b, 0xb0,
	0x7f, 0x76, 0x7e, 0x0f, 0xf0, 0x47, 0x10, 0x21, 0x91, 0x45, 0x47, 0x63, 0xbd, 0x30, 0xe3, 0x25,
	0x2c, 0x6a, 0xd7, 0xb6, 0x3c, 0x80, 0x59, 0x00, 0x07, 0x7e, 0xf1, 0x5c, 0x74, 0x94, 0x5f, 0xc0,
	0x7c, 0xfb, 0xd3, 0xca, 0xe9, 0x76, 0x3c, 0xf4, 0x63, 0x99, 0x86, 0x3f, 0xbd, 0xf9, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x2b, 0x6f, 0x01, 0xa0, 0x6b, 0x01, 0x00, 0x00,
}
