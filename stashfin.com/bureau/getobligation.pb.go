// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: bureau/getobligation.proto

package bureau

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

type ObligationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ObligationRequest) Reset() {
	*x = ObligationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getobligation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObligationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObligationRequest) ProtoMessage() {}

func (x *ObligationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getobligation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObligationRequest.ProtoReflect.Descriptor instead.
func (*ObligationRequest) Descriptor() ([]byte, []int) {
	return file_bureau_getobligation_proto_rawDescGZIP(), []int{0}
}

func (x *ObligationRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type ObligationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   *ObligationData `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *ObligationResponse) Reset() {
	*x = ObligationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getobligation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObligationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObligationResponse) ProtoMessage() {}

func (x *ObligationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getobligation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObligationResponse.ProtoReflect.Descriptor instead.
func (*ObligationResponse) Descriptor() ([]byte, []int) {
	return file_bureau_getobligation_proto_rawDescGZIP(), []int{1}
}

func (x *ObligationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ObligationResponse) GetData() *ObligationData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ObligationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxAllowedDbr          int32          `protobuf:"varint,1,opt,name=max_allowed_dbr,json=maxAllowedDbr,proto3" json:"max_allowed_dbr,omitempty"`
	ActiveLoansObligation  int32          `protobuf:"varint,2,opt,name=active_loans_obligation,json=activeLoansObligation,proto3" json:"active_loans_obligation,omitempty"`
	MonthlyObligations     int32          `protobuf:"varint,3,opt,name=monthly_obligations,json=monthlyObligations,proto3" json:"monthly_obligations,omitempty"`
	BureauMonthObligations int32          `protobuf:"varint,4,opt,name=bureau_month_obligations,json=bureauMonthObligations,proto3" json:"bureau_month_obligations,omitempty"`
	MaxObligation          int32          `protobuf:"varint,5,opt,name=max_obligation,json=maxObligation,proto3" json:"max_obligation,omitempty"`
	Salary                 int32          `protobuf:"varint,6,opt,name=salary,proto3" json:"salary,omitempty"`
	MaxEmi                 int32          `protobuf:"varint,7,opt,name=max_emi,json=maxEmi,proto3" json:"max_emi,omitempty"`
	Bureau                 *Bureau        `protobuf:"bytes,8,opt,name=bureau,proto3,oneof" json:"bureau,omitempty"`
	BureauObligationRows   []*AccountList `protobuf:"bytes,9,rep,name=bureau_obligation_rows,json=bureauObligationRows,proto3" json:"bureau_obligation_rows,omitempty"`
}

func (x *ObligationData) Reset() {
	*x = ObligationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getobligation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObligationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObligationData) ProtoMessage() {}

func (x *ObligationData) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getobligation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObligationData.ProtoReflect.Descriptor instead.
func (*ObligationData) Descriptor() ([]byte, []int) {
	return file_bureau_getobligation_proto_rawDescGZIP(), []int{2}
}

func (x *ObligationData) GetMaxAllowedDbr() int32 {
	if x != nil {
		return x.MaxAllowedDbr
	}
	return 0
}

func (x *ObligationData) GetActiveLoansObligation() int32 {
	if x != nil {
		return x.ActiveLoansObligation
	}
	return 0
}

func (x *ObligationData) GetMonthlyObligations() int32 {
	if x != nil {
		return x.MonthlyObligations
	}
	return 0
}

func (x *ObligationData) GetBureauMonthObligations() int32 {
	if x != nil {
		return x.BureauMonthObligations
	}
	return 0
}

func (x *ObligationData) GetMaxObligation() int32 {
	if x != nil {
		return x.MaxObligation
	}
	return 0
}

func (x *ObligationData) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *ObligationData) GetMaxEmi() int32 {
	if x != nil {
		return x.MaxEmi
	}
	return 0
}

func (x *ObligationData) GetBureau() *Bureau {
	if x != nil {
		return x.Bureau
	}
	return nil
}

func (x *ObligationData) GetBureauObligationRows() []*AccountList {
	if x != nil {
		return x.BureauObligationRows
	}
	return nil
}

type Bureau struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score    int32  `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	PullDate string `protobuf:"bytes,2,opt,name=pull_date,json=pullDate,proto3" json:"pull_date,omitempty"`
}

func (x *Bureau) Reset() {
	*x = Bureau{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getobligation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bureau) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bureau) ProtoMessage() {}

func (x *Bureau) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getobligation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bureau.ProtoReflect.Descriptor instead.
func (*Bureau) Descriptor() ([]byte, []int) {
	return file_bureau_getobligation_proto_rawDescGZIP(), []int{3}
}

func (x *Bureau) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Bureau) GetPullDate() string {
	if x != nil {
		return x.PullDate
	}
	return ""
}

type AccountList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountType      int32 `protobuf:"varint,2,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	CurrentBalance   int32 `protobuf:"varint,3,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
	SanctionedAmount int32 `protobuf:"varint,4,opt,name=sanctioned_amount,json=sanctionedAmount,proto3" json:"sanctioned_amount,omitempty"`
	EmiAmount        int32 `protobuf:"varint,5,opt,name=emi_amount,json=emiAmount,proto3" json:"emi_amount,omitempty"`
	Obligation       int32 `protobuf:"varint,6,opt,name=obligation,proto3" json:"obligation,omitempty"`
}

func (x *AccountList) Reset() {
	*x = AccountList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bureau_getobligation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountList) ProtoMessage() {}

func (x *AccountList) ProtoReflect() protoreflect.Message {
	mi := &file_bureau_getobligation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountList.ProtoReflect.Descriptor instead.
func (*AccountList) Descriptor() ([]byte, []int) {
	return file_bureau_getobligation_proto_rawDescGZIP(), []int{4}
}

func (x *AccountList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountList) GetAccountType() int32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *AccountList) GetCurrentBalance() int32 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

func (x *AccountList) GetSanctionedAmount() int32 {
	if x != nil {
		return x.SanctionedAmount
	}
	return 0
}

func (x *AccountList) GetEmiAmount() int32 {
	if x != nil {
		return x.EmiAmount
	}
	return 0
}

func (x *AccountList) GetObligation() int32 {
	if x != nil {
		return x.Obligation
	}
	return 0
}

var File_bureau_getobligation_proto protoreflect.FileDescriptor

var file_bureau_getobligation_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2f, 0x67, 0x65, 0x74, 0x6f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x75,
	0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x11, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x12, 0x6f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65,
	0x74, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd2,
	0x03, 0x0a, 0x0e, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x64, 0x62, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x44, 0x62, 0x72, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x5f, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x73, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x6f, 0x62, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x5f, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x5f, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x61, 0x78, 0x5f, 0x65, 0x6d, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61,
	0x78, 0x45, 0x6d, 0x69, 0x12, 0x39, 0x0a, 0x06, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65,
	0x74, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x72, 0x65,
	0x61, 0x75, 0x48, 0x00, 0x52, 0x06, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x88, 0x01, 0x01, 0x12,
	0x57, 0x0a, 0x16, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x5f, 0x6f, 0x62, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x6f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x14, 0x62, 0x75, 0x72, 0x65, 0x61, 0x75, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x77, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x75, 0x72,
	0x65, 0x61, 0x75, 0x22, 0x3b, 0x0a, 0x06, 0x42, 0x75, 0x72, 0x65, 0x61, 0x75, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65,
	0x22, 0xd5, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x69,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x6d, 0x69, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x62, 0x6c, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x62,
	0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bureau_getobligation_proto_rawDescOnce sync.Once
	file_bureau_getobligation_proto_rawDescData = file_bureau_getobligation_proto_rawDesc
)

func file_bureau_getobligation_proto_rawDescGZIP() []byte {
	file_bureau_getobligation_proto_rawDescOnce.Do(func() {
		file_bureau_getobligation_proto_rawDescData = protoimpl.X.CompressGZIP(file_bureau_getobligation_proto_rawDescData)
	})
	return file_bureau_getobligation_proto_rawDescData
}

var file_bureau_getobligation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bureau_getobligation_proto_goTypes = []interface{}{
	(*ObligationRequest)(nil),  // 0: bureau.getobligation.obligationRequest
	(*ObligationResponse)(nil), // 1: bureau.getobligation.obligationResponse
	(*ObligationData)(nil),     // 2: bureau.getobligation.ObligationData
	(*Bureau)(nil),             // 3: bureau.getobligation.Bureau
	(*AccountList)(nil),        // 4: bureau.getobligation.AccountList
}
var file_bureau_getobligation_proto_depIdxs = []int32{
	2, // 0: bureau.getobligation.obligationResponse.data:type_name -> bureau.getobligation.ObligationData
	3, // 1: bureau.getobligation.ObligationData.bureau:type_name -> bureau.getobligation.Bureau
	4, // 2: bureau.getobligation.ObligationData.bureau_obligation_rows:type_name -> bureau.getobligation.AccountList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bureau_getobligation_proto_init() }
func file_bureau_getobligation_proto_init() {
	if File_bureau_getobligation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bureau_getobligation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObligationRequest); i {
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
		file_bureau_getobligation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObligationResponse); i {
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
		file_bureau_getobligation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObligationData); i {
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
		file_bureau_getobligation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bureau); i {
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
		file_bureau_getobligation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountList); i {
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
	file_bureau_getobligation_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_bureau_getobligation_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bureau_getobligation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bureau_getobligation_proto_goTypes,
		DependencyIndexes: file_bureau_getobligation_proto_depIdxs,
		MessageInfos:      file_bureau_getobligation_proto_msgTypes,
	}.Build()
	File_bureau_getobligation_proto = out.File
	file_bureau_getobligation_proto_rawDesc = nil
	file_bureau_getobligation_proto_goTypes = nil
	file_bureau_getobligation_proto_depIdxs = nil
}
