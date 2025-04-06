// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: loans/fcppayment.proto

package loans

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

type FcpPaymentRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	BillAmount    int32                     `protobuf:"varint,1,opt,name=bill_amount,json=billAmount,proto3" json:"bill_amount,omitempty"`
	Customerid    int32                     `protobuf:"varint,2,opt,name=customerid,proto3" json:"customerid,omitempty"`
	LoanId        []int32                   `protobuf:"varint,3,rep,packed,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	Mode          string                    `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	AddOns        []*FcpPaymentRequestAddOn `protobuf:"bytes,5,rep,name=add_ons,json=addOns,proto3" json:"add_ons,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FcpPaymentRequest) Reset() {
	*x = FcpPaymentRequest{}
	mi := &file_loans_fcppayment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FcpPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FcpPaymentRequest) ProtoMessage() {}

func (x *FcpPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_fcppayment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FcpPaymentRequest.ProtoReflect.Descriptor instead.
func (*FcpPaymentRequest) Descriptor() ([]byte, []int) {
	return file_loans_fcppayment_proto_rawDescGZIP(), []int{0}
}

func (x *FcpPaymentRequest) GetBillAmount() int32 {
	if x != nil {
		return x.BillAmount
	}
	return 0
}

func (x *FcpPaymentRequest) GetCustomerid() int32 {
	if x != nil {
		return x.Customerid
	}
	return 0
}

func (x *FcpPaymentRequest) GetLoanId() []int32 {
	if x != nil {
		return x.LoanId
	}
	return nil
}

func (x *FcpPaymentRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *FcpPaymentRequest) GetAddOns() []*FcpPaymentRequestAddOn {
	if x != nil {
		return x.AddOns
	}
	return nil
}

type FcpPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RedirectUrl   string                 `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	OrderId       string                 `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FcpPaymentResponse) Reset() {
	*x = FcpPaymentResponse{}
	mi := &file_loans_fcppayment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FcpPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FcpPaymentResponse) ProtoMessage() {}

func (x *FcpPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_fcppayment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FcpPaymentResponse.ProtoReflect.Descriptor instead.
func (*FcpPaymentResponse) Descriptor() ([]byte, []int) {
	return file_loans_fcppayment_proto_rawDescGZIP(), []int{1}
}

func (x *FcpPaymentResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *FcpPaymentResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *FcpPaymentResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type FcpPaymentRequestAddOn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PlanId        string                 `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Amount        int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	MetaData      string                 `protobuf:"bytes,4,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FcpPaymentRequestAddOn) Reset() {
	*x = FcpPaymentRequestAddOn{}
	mi := &file_loans_fcppayment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FcpPaymentRequestAddOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FcpPaymentRequestAddOn) ProtoMessage() {}

func (x *FcpPaymentRequestAddOn) ProtoReflect() protoreflect.Message {
	mi := &file_loans_fcppayment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FcpPaymentRequestAddOn.ProtoReflect.Descriptor instead.
func (*FcpPaymentRequestAddOn) Descriptor() ([]byte, []int) {
	return file_loans_fcppayment_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FcpPaymentRequestAddOn) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FcpPaymentRequestAddOn) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *FcpPaymentRequestAddOn) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FcpPaymentRequestAddOn) GetMetaData() string {
	if x != nil {
		return x.MetaData
	}
	return ""
}

var File_loans_fcppayment_proto protoreflect.FileDescriptor

var file_loans_fcppayment_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x66, 0x63, 0x70, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e,
	0x66, 0x63, 0x70, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xb2, 0x02, 0x0a, 0x11, 0x66,
	0x63, 0x70, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x43,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x66, 0x63, 0x70, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x66, 0x63, 0x70, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x64, 0x64,
	0x4f, 0x6e, 0x73, 0x1a, 0x6a, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x7d, 0x0a, 0x12, 0x66, 0x63, 0x70, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_loans_fcppayment_proto_rawDescOnce sync.Once
	file_loans_fcppayment_proto_rawDescData []byte
)

func file_loans_fcppayment_proto_rawDescGZIP() []byte {
	file_loans_fcppayment_proto_rawDescOnce.Do(func() {
		file_loans_fcppayment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_loans_fcppayment_proto_rawDesc), len(file_loans_fcppayment_proto_rawDesc)))
	})
	return file_loans_fcppayment_proto_rawDescData
}

var file_loans_fcppayment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_loans_fcppayment_proto_goTypes = []any{
	(*FcpPaymentRequest)(nil),      // 0: loans.fcppayment.fcpPaymentRequest
	(*FcpPaymentResponse)(nil),     // 1: loans.fcppayment.fcpPaymentResponse
	(*FcpPaymentRequestAddOn)(nil), // 2: loans.fcppayment.fcpPaymentRequest.add_on
}
var file_loans_fcppayment_proto_depIdxs = []int32{
	2, // 0: loans.fcppayment.fcpPaymentRequest.add_ons:type_name -> loans.fcppayment.fcpPaymentRequest.add_on
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_loans_fcppayment_proto_init() }
func file_loans_fcppayment_proto_init() {
	if File_loans_fcppayment_proto != nil {
		return
	}
	file_loans_fcppayment_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_loans_fcppayment_proto_rawDesc), len(file_loans_fcppayment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_fcppayment_proto_goTypes,
		DependencyIndexes: file_loans_fcppayment_proto_depIdxs,
		MessageInfos:      file_loans_fcppayment_proto_msgTypes,
	}.Build()
	File_loans_fcppayment_proto = out.File
	file_loans_fcppayment_proto_goTypes = nil
	file_loans_fcppayment_proto_depIdxs = nil
}
