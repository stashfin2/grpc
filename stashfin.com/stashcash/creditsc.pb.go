// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: stashcash/creditsc.proto

package stashcash

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

type CreditScRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CampaignId    int64                  `protobuf:"varint,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Amount        float32                `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	ScType        string                 `protobuf:"bytes,4,opt,name=sc_type,json=scType,proto3" json:"sc_type,omitempty"`
	Expiry        *string                `protobuf:"bytes,5,opt,name=expiry,proto3,oneof" json:"expiry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreditScRequest) Reset() {
	*x = CreditScRequest{}
	mi := &file_stashcash_creditsc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreditScRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditScRequest) ProtoMessage() {}

func (x *CreditScRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stashcash_creditsc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditScRequest.ProtoReflect.Descriptor instead.
func (*CreditScRequest) Descriptor() ([]byte, []int) {
	return file_stashcash_creditsc_proto_rawDescGZIP(), []int{0}
}

func (x *CreditScRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreditScRequest) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *CreditScRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreditScRequest) GetScType() string {
	if x != nil {
		return x.ScType
	}
	return ""
}

func (x *CreditScRequest) GetExpiry() string {
	if x != nil && x.Expiry != nil {
		return *x.Expiry
	}
	return ""
}

type CreditScResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data          *CreditScResponse_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreditScResponse) Reset() {
	*x = CreditScResponse{}
	mi := &file_stashcash_creditsc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreditScResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditScResponse) ProtoMessage() {}

func (x *CreditScResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stashcash_creditsc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditScResponse.ProtoReflect.Descriptor instead.
func (*CreditScResponse) Descriptor() ([]byte, []int) {
	return file_stashcash_creditsc_proto_rawDescGZIP(), []int{1}
}

func (x *CreditScResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreditScResponse) GetData() *CreditScResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreditScResponse_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Locked        float32                `protobuf:"fixed32,1,opt,name=locked,proto3" json:"locked,omitempty"`
	Unlocked      float32                `protobuf:"fixed32,2,opt,name=unlocked,proto3" json:"unlocked,omitempty"`
	Balance       float32                `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	TxnId         int64                  `protobuf:"varint,4,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreditScResponse_Data) Reset() {
	*x = CreditScResponse_Data{}
	mi := &file_stashcash_creditsc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreditScResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditScResponse_Data) ProtoMessage() {}

func (x *CreditScResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_stashcash_creditsc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditScResponse_Data.ProtoReflect.Descriptor instead.
func (*CreditScResponse_Data) Descriptor() ([]byte, []int) {
	return file_stashcash_creditsc_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreditScResponse_Data) GetLocked() float32 {
	if x != nil {
		return x.Locked
	}
	return 0
}

func (x *CreditScResponse_Data) GetUnlocked() float32 {
	if x != nil {
		return x.Unlocked
	}
	return 0
}

func (x *CreditScResponse_Data) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *CreditScResponse_Data) GetTxnId() int64 {
	if x != nil {
		return x.TxnId
	}
	return 0
}

var File_stashcash_creditsc_proto protoreflect.FileDescriptor

var file_stashcash_creditsc_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x73, 0x74, 0x61, 0x73, 0x68, 0x63, 0x61, 0x73, 0x68, 0x2f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x74, 0x61, 0x73,
	0x68, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x63, 0x22, 0xac,
	0x01, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0xd6, 0x01,
	0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x61, 0x73, 0x68,
	0x63, 0x61, 0x73, 0x68, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x63, 0x2e, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x53, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x6b, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x78, 0x6e, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_stashcash_creditsc_proto_rawDescOnce sync.Once
	file_stashcash_creditsc_proto_rawDescData []byte
)

func file_stashcash_creditsc_proto_rawDescGZIP() []byte {
	file_stashcash_creditsc_proto_rawDescOnce.Do(func() {
		file_stashcash_creditsc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stashcash_creditsc_proto_rawDesc), len(file_stashcash_creditsc_proto_rawDesc)))
	})
	return file_stashcash_creditsc_proto_rawDescData
}

var file_stashcash_creditsc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_stashcash_creditsc_proto_goTypes = []any{
	(*CreditScRequest)(nil),       // 0: stashcash.creditsc.creditScRequest
	(*CreditScResponse)(nil),      // 1: stashcash.creditsc.creditScResponse
	(*CreditScResponse_Data)(nil), // 2: stashcash.creditsc.creditScResponse.Data
}
var file_stashcash_creditsc_proto_depIdxs = []int32{
	2, // 0: stashcash.creditsc.creditScResponse.data:type_name -> stashcash.creditsc.creditScResponse.Data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stashcash_creditsc_proto_init() }
func file_stashcash_creditsc_proto_init() {
	if File_stashcash_creditsc_proto != nil {
		return
	}
	file_stashcash_creditsc_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stashcash_creditsc_proto_rawDesc), len(file_stashcash_creditsc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stashcash_creditsc_proto_goTypes,
		DependencyIndexes: file_stashcash_creditsc_proto_depIdxs,
		MessageInfos:      file_stashcash_creditsc_proto_msgTypes,
	}.Build()
	File_stashcash_creditsc_proto = out.File
	file_stashcash_creditsc_proto_goTypes = nil
	file_stashcash_creditsc_proto_depIdxs = nil
}
