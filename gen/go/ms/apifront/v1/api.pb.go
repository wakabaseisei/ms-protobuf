// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: ms/apifront/v1/api.proto

package apifrontv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GreetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	mi := &file_ms_apifront_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ms_apifront_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_ms_apifront_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greeting      string                 `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	mi := &file_ms_apifront_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ms_apifront_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_ms_apifront_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

var File_ms_apifront_v1_api_proto protoreflect.FileDescriptor

var file_ms_apifront_v1_api_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b,
	0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x32, 0x56, 0x0a, 0x0c, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x61, 0x6b, 0x61, 0x62, 0x61, 0x73, 0x65, 0x69, 0x73, 0x65, 0x69, 0x2f, 0x6d,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x70, 0x69, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_ms_apifront_v1_api_proto_rawDescOnce sync.Once
	file_ms_apifront_v1_api_proto_rawDescData []byte
)

func file_ms_apifront_v1_api_proto_rawDescGZIP() []byte {
	file_ms_apifront_v1_api_proto_rawDescOnce.Do(func() {
		file_ms_apifront_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ms_apifront_v1_api_proto_rawDesc), len(file_ms_apifront_v1_api_proto_rawDesc)))
	})
	return file_ms_apifront_v1_api_proto_rawDescData
}

var file_ms_apifront_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ms_apifront_v1_api_proto_goTypes = []any{
	(*GreetRequest)(nil),  // 0: ms.apifront.v1.GreetRequest
	(*GreetResponse)(nil), // 1: ms.apifront.v1.GreetResponse
}
var file_ms_apifront_v1_api_proto_depIdxs = []int32{
	0, // 0: ms.apifront.v1.GreetService.Greet:input_type -> ms.apifront.v1.GreetRequest
	1, // 1: ms.apifront.v1.GreetService.Greet:output_type -> ms.apifront.v1.GreetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ms_apifront_v1_api_proto_init() }
func file_ms_apifront_v1_api_proto_init() {
	if File_ms_apifront_v1_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ms_apifront_v1_api_proto_rawDesc), len(file_ms_apifront_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ms_apifront_v1_api_proto_goTypes,
		DependencyIndexes: file_ms_apifront_v1_api_proto_depIdxs,
		MessageInfos:      file_ms_apifront_v1_api_proto_msgTypes,
	}.Build()
	File_ms_apifront_v1_api_proto = out.File
	file_ms_apifront_v1_api_proto_goTypes = nil
	file_ms_apifront_v1_api_proto_depIdxs = nil
}
