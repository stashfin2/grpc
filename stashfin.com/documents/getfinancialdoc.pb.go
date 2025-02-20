// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: documents/getfinancialdoc.proto

package documents

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

type Financialdocrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId         int32                        `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ApprovedDate       string                       `protobuf:"bytes,2,opt,name=approved_date,json=approvedDate,proto3" json:"approved_date,omitempty"`
	Tenure             int32                        `protobuf:"varint,3,opt,name=tenure,proto3" json:"tenure,omitempty"`
	ProcessingFees     float64                      `protobuf:"fixed64,4,opt,name=processing_fees,json=processingFees,proto3" json:"processing_fees,omitempty"`
	CreditReportFees   float64                      `protobuf:"fixed64,5,opt,name=credit_report_fees,json=creditReportFees,proto3" json:"credit_report_fees,omitempty"`
	InsuranceFees      float64                      `protobuf:"fixed64,6,opt,name=insurance_fees,json=insuranceFees,proto3" json:"insurance_fees,omitempty"`
	DisburedAmount     float64                      `protobuf:"fixed64,7,opt,name=disbured_amount,json=disburedAmount,proto3" json:"disbured_amount,omitempty"`
	SanctionedAmount   float64                      `protobuf:"fixed64,8,opt,name=sanctioned_amount,json=sanctionedAmount,proto3" json:"sanctioned_amount,omitempty"`
	AprValue           float64                      `protobuf:"fixed64,9,opt,name=apr_value,json=aprValue,proto3" json:"apr_value,omitempty"`
	Roi                float64                      `protobuf:"fixed64,10,opt,name=roi,proto3" json:"roi,omitempty"`
	TxnFees            float64                      `protobuf:"fixed64,11,opt,name=txn_fees,json=txnFees,proto3" json:"txn_fees,omitempty"`
	TxnFeesGst         float64                      `protobuf:"fixed64,12,opt,name=txn_fees_gst,json=txnFeesGst,proto3" json:"txn_fees_gst,omitempty"`
	TxnRate            float64                      `protobuf:"fixed64,13,opt,name=txn_rate,json=txnRate,proto3" json:"txn_rate,omitempty"`
	ProcessingFeesRate float64                      `protobuf:"fixed64,14,opt,name=processing_fees_rate,json=processingFeesRate,proto3" json:"processing_fees_rate,omitempty"`
	ProcessingFeesGst  float64                      `protobuf:"fixed64,15,opt,name=processing_fees_gst,json=processingFeesGst,proto3" json:"processing_fees_gst,omitempty"`
	EmiStartDate       string                       `protobuf:"bytes,16,opt,name=emi_start_date,json=emiStartDate,proto3" json:"emi_start_date,omitempty"`
	Installments       []*Financialdocrequest_Field `protobuf:"bytes,17,rep,name=installments,proto3" json:"installments,omitempty"`
}

func (x *Financialdocrequest) Reset() {
	*x = Financialdocrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_getfinancialdoc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Financialdocrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Financialdocrequest) ProtoMessage() {}

func (x *Financialdocrequest) ProtoReflect() protoreflect.Message {
	mi := &file_documents_getfinancialdoc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Financialdocrequest.ProtoReflect.Descriptor instead.
func (*Financialdocrequest) Descriptor() ([]byte, []int) {
	return file_documents_getfinancialdoc_proto_rawDescGZIP(), []int{0}
}

func (x *Financialdocrequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Financialdocrequest) GetApprovedDate() string {
	if x != nil {
		return x.ApprovedDate
	}
	return ""
}

func (x *Financialdocrequest) GetTenure() int32 {
	if x != nil {
		return x.Tenure
	}
	return 0
}

func (x *Financialdocrequest) GetProcessingFees() float64 {
	if x != nil {
		return x.ProcessingFees
	}
	return 0
}

func (x *Financialdocrequest) GetCreditReportFees() float64 {
	if x != nil {
		return x.CreditReportFees
	}
	return 0
}

func (x *Financialdocrequest) GetInsuranceFees() float64 {
	if x != nil {
		return x.InsuranceFees
	}
	return 0
}

func (x *Financialdocrequest) GetDisburedAmount() float64 {
	if x != nil {
		return x.DisburedAmount
	}
	return 0
}

func (x *Financialdocrequest) GetSanctionedAmount() float64 {
	if x != nil {
		return x.SanctionedAmount
	}
	return 0
}

func (x *Financialdocrequest) GetAprValue() float64 {
	if x != nil {
		return x.AprValue
	}
	return 0
}

func (x *Financialdocrequest) GetRoi() float64 {
	if x != nil {
		return x.Roi
	}
	return 0
}

func (x *Financialdocrequest) GetTxnFees() float64 {
	if x != nil {
		return x.TxnFees
	}
	return 0
}

func (x *Financialdocrequest) GetTxnFeesGst() float64 {
	if x != nil {
		return x.TxnFeesGst
	}
	return 0
}

func (x *Financialdocrequest) GetTxnRate() float64 {
	if x != nil {
		return x.TxnRate
	}
	return 0
}

func (x *Financialdocrequest) GetProcessingFeesRate() float64 {
	if x != nil {
		return x.ProcessingFeesRate
	}
	return 0
}

func (x *Financialdocrequest) GetProcessingFeesGst() float64 {
	if x != nil {
		return x.ProcessingFeesGst
	}
	return 0
}

func (x *Financialdocrequest) GetEmiStartDate() string {
	if x != nil {
		return x.EmiStartDate
	}
	return ""
}

func (x *Financialdocrequest) GetInstallments() []*Financialdocrequest_Field {
	if x != nil {
		return x.Installments
	}
	return nil
}

type FinancialAgreementdata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinancialDocUrl   string `protobuf:"bytes,1,opt,name=financial_doc_url,json=financialDocUrl,proto3" json:"financial_doc_url,omitempty"`
	SanctionLatterUrl string `protobuf:"bytes,2,opt,name=sanction_latter_url,json=sanctionLatterUrl,proto3" json:"sanction_latter_url,omitempty"`
}

func (x *FinancialAgreementdata) Reset() {
	*x = FinancialAgreementdata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_getfinancialdoc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinancialAgreementdata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinancialAgreementdata) ProtoMessage() {}

func (x *FinancialAgreementdata) ProtoReflect() protoreflect.Message {
	mi := &file_documents_getfinancialdoc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinancialAgreementdata.ProtoReflect.Descriptor instead.
func (*FinancialAgreementdata) Descriptor() ([]byte, []int) {
	return file_documents_getfinancialdoc_proto_rawDescGZIP(), []int{1}
}

func (x *FinancialAgreementdata) GetFinancialDocUrl() string {
	if x != nil {
		return x.FinancialDocUrl
	}
	return ""
}

func (x *FinancialAgreementdata) GetSanctionLatterUrl() string {
	if x != nil {
		return x.SanctionLatterUrl
	}
	return ""
}

type Financialdocresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string                  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32                   `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Data       *FinancialAgreementdata `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Message    string                  `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Financialdocresponse) Reset() {
	*x = Financialdocresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_getfinancialdoc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Financialdocresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Financialdocresponse) ProtoMessage() {}

func (x *Financialdocresponse) ProtoReflect() protoreflect.Message {
	mi := &file_documents_getfinancialdoc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Financialdocresponse.ProtoReflect.Descriptor instead.
func (*Financialdocresponse) Descriptor() ([]byte, []int) {
	return file_documents_getfinancialdoc_proto_rawDescGZIP(), []int{2}
}

func (x *Financialdocresponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Financialdocresponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Financialdocresponse) GetData() *FinancialAgreementdata {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Financialdocresponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Financialdocrequest_Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      int32  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	SeqNo           int32  `protobuf:"varint,2,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	Amount          int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	StartingBalance int32  `protobuf:"varint,4,opt,name=starting_balance,json=startingBalance,proto3" json:"starting_balance,omitempty"`
	EndingBalance   int32  `protobuf:"varint,5,opt,name=ending_balance,json=endingBalance,proto3" json:"ending_balance,omitempty"`
	Principal       int32  `protobuf:"varint,6,opt,name=principal,proto3" json:"principal,omitempty"`
	Interest        int32  `protobuf:"varint,7,opt,name=interest,proto3" json:"interest,omitempty"`
	DueDate         string `protobuf:"bytes,8,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	Status          string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Discount        int32  `protobuf:"varint,10,opt,name=discount,proto3" json:"discount,omitempty"`
	Penalty         int32  `protobuf:"varint,11,opt,name=penalty,proto3" json:"penalty,omitempty"`
}

func (x *Financialdocrequest_Field) Reset() {
	*x = Financialdocrequest_Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_getfinancialdoc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Financialdocrequest_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Financialdocrequest_Field) ProtoMessage() {}

func (x *Financialdocrequest_Field) ProtoReflect() protoreflect.Message {
	mi := &file_documents_getfinancialdoc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Financialdocrequest_Field.ProtoReflect.Descriptor instead.
func (*Financialdocrequest_Field) Descriptor() ([]byte, []int) {
	return file_documents_getfinancialdoc_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Financialdocrequest_Field) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Financialdocrequest_Field) GetSeqNo() int32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *Financialdocrequest_Field) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Financialdocrequest_Field) GetStartingBalance() int32 {
	if x != nil {
		return x.StartingBalance
	}
	return 0
}

func (x *Financialdocrequest_Field) GetEndingBalance() int32 {
	if x != nil {
		return x.EndingBalance
	}
	return 0
}

func (x *Financialdocrequest_Field) GetPrincipal() int32 {
	if x != nil {
		return x.Principal
	}
	return 0
}

func (x *Financialdocrequest_Field) GetInterest() int32 {
	if x != nil {
		return x.Interest
	}
	return 0
}

func (x *Financialdocrequest_Field) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

func (x *Financialdocrequest_Field) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Financialdocrequest_Field) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Financialdocrequest_Field) GetPenalty() int32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

var File_documents_getfinancialdoc_proto protoreflect.FileDescriptor

var file_documents_getfinancialdoc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x66,
	0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x64, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x65, 0x74,
	0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x64, 0x6f, 0x63, 0x22, 0xff, 0x07, 0x0a,
	0x13, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x64, 0x6f, 0x63, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x75,
	0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x65, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x65, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73,
	0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x65, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x62, 0x75,
	0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x61, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x72, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x70, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x72, 0x6f, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e, 0x5f, 0x66, 0x65, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x46, 0x65, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0c, 0x74, 0x78, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x5f, 0x67, 0x73, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x78, 0x6e, 0x46, 0x65, 0x65, 0x73, 0x47,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x73,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65,
	0x65, 0x73, 0x5f, 0x67, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x73, 0x47, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0e, 0x65, 0x6d, 0x69, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x69, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x58, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x69, 0x61, 0x6c, 0x64, 0x6f, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61,
	0x6c, 0x64, 0x6f, 0x63, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0xcc, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65,
	0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x22, 0x74,
	0x0a, 0x16, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x6f,
	0x63, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6c, 0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x73, 0x61, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x22, 0xaf, 0x01, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69,
	0x61, 0x6c, 0x64, 0x6f, 0x63, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x67, 0x65, 0x74, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x64, 0x6f, 0x63, 0x2e,
	0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_documents_getfinancialdoc_proto_rawDescOnce sync.Once
	file_documents_getfinancialdoc_proto_rawDescData = file_documents_getfinancialdoc_proto_rawDesc
)

func file_documents_getfinancialdoc_proto_rawDescGZIP() []byte {
	file_documents_getfinancialdoc_proto_rawDescOnce.Do(func() {
		file_documents_getfinancialdoc_proto_rawDescData = protoimpl.X.CompressGZIP(file_documents_getfinancialdoc_proto_rawDescData)
	})
	return file_documents_getfinancialdoc_proto_rawDescData
}

var file_documents_getfinancialdoc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_documents_getfinancialdoc_proto_goTypes = []interface{}{
	(*Financialdocrequest)(nil),       // 0: documents.getfinancialdoc.financialdocrequest
	(*FinancialAgreementdata)(nil),    // 1: documents.getfinancialdoc.financialAgreementdata
	(*Financialdocresponse)(nil),      // 2: documents.getfinancialdoc.financialdocresponse
	(*Financialdocrequest_Field)(nil), // 3: documents.getfinancialdoc.financialdocrequest.Field
}
var file_documents_getfinancialdoc_proto_depIdxs = []int32{
	3, // 0: documents.getfinancialdoc.financialdocrequest.installments:type_name -> documents.getfinancialdoc.financialdocrequest.Field
	1, // 1: documents.getfinancialdoc.financialdocresponse.data:type_name -> documents.getfinancialdoc.financialAgreementdata
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_documents_getfinancialdoc_proto_init() }
func file_documents_getfinancialdoc_proto_init() {
	if File_documents_getfinancialdoc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_documents_getfinancialdoc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Financialdocrequest); i {
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
		file_documents_getfinancialdoc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinancialAgreementdata); i {
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
		file_documents_getfinancialdoc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Financialdocresponse); i {
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
		file_documents_getfinancialdoc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Financialdocrequest_Field); i {
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
			RawDescriptor: file_documents_getfinancialdoc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_documents_getfinancialdoc_proto_goTypes,
		DependencyIndexes: file_documents_getfinancialdoc_proto_depIdxs,
		MessageInfos:      file_documents_getfinancialdoc_proto_msgTypes,
	}.Build()
	File_documents_getfinancialdoc_proto = out.File
	file_documents_getfinancialdoc_proto_rawDesc = nil
	file_documents_getfinancialdoc_proto_goTypes = nil
	file_documents_getfinancialdoc_proto_depIdxs = nil
}
