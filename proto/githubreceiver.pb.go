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
	Name                 string      `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	CheckRun             *CheckRun   `protobuf:"bytes,8,opt,name=check_run,json=checkRun,proto3" json:"check_run,omitempty"`
	CheckSuite           *CheckSuite `protobuf:"bytes,7,opt,name=check_suite,json=checkSuite,proto3" json:"check_suite,omitempty"`
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

func (m *Ping) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ping) GetCheckRun() *CheckRun {
	if m != nil {
		return m.CheckRun
	}
	return nil
}

func (m *Ping) GetCheckSuite() *CheckSuite {
	if m != nil {
		return m.CheckSuite
	}
	return nil
}

type CheckSuite struct {
	PullRequests         []*PullRequest `protobuf:"bytes,1,rep,name=pull_requests,json=pullRequests,proto3" json:"pull_requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CheckSuite) Reset()         { *m = CheckSuite{} }
func (m *CheckSuite) String() string { return proto.CompactTextString(m) }
func (*CheckSuite) ProtoMessage()    {}
func (*CheckSuite) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7162ad0e80472c, []int{2}
}

func (m *CheckSuite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSuite.Unmarshal(m, b)
}
func (m *CheckSuite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSuite.Marshal(b, m, deterministic)
}
func (m *CheckSuite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSuite.Merge(m, src)
}
func (m *CheckSuite) XXX_Size() int {
	return xxx_messageInfo_CheckSuite.Size(m)
}
func (m *CheckSuite) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSuite.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSuite proto.InternalMessageInfo

func (m *CheckSuite) GetPullRequests() []*PullRequest {
	if m != nil {
		return m.PullRequests
	}
	return nil
}

type PullRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullRequest) Reset()         { *m = PullRequest{} }
func (m *PullRequest) String() string { return proto.CompactTextString(m) }
func (*PullRequest) ProtoMessage()    {}
func (*PullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7162ad0e80472c, []int{3}
}

func (m *PullRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullRequest.Unmarshal(m, b)
}
func (m *PullRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullRequest.Marshal(b, m, deterministic)
}
func (m *PullRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullRequest.Merge(m, src)
}
func (m *PullRequest) XXX_Size() int {
	return xxx_messageInfo_PullRequest.Size(m)
}
func (m *PullRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullRequest proto.InternalMessageInfo

func (m *PullRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type CheckRun struct {
	Conclusion           string   `protobuf:"bytes,1,opt,name=conclusion,proto3" json:"conclusion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRun) Reset()         { *m = CheckRun{} }
func (m *CheckRun) String() string { return proto.CompactTextString(m) }
func (*CheckRun) ProtoMessage()    {}
func (*CheckRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7162ad0e80472c, []int{4}
}

func (m *CheckRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRun.Unmarshal(m, b)
}
func (m *CheckRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRun.Marshal(b, m, deterministic)
}
func (m *CheckRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRun.Merge(m, src)
}
func (m *CheckRun) XXX_Size() int {
	return xxx_messageInfo_CheckRun.Size(m)
}
func (m *CheckRun) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRun.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRun proto.InternalMessageInfo

func (m *CheckRun) GetConclusion() string {
	if m != nil {
		return m.Conclusion
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
	return fileDescriptor_cc7162ad0e80472c, []int{5}
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
	return fileDescriptor_cc7162ad0e80472c, []int{6}
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
	proto.RegisterType((*CheckSuite)(nil), "githubreceiver.CheckSuite")
	proto.RegisterType((*PullRequest)(nil), "githubreceiver.PullRequest")
	proto.RegisterType((*CheckRun)(nil), "githubreceiver.CheckRun")
	proto.RegisterType((*Repository)(nil), "githubreceiver.Repository")
	proto.RegisterType((*Issue)(nil), "githubreceiver.Issue")
}

func init() { proto.RegisterFile("githubreceiver.proto", fileDescriptor_cc7162ad0e80472c) }

var fileDescriptor_cc7162ad0e80472c = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x95, 0x38, 0x76, 0xed, 0x31, 0x20, 0xb4, 0x2a, 0xb0, 0xb4, 0x08, 0x90, 0x4f, 0x15,
	0x48, 0x45, 0x2a, 0xe2, 0x42, 0x85, 0x84, 0xe8, 0xa9, 0x97, 0x28, 0x2c, 0xdc, 0xad, 0x64, 0x19,
	0x27, 0x2b, 0x9c, 0x5d, 0x67, 0xff, 0x80, 0xf2, 0x6e, 0x3c, 0x1c, 0xda, 0xb5, 0x13, 0x3b, 0x56,
	0x6f, 0x33, 0xf3, 0xfd, 0xe6, 0xdb, 0xf1, 0x8c, 0xe1, 0x7c, 0x2d, 0xec, 0xc6, 0xad, 0x34, 0x72,
	0x14, 0x7f, 0x50, 0x5f, 0x37, 0x5a, 0x59, 0x45, 0x9e, 0x9c, 0x56, 0x0b, 0x0e, 0xc9, 0x9d, 0x92,
	0x95, 0x58, 0x93, 0x5b, 0xb8, 0xb0, 0x62, 0x8b, 0xe5, 0x0a, 0xed, 0x5f, 0x44, 0x59, 0xee, 0x1c,
	0x3a, 0x2c, 0x1b, 0xad, 0x38, 0x1a, 0x43, 0x27, 0x6f, 0x27, 0x57, 0x11, 0x7b, 0xe1, 0x89, 0x6f,
	0x2d, 0xf0, 0xdd, 0xeb, 0x8b, 0x56, 0x26, 0xaf, 0x20, 0xeb, 0x48, 0xfc, 0x45, 0xa7, 0x81, 0xed,
	0x0b, 0xc5, 0xbf, 0x29, 0xcc, 0x16, 0x42, 0xae, 0xc9, 0x73, 0x48, 0x96, 0xdc, 0x0a, 0x25, 0x83,
	0x5f, 0xc6, 0xba, 0x8c, 0xbc, 0x87, 0x58, 0x18, 0xe3, 0x30, 0xb4, 0xe6, 0x37, 0xcf, 0xae, 0x47,
	0xb3, 0xdf, 0x7b, 0x91, 0xb5, 0x0c, 0x79, 0x0a, 0x91, 0xc6, 0x8a, 0x46, 0xc1, 0xc1, 0x87, 0xe4,
	0x33, 0x80, 0xc6, 0x46, 0x19, 0x61, 0x95, 0xde, 0xd3, 0x59, 0xf0, 0xb8, 0x18, 0x7b, 0xb0, 0x23,
	0xc1, 0x06, 0x34, 0x79, 0x09, 0xa9, 0xc6, 0xaa, 0xb4, 0xfb, 0x06, 0x69, 0x1c, 0x2c, 0xcf, 0x34,
	0x56, 0x3f, 0xf7, 0x0d, 0x12, 0x02, 0x33, 0xb9, 0xdc, 0x22, 0x4d, 0x42, 0x39, 0xc4, 0xe4, 0x13,
	0x64, 0x7c, 0x83, 0xfc, 0x77, 0xa9, 0x9d, 0xa4, 0x69, 0x78, 0x89, 0x8e, 0x5f, 0xba, 0xf3, 0x00,
	0x73, 0x92, 0xa5, 0xbc, 0x8b, 0xc8, 0x2d, 0xe4, 0x6d, 0x9b, 0x71, 0xc2, 0x22, 0x3d, 0x7b, 0x78,
	0xc4, 0xd0, 0xf8, 0xc3, 0x13, 0x0c, 0xf8, 0x31, 0x2e, 0xe6, 0x00, 0xbd, 0x42, 0xbe, 0xc2, 0xe3,
	0xc6, 0xd5, 0x75, 0xa9, 0x71, 0xe7, 0xd0, 0x58, 0x7f, 0x9a, 0xe8, 0x2a, 0xbf, 0xb9, 0x1c, 0x9b,
	0x2d, 0x5c, 0x5d, 0xb3, 0x96, 0x61, 0x8f, 0x9a, 0x3e, 0x31, 0xc5, 0x1b, 0xc8, 0x07, 0xa2, 0xdf,
	0xa7, 0xd3, 0x75, 0x77, 0x11, 0x1f, 0x16, 0xef, 0x20, 0x3d, 0x7c, 0x03, 0x79, 0x0d, 0xc0, 0x95,
	0xe4, 0xb5, 0x33, 0xfd, 0xd9, 0x06, 0x95, 0xe2, 0x0b, 0x40, 0xbf, 0xd9, 0xe3, 0xca, 0x26, 0x83,
	0x95, 0x5d, 0x42, 0x56, 0xf9, 0x81, 0x83, 0x30, 0x0d, 0x42, 0xea, 0x0b, 0xf3, 0xe5, 0x16, 0x8b,
	0x0f, 0x10, 0xdf, 0x1f, 0xae, 0x7a, 0x3a, 0x05, 0x39, 0x87, 0xd8, 0x0a, 0x5b, 0x1f, 0x7a, 0xda,
	0x64, 0x95, 0x84, 0xff, 0xf8, 0xe3, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x38, 0xae, 0x2e,
	0xdf, 0x02, 0x00, 0x00,
}
