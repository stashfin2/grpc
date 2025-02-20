// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: loans/getcbloandetails.proto

package loans

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

type GetCBloanDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetCBloanDetailsRequest) Reset() {
	*x = GetCBloanDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getcbloandetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCBloanDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCBloanDetailsRequest) ProtoMessage() {}

func (x *GetCBloanDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getcbloandetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCBloanDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetCBloanDetailsRequest) Descriptor() ([]byte, []int) {
	return file_loans_getcbloandetails_proto_rawDescGZIP(), []int{0}
}

func (x *GetCBloanDetailsRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetCBloanDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisbursedAmount  int32 `protobuf:"varint,1,opt,name=disbursed_amount,json=disbursedAmount,proto3" json:"disbursed_amount,omitempty"`
	RemainingEmi     int32 `protobuf:"varint,2,opt,name=remaining_emi,json=remainingEmi,proto3" json:"remaining_emi,omitempty"`
	BalanceEmiAmount int32 `protobuf:"varint,3,opt,name=balance_emi_amount,json=balanceEmiAmount,proto3" json:"balance_emi_amount,omitempty"`
}

func (x *GetCBloanDetailsResponse) Reset() {
	*x = GetCBloanDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getcbloandetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCBloanDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCBloanDetailsResponse) ProtoMessage() {}

func (x *GetCBloanDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getcbloandetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCBloanDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetCBloanDetailsResponse) Descriptor() ([]byte, []int) {
	return file_loans_getcbloandetails_proto_rawDescGZIP(), []int{1}
}

func (x *GetCBloanDetailsResponse) GetDisbursedAmount() int32 {
	if x != nil {
		return x.DisbursedAmount
	}
	return 0
}

func (x *GetCBloanDetailsResponse) GetRemainingEmi() int32 {
	if x != nil {
		return x.RemainingEmi
	}
	return 0
}

func (x *GetCBloanDetailsResponse) GetBalanceEmiAmount() int32 {
	if x != nil {
		return x.BalanceEmiAmount
	}
	return 0
}

var File_loans_getcbloandetails_proto protoreflect.FileDescriptor

var file_loans_getcbloandetails_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x62, 0x6c, 0x6f, 0x61,
	0x6e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x62, 0x6c, 0x6f, 0x61, 0x6e, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x3a, 0x0a, 0x17, 0x67, 0x65, 0x74, 0x43, 0x42, 0x6c,
	0x6f, 0x61, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x43, 0x42, 0x6c, 0x6f, 0x61, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x62, 0x75,
	0x72, 0x73, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6d, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x69, 0x12,
	0x2c, 0x0a, 0x12, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6d, 0x69, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x45, 0x6d, 0x69, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loans_getcbloandetails_proto_rawDescOnce sync.Once
	file_loans_getcbloandetails_proto_rawDescData = file_loans_getcbloandetails_proto_rawDesc
)

func file_loans_getcbloandetails_proto_rawDescGZIP() []byte {
	file_loans_getcbloandetails_proto_rawDescOnce.Do(func() {
		file_loans_getcbloandetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_getcbloandetails_proto_rawDescData)
	})
	return file_loans_getcbloandetails_proto_rawDescData
}

var file_loans_getcbloandetails_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loans_getcbloandetails_proto_goTypes = []interface{}{
	(*GetCBloanDetailsRequest)(nil),  // 0: loans.getcbloandetails.getCBloanDetailsRequest
	(*GetCBloanDetailsResponse)(nil), // 1: loans.getcbloandetails.getCBloanDetailsResponse
}
var file_loans_getcbloandetails_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loans_getcbloandetails_proto_init() }
func file_loans_getcbloandetails_proto_init() {
	if File_loans_getcbloandetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_getcbloandetails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCBloanDetailsRequest); i {
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
		file_loans_getcbloandetails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCBloanDetailsResponse); i {
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
			RawDescriptor: file_loans_getcbloandetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_getcbloandetails_proto_goTypes,
		DependencyIndexes: file_loans_getcbloandetails_proto_depIdxs,
		MessageInfos:      file_loans_getcbloandetails_proto_msgTypes,
	}.Build()
	File_loans_getcbloandetails_proto = out.File
	file_loans_getcbloandetails_proto_rawDesc = nil
	file_loans_getcbloandetails_proto_goTypes = nil
	file_loans_getcbloandetails_proto_depIdxs = nil
}
