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
	Processed               int64    `protobuf:"varint,2,opt,name=processed,proto3" json:"processed,omitempty"`
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

func (m *Config) GetProcessed() int64 {
	if m != nil {
		return m.Processed
	}
	return 0
}

type Ping struct {
	Action               string      `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Issue                *Issue      `protobuf:"bytes,2,opt,name=issue,proto3" json:"issue,omitempty"`
	Ref                  string      `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	Repository           *Repository `protobuf:"bytes,4,opt,name=repository,proto3" json:"repository,omitempty"`
	RefType              string      `protobuf:"bytes,5,opt,name=ref_type,json=refType,proto3" json:"ref_type,omitempty"`
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

func (m *Ping) GetRefType() string {
	if m != nil {
		return m.RefType
	}
	return ""
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
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
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

func (m *Issue) GetTitle() string {
	if m != nil {
		return m.Title
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
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0xf3, 0xc7, 0x66, 0x04, 0x91, 0xa1, 0x6a, 0xac, 0x1e, 0x24, 0x27, 0x41, 0xa8,
	0xa0, 0x37, 0xc5, 0x8b, 0x9e, 0x7a, 0x91, 0x1a, 0xbc, 0x87, 0x24, 0x4e, 0xe2, 0x42, 0x92, 0x8d,
	0x9b, 0x5d, 0x25, 0xdf, 0xcb, 0x0f, 0x28, 0x3b, 0x49, 0xad, 0xf6, 0x36, 0xfb, 0xde, 0xef, 0x3d,
	0x1e, 0x2c, 0xcc, 0x4b, 0xa1, 0xdf, 0x4d, 0xa6, 0x28, 0x27, 0xf1, 0x49, 0x6a, 0xd9, 0x2a, 0xa9,
	0x25, 0x1e, 0xfc, 0x57, 0xa3, 0x1c, 0xfc, 0x27, 0xd9, 0x14, 0xa2, 0xc4, 0x7b, 0x58, 0x68, 0x51,
	0x53, 0x92, 0x91, 0xfe, 0x22, 0x6a, 0x92, 0x0f, 0x43, 0x86, 0x92, 0x56, 0xc9, 0x9c, 0xba, 0x2e,
	0x9c, 0x5c, 0x4c, 0x2e, 0x9d, 0xf8, 0xc4, 0x12, 0x8f, 0x03, 0xf0, 0x62, 0xfd, 0xf5, 0x60, 0xe3,
	0x39, 0x04, 0x23, 0x49, 0x6f, 0xe1, 0x94, 0xd9, 0xad, 0x10, 0x7d, 0x4f, 0xc0, 0x5d, 0x8b, 0xa6,
	0xc4, 0x63, 0xf0, 0xd3, 0x5c, 0x0b, 0xd9, 0x70, 0x5f, 0x10, 0x8f, 0x2f, 0xbc, 0x02, 0x4f, 0x74,
	0x9d, 0x21, 0x8e, 0xee, 0xdf, 0x1c, 0x2d, 0x77, 0xb6, 0xaf, 0xac, 0x19, 0x0f, 0x0c, 0x1e, 0x82,
	0xa3, 0xa8, 0x08, 0x1d, 0x6e, 0xb0, 0x27, 0xde, 0x01, 0x28, 0x6a, 0x65, 0x27, 0xb4, 0x54, 0x7d,
	0xe8, 0x72, 0xc7, 0x62, 0xb7, 0x23, 0xfe, 0x25, 0xe2, 0x3f, 0x34, 0x9e, 0xc2, 0x4c, 0x51, 0x91,
	0xe8, 0xbe, 0xa5, 0xd0, 0xe3, 0xca, 0x3d, 0x45, 0xc5, 0x6b, 0xdf, 0x52, 0xf4, 0x00, 0xb0, 0x0d,
	0x21, 0x82, 0xdb, 0xa4, 0x35, 0x8d, 0xcb, 0xf9, 0xc6, 0x33, 0x08, 0x0a, 0x53, 0x55, 0x09, 0x1b,
	0x53, 0x36, 0x66, 0x56, 0x78, 0x4e, 0x6b, 0x8a, 0xae, 0xc1, 0x5b, 0x6d, 0x06, 0x1b, 0x55, 0x8d,
	0x41, 0x7b, 0xe2, 0x1c, 0x3c, 0x2d, 0x74, 0xb5, 0xc9, 0x0c, 0x8f, 0xcc, 0xe7, 0x2f, 0xba, 0xfd,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x06, 0xe4, 0x50, 0x4a, 0xba, 0x01, 0x00, 0x00,
}
