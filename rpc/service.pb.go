// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: service.proto

package rpc

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScanResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScanIP    string               `protobuf:"bytes,1,opt,name=ScanIP,proto3" json:"ScanIP,omitempty"`
	Host      string               `protobuf:"bytes,2,opt,name=Host,proto3" json:"Host,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Ports     []*PortInfo          `protobuf:"bytes,4,rep,name=Ports,proto3" json:"Ports,omitempty"`
	ScanError string               `protobuf:"bytes,5,opt,name=ScanError,proto3" json:"ScanError,omitempty"`
	Platform  string               `protobuf:"bytes,6,opt,name=Platform,proto3" json:"Platform,omitempty"`
	PingInfo  *PingInfo            `protobuf:"bytes,7,opt,name=PingInfo,proto3" json:"PingInfo,omitempty"`
	AuthToken *AuthToken           `protobuf:"bytes,8,opt,name=AuthToken,proto3" json:"AuthToken,omitempty"`
}

func (x *ScanResults) Reset() {
	*x = ScanResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResults) ProtoMessage() {}

func (x *ScanResults) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResults.ProtoReflect.Descriptor instead.
func (*ScanResults) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *ScanResults) GetScanIP() string {
	if x != nil {
		return x.ScanIP
	}
	return ""
}

func (x *ScanResults) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ScanResults) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ScanResults) GetPorts() []*PortInfo {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *ScanResults) GetScanError() string {
	if x != nil {
		return x.ScanError
	}
	return ""
}

func (x *ScanResults) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ScanResults) GetPingInfo() *PingInfo {
	if x != nil {
		return x.PingInfo
	}
	return nil
}

func (x *ScanResults) GetAuthToken() *AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

type PortInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32 `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
	Open bool  `protobuf:"varint,2,opt,name=Open,proto3" json:"Open,omitempty"`
}

func (x *PortInfo) Reset() {
	*x = PortInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortInfo) ProtoMessage() {}

func (x *PortInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortInfo.ProtoReflect.Descriptor instead.
func (*PortInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *PortInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PortInfo) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *AuthToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transmitted int32  `protobuf:"varint,1,opt,name=Transmitted,proto3" json:"Transmitted,omitempty"`
	Received    int32  `protobuf:"varint,2,opt,name=Received,proto3" json:"Received,omitempty"`
	Loss        int32  `protobuf:"varint,3,opt,name=Loss,proto3" json:"Loss,omitempty"`
	Output      string `protobuf:"bytes,4,opt,name=Output,proto3" json:"Output,omitempty"`
}

func (x *PingInfo) Reset() {
	*x = PingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingInfo) ProtoMessage() {}

func (x *PingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingInfo.ProtoReflect.Descriptor instead.
func (*PingInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *PingInfo) GetTransmitted() int32 {
	if x != nil {
		return x.Transmitted
	}
	return 0
}

func (x *PingInfo) GetReceived() int32 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *PingInfo) GetLoss() int32 {
	if x != nil {
		return x.Loss
	}
	return 0
}

func (x *PingInfo) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x50, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x50, 0x12, 0x12, 0x0a,
	0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x50,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x63, 0x61, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x63, 0x61, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x50, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x22, 0x21, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x74, 0x0a, 0x08, 0x50, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x3c, 0x0a, 0x09, 0x50, 0x6f, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x0a, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(*ScanResults)(nil),         // 0: rpc.ScanResults
	(*PortInfo)(nil),            // 1: rpc.PortInfo
	(*AuthToken)(nil),           // 2: rpc.AuthToken
	(*PingInfo)(nil),            // 3: rpc.PingInfo
	(*Empty)(nil),               // 4: rpc.Empty
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_service_proto_depIdxs = []int32{
	5, // 0: rpc.ScanResults.CreatedAt:type_name -> google.protobuf.Timestamp
	1, // 1: rpc.ScanResults.Ports:type_name -> rpc.PortInfo
	3, // 2: rpc.ScanResults.PingInfo:type_name -> rpc.PingInfo
	2, // 3: rpc.ScanResults.AuthToken:type_name -> rpc.AuthToken
	0, // 4: rpc.PoeStatus.SaveScanResults:input_type -> rpc.ScanResults
	4, // 5: rpc.PoeStatus.SaveScanResults:output_type -> rpc.Empty
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResults); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
