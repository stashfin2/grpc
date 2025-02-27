// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: colender/disburseLoan.proto

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

type DisburalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      int32   `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	LoanId          int32   `protobuf:"varint,2,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	DisbursalAmount int32   `protobuf:"varint,3,opt,name=disbursal_amount,json=disbursalAmount,proto3" json:"disbursal_amount,omitempty"`
	ApprovalAmount  int32   `protobuf:"varint,4,opt,name=approval_amount,json=approvalAmount,proto3" json:"approval_amount,omitempty"`
	DisbursalDate   string  `protobuf:"bytes,5,opt,name=disbursal_date,json=disbursalDate,proto3" json:"disbursal_date,omitempty"`
	DisbursalMode   string  `protobuf:"bytes,6,opt,name=disbursal_mode,json=disbursalMode,proto3" json:"disbursal_mode,omitempty"`
	Tenure          int32   `protobuf:"varint,7,opt,name=tenure,proto3" json:"tenure,omitempty"`
	Roi             int32   `protobuf:"varint,8,opt,name=roi,proto3" json:"roi,omitempty"`
	DailyRoi        int32   `protobuf:"varint,9,opt,name=daily_roi,json=dailyRoi,proto3" json:"daily_roi,omitempty"`
	FirstEmiDate    string  `protobuf:"bytes,10,opt,name=first_emi_date,json=firstEmiDate,proto3" json:"first_emi_date,omitempty"`
	LoanCycle       int32   `protobuf:"varint,11,opt,name=loan_cycle,json=loanCycle,proto3" json:"loan_cycle,omitempty"`
	NachUrmn        *string `protobuf:"bytes,12,opt,name=nach_urmn,json=nachUrmn,proto3,oneof" json:"nach_urmn,omitempty"`
	Utr             string  `protobuf:"bytes,13,opt,name=utr,proto3" json:"utr,omitempty"`
}

func (x *DisburalRequest) Reset() {
	*x = DisburalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colender_disburseLoan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisburalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisburalRequest) ProtoMessage() {}

func (x *DisburalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_colender_disburseLoan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisburalRequest.ProtoReflect.Descriptor instead.
func (*DisburalRequest) Descriptor() ([]byte, []int) {
	return file_colender_disburseLoan_proto_rawDescGZIP(), []int{0}
}

func (x *DisburalRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *DisburalRequest) GetLoanId() int32 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

func (x *DisburalRequest) GetDisbursalAmount() int32 {
	if x != nil {
		return x.DisbursalAmount
	}
	return 0
}

func (x *DisburalRequest) GetApprovalAmount() int32 {
	if x != nil {
		return x.ApprovalAmount
	}
	return 0
}

func (x *DisburalRequest) GetDisbursalDate() string {
	if x != nil {
		return x.DisbursalDate
	}
	return ""
}

func (x *DisburalRequest) GetDisbursalMode() string {
	if x != nil {
		return x.DisbursalMode
	}
	return ""
}

func (x *DisburalRequest) GetTenure() int32 {
	if x != nil {
		return x.Tenure
	}
	return 0
}

func (x *DisburalRequest) GetRoi() int32 {
	if x != nil {
		return x.Roi
	}
	return 0
}

func (x *DisburalRequest) GetDailyRoi() int32 {
	if x != nil {
		return x.DailyRoi
	}
	return 0
}

func (x *DisburalRequest) GetFirstEmiDate() string {
	if x != nil {
		return x.FirstEmiDate
	}
	return ""
}

func (x *DisburalRequest) GetLoanCycle() int32 {
	if x != nil {
		return x.LoanCycle
	}
	return 0
}

func (x *DisburalRequest) GetNachUrmn() string {
	if x != nil && x.NachUrmn != nil {
		return *x.NachUrmn
	}
	return ""
}

func (x *DisburalRequest) GetUtr() string {
	if x != nil {
		return x.Utr
	}
	return ""
}

type DisburalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ColenderId              int32  `protobuf:"varint,2,opt,name=colender_id,json=colenderId,proto3" json:"colender_id,omitempty"`
	ColenderDisbursalAmount int32  `protobuf:"varint,3,opt,name=colender_disbursal_amount,json=colenderDisbursalAmount,proto3" json:"colender_disbursal_amount,omitempty"`
	DisbursalDate           string `protobuf:"bytes,4,opt,name=disbursal_date,json=disbursalDate,proto3" json:"disbursal_date,omitempty"`
	DisbursalMode           string `protobuf:"bytes,5,opt,name=disbursal_mode,json=disbursalMode,proto3" json:"disbursal_mode,omitempty"`
	ColenderPercentage      int32  `protobuf:"varint,6,opt,name=colender_percentage,json=colenderPercentage,proto3" json:"colender_percentage,omitempty"`
	ColenderApprovalAmount  int32  `protobuf:"varint,7,opt,name=colender_approval_amount,json=colenderApprovalAmount,proto3" json:"colender_approval_amount,omitempty"`
}

func (x *DisburalResponse) Reset() {
	*x = DisburalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colender_disburseLoan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisburalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisburalResponse) ProtoMessage() {}

func (x *DisburalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colender_disburseLoan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisburalResponse.ProtoReflect.Descriptor instead.
func (*DisburalResponse) Descriptor() ([]byte, []int) {
	return file_colender_disburseLoan_proto_rawDescGZIP(), []int{1}
}

func (x *DisburalResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DisburalResponse) GetColenderId() int32 {
	if x != nil {
		return x.ColenderId
	}
	return 0
}

func (x *DisburalResponse) GetColenderDisbursalAmount() int32 {
	if x != nil {
		return x.ColenderDisbursalAmount
	}
	return 0
}

func (x *DisburalResponse) GetDisbursalDate() string {
	if x != nil {
		return x.DisbursalDate
	}
	return ""
}

func (x *DisburalResponse) GetDisbursalMode() string {
	if x != nil {
		return x.DisbursalMode
	}
	return ""
}

func (x *DisburalResponse) GetColenderPercentage() int32 {
	if x != nil {
		return x.ColenderPercentage
	}
	return 0
}

func (x *DisburalResponse) GetColenderApprovalAmount() int32 {
	if x != nil {
		return x.ColenderApprovalAmount
	}
	return 0
}

var File_colender_disburseLoan_proto protoreflect.FileDescriptor

var file_colender_disburseLoan_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x64, 0x69, 0x73, 0x62, 0x75,
	0x72, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x65,
	0x4c, 0x6f, 0x61, 0x6e, 0x22, 0xbb, 0x03, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x69,
	0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72,
	0x73, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x75, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x6f, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f, 0x69, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x72, 0x6f, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x6f, 0x69, 0x12, 0x24, 0x0a, 0x0e, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x65, 0x6d, 0x69, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x45, 0x6d, 0x69, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x61, 0x6e, 0x43, 0x79, 0x63, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x09, 0x6e, 0x61, 0x63, 0x68, 0x5f, 0x75, 0x72, 0x6d, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x61, 0x63, 0x68, 0x55, 0x72, 0x6d, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x74, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x74, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x63, 0x68, 0x5f, 0x75, 0x72,
	0x6d, 0x6e, 0x22, 0xc0, 0x02, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x3a, 0x0a, 0x19, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73,
	0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x17, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x69, 0x73,
	0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x61, 0x6c,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x62, 0x75, 0x72, 0x73, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f,
	0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x63,
	0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x63,
	0x6f, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_colender_disburseLoan_proto_rawDescOnce sync.Once
	file_colender_disburseLoan_proto_rawDescData = file_colender_disburseLoan_proto_rawDesc
)

func file_colender_disburseLoan_proto_rawDescGZIP() []byte {
	file_colender_disburseLoan_proto_rawDescOnce.Do(func() {
		file_colender_disburseLoan_proto_rawDescData = protoimpl.X.CompressGZIP(file_colender_disburseLoan_proto_rawDescData)
	})
	return file_colender_disburseLoan_proto_rawDescData
}

var file_colender_disburseLoan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_colender_disburseLoan_proto_goTypes = []interface{}{
	(*DisburalRequest)(nil),  // 0: colender.disburseLoan.disburalRequest
	(*DisburalResponse)(nil), // 1: colender.disburseLoan.disburalResponse
}
var file_colender_disburseLoan_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_colender_disburseLoan_proto_init() }
func file_colender_disburseLoan_proto_init() {
	if File_colender_disburseLoan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_colender_disburseLoan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisburalRequest); i {
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
		file_colender_disburseLoan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisburalResponse); i {
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
	file_colender_disburseLoan_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_colender_disburseLoan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_colender_disburseLoan_proto_goTypes,
		DependencyIndexes: file_colender_disburseLoan_proto_depIdxs,
		MessageInfos:      file_colender_disburseLoan_proto_msgTypes,
	}.Build()
	File_colender_disburseLoan_proto = out.File
	file_colender_disburseLoan_proto_rawDesc = nil
	file_colender_disburseLoan_proto_goTypes = nil
	file_colender_disburseLoan_proto_depIdxs = nil
}
