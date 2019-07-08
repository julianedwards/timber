// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/budget_period.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible period of a Budget.
type BudgetPeriodEnum_BudgetPeriod int32

const (
	// Not specified.
	BudgetPeriodEnum_UNSPECIFIED BudgetPeriodEnum_BudgetPeriod = 0
	// Used for return value only. Represents value unknown in this version.
	BudgetPeriodEnum_UNKNOWN BudgetPeriodEnum_BudgetPeriod = 1
	// Daily budget.
	BudgetPeriodEnum_DAILY BudgetPeriodEnum_BudgetPeriod = 2
	// Custom budget.
	BudgetPeriodEnum_CUSTOM BudgetPeriodEnum_BudgetPeriod = 3
	// Fixed daily budget.
	BudgetPeriodEnum_FIXED_DAILY BudgetPeriodEnum_BudgetPeriod = 4
)

var BudgetPeriodEnum_BudgetPeriod_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DAILY",
	3: "CUSTOM",
	4: "FIXED_DAILY",
}
var BudgetPeriodEnum_BudgetPeriod_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"DAILY":       2,
	"CUSTOM":      3,
	"FIXED_DAILY": 4,
}

func (x BudgetPeriodEnum_BudgetPeriod) String() string {
	return proto.EnumName(BudgetPeriodEnum_BudgetPeriod_name, int32(x))
}
func (BudgetPeriodEnum_BudgetPeriod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_budget_period_f023e2e813969a9a, []int{0, 0}
}

// Message describing Budget period.
type BudgetPeriodEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BudgetPeriodEnum) Reset()         { *m = BudgetPeriodEnum{} }
func (m *BudgetPeriodEnum) String() string { return proto.CompactTextString(m) }
func (*BudgetPeriodEnum) ProtoMessage()    {}
func (*BudgetPeriodEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_budget_period_f023e2e813969a9a, []int{0}
}
func (m *BudgetPeriodEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BudgetPeriodEnum.Unmarshal(m, b)
}
func (m *BudgetPeriodEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BudgetPeriodEnum.Marshal(b, m, deterministic)
}
func (dst *BudgetPeriodEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BudgetPeriodEnum.Merge(dst, src)
}
func (m *BudgetPeriodEnum) XXX_Size() int {
	return xxx_messageInfo_BudgetPeriodEnum.Size(m)
}
func (m *BudgetPeriodEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BudgetPeriodEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BudgetPeriodEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BudgetPeriodEnum)(nil), "google.ads.googleads.v0.enums.BudgetPeriodEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.BudgetPeriodEnum_BudgetPeriod", BudgetPeriodEnum_BudgetPeriod_name, BudgetPeriodEnum_BudgetPeriod_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/budget_period.proto", fileDescriptor_budget_period_f023e2e813969a9a)
}

var fileDescriptor_budget_period_f023e2e813969a9a = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xa4, 0xd2, 0x94, 0xf4, 0xd4, 0x92, 0xf8, 0x82, 0xd4, 0xa2, 0xcc, 0xfc,
	0x14, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x3a, 0xbd, 0xc4, 0x94, 0x62, 0x3d,
	0xb8, 0x16, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x16, 0xa5, 0x0c, 0x2e, 0x01, 0x27, 0xb0, 0xae, 0x00,
	0xb0, 0x26, 0xd7, 0xbc, 0xd2, 0x5c, 0xa5, 0x10, 0x2e, 0x1e, 0x64, 0x31, 0x21, 0x7e, 0x2e, 0xee,
	0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e,
	0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x21, 0x4e, 0x2e, 0x56, 0x17, 0x47,
	0x4f, 0x9f, 0x48, 0x01, 0x26, 0x21, 0x2e, 0x2e, 0x36, 0xe7, 0xd0, 0xe0, 0x10, 0x7f, 0x5f, 0x01,
	0x66, 0x90, 0x26, 0x37, 0xcf, 0x08, 0x57, 0x97, 0x78, 0x88, 0x24, 0x8b, 0xd3, 0x33, 0x46, 0x2e,
	0xc5, 0xe4, 0xfc, 0x5c, 0x3d, 0xbc, 0xee, 0x71, 0x12, 0x44, 0xb6, 0x39, 0x00, 0xe4, 0x83, 0x00,
	0xc6, 0x28, 0x27, 0xa8, 0x9e, 0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0xbd, 0xfc, 0xa2, 0x74, 0xfd,
	0xf4, 0xd4, 0x3c, 0xb0, 0xff, 0x60, 0xc1, 0x50, 0x90, 0x59, 0x8c, 0x23, 0x54, 0xac, 0xc1, 0xe4,
	0x22, 0x26, 0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c, 0xb2, 0xee, 0x10, 0xa3, 0x1c, 0x53, 0x8a, 0xf5,
	0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40, 0x0f, 0xe4, 0xf3, 0xe2, 0x53, 0x30, 0xf9, 0x18, 0xc7, 0x94,
	0xe2, 0x18, 0xb8, 0x7c, 0x4c, 0x98, 0x41, 0x0c, 0x58, 0xfe, 0x15, 0x93, 0x22, 0x44, 0xd0, 0xca,
	0xca, 0x31, 0xa5, 0xd8, 0xca, 0x0a, 0xae, 0xc2, 0xca, 0x2a, 0xcc, 0xc0, 0xca, 0x0a, 0xac, 0x26,
	0x89, 0x0d, 0xec, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x8d, 0x03, 0x9f, 0xad,
	0x01, 0x00, 0x00,
}