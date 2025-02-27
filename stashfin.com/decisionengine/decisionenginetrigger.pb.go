// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: decisionengine/decisionenginetrigger.proto

package decisionengine

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

type DecisionEngineTriggerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int32                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	BannerCode    *string                `protobuf:"bytes,2,opt,name=banner_code,json=bannerCode,proto3,oneof" json:"banner_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecisionEngineTriggerRequest) Reset() {
	*x = DecisionEngineTriggerRequest{}
	mi := &file_decisionengine_decisionenginetrigger_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecisionEngineTriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionEngineTriggerRequest) ProtoMessage() {}

func (x *DecisionEngineTriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_decisionengine_decisionenginetrigger_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionEngineTriggerRequest.ProtoReflect.Descriptor instead.
func (*DecisionEngineTriggerRequest) Descriptor() ([]byte, []int) {
	return file_decisionengine_decisionenginetrigger_proto_rawDescGZIP(), []int{0}
}

func (x *DecisionEngineTriggerRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *DecisionEngineTriggerRequest) GetBannerCode() string {
	if x != nil && x.BannerCode != nil {
		return *x.BannerCode
	}
	return ""
}

type DecisionEngineTriggerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	ErrorCode     *string                `protobuf:"bytes,3,opt,name=error_code,json=errorCode,proto3,oneof" json:"error_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecisionEngineTriggerResponse) Reset() {
	*x = DecisionEngineTriggerResponse{}
	mi := &file_decisionengine_decisionenginetrigger_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecisionEngineTriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionEngineTriggerResponse) ProtoMessage() {}

func (x *DecisionEngineTriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_decisionengine_decisionenginetrigger_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionEngineTriggerResponse.ProtoReflect.Descriptor instead.
func (*DecisionEngineTriggerResponse) Descriptor() ([]byte, []int) {
	return file_decisionengine_decisionenginetrigger_proto_rawDescGZIP(), []int{1}
}

func (x *DecisionEngineTriggerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DecisionEngineTriggerResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *DecisionEngineTriggerResponse) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

var File_decisionengine_decisionenginetrigger_proto protoreflect.FileDescriptor

var file_decisionengine_decisionenginetrigger_proto_rawDesc = string([]byte{
	0x0a, 0x2a, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x22, 0x75, 0x0a, 0x1c, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x1d, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_decisionengine_decisionenginetrigger_proto_rawDescOnce sync.Once
	file_decisionengine_decisionenginetrigger_proto_rawDescData []byte
)

func file_decisionengine_decisionenginetrigger_proto_rawDescGZIP() []byte {
	file_decisionengine_decisionenginetrigger_proto_rawDescOnce.Do(func() {
		file_decisionengine_decisionenginetrigger_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_decisionengine_decisionenginetrigger_proto_rawDesc), len(file_decisionengine_decisionenginetrigger_proto_rawDesc)))
	})
	return file_decisionengine_decisionenginetrigger_proto_rawDescData
}

var file_decisionengine_decisionenginetrigger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_decisionengine_decisionenginetrigger_proto_goTypes = []any{
	(*DecisionEngineTriggerRequest)(nil),  // 0: decisionengine.decisionEngineTrigger.decisionEngineTriggerRequest
	(*DecisionEngineTriggerResponse)(nil), // 1: decisionengine.decisionEngineTrigger.decisionEngineTriggerResponse
}
var file_decisionengine_decisionenginetrigger_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_decisionengine_decisionenginetrigger_proto_init() }
func file_decisionengine_decisionenginetrigger_proto_init() {
	if File_decisionengine_decisionenginetrigger_proto != nil {
		return
	}
	file_decisionengine_decisionenginetrigger_proto_msgTypes[0].OneofWrappers = []any{}
	file_decisionengine_decisionenginetrigger_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_decisionengine_decisionenginetrigger_proto_rawDesc), len(file_decisionengine_decisionenginetrigger_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_decisionengine_decisionenginetrigger_proto_goTypes,
		DependencyIndexes: file_decisionengine_decisionenginetrigger_proto_depIdxs,
		MessageInfos:      file_decisionengine_decisionenginetrigger_proto_msgTypes,
	}.Build()
	File_decisionengine_decisionenginetrigger_proto = out.File
	file_decisionengine_decisionenginetrigger_proto_goTypes = nil
	file_decisionengine_decisionenginetrigger_proto_depIdxs = nil
}
