// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vendor/cedar.proto

package internal

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LogStorage int32

const (
	LogStorage_LOG_STORAGE_S3     LogStorage = 0
	LogStorage_LOG_STORAGE_GRIDFS LogStorage = 1
	LogStorage_LOG_STORAGE_LOCAL  LogStorage = 2
)

var LogStorage_name = map[int32]string{
	0: "LOG_STORAGE_S3",
	1: "LOG_STORAGE_GRIDFS",
	2: "LOG_STORAGE_LOCAL",
}

var LogStorage_value = map[string]int32{
	"LOG_STORAGE_S3":     0,
	"LOG_STORAGE_GRIDFS": 1,
	"LOG_STORAGE_LOCAL":  2,
}

func (x LogStorage) String() string {
	return proto.EnumName(LogStorage_name, int32(x))
}

func (LogStorage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{0}
}

type LogFormat int32

const (
	LogFormat_LOG_FORMAT_UNKNOWN LogFormat = 0
	LogFormat_LOG_FORMAT_TEXT    LogFormat = 1
	LogFormat_LOG_FORMAT_JSON    LogFormat = 2
	LogFormat_LOG_FORMAT_BSON    LogFormat = 3
)

var LogFormat_name = map[int32]string{
	0: "LOG_FORMAT_UNKNOWN",
	1: "LOG_FORMAT_TEXT",
	2: "LOG_FORMAT_JSON",
	3: "LOG_FORMAT_BSON",
}

var LogFormat_value = map[string]int32{
	"LOG_FORMAT_UNKNOWN": 0,
	"LOG_FORMAT_TEXT":    1,
	"LOG_FORMAT_JSON":    2,
	"LOG_FORMAT_BSON":    3,
}

func (x LogFormat) String() string {
	return proto.EnumName(LogFormat_name, int32(x))
}

func (LogFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{1}
}

type LogData struct {
	Info                 *LogInfo             `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Storage              LogStorage           `protobuf:"varint,2,opt,name=storage,proto3,enum=cedar.LogStorage" json:"storage,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogData) Reset()         { *m = LogData{} }
func (m *LogData) String() string { return proto.CompactTextString(m) }
func (*LogData) ProtoMessage()    {}
func (*LogData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{0}
}

func (m *LogData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogData.Unmarshal(m, b)
}
func (m *LogData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogData.Marshal(b, m, deterministic)
}
func (m *LogData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogData.Merge(m, src)
}
func (m *LogData) XXX_Size() int {
	return xxx_messageInfo_LogData.Size(m)
}
func (m *LogData) XXX_DiscardUnknown() {
	xxx_messageInfo_LogData.DiscardUnknown(m)
}

var xxx_messageInfo_LogData proto.InternalMessageInfo

func (m *LogData) GetInfo() *LogInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *LogData) GetStorage() LogStorage {
	if m != nil {
		return m.Storage
	}
	return LogStorage_LOG_STORAGE_S3
}

func (m *LogData) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type LogInfo struct {
	Project              string            `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Version              string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Variant              string            `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	TaskName             string            `protobuf:"bytes,4,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskId               string            `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Execution            int32             `protobuf:"varint,6,opt,name=execution,proto3" json:"execution,omitempty"`
	TestName             string            `protobuf:"bytes,7,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	Trial                int32             `protobuf:"varint,8,opt,name=trial,proto3" json:"trial,omitempty"`
	ProcName             string            `protobuf:"bytes,9,opt,name=proc_name,json=procName,proto3" json:"proc_name,omitempty"`
	Format               LogFormat         `protobuf:"varint,10,opt,name=format,proto3,enum=cedar.LogFormat" json:"format,omitempty"`
	Arguments            map[string]string `protobuf:"bytes,11,rep,name=arguments,proto3" json:"arguments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mainline             bool              `protobuf:"varint,12,opt,name=mainline,proto3" json:"mainline,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LogInfo) Reset()         { *m = LogInfo{} }
func (m *LogInfo) String() string { return proto.CompactTextString(m) }
func (*LogInfo) ProtoMessage()    {}
func (*LogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{1}
}

func (m *LogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInfo.Unmarshal(m, b)
}
func (m *LogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInfo.Marshal(b, m, deterministic)
}
func (m *LogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInfo.Merge(m, src)
}
func (m *LogInfo) XXX_Size() int {
	return xxx_messageInfo_LogInfo.Size(m)
}
func (m *LogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogInfo proto.InternalMessageInfo

func (m *LogInfo) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *LogInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LogInfo) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func (m *LogInfo) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *LogInfo) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *LogInfo) GetExecution() int32 {
	if m != nil {
		return m.Execution
	}
	return 0
}

func (m *LogInfo) GetTestName() string {
	if m != nil {
		return m.TestName
	}
	return ""
}

func (m *LogInfo) GetTrial() int32 {
	if m != nil {
		return m.Trial
	}
	return 0
}

func (m *LogInfo) GetProcName() string {
	if m != nil {
		return m.ProcName
	}
	return ""
}

func (m *LogInfo) GetFormat() LogFormat {
	if m != nil {
		return m.Format
	}
	return LogFormat_LOG_FORMAT_UNKNOWN
}

func (m *LogInfo) GetArguments() map[string]string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *LogInfo) GetMainline() bool {
	if m != nil {
		return m.Mainline
	}
	return false
}

type LogLines struct {
	LogId                string     `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Lines                []*LogLine `protobuf:"bytes,2,rep,name=lines,proto3" json:"lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LogLines) Reset()         { *m = LogLines{} }
func (m *LogLines) String() string { return proto.CompactTextString(m) }
func (*LogLines) ProtoMessage()    {}
func (*LogLines) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{2}
}

func (m *LogLines) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLines.Unmarshal(m, b)
}
func (m *LogLines) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLines.Marshal(b, m, deterministic)
}
func (m *LogLines) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLines.Merge(m, src)
}
func (m *LogLines) XXX_Size() int {
	return xxx_messageInfo_LogLines.Size(m)
}
func (m *LogLines) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLines.DiscardUnknown(m)
}

var xxx_messageInfo_LogLines proto.InternalMessageInfo

func (m *LogLines) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

func (m *LogLines) GetLines() []*LogLine {
	if m != nil {
		return m.Lines
	}
	return nil
}

type LogLine struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data                 string               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogLine) Reset()         { *m = LogLine{} }
func (m *LogLine) String() string { return proto.CompactTextString(m) }
func (*LogLine) ProtoMessage()    {}
func (*LogLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{3}
}

func (m *LogLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLine.Unmarshal(m, b)
}
func (m *LogLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLine.Marshal(b, m, deterministic)
}
func (m *LogLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLine.Merge(m, src)
}
func (m *LogLine) XXX_Size() int {
	return xxx_messageInfo_LogLine.Size(m)
}
func (m *LogLine) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLine.DiscardUnknown(m)
}

var xxx_messageInfo_LogLine proto.InternalMessageInfo

func (m *LogLine) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogLine) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type LogEndInfo struct {
	LogId                string               `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	ExitCode             int32                `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	CompletedAt          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogEndInfo) Reset()         { *m = LogEndInfo{} }
func (m *LogEndInfo) String() string { return proto.CompactTextString(m) }
func (*LogEndInfo) ProtoMessage()    {}
func (*LogEndInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{4}
}

func (m *LogEndInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEndInfo.Unmarshal(m, b)
}
func (m *LogEndInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEndInfo.Marshal(b, m, deterministic)
}
func (m *LogEndInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEndInfo.Merge(m, src)
}
func (m *LogEndInfo) XXX_Size() int {
	return xxx_messageInfo_LogEndInfo.Size(m)
}
func (m *LogEndInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEndInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogEndInfo proto.InternalMessageInfo

func (m *LogEndInfo) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

func (m *LogEndInfo) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *LogEndInfo) GetCompletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CompletedAt
	}
	return nil
}

type BuildloggerResponse struct {
	LogId                string   `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildloggerResponse) Reset()         { *m = BuildloggerResponse{} }
func (m *BuildloggerResponse) String() string { return proto.CompactTextString(m) }
func (*BuildloggerResponse) ProtoMessage()    {}
func (*BuildloggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f43d153eecc12f, []int{5}
}

func (m *BuildloggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildloggerResponse.Unmarshal(m, b)
}
func (m *BuildloggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildloggerResponse.Marshal(b, m, deterministic)
}
func (m *BuildloggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildloggerResponse.Merge(m, src)
}
func (m *BuildloggerResponse) XXX_Size() int {
	return xxx_messageInfo_BuildloggerResponse.Size(m)
}
func (m *BuildloggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildloggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildloggerResponse proto.InternalMessageInfo

func (m *BuildloggerResponse) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

func init() {
	proto.RegisterEnum("cedar.LogStorage", LogStorage_name, LogStorage_value)
	proto.RegisterEnum("cedar.LogFormat", LogFormat_name, LogFormat_value)
	proto.RegisterType((*LogData)(nil), "cedar.LogData")
	proto.RegisterType((*LogInfo)(nil), "cedar.LogInfo")
	proto.RegisterMapType((map[string]string)(nil), "cedar.LogInfo.ArgumentsEntry")
	proto.RegisterType((*LogLines)(nil), "cedar.LogLines")
	proto.RegisterType((*LogLine)(nil), "cedar.LogLine")
	proto.RegisterType((*LogEndInfo)(nil), "cedar.LogEndInfo")
	proto.RegisterType((*BuildloggerResponse)(nil), "cedar.BuildloggerResponse")
}

func init() { proto.RegisterFile("vendor/cedar.proto", fileDescriptor_f2f43d153eecc12f) }

var fileDescriptor_f2f43d153eecc12f = []byte{
	// 723 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0xc5, 0x09, 0x49, 0xec, 0x1b, 0x14, 0xcc, 0xf0, 0x78, 0xcf, 0x0a, 0xaf, 0x6a, 0x14, 0x75,
	0x11, 0xd1, 0x2a, 0x48, 0x61, 0x51, 0xa0, 0x65, 0x11, 0x42, 0x88, 0xd2, 0xba, 0x89, 0x34, 0x49,
	0x45, 0xd5, 0x4d, 0x34, 0xc4, 0x13, 0xcb, 0xc5, 0x9e, 0xb1, 0xc6, 0x13, 0x54, 0x56, 0xfd, 0x15,
	0xfd, 0x37, 0xfd, 0x69, 0x5d, 0x54, 0x33, 0x76, 0x3e, 0x40, 0x05, 0xb5, 0x3b, 0xdf, 0x73, 0xee,
	0xb9, 0x73, 0x7c, 0xe7, 0xd8, 0x80, 0x6e, 0x29, 0xf3, 0xb8, 0x38, 0x9c, 0x52, 0x8f, 0x88, 0x66,
	0x2c, 0xb8, 0xe4, 0xa8, 0xa0, 0x8b, 0xea, 0x73, 0x9f, 0x73, 0x3f, 0xa4, 0x87, 0x1a, 0xbc, 0x9e,
	0xcf, 0x0e, 0x65, 0x10, 0xd1, 0x44, 0x92, 0x28, 0x4e, 0xfb, 0xea, 0xdf, 0x0d, 0x28, 0xb9, 0xdc,
	0xbf, 0x20, 0x92, 0xa0, 0x3a, 0x6c, 0x06, 0x6c, 0xc6, 0x1d, 0xa3, 0x66, 0x34, 0xca, 0xad, 0x4a,
	0x33, 0x9d, 0xe7, 0x72, 0xbf, 0xcf, 0x66, 0x1c, 0x6b, 0x0e, 0xbd, 0x84, 0x52, 0x22, 0xb9, 0x20,
	0x3e, 0x75, 0x72, 0x35, 0xa3, 0x51, 0x69, 0xed, 0xac, 0xda, 0x46, 0x29, 0x81, 0x17, 0x1d, 0xe8,
	0x04, 0x60, 0x2a, 0x28, 0x91, 0xd4, 0x9b, 0x10, 0xe9, 0xe4, 0xf5, 0xd8, 0x6a, 0x33, 0xb5, 0xd4,
	0x5c, 0x58, 0x6a, 0x8e, 0x17, 0x96, 0xb0, 0x95, 0x75, 0xb7, 0x65, 0xfd, 0x47, 0x5e, 0xfb, 0x52,
	0x27, 0x23, 0x07, 0x4a, 0xb1, 0xe0, 0x5f, 0xe8, 0x54, 0x6a, 0x6b, 0x16, 0x5e, 0x94, 0x8a, 0xb9,
	0xa5, 0x22, 0x09, 0x38, 0xd3, 0x6e, 0x2c, 0xbc, 0x28, 0x35, 0x43, 0x44, 0x40, 0x58, 0x7a, 0xae,
	0x62, 0xd2, 0x12, 0xed, 0x83, 0x25, 0x49, 0x72, 0x33, 0x61, 0x24, 0xa2, 0xce, 0xa6, 0xe6, 0x4c,
	0x05, 0x0c, 0x48, 0x44, 0xd1, 0x7f, 0x50, 0xd2, 0x64, 0xe0, 0x39, 0x05, 0x4d, 0x15, 0x55, 0xd9,
	0xf7, 0xd0, 0xff, 0x60, 0xd1, 0xaf, 0x74, 0x3a, 0x97, 0xea, 0xac, 0x62, 0xcd, 0x68, 0x14, 0xf0,
	0x0a, 0xd0, 0x33, 0x69, 0x22, 0xd3, 0x99, 0xa5, 0x6c, 0x26, 0x4d, 0xa4, 0x9e, 0xf9, 0x0f, 0x14,
	0xa4, 0x08, 0x48, 0xe8, 0x98, 0x5a, 0x96, 0x16, 0x4a, 0x12, 0x0b, 0x3e, 0x4d, 0x25, 0x56, 0x2a,
	0x51, 0x80, 0x96, 0x34, 0xa0, 0x38, 0xe3, 0x22, 0x22, 0xd2, 0x01, 0xbd, 0x64, 0x7b, 0xb5, 0xe4,
	0x4b, 0x8d, 0xe3, 0x8c, 0x47, 0x6f, 0xc0, 0x22, 0xc2, 0x9f, 0x47, 0x94, 0xc9, 0xc4, 0x29, 0xd7,
	0xf2, 0x8d, 0x72, 0xeb, 0xd9, 0xfd, 0x8b, 0x6b, 0xb6, 0x17, 0x7c, 0x97, 0x49, 0x71, 0x87, 0x57,
	0xfd, 0xa8, 0x0a, 0x66, 0x44, 0x02, 0x16, 0x06, 0x8c, 0x3a, 0x5b, 0x35, 0xa3, 0x61, 0xe2, 0x65,
	0x5d, 0x7d, 0x0b, 0x95, 0xfb, 0x42, 0x64, 0x43, 0xfe, 0x86, 0xde, 0x65, 0x57, 0xa0, 0x1e, 0xd5,
	0x9b, 0xdd, 0x92, 0x70, 0x4e, 0xb3, 0xe5, 0xa7, 0xc5, 0x69, 0xee, 0xd8, 0xa8, 0xf7, 0xc0, 0x74,
	0xb9, 0xef, 0x06, 0x8c, 0x26, 0x68, 0x0f, 0x8a, 0x21, 0xf7, 0xd5, 0x4a, 0x53, 0x69, 0x21, 0xe4,
	0x7e, 0xdf, 0x43, 0x2f, 0xa0, 0xa0, 0x0e, 0x4a, 0x9c, 0x9c, 0x76, 0xbd, 0x16, 0x37, 0x25, 0xc3,
	0x29, 0x59, 0xbf, 0xd2, 0x31, 0x50, 0x08, 0x3a, 0x06, 0x6b, 0x99, 0xde, 0x2c, 0xa3, 0x4f, 0x86,
	0x69, 0xd9, 0x8c, 0x10, 0x6c, 0x7a, 0x44, 0x92, 0xcc, 0xa6, 0x7e, 0xae, 0x7f, 0x03, 0x70, 0xb9,
	0xdf, 0x65, 0x9e, 0x8e, 0xd8, 0x23, 0x1e, 0xf7, 0xd5, 0xad, 0x07, 0x72, 0x32, 0xe5, 0x5e, 0xfa,
	0x92, 0x05, 0x6c, 0x2a, 0xa0, 0xc3, 0x3d, 0x8a, 0xce, 0x60, 0x6b, 0xca, 0xa3, 0x38, 0xa4, 0x7f,
	0x9c, 0xef, 0xf2, 0xb2, 0xbf, 0x2d, 0xeb, 0xaf, 0x60, 0xf7, 0x7c, 0x1e, 0x84, 0x5e, 0xc8, 0x7d,
	0x9f, 0x0a, 0x4c, 0x93, 0x98, 0xb3, 0x84, 0x3e, 0xe2, 0xe4, 0x60, 0xa8, 0xed, 0x66, 0x5f, 0x18,
	0x42, 0x50, 0x71, 0x87, 0xbd, 0xc9, 0x68, 0x3c, 0xc4, 0xed, 0x5e, 0x77, 0x32, 0x3a, 0xb2, 0x37,
	0xd0, 0xbf, 0x80, 0xd6, 0xb1, 0x1e, 0xee, 0x5f, 0x5c, 0x8e, 0x6c, 0x03, 0xed, 0xc1, 0xce, 0x3a,
	0xee, 0x0e, 0x3b, 0x6d, 0xd7, 0xce, 0x1d, 0x5c, 0x83, 0xb5, 0x4c, 0xd3, 0x42, 0x7b, 0x39, 0xc4,
	0x1f, 0xda, 0xe3, 0xc9, 0xc7, 0xc1, 0xfb, 0xc1, 0xf0, 0x6a, 0x60, 0x6f, 0xa0, 0x5d, 0xd8, 0x5e,
	0xc3, 0xc7, 0xdd, 0x4f, 0x63, 0xdb, 0x78, 0x00, 0xbe, 0x1b, 0x0d, 0x07, 0x76, 0xee, 0x01, 0x78,
	0xae, 0xc0, 0x7c, 0xeb, 0xa7, 0x01, 0xe5, 0xb5, 0x77, 0x44, 0xaf, 0xc1, 0xea, 0xe8, 0x2f, 0xdc,
	0xe5, 0x3e, 0x5a, 0xbb, 0x70, 0xf5, 0xf7, 0xa9, 0x56, 0xb3, 0xfa, 0x77, 0x4b, 0x39, 0x83, 0x4a,
	0x3b, 0x8e, 0x29, 0xf3, 0x96, 0xa1, 0xda, 0xbe, 0x1f, 0x97, 0xe4, 0x49, 0xf9, 0x29, 0x58, 0x23,
	0x29, 0x28, 0x89, 0xd4, 0xb9, 0x7f, 0xa3, 0x6c, 0x18, 0xe8, 0x04, 0xcc, 0x4e, 0xc8, 0x13, 0x6d,
	0x79, 0xed, 0x5f, 0x97, 0x05, 0xe7, 0x29, 0xf1, 0x39, 0x7c, 0x36, 0x03, 0x26, 0xa9, 0x60, 0x24,
	0xbc, 0x2e, 0xea, 0x38, 0x1c, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x05, 0x5f, 0x43, 0x4f, 0xac,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildloggerClient is the client API for Buildlogger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildloggerClient interface {
	CreateLog(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*BuildloggerResponse, error)
	AppendLogLines(ctx context.Context, in *LogLines, opts ...grpc.CallOption) (*BuildloggerResponse, error)
	StreamLog(ctx context.Context, opts ...grpc.CallOption) (Buildlogger_StreamLogClient, error)
	CloseLog(ctx context.Context, in *LogEndInfo, opts ...grpc.CallOption) (*BuildloggerResponse, error)
}

type buildloggerClient struct {
	cc *grpc.ClientConn
}

func NewBuildloggerClient(cc *grpc.ClientConn) BuildloggerClient {
	return &buildloggerClient{cc}
}

func (c *buildloggerClient) CreateLog(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*BuildloggerResponse, error) {
	out := new(BuildloggerResponse)
	err := c.cc.Invoke(ctx, "/cedar.Buildlogger/CreateLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildloggerClient) AppendLogLines(ctx context.Context, in *LogLines, opts ...grpc.CallOption) (*BuildloggerResponse, error) {
	out := new(BuildloggerResponse)
	err := c.cc.Invoke(ctx, "/cedar.Buildlogger/AppendLogLines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildloggerClient) StreamLog(ctx context.Context, opts ...grpc.CallOption) (Buildlogger_StreamLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Buildlogger_serviceDesc.Streams[0], "/cedar.Buildlogger/StreamLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildloggerStreamLogClient{stream}
	return x, nil
}

type Buildlogger_StreamLogClient interface {
	Send(*LogLines) error
	CloseAndRecv() (*BuildloggerResponse, error)
	grpc.ClientStream
}

type buildloggerStreamLogClient struct {
	grpc.ClientStream
}

func (x *buildloggerStreamLogClient) Send(m *LogLines) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildloggerStreamLogClient) CloseAndRecv() (*BuildloggerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BuildloggerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildloggerClient) CloseLog(ctx context.Context, in *LogEndInfo, opts ...grpc.CallOption) (*BuildloggerResponse, error) {
	out := new(BuildloggerResponse)
	err := c.cc.Invoke(ctx, "/cedar.Buildlogger/CloseLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildloggerServer is the server API for Buildlogger service.
type BuildloggerServer interface {
	CreateLog(context.Context, *LogData) (*BuildloggerResponse, error)
	AppendLogLines(context.Context, *LogLines) (*BuildloggerResponse, error)
	StreamLog(Buildlogger_StreamLogServer) error
	CloseLog(context.Context, *LogEndInfo) (*BuildloggerResponse, error)
}

// UnimplementedBuildloggerServer can be embedded to have forward compatible implementations.
type UnimplementedBuildloggerServer struct {
}

func (*UnimplementedBuildloggerServer) CreateLog(ctx context.Context, req *LogData) (*BuildloggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLog not implemented")
}
func (*UnimplementedBuildloggerServer) AppendLogLines(ctx context.Context, req *LogLines) (*BuildloggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendLogLines not implemented")
}
func (*UnimplementedBuildloggerServer) StreamLog(srv Buildlogger_StreamLogServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLog not implemented")
}
func (*UnimplementedBuildloggerServer) CloseLog(ctx context.Context, req *LogEndInfo) (*BuildloggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseLog not implemented")
}

func RegisterBuildloggerServer(s *grpc.Server, srv BuildloggerServer) {
	s.RegisterService(&_Buildlogger_serviceDesc, srv)
}

func _Buildlogger_CreateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildloggerServer).CreateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.Buildlogger/CreateLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildloggerServer).CreateLog(ctx, req.(*LogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildlogger_AppendLogLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLines)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildloggerServer).AppendLogLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.Buildlogger/AppendLogLines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildloggerServer).AppendLogLines(ctx, req.(*LogLines))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildlogger_StreamLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildloggerServer).StreamLog(&buildloggerStreamLogServer{stream})
}

type Buildlogger_StreamLogServer interface {
	SendAndClose(*BuildloggerResponse) error
	Recv() (*LogLines, error)
	grpc.ServerStream
}

type buildloggerStreamLogServer struct {
	grpc.ServerStream
}

func (x *buildloggerStreamLogServer) SendAndClose(m *BuildloggerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildloggerStreamLogServer) Recv() (*LogLines, error) {
	m := new(LogLines)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Buildlogger_CloseLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogEndInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildloggerServer).CloseLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.Buildlogger/CloseLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildloggerServer).CloseLog(ctx, req.(*LogEndInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Buildlogger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cedar.Buildlogger",
	HandlerType: (*BuildloggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLog",
			Handler:    _Buildlogger_CreateLog_Handler,
		},
		{
			MethodName: "AppendLogLines",
			Handler:    _Buildlogger_AppendLogLines_Handler,
		},
		{
			MethodName: "CloseLog",
			Handler:    _Buildlogger_CloseLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLog",
			Handler:       _Buildlogger_StreamLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "vendor/cedar.proto",
}
