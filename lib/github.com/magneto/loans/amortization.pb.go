// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: loans/amortization.proto

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

type AmortizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoanId int32 `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
}

func (x *AmortizationRequest) Reset() {
	*x = AmortizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_amortization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmortizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmortizationRequest) ProtoMessage() {}

func (x *AmortizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_amortization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmortizationRequest.ProtoReflect.Descriptor instead.
func (*AmortizationRequest) Descriptor() ([]byte, []int) {
	return file_loans_amortization_proto_rawDescGZIP(), []int{0}
}

func (x *AmortizationRequest) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

type AmortizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installments []*AmortizationResponse_Field `protobuf:"bytes,1,rep,name=installments,proto3" json:"installments,omitempty"`
}

func (x *AmortizationResponse) Reset() {
	*x = AmortizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_amortization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmortizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmortizationResponse) ProtoMessage() {}

func (x *AmortizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_amortization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmortizationResponse.ProtoReflect.Descriptor instead.
func (*AmortizationResponse) Descriptor() ([]byte, []int) {
	return file_loans_amortization_proto_rawDescGZIP(), []int{1}
}

func (x *AmortizationResponse) GetInstallments() []*AmortizationResponse_Field {
	if x != nil {
		return x.Installments
	}
	return nil
}

type AmortizationResponse_Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoanId          int32  `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	CustomerId      int32  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	SeqNo           int32  `protobuf:"varint,3,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	Amount          int32  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	StartingBalance int32  `protobuf:"varint,5,opt,name=starting_balance,json=startingBalance,proto3" json:"starting_balance,omitempty"`
	EndingBalance   int32  `protobuf:"varint,6,opt,name=ending_balance,json=endingBalance,proto3" json:"ending_balance,omitempty"`
	Principal       int32  `protobuf:"varint,7,opt,name=principal,proto3" json:"principal,omitempty"`
	Interest        int32  `protobuf:"varint,8,opt,name=interest,proto3" json:"interest,omitempty"`
	DueDate         string `protobuf:"bytes,9,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	Status          string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Discount        int32  `protobuf:"varint,11,opt,name=discount,proto3" json:"discount,omitempty"`
	Penalty         int32  `protobuf:"varint,12,opt,name=penalty,proto3" json:"penalty,omitempty"`
}

func (x *AmortizationResponse_Field) Reset() {
	*x = AmortizationResponse_Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_amortization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmortizationResponse_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmortizationResponse_Field) ProtoMessage() {}

func (x *AmortizationResponse_Field) ProtoReflect() protoreflect.Message {
	mi := &file_loans_amortization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmortizationResponse_Field.ProtoReflect.Descriptor instead.
func (*AmortizationResponse_Field) Descriptor() ([]byte, []int) {
	return file_loans_amortization_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AmortizationResponse_Field) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

func (x *AmortizationResponse_Field) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *AmortizationResponse_Field) GetSeqNo() int32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *AmortizationResponse_Field) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AmortizationResponse_Field) GetStartingBalance() int32 {
	if x != nil {
		return x.StartingBalance
	}
	return 0
}

func (x *AmortizationResponse_Field) GetEndingBalance() int32 {
	if x != nil {
		return x.EndingBalance
	}
	return 0
}

func (x *AmortizationResponse_Field) GetPrincipal() int32 {
	if x != nil {
		return x.Principal
	}
	return 0
}

func (x *AmortizationResponse_Field) GetInterest() int32 {
	if x != nil {
		return x.Interest
	}
	return 0
}

func (x *AmortizationResponse_Field) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

func (x *AmortizationResponse_Field) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AmortizationResponse_Field) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *AmortizationResponse_Field) GetPenalty() int32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

var File_loans_amortization_proto protoreflect.FileDescriptor

var file_loans_amortization_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6c, 0x6f, 0x61, 0x6e,
	0x73, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e,
	0x0a, 0x13, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x22, 0xd2,
	0x03, 0x0a, 0x14, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xe5, 0x02, 0x0a, 0x05,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e,
	0x61, 0x6c, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61,
	0x6c, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loans_amortization_proto_rawDescOnce sync.Once
	file_loans_amortization_proto_rawDescData = file_loans_amortization_proto_rawDesc
)

func file_loans_amortization_proto_rawDescGZIP() []byte {
	file_loans_amortization_proto_rawDescOnce.Do(func() {
		file_loans_amortization_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_amortization_proto_rawDescData)
	})
	return file_loans_amortization_proto_rawDescData
}

var file_loans_amortization_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_loans_amortization_proto_goTypes = []interface{}{
	(*AmortizationRequest)(nil),        // 0: loans.amortization.amortizationRequest
	(*AmortizationResponse)(nil),       // 1: loans.amortization.amortizationResponse
	(*AmortizationResponse_Field)(nil), // 2: loans.amortization.amortizationResponse.Field
}
var file_loans_amortization_proto_depIdxs = []int32{
	2, // 0: loans.amortization.amortizationResponse.installments:type_name -> loans.amortization.amortizationResponse.Field
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_loans_amortization_proto_init() }
func file_loans_amortization_proto_init() {
	if File_loans_amortization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_amortization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmortizationRequest); i {
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
		file_loans_amortization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmortizationResponse); i {
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
		file_loans_amortization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmortizationResponse_Field); i {
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
			RawDescriptor: file_loans_amortization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_amortization_proto_goTypes,
		DependencyIndexes: file_loans_amortization_proto_depIdxs,
		MessageInfos:      file_loans_amortization_proto_msgTypes,
	}.Build()
	File_loans_amortization_proto = out.File
	file_loans_amortization_proto_rawDesc = nil
	file_loans_amortization_proto_goTypes = nil
	file_loans_amortization_proto_depIdxs = nil
}
