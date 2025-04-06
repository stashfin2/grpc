// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: stashcash/getschistory.proto

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

type GetschistoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Page          string                 `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit         string                 `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	FilterBy      *string                `protobuf:"bytes,4,opt,name=filter_by,json=filterBy,proto3,oneof" json:"filter_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetschistoryRequest) Reset() {
	*x = GetschistoryRequest{}
	mi := &file_stashcash_getschistory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetschistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetschistoryRequest) ProtoMessage() {}

func (x *GetschistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stashcash_getschistory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetschistoryRequest.ProtoReflect.Descriptor instead.
func (*GetschistoryRequest) Descriptor() ([]byte, []int) {
	return file_stashcash_getschistory_proto_rawDescGZIP(), []int{0}
}

func (x *GetschistoryRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GetschistoryRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GetschistoryRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *GetschistoryRequest) GetFilterBy() string {
	if x != nil && x.FilterBy != nil {
		return *x.FilterBy
	}
	return ""
}

type GetschistoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data          *ScHistoryData         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetschistoryResponse) Reset() {
	*x = GetschistoryResponse{}
	mi := &file_stashcash_getschistory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetschistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetschistoryResponse) ProtoMessage() {}

func (x *GetschistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stashcash_getschistory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetschistoryResponse.ProtoReflect.Descriptor instead.
func (*GetschistoryResponse) Descriptor() ([]byte, []int) {
	return file_stashcash_getschistory_proto_rawDescGZIP(), []int{1}
}

func (x *GetschistoryResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetschistoryResponse) GetData() *ScHistoryData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ScHistoryData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	History       []*SCtransaction       `protobuf:"bytes,1,rep,name=history,proto3" json:"history,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScHistoryData) Reset() {
	*x = ScHistoryData{}
	mi := &file_stashcash_getschistory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScHistoryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScHistoryData) ProtoMessage() {}

func (x *ScHistoryData) ProtoReflect() protoreflect.Message {
	mi := &file_stashcash_getschistory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScHistoryData.ProtoReflect.Descriptor instead.
func (*ScHistoryData) Descriptor() ([]byte, []int) {
	return file_stashcash_getschistory_proto_rawDescGZIP(), []int{2}
}

func (x *ScHistoryData) GetHistory() []*SCtransaction {
	if x != nil {
		return x.History
	}
	return nil
}

type SCtransaction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    int64                  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CampaignId    int64                  `protobuf:"varint,3,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	TxnId         string                 `protobuf:"bytes,4,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	ScType        string                 `protobuf:"bytes,5,opt,name=sc_type,json=scType,proto3" json:"sc_type,omitempty"`
	Amount        float64                `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Remaining     float64                `protobuf:"fixed64,7,opt,name=remaining,proto3" json:"remaining,omitempty"`
	TxnType       string                 `protobuf:"bytes,8,opt,name=txn_type,json=txnType,proto3" json:"txn_type,omitempty"`
	Status        string                 `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Expiry        *string                `protobuf:"bytes,10,opt,name=expiry,proto3,oneof" json:"expiry,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	IsActive      bool                   `protobuf:"varint,12,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	IsDeleted     bool                   `protobuf:"varint,13,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SCtransaction) Reset() {
	*x = SCtransaction{}
	mi := &file_stashcash_getschistory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SCtransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCtransaction) ProtoMessage() {}

func (x *SCtransaction) ProtoReflect() protoreflect.Message {
	mi := &file_stashcash_getschistory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCtransaction.ProtoReflect.Descriptor instead.
func (*SCtransaction) Descriptor() ([]byte, []int) {
	return file_stashcash_getschistory_proto_rawDescGZIP(), []int{3}
}

func (x *SCtransaction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SCtransaction) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SCtransaction) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *SCtransaction) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

func (x *SCtransaction) GetScType() string {
	if x != nil {
		return x.ScType
	}
	return ""
}

func (x *SCtransaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SCtransaction) GetRemaining() float64 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

func (x *SCtransaction) GetTxnType() string {
	if x != nil {
		return x.TxnType
	}
	return ""
}

func (x *SCtransaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SCtransaction) GetExpiry() string {
	if x != nil && x.Expiry != nil {
		return *x.Expiry
	}
	return ""
}

func (x *SCtransaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SCtransaction) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *SCtransaction) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_stashcash_getschistory_proto protoreflect.FileDescriptor

var file_stashcash_getschistory_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x73, 0x74, 0x61, 0x73, 0x68, 0x63, 0x61, 0x73, 0x68, 0x2f, 0x67, 0x65, 0x74, 0x73,
	0x63, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x73, 0x74, 0x61, 0x73, 0x68, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x63, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x73, 0x63,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x22, 0x69, 0x0a, 0x14, 0x67, 0x65, 0x74,
	0x73, 0x63, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x61, 0x73, 0x68, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x63, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x53, 0x63, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x50, 0x0a, 0x0d, 0x53, 0x63, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x61, 0x73, 0x68, 0x63, 0x61,
	0x73, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x63, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x53, 0x43, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xfd, 0x02, 0x0a, 0x0d, 0x53, 0x43, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_stashcash_getschistory_proto_rawDescOnce sync.Once
	file_stashcash_getschistory_proto_rawDescData []byte
)

func file_stashcash_getschistory_proto_rawDescGZIP() []byte {
	file_stashcash_getschistory_proto_rawDescOnce.Do(func() {
		file_stashcash_getschistory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stashcash_getschistory_proto_rawDesc), len(file_stashcash_getschistory_proto_rawDesc)))
	})
	return file_stashcash_getschistory_proto_rawDescData
}

var file_stashcash_getschistory_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stashcash_getschistory_proto_goTypes = []any{
	(*GetschistoryRequest)(nil),  // 0: stashcash.getschistory.getschistoryRequest
	(*GetschistoryResponse)(nil), // 1: stashcash.getschistory.getschistoryResponse
	(*ScHistoryData)(nil),        // 2: stashcash.getschistory.ScHistoryData
	(*SCtransaction)(nil),        // 3: stashcash.getschistory.SCtransaction
}
var file_stashcash_getschistory_proto_depIdxs = []int32{
	2, // 0: stashcash.getschistory.getschistoryResponse.data:type_name -> stashcash.getschistory.ScHistoryData
	3, // 1: stashcash.getschistory.ScHistoryData.history:type_name -> stashcash.getschistory.SCtransaction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stashcash_getschistory_proto_init() }
func file_stashcash_getschistory_proto_init() {
	if File_stashcash_getschistory_proto != nil {
		return
	}
	file_stashcash_getschistory_proto_msgTypes[0].OneofWrappers = []any{}
	file_stashcash_getschistory_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stashcash_getschistory_proto_rawDesc), len(file_stashcash_getschistory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stashcash_getschistory_proto_goTypes,
		DependencyIndexes: file_stashcash_getschistory_proto_depIdxs,
		MessageInfos:      file_stashcash_getschistory_proto_msgTypes,
	}.Build()
	File_stashcash_getschistory_proto = out.File
	file_stashcash_getschistory_proto_goTypes = nil
	file_stashcash_getschistory_proto_depIdxs = nil
}
