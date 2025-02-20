// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: colender/forecloseLoan.proto

package colender

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

type ForecloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  int32  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	LoanId      int32  `protobuf:"varint,2,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	SeqNo       int32  `protobuf:"varint,3,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	PaymentMode string `protobuf:"bytes,4,opt,name=payment_mode,json=paymentMode,proto3" json:"payment_mode,omitempty"`
	Utr         string `protobuf:"bytes,5,opt,name=utr,proto3" json:"utr,omitempty"`
	PaymentDate string `protobuf:"bytes,6,opt,name=payment_date,json=paymentDate,proto3" json:"payment_date,omitempty"`
	PaymentId   int32  `protobuf:"varint,7,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Amount      int32  `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ForecloseRequest) Reset() {
	*x = ForecloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colender_forecloseLoan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecloseRequest) ProtoMessage() {}

func (x *ForecloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_colender_forecloseLoan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecloseRequest.ProtoReflect.Descriptor instead.
func (*ForecloseRequest) Descriptor() ([]byte, []int) {
	return file_colender_forecloseLoan_proto_rawDescGZIP(), []int{0}
}

func (x *ForecloseRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ForecloseRequest) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

func (x *ForecloseRequest) GetSeqNo() int32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *ForecloseRequest) GetPaymentMode() string {
	if x != nil {
		return x.PaymentMode
	}
	return ""
}

func (x *ForecloseRequest) GetUtr() string {
	if x != nil {
		return x.Utr
	}
	return ""
}

func (x *ForecloseRequest) GetPaymentDate() string {
	if x != nil {
		return x.PaymentDate
	}
	return ""
}

func (x *ForecloseRequest) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *ForecloseRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ForecloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColenderId       int32                `protobuf:"varint,1,opt,name=colender_id,json=colenderId,proto3" json:"colender_id,omitempty"`
	Query            *string              `protobuf:"bytes,2,opt,name=query,proto3,oneof" json:"query,omitempty"`
	ActionStatus     *ForecloseLoanStatus `protobuf:"bytes,3,opt,name=actionStatus,proto3" json:"actionStatus,omitempty"`
	InstallmentQuery []string             `protobuf:"bytes,4,rep,name=installmentQuery,proto3" json:"installmentQuery,omitempty"`
	PaymentsQuery    []string             `protobuf:"bytes,5,rep,name=paymentsQuery,proto3" json:"paymentsQuery,omitempty"`
	ForecloseQuery   string               `protobuf:"bytes,6,opt,name=forecloseQuery,proto3" json:"forecloseQuery,omitempty"`
	CloseLoanQuery   *string              `protobuf:"bytes,7,opt,name=closeLoanQuery,proto3,oneof" json:"closeLoanQuery,omitempty"`
}

func (x *ForecloseResponse) Reset() {
	*x = ForecloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colender_forecloseLoan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecloseResponse) ProtoMessage() {}

func (x *ForecloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colender_forecloseLoan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecloseResponse.ProtoReflect.Descriptor instead.
func (*ForecloseResponse) Descriptor() ([]byte, []int) {
	return file_colender_forecloseLoan_proto_rawDescGZIP(), []int{1}
}

func (x *ForecloseResponse) GetColenderId() int32 {
	if x != nil {
		return x.ColenderId
	}
	return 0
}

func (x *ForecloseResponse) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *ForecloseResponse) GetActionStatus() *ForecloseLoanStatus {
	if x != nil {
		return x.ActionStatus
	}
	return nil
}

func (x *ForecloseResponse) GetInstallmentQuery() []string {
	if x != nil {
		return x.InstallmentQuery
	}
	return nil
}

func (x *ForecloseResponse) GetPaymentsQuery() []string {
	if x != nil {
		return x.PaymentsQuery
	}
	return nil
}

func (x *ForecloseResponse) GetForecloseQuery() string {
	if x != nil {
		return x.ForecloseQuery
	}
	return ""
}

func (x *ForecloseResponse) GetCloseLoanQuery() string {
	if x != nil && x.CloseLoanQuery != nil {
		return *x.CloseLoanQuery
	}
	return ""
}

type ForecloseLoanStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error   *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ForecloseLoanStatus) Reset() {
	*x = ForecloseLoanStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colender_forecloseLoan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecloseLoanStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecloseLoanStatus) ProtoMessage() {}

func (x *ForecloseLoanStatus) ProtoReflect() protoreflect.Message {
	mi := &file_colender_forecloseLoan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecloseLoanStatus.ProtoReflect.Descriptor instead.
func (*ForecloseLoanStatus) Descriptor() ([]byte, []int) {
	return file_colender_forecloseLoan_proto_rawDescGZIP(), []int{2}
}

func (x *ForecloseLoanStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ForecloseLoanStatus) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *ForecloseLoanStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_colender_forecloseLoan_proto protoreflect.FileDescriptor

var file_colender_forecloseLoan_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x22, 0xf2, 0x01, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x65, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c,
	0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x74,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xe4, 0x02, 0x0a, 0x11,
	0x66, 0x6f, 0x72, 0x65, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a,
	0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x66,
	0x6f, 0x72, 0x65, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x2e, 0x66, 0x6f, 0x72,
	0x65, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a,
	0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x22, 0x6c, 0x0a, 0x13, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x4c,
	0x6f, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_colender_forecloseLoan_proto_rawDescOnce sync.Once
	file_colender_forecloseLoan_proto_rawDescData = file_colender_forecloseLoan_proto_rawDesc
)

func file_colender_forecloseLoan_proto_rawDescGZIP() []byte {
	file_colender_forecloseLoan_proto_rawDescOnce.Do(func() {
		file_colender_forecloseLoan_proto_rawDescData = protoimpl.X.CompressGZIP(file_colender_forecloseLoan_proto_rawDescData)
	})
	return file_colender_forecloseLoan_proto_rawDescData
}

var file_colender_forecloseLoan_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_colender_forecloseLoan_proto_goTypes = []interface{}{
	(*ForecloseRequest)(nil),    // 0: colender.forecloseLoan.forecloseRequest
	(*ForecloseResponse)(nil),   // 1: colender.forecloseLoan.forecloseResponse
	(*ForecloseLoanStatus)(nil), // 2: colender.forecloseLoan.forecloseLoanStatus
}
var file_colender_forecloseLoan_proto_depIdxs = []int32{
	2, // 0: colender.forecloseLoan.forecloseResponse.actionStatus:type_name -> colender.forecloseLoan.forecloseLoanStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_colender_forecloseLoan_proto_init() }
func file_colender_forecloseLoan_proto_init() {
	if File_colender_forecloseLoan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_colender_forecloseLoan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecloseRequest); i {
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
		file_colender_forecloseLoan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecloseResponse); i {
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
		file_colender_forecloseLoan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecloseLoanStatus); i {
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
	file_colender_forecloseLoan_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_colender_forecloseLoan_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_colender_forecloseLoan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_colender_forecloseLoan_proto_goTypes,
		DependencyIndexes: file_colender_forecloseLoan_proto_depIdxs,
		MessageInfos:      file_colender_forecloseLoan_proto_msgTypes,
	}.Build()
	File_colender_forecloseLoan_proto = out.File
	file_colender_forecloseLoan_proto_rawDesc = nil
	file_colender_forecloseLoan_proto_goTypes = nil
	file_colender_forecloseLoan_proto_depIdxs = nil
}
