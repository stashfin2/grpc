// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: loans/paynow.proto

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

type PayNowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId    int32                 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Mode          string                `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	AddOns        []*PayNowRequestAddOn `protobuf:"bytes,3,rep,name=add_ons,json=addOns,proto3" json:"add_ons,omitempty"`
	Amount        int32                 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	LoanId        []int32               `protobuf:"varint,5,rep,packed,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	InstallmentId []int32               `protobuf:"varint,6,rep,packed,name=installment_id,json=installmentId,proto3" json:"installment_id,omitempty"`
	BillId        []int32               `protobuf:"varint,7,rep,packed,name=bill_id,json=billId,proto3" json:"bill_id,omitempty"`
}

func (x *PayNowRequest) Reset() {
	*x = PayNowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_paynow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNowRequest) ProtoMessage() {}

func (x *PayNowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loans_paynow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNowRequest.ProtoReflect.Descriptor instead.
func (*PayNowRequest) Descriptor() ([]byte, []int) {
	return file_loans_paynow_proto_rawDescGZIP(), []int{0}
}

func (x *PayNowRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *PayNowRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *PayNowRequest) GetAddOns() []*PayNowRequestAddOn {
	if x != nil {
		return x.AddOns
	}
	return nil
}

func (x *PayNowRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayNowRequest) GetLoanId() []int32 {
	if x != nil {
		return x.LoanId
	}
	return nil
}

func (x *PayNowRequest) GetInstallmentId() []int32 {
	if x != nil {
		return x.InstallmentId
	}
	return nil
}

func (x *PayNowRequest) GetBillId() []int32 {
	if x != nil {
		return x.BillId
	}
	return nil
}

type PayNowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	OrderId     string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *PayNowResponse) Reset() {
	*x = PayNowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_paynow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNowResponse) ProtoMessage() {}

func (x *PayNowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loans_paynow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNowResponse.ProtoReflect.Descriptor instead.
func (*PayNowResponse) Descriptor() ([]byte, []int) {
	return file_loans_paynow_proto_rawDescGZIP(), []int{1}
}

func (x *PayNowResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *PayNowResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type PayNowRequestAddOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PlanId   string `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Amount   int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	MetaData string `protobuf:"bytes,4,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
}

func (x *PayNowRequestAddOn) Reset() {
	*x = PayNowRequestAddOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loans_paynow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNowRequestAddOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNowRequestAddOn) ProtoMessage() {}

func (x *PayNowRequestAddOn) ProtoReflect() protoreflect.Message {
	mi := &file_loans_paynow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNowRequestAddOn.ProtoReflect.Descriptor instead.
func (*PayNowRequestAddOn) Descriptor() ([]byte, []int) {
	return file_loans_paynow_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PayNowRequestAddOn) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PayNowRequestAddOn) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *PayNowRequestAddOn) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayNowRequestAddOn) GetMetaData() string {
	if x != nil {
		return x.MetaData
	}
	return ""
}

var File_loans_paynow_proto protoreflect.FileDescriptor

var file_loans_paynow_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6e, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6e,
	0x6f, 0x77, 0x22, 0xde, 0x02, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x5f, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x6f, 0x61,
	0x6e, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6e, 0x6f, 0x77, 0x2e, 0x70, 0x61, 0x79, 0x4e, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x64, 0x64, 0x4f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x06, 0x62, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x1a, 0x6a, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x5f, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loans_paynow_proto_rawDescOnce sync.Once
	file_loans_paynow_proto_rawDescData = file_loans_paynow_proto_rawDesc
)

func file_loans_paynow_proto_rawDescGZIP() []byte {
	file_loans_paynow_proto_rawDescOnce.Do(func() {
		file_loans_paynow_proto_rawDescData = protoimpl.X.CompressGZIP(file_loans_paynow_proto_rawDescData)
	})
	return file_loans_paynow_proto_rawDescData
}

var file_loans_paynow_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_loans_paynow_proto_goTypes = []interface{}{
	(*PayNowRequest)(nil),      // 0: loans.paynow.payNowRequest
	(*PayNowResponse)(nil),     // 1: loans.paynow.payNowResponse
	(*PayNowRequestAddOn)(nil), // 2: loans.paynow.payNowRequest.add_on
}
var file_loans_paynow_proto_depIdxs = []int32{
	2, // 0: loans.paynow.payNowRequest.add_ons:type_name -> loans.paynow.payNowRequest.add_on
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_loans_paynow_proto_init() }
func file_loans_paynow_proto_init() {
	if File_loans_paynow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loans_paynow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNowRequest); i {
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
		file_loans_paynow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNowResponse); i {
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
		file_loans_paynow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNowRequestAddOn); i {
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
			RawDescriptor: file_loans_paynow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loans_paynow_proto_goTypes,
		DependencyIndexes: file_loans_paynow_proto_depIdxs,
		MessageInfos:      file_loans_paynow_proto_msgTypes,
	}.Build()
	File_loans_paynow_proto = out.File
	file_loans_paynow_proto_rawDesc = nil
	file_loans_paynow_proto_goTypes = nil
	file_loans_paynow_proto_depIdxs = nil
}
