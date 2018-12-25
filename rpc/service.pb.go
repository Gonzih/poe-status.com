// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ScanResults struct {
	ScanIP               string               `protobuf:"bytes,1,opt,name=ScanIP,proto3" json:"ScanIP,omitempty"`
	Host                 string               `protobuf:"bytes,2,opt,name=Host,proto3" json:"Host,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Ports                []*PortInfo          `protobuf:"bytes,4,rep,name=Ports,proto3" json:"Ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ScanResults) Reset()         { *m = ScanResults{} }
func (m *ScanResults) String() string { return proto.CompactTextString(m) }
func (*ScanResults) ProtoMessage()    {}
func (*ScanResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1e589d75f18b77b9, []int{0}
}
func (m *ScanResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanResults.Unmarshal(m, b)
}
func (m *ScanResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanResults.Marshal(b, m, deterministic)
}
func (dst *ScanResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanResults.Merge(dst, src)
}
func (m *ScanResults) XXX_Size() int {
	return xxx_messageInfo_ScanResults.Size(m)
}
func (m *ScanResults) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanResults.DiscardUnknown(m)
}

var xxx_messageInfo_ScanResults proto.InternalMessageInfo

func (m *ScanResults) GetScanIP() string {
	if m != nil {
		return m.ScanIP
	}
	return ""
}

func (m *ScanResults) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ScanResults) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ScanResults) GetPorts() []*PortInfo {
	if m != nil {
		return m.Ports
	}
	return nil
}

type PortInfo struct {
	Port                 int32    `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
	Open                 bool     `protobuf:"varint,2,opt,name=Open,proto3" json:"Open,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortInfo) Reset()         { *m = PortInfo{} }
func (m *PortInfo) String() string { return proto.CompactTextString(m) }
func (*PortInfo) ProtoMessage()    {}
func (*PortInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1e589d75f18b77b9, []int{1}
}
func (m *PortInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortInfo.Unmarshal(m, b)
}
func (m *PortInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortInfo.Marshal(b, m, deterministic)
}
func (dst *PortInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortInfo.Merge(dst, src)
}
func (m *PortInfo) XXX_Size() int {
	return xxx_messageInfo_PortInfo.Size(m)
}
func (m *PortInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PortInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PortInfo proto.InternalMessageInfo

func (m *PortInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PortInfo) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1e589d75f18b77b9, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScanResults)(nil), "rpc.ScanResults")
	proto.RegisterType((*PortInfo)(nil), "rpc.PortInfo")
	proto.RegisterType((*Empty)(nil), "rpc.Empty")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_1e589d75f18b77b9) }

var fileDescriptor_service_1e589d75f18b77b9 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xd2, 0x94, 0xe6, 0xa2, 0x0a, 0xe4, 0x01, 0x45, 0x59, 0x88, 0xc2, 0x92, 0xc9,
	0x91, 0xc2, 0xc2, 0xc0, 0x02, 0x08, 0x89, 0x4e, 0x44, 0x0e, 0x13, 0x9b, 0x1b, 0xdc, 0xaa, 0x52,
	0x13, 0x5b, 0xf6, 0xa5, 0x12, 0x6f, 0xc2, 0xe3, 0x22, 0x5f, 0x88, 0xe8, 0xf6, 0xdf, 0xe7, 0xb3,
	0xbe, 0xbb, 0x83, 0xb5, 0x53, 0xf6, 0x74, 0xe8, 0x14, 0x37, 0x56, 0xa3, 0x66, 0xa1, 0x35, 0x5d,
	0x76, 0xbb, 0xd7, 0x7a, 0x7f, 0x54, 0x15, 0xa1, 0xed, 0xb8, 0xab, 0xf0, 0xd0, 0x2b, 0x87, 0xb2,
	0x37, 0x53, 0x57, 0xf1, 0x13, 0x40, 0xd2, 0x76, 0x72, 0x10, 0xca, 0x8d, 0x47, 0x74, 0xec, 0x06,
	0x96, 0xbe, 0xdc, 0x34, 0x69, 0x90, 0x07, 0x65, 0x2c, 0xfe, 0x2a, 0xc6, 0x60, 0xf1, 0xa6, 0x1d,
	0xa6, 0x17, 0x44, 0x29, 0xb3, 0x07, 0x88, 0x5f, 0xac, 0x92, 0xa8, 0xbe, 0x9e, 0x30, 0x0d, 0xf3,
	0xa0, 0x4c, 0xea, 0x8c, 0x4f, 0x42, 0x3e, 0x0b, 0xf9, 0xc7, 0x2c, 0x14, 0xff, 0xcd, 0xec, 0x0e,
	0xa2, 0x46, 0x5b, 0x74, 0xe9, 0x22, 0x0f, 0xcb, 0xa4, 0x5e, 0x73, 0x6b, 0x3a, 0xee, 0xc9, 0x66,
	0xd8, 0x69, 0x31, 0xbd, 0x15, 0x35, 0xac, 0x66, 0xe4, 0xf5, 0x3e, 0xd3, 0x50, 0x91, 0xa0, 0xec,
	0xd9, 0xbb, 0x51, 0x03, 0x8d, 0xb4, 0x12, 0x94, 0x8b, 0x4b, 0x88, 0x5e, 0x7b, 0x83, 0xdf, 0xf5,
	0x23, 0xc4, 0x8d, 0x56, 0x2d, 0x4a, 0x1c, 0x1d, 0xab, 0xe0, 0xaa, 0x95, 0x27, 0x75, 0xbe, 0xe7,
	0x35, 0x29, 0xcf, 0x48, 0x06, 0x44, 0xe8, 0xf7, 0x73, 0xf4, 0xe9, 0xaf, 0xb7, 0x5d, 0xd2, 0x16,
	0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x51, 0x03, 0xe9, 0x5c, 0x5a, 0x01, 0x00, 0x00,
}