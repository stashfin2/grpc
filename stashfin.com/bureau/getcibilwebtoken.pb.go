// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: bureau/getcibilwebtoken.proto

package bureau

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

type WebTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *WebTokenRequest) Reset() {
	*x = WebTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getcibilwebtoken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebTokenRequest) ProtoMessage() {}

func (x *WebTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getcibilwebtoken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebTokenRequest.ProtoReflect.Descriptor instead.
func (*WebTokenRequest) Descriptor() ([]byte, []int) {
	return file_bureau_getcibilwebtoken_proto_rawDescGZIP(), []int{0}
}

func (x *WebTokenRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type WebTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   *Data  `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *WebTokenResponse) Reset() {
	*x = WebTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getcibilwebtoken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebTokenResponse) ProtoMessage() {}

func (x *WebTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getcibilwebtoken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebTokenResponse.ProtoReflect.Descriptor instead.
func (*WebTokenResponse) Descriptor() ([]byte, []int) {
	return file_bureau_getcibilwebtoken_proto_rawDescGZIP(), []int{1}
}

func (x *WebTokenResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WebTokenResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ReportUrl string `protobuf:"bytes,2,opt,name=report_url,json=reportUrl,proto3" json:"report_url,omitempty"`
	WebToken  string `protobuf:"bytes,3,opt,name=web_token,json=webToken,proto3" json:"web_token,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getcibilwebtoken_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getcibilwebtoken_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_bureau_getcibilwebtoken_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Data) GetReportUrl() string {
	if x != nil {
		return x.ReportUrl
	}
	return ""
}

func (x *Data) GetWebToken() string {
	if x != nil {
		return x.WebToken
	}
	return ""
}

var File_bureau_getcibilwebtoken_proto protoreflect.FileDescriptor

var file_bureau_getcibilwebtoken_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x69, 0x62, 0x69,
	0x6c, 0x77, 0x65, 0x62, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x69, 0x62, 0x69, 0x6c,
	0x77, 0x65, 0x62, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x0f, 0x77, 0x65, 0x62, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x10,
	0x77, 0x65, 0x62, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e,
	0x67, 0x65, 0x74, 0x63, 0x69, 0x62, 0x69, 0x6c, 0x77, 0x65, 0x62, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65,
	0x62, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x65, 0x62, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bureau_getcibilwebtoken_proto_rawDescOnce sync.Once
	file_bureau_getcibilwebtoken_proto_rawDescData = file_bureau_getcibilwebtoken_proto_rawDesc
)

func file_bureau_getcibilwebtoken_proto_rawDescGZIP() []byte {
	file_bureau_getcibilwebtoken_proto_rawDescOnce.Do(func() {
		file_bureau_getcibilwebtoken_proto_rawDescData = protoimpl.X.CompressGZIP(file_bureau_getcibilwebtoken_proto_rawDescData)
	})
	return file_bureau_getcibilwebtoken_proto_rawDescData
}

var file_bureau_getcibilwebtoken_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bureau_getcibilwebtoken_proto_goTypes = []interface{}{
	(*WebTokenRequest)(nil),  // 0: bureau.getcibilwebtoken.webTokenRequest
	(*WebTokenResponse)(nil), // 1: bureau.getcibilwebtoken.webTokenResponse
	(*Data)(nil),             // 2: bureau.getcibilwebtoken.Data
}
var file_bureau_getcibilwebtoken_proto_depIdxs = []int32{
	2, // 0: bureau.getcibilwebtoken.webTokenResponse.data:type_name -> bureau.getcibilwebtoken.Data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bureau_getcibilwebtoken_proto_init() }
func file_bureau_getcibilwebtoken_proto_init() {
	if File_bureau_getcibilwebtoken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bureau_getcibilwebtoken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebTokenRequest); i {
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
		file_bureau_getcibilwebtoken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebTokenResponse); i {
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
		file_bureau_getcibilwebtoken_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
	file_bureau_getcibilwebtoken_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bureau_getcibilwebtoken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bureau_getcibilwebtoken_proto_goTypes,
		DependencyIndexes: file_bureau_getcibilwebtoken_proto_depIdxs,
		MessageInfos:      file_bureau_getcibilwebtoken_proto_msgTypes,
	}.Build()
	File_bureau_getcibilwebtoken_proto = out.File
	file_bureau_getcibilwebtoken_proto_rawDesc = nil
	file_bureau_getcibilwebtoken_proto_goTypes = nil
	file_bureau_getcibilwebtoken_proto_depIdxs = nil
}
