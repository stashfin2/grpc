// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: upi/getrecentcontacts.proto

package upi

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

type GetRecentContactsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         *int32                 `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	OffSet        *int32                 `protobuf:"varint,2,opt,name=off_set,json=offSet,proto3,oneof" json:"off_set,omitempty"`
	TxnType       *string                `protobuf:"bytes,3,opt,name=txn_type,json=txnType,proto3,oneof" json:"txn_type,omitempty"`
	SearchBy      *string                `protobuf:"bytes,4,opt,name=search_by,json=searchBy,proto3,oneof" json:"search_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecentContactsRequest) Reset() {
	*x = GetRecentContactsRequest{}
	mi := &file_upi_getrecentcontacts_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecentContactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentContactsRequest) ProtoMessage() {}

func (x *GetRecentContactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getrecentcontacts_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecentContactsRequest.ProtoReflect.Descriptor instead.
func (*GetRecentContactsRequest) Descriptor() ([]byte, []int) {
	return file_upi_getrecentcontacts_proto_rawDescGZIP(), []int{0}
}

func (x *GetRecentContactsRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetRecentContactsRequest) GetOffSet() int32 {
	if x != nil && x.OffSet != nil {
		return *x.OffSet
	}
	return 0
}

func (x *GetRecentContactsRequest) GetTxnType() string {
	if x != nil && x.TxnType != nil {
		return *x.TxnType
	}
	return ""
}

func (x *GetRecentContactsRequest) GetSearchBy() string {
	if x != nil && x.SearchBy != nil {
		return *x.SearchBy
	}
	return ""
}

type RecentContacts struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	PayeeVpa         *string                `protobuf:"bytes,1,opt,name=payee_vpa,json=payeeVpa,proto3,oneof" json:"payee_vpa,omitempty"`    // Payee's VPA (Virtual Payment Address)
	PayeeName        *string                `protobuf:"bytes,2,opt,name=payee_name,json=payeeName,proto3,oneof" json:"payee_name,omitempty"` // Name of the payee
	PayeeAccNumber   *string                `protobuf:"bytes,3,opt,name=payee_acc_number,json=payeeAccNumber,proto3,oneof" json:"payee_acc_number,omitempty"`
	PayeeAccProvider *string                `protobuf:"bytes,4,opt,name=payee_acc_provider,json=payeeAccProvider,proto3,oneof" json:"payee_acc_provider,omitempty"`
	CreatedAt        *string                `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	Amount           *string                `protobuf:"bytes,6,opt,name=amount,proto3,oneof" json:"amount,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *RecentContacts) Reset() {
	*x = RecentContacts{}
	mi := &file_upi_getrecentcontacts_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecentContacts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecentContacts) ProtoMessage() {}

func (x *RecentContacts) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getrecentcontacts_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecentContacts.ProtoReflect.Descriptor instead.
func (*RecentContacts) Descriptor() ([]byte, []int) {
	return file_upi_getrecentcontacts_proto_rawDescGZIP(), []int{1}
}

func (x *RecentContacts) GetPayeeVpa() string {
	if x != nil && x.PayeeVpa != nil {
		return *x.PayeeVpa
	}
	return ""
}

func (x *RecentContacts) GetPayeeName() string {
	if x != nil && x.PayeeName != nil {
		return *x.PayeeName
	}
	return ""
}

func (x *RecentContacts) GetPayeeAccNumber() string {
	if x != nil && x.PayeeAccNumber != nil {
		return *x.PayeeAccNumber
	}
	return ""
}

func (x *RecentContacts) GetPayeeAccProvider() string {
	if x != nil && x.PayeeAccProvider != nil {
		return *x.PayeeAccProvider
	}
	return ""
}

func (x *RecentContacts) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

func (x *RecentContacts) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Contacts      []*RecentContacts      `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_upi_getrecentcontacts_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getrecentcontacts_proto_msgTypes[2]
	if x != nil {
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
	return file_upi_getrecentcontacts_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetContacts() []*RecentContacts {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type GetRecentContactsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *Data                  `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecentContactsResponse) Reset() {
	*x = GetRecentContactsResponse{}
	mi := &file_upi_getrecentcontacts_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecentContactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentContactsResponse) ProtoMessage() {}

func (x *GetRecentContactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_getrecentcontacts_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecentContactsResponse.ProtoReflect.Descriptor instead.
func (*GetRecentContactsResponse) Descriptor() ([]byte, []int) {
	return file_upi_getrecentcontacts_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecentContactsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetRecentContactsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRecentContactsResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_upi_getrecentcontacts_proto protoreflect.FileDescriptor

var file_upi_getrecentcontacts_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x75, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x75,
	0x70, 0x69, 0x2e, 0x67, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07,
	0x6f, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x53, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x74, 0x78,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07,
	0x74, 0x78, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x66, 0x66, 0x5f, 0x73,
	0x65, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x62, 0x79, 0x22, 0xdc, 0x02,
	0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x76, 0x70, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x79, 0x65, 0x65, 0x56, 0x70, 0x61, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x65, 0x65, 0x41, 0x63, 0x63, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x10, 0x70, 0x61, 0x79, 0x65, 0x65, 0x41, 0x63, 0x63, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61,
	0x79, 0x65, 0x65, 0x5f, 0x76, 0x70, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x79, 0x65,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x15, 0x0a, 0x13, 0x5f,
	0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x19, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_upi_getrecentcontacts_proto_rawDescOnce sync.Once
	file_upi_getrecentcontacts_proto_rawDescData []byte
)

func file_upi_getrecentcontacts_proto_rawDescGZIP() []byte {
	file_upi_getrecentcontacts_proto_rawDescOnce.Do(func() {
		file_upi_getrecentcontacts_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upi_getrecentcontacts_proto_rawDesc), len(file_upi_getrecentcontacts_proto_rawDesc)))
	})
	return file_upi_getrecentcontacts_proto_rawDescData
}

var file_upi_getrecentcontacts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_upi_getrecentcontacts_proto_goTypes = []any{
	(*GetRecentContactsRequest)(nil),  // 0: upi.getRecentContacts.getRecentContactsRequest
	(*RecentContacts)(nil),            // 1: upi.getRecentContacts.RecentContacts
	(*Data)(nil),                      // 2: upi.getRecentContacts.Data
	(*GetRecentContactsResponse)(nil), // 3: upi.getRecentContacts.getRecentContactsResponse
}
var file_upi_getrecentcontacts_proto_depIdxs = []int32{
	1, // 0: upi.getRecentContacts.Data.contacts:type_name -> upi.getRecentContacts.RecentContacts
	2, // 1: upi.getRecentContacts.getRecentContactsResponse.data:type_name -> upi.getRecentContacts.Data
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upi_getrecentcontacts_proto_init() }
func file_upi_getrecentcontacts_proto_init() {
	if File_upi_getrecentcontacts_proto != nil {
		return
	}
	file_upi_getrecentcontacts_proto_msgTypes[0].OneofWrappers = []any{}
	file_upi_getrecentcontacts_proto_msgTypes[1].OneofWrappers = []any{}
	file_upi_getrecentcontacts_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upi_getrecentcontacts_proto_rawDesc), len(file_upi_getrecentcontacts_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_getrecentcontacts_proto_goTypes,
		DependencyIndexes: file_upi_getrecentcontacts_proto_depIdxs,
		MessageInfos:      file_upi_getrecentcontacts_proto_msgTypes,
	}.Build()
	File_upi_getrecentcontacts_proto = out.File
	file_upi_getrecentcontacts_proto_goTypes = nil
	file_upi_getrecentcontacts_proto_depIdxs = nil
}
