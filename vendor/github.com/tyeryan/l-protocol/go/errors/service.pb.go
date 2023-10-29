// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: errors/service.proto

package errors

import (
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

type Caller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	FuncName    string `protobuf:"bytes,2,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	FileName    string `protobuf:"bytes,3,opt,name=FileName,proto3" json:"FileName,omitempty"`
	LineNum     int32  `protobuf:"varint,4,opt,name=lineNum,proto3" json:"lineNum,omitempty"`
}

func (x *Caller) Reset() {
	*x = Caller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Caller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Caller) ProtoMessage() {}

func (x *Caller) ProtoReflect() protoreflect.Message {
	mi := &file_errors_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Caller.ProtoReflect.Descriptor instead.
func (*Caller) Descriptor() ([]byte, []int) {
	return file_errors_service_proto_rawDescGZIP(), []int{0}
}

func (x *Caller) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Caller) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *Caller) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Caller) GetLineNum() int32 {
	if x != nil {
		return x.LineNum
	}
	return 0
}

type ErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stan       string  `protobuf:"bytes,1,opt,name=Stan,proto3" json:"Stan,omitempty"`
	Code       string  `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	GrpcCode   uint32  `protobuf:"varint,3,opt,name=GrpcCode,proto3" json:"GrpcCode,omitempty"`
	Desc       string  `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Error      string  `protobuf:"bytes,5,opt,name=Error,proto3" json:"Error,omitempty"`
	CallerInfo *Caller `protobuf:"bytes,6,opt,name=CallerInfo,proto3" json:"CallerInfo,omitempty"`
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_errors_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetail.ProtoReflect.Descriptor instead.
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return file_errors_service_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorDetail) GetStan() string {
	if x != nil {
		return x.Stan
	}
	return ""
}

func (x *ErrorDetail) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ErrorDetail) GetGrpcCode() uint32 {
	if x != nil {
		return x.GrpcCode
	}
	return 0
}

func (x *ErrorDetail) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ErrorDetail) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ErrorDetail) GetCallerInfo() *Caller {
	if x != nil {
		return x.CallerInfo
	}
	return nil
}

var File_errors_service_proto protoreflect.FileDescriptor

var file_errors_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x7c,
	0x0a, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75,
	0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75,
	0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0xab, 0x01, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x74, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x74, 0x61, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x47, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0a,
	0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x65, 0x72, 0x79, 0x61, 0x6e,
	0x2f, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_service_proto_rawDescOnce sync.Once
	file_errors_service_proto_rawDescData = file_errors_service_proto_rawDesc
)

func file_errors_service_proto_rawDescGZIP() []byte {
	file_errors_service_proto_rawDescOnce.Do(func() {
		file_errors_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_service_proto_rawDescData)
	})
	return file_errors_service_proto_rawDescData
}

var file_errors_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_errors_service_proto_goTypes = []interface{}{
	(*Caller)(nil),      // 0: errors.Caller
	(*ErrorDetail)(nil), // 1: errors.ErrorDetail
}
var file_errors_service_proto_depIdxs = []int32{
	0, // 0: errors.ErrorDetail.CallerInfo:type_name -> errors.Caller
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errors_service_proto_init() }
func file_errors_service_proto_init() {
	if File_errors_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Caller); i {
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
		file_errors_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetail); i {
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
			RawDescriptor: file_errors_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_service_proto_goTypes,
		DependencyIndexes: file_errors_service_proto_depIdxs,
		MessageInfos:      file_errors_service_proto_msgTypes,
	}.Build()
	File_errors_service_proto = out.File
	file_errors_service_proto_rawDesc = nil
	file_errors_service_proto_goTypes = nil
	file_errors_service_proto_depIdxs = nil
}
