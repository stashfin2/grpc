// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: upi/gettransactions.proto

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

type GetTransactionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxnStatus     []string               `protobuf:"bytes,1,rep,name=txn_status,json=txnStatus,proto3" json:"txn_status,omitempty"`
	TxnType       []string               `protobuf:"bytes,2,rep,name=txn_type,json=txnType,proto3" json:"txn_type,omitempty"`
	Limit         *int32                 `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	OffSet        *int32                 `protobuf:"varint,4,opt,name=off_set,json=offSet,proto3,oneof" json:"off_set,omitempty"`
	SearchBy      *string                `protobuf:"bytes,5,opt,name=search_by,json=searchBy,proto3,oneof" json:"search_by,omitempty"`
	Months        *string                `protobuf:"bytes,6,opt,name=months,proto3,oneof" json:"months,omitempty"`
	ExternalTxnId *string                `protobuf:"bytes,7,opt,name=external_txn_id,json=externalTxnId,proto3,oneof" json:"external_txn_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionsRequest) Reset() {
	*x = GetTransactionsRequest{}
	mi := &file_upi_gettransactions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsRequest) ProtoMessage() {}

func (x *GetTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upi_gettransactions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_upi_gettransactions_proto_rawDescGZIP(), []int{0}
}

func (x *GetTransactionsRequest) GetTxnStatus() []string {
	if x != nil {
		return x.TxnStatus
	}
	return nil
}

func (x *GetTransactionsRequest) GetTxnType() []string {
	if x != nil {
		return x.TxnType
	}
	return nil
}

func (x *GetTransactionsRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetTransactionsRequest) GetOffSet() int32 {
	if x != nil && x.OffSet != nil {
		return *x.OffSet
	}
	return 0
}

func (x *GetTransactionsRequest) GetSearchBy() string {
	if x != nil && x.SearchBy != nil {
		return *x.SearchBy
	}
	return ""
}

func (x *GetTransactionsRequest) GetMonths() string {
	if x != nil && x.Months != nil {
		return *x.Months
	}
	return ""
}

func (x *GetTransactionsRequest) GetExternalTxnId() string {
	if x != nil && x.ExternalTxnId != nil {
		return *x.ExternalTxnId
	}
	return ""
}

type Transactions struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ExternalTxnId    *string                `protobuf:"bytes,1,opt,name=external_txn_id,json=externalTxnId,proto3,oneof" json:"external_txn_id,omitempty"`           // External transaction ID
	TxnType          *string                `protobuf:"bytes,2,opt,name=txn_type,json=txnType,proto3,oneof" json:"txn_type,omitempty"`                               // Type of transaction
	Note             *string                `protobuf:"bytes,3,opt,name=note,proto3,oneof" json:"note,omitempty"`                                                    // Note for the transaction
	Mobile           *string                `protobuf:"bytes,4,opt,name=mobile,proto3,oneof" json:"mobile,omitempty"`                                                // Mobile number
	Amount           *float64               `protobuf:"fixed64,5,opt,name=amount,proto3,oneof" json:"amount,omitempty"`                                              // Transaction amount
	TxnStatus        *string                `protobuf:"bytes,6,opt,name=txn_status,json=txnStatus,proto3,oneof" json:"txn_status,omitempty"`                         // Status of the transaction
	Rrn              *string                `protobuf:"bytes,7,opt,name=rrn,proto3,oneof" json:"rrn,omitempty"`                                                      // Retrieval reference number (RRN)
	PaidAt           *string                `protobuf:"bytes,8,opt,name=paid_at,json=paidAt,proto3,oneof" json:"paid_at,omitempty"`                                  // Payment timestamp in ISO format
	PayerVpa         *string                `protobuf:"bytes,9,opt,name=payer_vpa,json=payerVpa,proto3,oneof" json:"payer_vpa,omitempty"`                            // Payer's VPA (Virtual Payment Address)
	PayeeVpa         *string                `protobuf:"bytes,10,opt,name=payee_vpa,json=payeeVpa,proto3,oneof" json:"payee_vpa,omitempty"`                           // Payee's VPA (Virtual Payment Address)
	PayeeName        *string                `protobuf:"bytes,11,opt,name=payee_name,json=payeeName,proto3,oneof" json:"payee_name,omitempty"`                        // Name of the payee
	PayeeAccNumber   *string                `protobuf:"bytes,12,opt,name=payee_acc_number,json=payeeAccNumber,proto3,oneof" json:"payee_acc_number,omitempty"`       // Payee's account number
	PayeeIfsc        *string                `protobuf:"bytes,13,opt,name=payee_ifsc,json=payeeIfsc,proto3,oneof" json:"payee_ifsc,omitempty"`                        // Payee's IFSC code
	PayeeAccProvider *string                `protobuf:"bytes,14,opt,name=payee_acc_provider,json=payeeAccProvider,proto3,oneof" json:"payee_acc_provider,omitempty"` // Name of the payee's account provider
	Umn              *string                `protobuf:"bytes,15,opt,name=umn,proto3,oneof" json:"umn,omitempty"`
	MandateTxnSeqNo  *int32                 `protobuf:"varint,17,opt,name=mandate_txn_seq_no,json=mandateTxnSeqNo,proto3,oneof" json:"mandate_txn_seq_no,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Transactions) Reset() {
	*x = Transactions{}
	mi := &file_upi_gettransactions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transactions) ProtoMessage() {}

func (x *Transactions) ProtoReflect() protoreflect.Message {
	mi := &file_upi_gettransactions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transactions.ProtoReflect.Descriptor instead.
func (*Transactions) Descriptor() ([]byte, []int) {
	return file_upi_gettransactions_proto_rawDescGZIP(), []int{1}
}

func (x *Transactions) GetExternalTxnId() string {
	if x != nil && x.ExternalTxnId != nil {
		return *x.ExternalTxnId
	}
	return ""
}

func (x *Transactions) GetTxnType() string {
	if x != nil && x.TxnType != nil {
		return *x.TxnType
	}
	return ""
}

func (x *Transactions) GetNote() string {
	if x != nil && x.Note != nil {
		return *x.Note
	}
	return ""
}

func (x *Transactions) GetMobile() string {
	if x != nil && x.Mobile != nil {
		return *x.Mobile
	}
	return ""
}

func (x *Transactions) GetAmount() float64 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *Transactions) GetTxnStatus() string {
	if x != nil && x.TxnStatus != nil {
		return *x.TxnStatus
	}
	return ""
}

func (x *Transactions) GetRrn() string {
	if x != nil && x.Rrn != nil {
		return *x.Rrn
	}
	return ""
}

func (x *Transactions) GetPaidAt() string {
	if x != nil && x.PaidAt != nil {
		return *x.PaidAt
	}
	return ""
}

func (x *Transactions) GetPayerVpa() string {
	if x != nil && x.PayerVpa != nil {
		return *x.PayerVpa
	}
	return ""
}

func (x *Transactions) GetPayeeVpa() string {
	if x != nil && x.PayeeVpa != nil {
		return *x.PayeeVpa
	}
	return ""
}

func (x *Transactions) GetPayeeName() string {
	if x != nil && x.PayeeName != nil {
		return *x.PayeeName
	}
	return ""
}

func (x *Transactions) GetPayeeAccNumber() string {
	if x != nil && x.PayeeAccNumber != nil {
		return *x.PayeeAccNumber
	}
	return ""
}

func (x *Transactions) GetPayeeIfsc() string {
	if x != nil && x.PayeeIfsc != nil {
		return *x.PayeeIfsc
	}
	return ""
}

func (x *Transactions) GetPayeeAccProvider() string {
	if x != nil && x.PayeeAccProvider != nil {
		return *x.PayeeAccProvider
	}
	return ""
}

func (x *Transactions) GetUmn() string {
	if x != nil && x.Umn != nil {
		return *x.Umn
	}
	return ""
}

func (x *Transactions) GetMandateTxnSeqNo() int32 {
	if x != nil && x.MandateTxnSeqNo != nil {
		return *x.MandateTxnSeqNo
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transactions  []*Transactions        `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Count         *int32                 `protobuf:"varint,2,opt,name=count,proto3,oneof" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_upi_gettransactions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_upi_gettransactions_proto_msgTypes[2]
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
	return file_upi_gettransactions_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetTransactions() []*Transactions {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Data) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type GetTransactionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *Data                  `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionsResponse) Reset() {
	*x = GetTransactionsResponse{}
	mi := &file_upi_gettransactions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsResponse) ProtoMessage() {}

func (x *GetTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upi_gettransactions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_upi_gettransactions_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransactionsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetTransactionsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTransactionsResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_upi_gettransactions_proto protoreflect.FileDescriptor

var file_upi_gettransactions_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x75, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x75, 0x70, 0x69,
	0x2e, 0x67, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xba, 0x02, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x78, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x78, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1c, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x53, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x06, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a,
	0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x78, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x74,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x62, 0x79, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0xa6, 0x06,
	0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b,
	0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x78, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x74,
	0x78, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x07, 0x74, 0x78, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x05, 0x52, 0x09, 0x74, 0x78, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x72, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x06, 0x52, 0x03, 0x72, 0x72, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x70, 0x61, 0x69,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x06, 0x70, 0x61,
	0x69, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x76, 0x70, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x08, 0x70, 0x61,
	0x79, 0x65, 0x72, 0x56, 0x70, 0x61, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x79,
	0x65, 0x65, 0x5f, 0x76, 0x70, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x79, 0x65, 0x65, 0x56, 0x70, 0x61, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x70,
	0x61, 0x79, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x0a, 0x52, 0x09, 0x70, 0x61, 0x79, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2d, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x0e, 0x70, 0x61, 0x79,
	0x65, 0x65, 0x41, 0x63, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x66, 0x73, 0x63, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x79, 0x65, 0x65, 0x49, 0x66, 0x73, 0x63, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0d,
	0x52, 0x10, 0x70, 0x61, 0x79, 0x65, 0x65, 0x41, 0x63, 0x63, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x6d, 0x6e, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0e, 0x52, 0x03, 0x75, 0x6d, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x12,
	0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x65, 0x71, 0x5f,
	0x6e, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0f, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x78, 0x6e, 0x53, 0x65, 0x71, 0x4e, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x78, 0x6e, 0x5f,
	0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x72, 0x72, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x61,
	0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x76, 0x70, 0x61, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x76, 0x70, 0x61, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x13, 0x0a, 0x11,
	0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x66, 0x73, 0x63,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x6d, 0x6e, 0x42,
	0x15, 0x0a, 0x13, 0x5f, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x78, 0x6e, 0x5f,
	0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x22, 0x72, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x45,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x17, 0x67,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x67, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_upi_gettransactions_proto_rawDescOnce sync.Once
	file_upi_gettransactions_proto_rawDescData []byte
)

func file_upi_gettransactions_proto_rawDescGZIP() []byte {
	file_upi_gettransactions_proto_rawDescOnce.Do(func() {
		file_upi_gettransactions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upi_gettransactions_proto_rawDesc), len(file_upi_gettransactions_proto_rawDesc)))
	})
	return file_upi_gettransactions_proto_rawDescData
}

var file_upi_gettransactions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_upi_gettransactions_proto_goTypes = []any{
	(*GetTransactionsRequest)(nil),  // 0: upi.getTransactions.getTransactionsRequest
	(*Transactions)(nil),            // 1: upi.getTransactions.Transactions
	(*Data)(nil),                    // 2: upi.getTransactions.Data
	(*GetTransactionsResponse)(nil), // 3: upi.getTransactions.getTransactionsResponse
}
var file_upi_gettransactions_proto_depIdxs = []int32{
	1, // 0: upi.getTransactions.Data.transactions:type_name -> upi.getTransactions.Transactions
	2, // 1: upi.getTransactions.getTransactionsResponse.data:type_name -> upi.getTransactions.Data
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upi_gettransactions_proto_init() }
func file_upi_gettransactions_proto_init() {
	if File_upi_gettransactions_proto != nil {
		return
	}
	file_upi_gettransactions_proto_msgTypes[0].OneofWrappers = []any{}
	file_upi_gettransactions_proto_msgTypes[1].OneofWrappers = []any{}
	file_upi_gettransactions_proto_msgTypes[2].OneofWrappers = []any{}
	file_upi_gettransactions_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upi_gettransactions_proto_rawDesc), len(file_upi_gettransactions_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upi_gettransactions_proto_goTypes,
		DependencyIndexes: file_upi_gettransactions_proto_depIdxs,
		MessageInfos:      file_upi_gettransactions_proto_msgTypes,
	}.Build()
	File_upi_gettransactions_proto = out.File
	file_upi_gettransactions_proto_goTypes = nil
	file_upi_gettransactions_proto_depIdxs = nil
}
