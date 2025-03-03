// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: decisionengine/locdisbursalapproval.proto

package decisionengine

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

type ValidateDisbursalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId         int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	LoanId             int32 `protobuf:"varint,2,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	StashfinEverDpd    int32 `protobuf:"varint,3,opt,name=stashfin_ever_dpd,json=stashfinEverDpd,proto3" json:"stashfin_ever_dpd,omitempty"`
	StashfinCurrentDpd int32 `protobuf:"varint,4,opt,name=stashfin_current_dpd,json=stashfinCurrentDpd,proto3" json:"stashfin_current_dpd,omitempty"`
}

func (x *ValidateDisbursalRequest) Reset() {
	*x = ValidateDisbursalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_decisionengine_locdisbursalapproval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateDisbursalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateDisbursalRequest) ProtoMessage() {}

func (x *ValidateDisbursalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_decisionengine_locdisbursalapproval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateDisbursalRequest.ProtoReflect.Descriptor instead.
func (*ValidateDisbursalRequest) Descriptor() ([]byte, []int) {
	return file_decisionengine_locdisbursalapproval_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateDisbursalRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ValidateDisbursalRequest) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

func (x *ValidateDisbursalRequest) GetStashfinEverDpd() int32 {
	if x != nil {
		return x.StashfinEverDpd
	}
	return 0
}

func (x *ValidateDisbursalRequest) GetStashfinCurrentDpd() int32 {
	if x != nil {
		return x.StashfinCurrentDpd
	}
	return 0
}

type ValidateDisbursalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message   *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	ErrorCode *string `protobuf:"bytes,3,opt,name=error_code,json=errorCode,proto3,oneof" json:"error_code,omitempty"`
}

func (x *ValidateDisbursalResponse) Reset() {
	*x = ValidateDisbursalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_decisionengine_locdisbursalapproval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateDisbursalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateDisbursalResponse) ProtoMessage() {}

func (x *ValidateDisbursalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_decisionengine_locdisbursalapproval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateDisbursalResponse.ProtoReflect.Descriptor instead.
func (*ValidateDisbursalResponse) Descriptor() ([]byte, []int) {
	return file_decisionengine_locdisbursalapproval_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateDisbursalResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ValidateDisbursalResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *ValidateDisbursalResponse) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

var File_decisionengine_locdisbursalapproval_proto protoreflect.FileDescriptor

var file_decisionengine_locdisbursalapproval_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x6c, 0x6f, 0x63, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x44, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c,
	0x22, 0xb2, 0x01, 0x0a, 0x18, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73,
	0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x73, 0x68,
	0x66, 0x69, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x70, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x73, 0x68, 0x66, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x72,
	0x44, 0x70, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x73, 0x68, 0x66, 0x69, 0x6e, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x70, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x73, 0x74, 0x61, 0x73, 0x68, 0x66, 0x69, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x44, 0x70, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_decisionengine_locdisbursalapproval_proto_rawDescOnce sync.Once
	file_decisionengine_locdisbursalapproval_proto_rawDescData = file_decisionengine_locdisbursalapproval_proto_rawDesc
)

func file_decisionengine_locdisbursalapproval_proto_rawDescGZIP() []byte {
	file_decisionengine_locdisbursalapproval_proto_rawDescOnce.Do(func() {
		file_decisionengine_locdisbursalapproval_proto_rawDescData = protoimpl.X.CompressGZIP(file_decisionengine_locdisbursalapproval_proto_rawDescData)
	})
	return file_decisionengine_locdisbursalapproval_proto_rawDescData
}

var file_decisionengine_locdisbursalapproval_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_decisionengine_locdisbursalapproval_proto_goTypes = []interface{}{
	(*ValidateDisbursalRequest)(nil),  // 0: decisionengine.validateLocDisbursal.validateDisbursalRequest
	(*ValidateDisbursalResponse)(nil), // 1: decisionengine.validateLocDisbursal.validateDisbursalResponse
}
var file_decisionengine_locdisbursalapproval_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_decisionengine_locdisbursalapproval_proto_init() }
func file_decisionengine_locdisbursalapproval_proto_init() {
	if File_decisionengine_locdisbursalapproval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_decisionengine_locdisbursalapproval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateDisbursalRequest); i {
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
		file_decisionengine_locdisbursalapproval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateDisbursalResponse); i {
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
	file_decisionengine_locdisbursalapproval_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_decisionengine_locdisbursalapproval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_decisionengine_locdisbursalapproval_proto_goTypes,
		DependencyIndexes: file_decisionengine_locdisbursalapproval_proto_depIdxs,
		MessageInfos:      file_decisionengine_locdisbursalapproval_proto_msgTypes,
	}.Build()
	File_decisionengine_locdisbursalapproval_proto = out.File
	file_decisionengine_locdisbursalapproval_proto_rawDesc = nil
	file_decisionengine_locdisbursalapproval_proto_goTypes = nil
	file_decisionengine_locdisbursalapproval_proto_depIdxs = nil
}
