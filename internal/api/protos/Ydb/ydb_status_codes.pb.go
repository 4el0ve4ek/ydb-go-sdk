// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_status_codes.proto

package Ydb

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

// reserved range [400000, 400999]
type StatusIds_StatusCode int32

const (
	StatusIds_STATUS_CODE_UNSPECIFIED StatusIds_StatusCode = 0
	StatusIds_SUCCESS                 StatusIds_StatusCode = 400000
	StatusIds_BAD_REQUEST             StatusIds_StatusCode = 400010
	StatusIds_UNAUTHORIZED            StatusIds_StatusCode = 400020
	StatusIds_INTERNAL_ERROR          StatusIds_StatusCode = 400030
	StatusIds_ABORTED                 StatusIds_StatusCode = 400040
	StatusIds_UNAVAILABLE             StatusIds_StatusCode = 400050
	StatusIds_OVERLOADED              StatusIds_StatusCode = 400060
	StatusIds_SCHEME_ERROR            StatusIds_StatusCode = 400070
	StatusIds_GENERIC_ERROR           StatusIds_StatusCode = 400080
	StatusIds_TIMEOUT                 StatusIds_StatusCode = 400090
	StatusIds_BAD_SESSION             StatusIds_StatusCode = 400100
	StatusIds_PRECONDITION_FAILED     StatusIds_StatusCode = 400120
	StatusIds_ALREADY_EXISTS          StatusIds_StatusCode = 400130
	StatusIds_NOT_FOUND               StatusIds_StatusCode = 400140
	StatusIds_SESSION_EXPIRED         StatusIds_StatusCode = 400150
	StatusIds_CANCELLED               StatusIds_StatusCode = 400160
)

var StatusIds_StatusCode_name = map[int32]string{
	0:      "STATUS_CODE_UNSPECIFIED",
	400000: "SUCCESS",
	400010: "BAD_REQUEST",
	400020: "UNAUTHORIZED",
	400030: "INTERNAL_ERROR",
	400040: "ABORTED",
	400050: "UNAVAILABLE",
	400060: "OVERLOADED",
	400070: "SCHEME_ERROR",
	400080: "GENERIC_ERROR",
	400090: "TIMEOUT",
	400100: "BAD_SESSION",
	400120: "PRECONDITION_FAILED",
	400130: "ALREADY_EXISTS",
	400140: "NOT_FOUND",
	400150: "SESSION_EXPIRED",
	400160: "CANCELLED",
}

var StatusIds_StatusCode_value = map[string]int32{
	"STATUS_CODE_UNSPECIFIED": 0,
	"SUCCESS":                 400000,
	"BAD_REQUEST":             400010,
	"UNAUTHORIZED":            400020,
	"INTERNAL_ERROR":          400030,
	"ABORTED":                 400040,
	"UNAVAILABLE":             400050,
	"OVERLOADED":              400060,
	"SCHEME_ERROR":            400070,
	"GENERIC_ERROR":           400080,
	"TIMEOUT":                 400090,
	"BAD_SESSION":             400100,
	"PRECONDITION_FAILED":     400120,
	"ALREADY_EXISTS":          400130,
	"NOT_FOUND":               400140,
	"SESSION_EXPIRED":         400150,
	"CANCELLED":               400160,
}

func (x StatusIds_StatusCode) String() string {
	return proto.EnumName(StatusIds_StatusCode_name, int32(x))
}

func (StatusIds_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f81e45819472f2bf, []int{0, 0}
}

type StatusIds struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusIds) Reset()         { *m = StatusIds{} }
func (m *StatusIds) String() string { return proto.CompactTextString(m) }
func (*StatusIds) ProtoMessage()    {}
func (*StatusIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_f81e45819472f2bf, []int{0}
}

func (m *StatusIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusIds.Unmarshal(m, b)
}
func (m *StatusIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusIds.Marshal(b, m, deterministic)
}
func (m *StatusIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusIds.Merge(m, src)
}
func (m *StatusIds) XXX_Size() int {
	return xxx_messageInfo_StatusIds.Size(m)
}
func (m *StatusIds) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusIds.DiscardUnknown(m)
}

var xxx_messageInfo_StatusIds proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("Ydb.StatusIds_StatusCode", StatusIds_StatusCode_name, StatusIds_StatusCode_value)
	proto.RegisterType((*StatusIds)(nil), "Ydb.StatusIds")
}

func init() { proto.RegisterFile("ydb_status_codes.proto", fileDescriptor_f81e45819472f2bf) }

var fileDescriptor_f81e45819472f2bf = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd1, 0xbd, 0x8e, 0x13, 0x31,
	0x10, 0x07, 0x70, 0x14, 0x24, 0x50, 0x06, 0x42, 0x1c, 0x87, 0x0f, 0x23, 0xba, 0x3c, 0x40, 0x1a,
	0x9e, 0xc0, 0x6b, 0x4f, 0x88, 0xa5, 0x8d, 0xbd, 0xf8, 0x23, 0x4a, 0x68, 0xac, 0x2c, 0x9b, 0x96,
	0x45, 0xd9, 0x44, 0x22, 0x1d, 0xa2, 0x44, 0x94, 0x88, 0x12, 0x5d, 0x79, 0xf5, 0x49, 0x9b, 0xea,
	0xea, 0xab, 0xaf, 0xbe, 0xe2, 0xaa, 0x2b, 0xef, 0x01, 0xae, 0x3c, 0x25, 0x1b, 0x29, 0xa5, 0x3d,
	0xa3, 0x9f, 0xfe, 0x9a, 0x3f, 0xbc, 0xdd, 0x16, 0x79, 0xac, 0xd6, 0x8b, 0xf5, 0xa6, 0x8a, 0x5f,
	0xcb, 0x62, 0x59, 0x0d, 0xbf, 0xaf, 0xca, 0x75, 0x49, 0x9f, 0xce, 0x8b, 0x7c, 0x70, 0xdf, 0x82,
	0xb6, 0x3b, 0xcc, 0x54, 0x51, 0x0d, 0x6e, 0x5b, 0x00, 0xcd, 0x4b, 0x94, 0xc5, 0x92, 0x7e, 0x80,
	0x77, 0xce, 0x73, 0x1f, 0x5c, 0x14, 0x46, 0x62, 0x0c, 0xda, 0x65, 0x28, 0xd4, 0x48, 0xa1, 0x24,
	0x4f, 0x68, 0x07, 0x9e, 0xbb, 0x20, 0x04, 0x3a, 0x47, 0x7e, 0xd6, 0x8c, 0xf6, 0xe0, 0x45, 0xc2,
	0x65, 0xb4, 0xf8, 0x39, 0xa0, 0xf3, 0xe4, 0x77, 0xcd, 0x28, 0x85, 0x97, 0x41, 0xf3, 0xe0, 0xc7,
	0xc6, 0xaa, 0x2f, 0x28, 0xc9, 0xdf, 0x9a, 0xd1, 0xd7, 0xf0, 0x4a, 0x69, 0x8f, 0x56, 0xf3, 0x34,
	0xa2, 0xb5, 0xc6, 0x92, 0xff, 0x35, 0xdb, 0x5b, 0x3c, 0x31, 0xd6, 0xa3, 0x24, 0xe7, 0x8d, 0x15,
	0x34, 0x9f, 0x72, 0x95, 0xf2, 0x24, 0x45, 0x72, 0x51, 0x33, 0x4a, 0x00, 0xcc, 0x14, 0x6d, 0x6a,
	0xb8, 0x44, 0x49, 0x2e, 0x1b, 0xdd, 0x89, 0x31, 0x4e, 0xf0, 0xe8, 0x5c, 0xd5, 0x8c, 0xf6, 0xa1,
	0xf3, 0x09, 0x35, 0x5a, 0x25, 0x8e, 0x9f, 0xd7, 0x0d, 0xee, 0xd5, 0x04, 0x4d, 0xf0, 0xe4, 0xe6,
	0x14, 0xd4, 0xa1, 0x73, 0xca, 0x68, 0x72, 0x57, 0x33, 0xfa, 0x1e, 0xfa, 0x99, 0x45, 0x61, 0xb4,
	0x54, 0x5e, 0x19, 0x1d, 0x47, 0x5c, 0xa5, 0x28, 0xc9, 0x43, 0x93, 0x97, 0xa7, 0x16, 0xb9, 0x9c,
	0x47, 0x9c, 0x29, 0xe7, 0x1d, 0xf9, 0xb5, 0x63, 0xb4, 0x0b, 0x6d, 0x6d, 0x7c, 0x1c, 0x99, 0xa0,
	0x25, 0xf9, 0xb3, 0x63, 0xf4, 0x0d, 0x74, 0x8f, 0x60, 0xc4, 0x59, 0xa6, 0x2c, 0x4a, 0xf2, 0xaf,
	0xd9, 0x13, 0x5c, 0x0b, 0x4c, 0xf7, 0xdc, 0xd9, 0x8e, 0x25, 0x03, 0xe8, 0xac, 0x36, 0xc3, 0xed,
	0xe2, 0x5b, 0xb1, 0xfc, 0x31, 0xdc, 0x16, 0x79, 0xd2, 0x3b, 0x9d, 0xbb, 0xca, 0xf6, 0xb5, 0x54,
	0xf9, 0xb3, 0x43, 0x3d, 0x1f, 0x1f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xac, 0x13, 0x5a, 0xb8,
	0x01, 0x00, 0x00,
}

const ()
