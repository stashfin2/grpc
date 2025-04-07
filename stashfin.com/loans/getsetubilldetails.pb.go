// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: loans/getsetubilldetails.proto

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

type GetSetuBillDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerIdentifiers []*GetSetuBillDetailsRequest_Field `protobuf:"bytes,1,rep,name=customerIdentifiers,proto3" json:"customerIdentifiers,omitempty"`
}

func (x *GetSetuBillDetailsRequest) Reset() {
	*x = GetSetuBillDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getsetubilldetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSetuBillDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSetuBillDetailsRequest) ProtoMessage() {}

func (x *GetSetuBillDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getsetubilldetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSetuBillDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetSetuBillDetailsRequest) Descriptor() ([]byte, []int) {
	return file_loans_getsetubilldetails_proto_rawDescGZIP(), []int{0}
}

func (x *GetSetuBillDetailsRequest) GetCustomerIdentifiers() []*GetSetuBillDetailsRequest_Field {
	if x != nil {
		return x.CustomerIdentifiers
	}
	return nil
}

type GetSetuBillDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         []*GetSetuBillDetailsResponse_Field `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalAmount  int32                               `protobuf:"varint,2,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	OrderId      string                              `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId   int32                               `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CustomerName string                              `protobuf:"bytes,5,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
}

func (x *GetSetuBillDetailsResponse) Reset() {
	*x = GetSetuBillDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getsetubilldetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSetuBillDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSetuBillDetailsResponse) ProtoMessage() {}

func (x *GetSetuBillDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getsetubilldetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSetuBillDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetSetuBillDetailsResponse) Descriptor() ([]byte, []int) {
	return file_loans_getsetubilldetails_proto_rawDescGZIP(), []int{1}
}

func (x *GetSetuBillDetailsResponse) GetData() []*GetSetuBillDetailsResponse_Field {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetSetuBillDetailsResponse) GetTotalAmount() int32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *GetSetuBillDetailsResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetSetuBillDetailsResponse) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GetSetuBillDetailsResponse) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

type GetSetuBillDetailsRequest_Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeName  string `protobuf:"bytes,1,opt,name=attributeName,proto3" json:"attributeName,omitempty"`
	AttributeValue string `protobuf:"bytes,2,opt,name=attributeValue,proto3" json:"attributeValue,omitempty"`
}

func (x *GetSetuBillDetailsRequest_Field) Reset() {
	*x = GetSetuBillDetailsRequest_Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getsetubilldetails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSetuBillDetailsRequest_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSetuBillDetailsRequest_Field) ProtoMessage() {}

func (x *GetSetuBillDetailsRequest_Field) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getsetubilldetails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSetuBillDetailsRequest_Field.ProtoReflect.Descriptor instead.
func (*GetSetuBillDetailsRequest_Field) Descriptor() ([]byte, []int) {
	return file_loans_getsetubilldetails_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetSetuBillDetailsRequest_Field) GetAttributeName() string {
	if x != nil {
		return x.AttributeName
	}
	return ""
}

func (x *GetSetuBillDetailsRequest_Field) GetAttributeValue() string {
	if x != nil {
		return x.AttributeValue
	}
	return ""
}

type GetSetuBillDetailsResponse_Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillAmount int32  `protobuf:"varint,1,opt,name=bill_amount,json=billAmount,proto3" json:"bill_amount,omitempty"`
	DueDate    string `protobuf:"bytes,2,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	Principal  int32  `protobuf:"varint,3,opt,name=principal,proto3" json:"principal,omitempty"`
	Interest   int32  `protobuf:"varint,4,opt,name=interest,proto3" json:"interest,omitempty"`
	InstId     int32  `protobuf:"varint,5,opt,name=inst_id,json=instId,proto3" json:"inst_id,omitempty"`
	Penalty    int32  `protobuf:"varint,6,opt,name=penalty,proto3" json:"penalty,omitempty"`
	LoanId     int32  `protobuf:"varint,7,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
}

func (x *GetSetuBillDetailsResponse_Field) Reset() {
	*x = GetSetuBillDetailsResponse_Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_getsetubilldetails_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSetuBillDetailsResponse_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSetuBillDetailsResponse_Field) ProtoMessage() {}

func (x *GetSetuBillDetailsResponse_Field) ProtoReflect() protoreflect.Message {
	mi := &file_loans_getsetubilldetails_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSetuBillDetailsResponse_Field.ProtoReflect.Descriptor instead.
func (*GetSetuBillDetailsResponse_Field) Descriptor() ([]byte, []int) {
	return file_loans_getsetubilldetails_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetSetuBillDetailsResponse_Field) GetBillAmount() int32 {
	if x != nil {
		return x.BillAmount
	}
	return 0
}

func (x *GetSetuBillDetailsResponse_Field) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

func (x *GetSetuBillDetailsResponse_Field) GetPrincipal() int32 {
	if x != nil {
		return x.Principal
	}
	return 0
}

func (x *GetSetuBillDetailsResponse_Field) GetInterest() int32 {
	if x != nil {
		return x.Interest
	}
	return 0
}

func (x *GetSetuBillDetailsResponse_Field) GetInstId() int32 {
	if x != nil {
		return x.InstId
	}
	return 0
}

func (x *GetSetuBillDetailsResponse_Field) GetPenalty() int32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *GetSetuBillDetailsResponse_Field) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

var File_loans_getsetubilldetails_proto protoreflect.FileDescriptor

var file_loans_getsetubilldetails_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x65, 0x74, 0x75, 0x62,
	0x69, 0x6c, 0x6c, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x65, 0x74, 0x75, 0x62,
	0x69, 0x6c, 0x6c, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x19, 0x67,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x75, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6b, 0x0a, 0x13, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x67, 0x65,
	0x74, 0x73, 0x65, 0x74, 0x75, 0x62, 0x69, 0x6c, 0x6c, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x75, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x55, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xbc, 0x03, 0x0a,
	0x1a, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x75, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6c, 0x6f, 0x61, 0x6e,
	0x73, 0x2e, 0x67, 0x65, 0x74, 0x73, 0x65, 0x74, 0x75, 0x62, 0x69, 0x6c, 0x6c, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x75, 0x42, 0x69, 0x6c, 0x6c,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0xc9, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69, 0x6c,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x62, 0x69, 0x6c, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x6e, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x69, 0x6e, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61,
	0x6c, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c,
	0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_loans_getsetubilldetails_proto_rawDescOnce sync.Once
	file_loans_getsetubilldetails_proto_rawDescData = file_loans_getsetubilldetails_proto_rawDesc
)

func file_loans_getsetubilldetails_proto_rawDescGZIP() []byte {
	file_loans_getsetubilldetails_proto_rawDescOnce.Do(func() {
		file_loans_getsetubilldetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_getsetubilldetails_proto_rawDescData)
	})
	return file_loans_getsetubilldetails_proto_rawDescData
}

var file_loans_getsetubilldetails_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_loans_getsetubilldetails_proto_goTypes = []interface{}{
	(*GetSetuBillDetailsRequest)(nil),        // 0: loans.getsetubilldetails.getSetuBillDetailsRequest
	(*GetSetuBillDetailsResponse)(nil),       // 1: loans.getsetubilldetails.getSetuBillDetailsResponse
	(*GetSetuBillDetailsRequest_Field)(nil),  // 2: loans.getsetubilldetails.getSetuBillDetailsRequest.Field
	(*GetSetuBillDetailsResponse_Field)(nil), // 3: loans.getsetubilldetails.getSetuBillDetailsResponse.Field
}
var file_loans_getsetubilldetails_proto_depIdxs = []int32{
	2, // 0: loans.getsetubilldetails.getSetuBillDetailsRequest.customerIdentifiers:type_name -> loans.getsetubilldetails.getSetuBillDetailsRequest.Field
	3, // 1: loans.getsetubilldetails.getSetuBillDetailsResponse.data:type_name -> loans.getsetubilldetails.getSetuBillDetailsResponse.Field
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_loans_getsetubilldetails_proto_init() }
func file_loans_getsetubilldetails_proto_init() {
	if File_loans_getsetubilldetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_getsetubilldetails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSetuBillDetailsRequest); i {
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
		file_loans_getsetubilldetails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSetuBillDetailsResponse); i {
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
		file_loans_getsetubilldetails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSetuBillDetailsRequest_Field); i {
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
		file_loans_getsetubilldetails_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSetuBillDetailsResponse_Field); i {
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
			RawDescriptor: file_loans_getsetubilldetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_getsetubilldetails_proto_goTypes,
		DependencyIndexes: file_loans_getsetubilldetails_proto_depIdxs,
		MessageInfos:      file_loans_getsetubilldetails_proto_msgTypes,
	}.Build()
	File_loans_getsetubilldetails_proto = out.File
	file_loans_getsetubilldetails_proto_rawDesc = nil
	file_loans_getsetubilldetails_proto_goTypes = nil
	file_loans_getsetubilldetails_proto_depIdxs = nil
}
