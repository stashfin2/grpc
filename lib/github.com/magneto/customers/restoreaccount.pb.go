// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: customers/restoreaccount.proto

package customers

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

type RestoreAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile       string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	DeletedToken string `protobuf:"bytes,2,opt,name=deleted_token,json=deletedToken,proto3" json:"deleted_token,omitempty"`
}

func (x *RestoreAccountRequest) Reset() {
	*x = RestoreAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_restoreaccount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreAccountRequest) ProtoMessage() {}

func (x *RestoreAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_restoreaccount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreAccountRequest.ProtoReflect.Descriptor instead.
func (*RestoreAccountRequest) Descriptor() ([]byte, []int) {
	return file_customers_restoreaccount_proto_rawDescGZIP(), []int{0}
}

func (x *RestoreAccountRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RestoreAccountRequest) GetDeletedToken() string {
	if x != nil {
		return x.DeletedToken
	}
	return ""
}

type RestoreAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwttoken string `protobuf:"bytes,1,opt,name=jwttoken,proto3" json:"jwttoken,omitempty"`
}

func (x *RestoreAccountResponse) Reset() {
	*x = RestoreAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_restoreaccount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreAccountResponse) ProtoMessage() {}

func (x *RestoreAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customers_restoreaccount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreAccountResponse.ProtoReflect.Descriptor instead.
func (*RestoreAccountResponse) Descriptor() ([]byte, []int) {
	return file_customers_restoreaccount_proto_rawDescGZIP(), []int{1}
}

func (x *RestoreAccountResponse) GetJwttoken() string {
	if x != nil {
		return x.Jwttoken
	}
	return ""
}

var File_customers_restoreaccount_proto protoreflect.FileDescriptor

var file_customers_restoreaccount_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x15, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x77,
	0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x77,
	0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_restoreaccount_proto_rawDescOnce sync.Once
	file_customers_restoreaccount_proto_rawDescData = file_customers_restoreaccount_proto_rawDesc
)

func file_customers_restoreaccount_proto_rawDescGZIP() []byte {
	file_customers_restoreaccount_proto_rawDescOnce.Do(func() {
		file_customers_restoreaccount_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_restoreaccount_proto_rawDescData)
	})
	return file_customers_restoreaccount_proto_rawDescData
}

var file_customers_restoreaccount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customers_restoreaccount_proto_goTypes = []interface{}{
	(*RestoreAccountRequest)(nil),  // 0: customers.restoreaccount.restoreAccountRequest
	(*RestoreAccountResponse)(nil), // 1: customers.restoreaccount.restoreAccountResponse
}
var file_customers_restoreaccount_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_restoreaccount_proto_init() }
func file_customers_restoreaccount_proto_init() {
	if File_customers_restoreaccount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_restoreaccount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreAccountRequest); i {
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
		file_customers_restoreaccount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreAccountResponse); i {
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
			RawDescriptor: file_customers_restoreaccount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customers_restoreaccount_proto_goTypes,
		DependencyIndexes: file_customers_restoreaccount_proto_depIdxs,
		MessageInfos:      file_customers_restoreaccount_proto_msgTypes,
	}.Build()
	File_customers_restoreaccount_proto = out.File
	file_customers_restoreaccount_proto_rawDesc = nil
	file_customers_restoreaccount_proto_goTypes = nil
	file_customers_restoreaccount_proto_depIdxs = nil
}
