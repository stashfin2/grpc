// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: payments/fundtransfer.proto

package payments

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

type FundTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId          int32                             `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ClientId            string                            `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ExternalOrderId     string                            `protobuf:"bytes,3,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	Amount              float64                           `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	TransferSource      string                            `protobuf:"bytes,5,opt,name=transfer_source,json=transferSource,proto3" json:"transfer_source,omitempty"`
	TransferAccountId   string                            `protobuf:"bytes,6,opt,name=transfer_account_id,json=transferAccountId,proto3" json:"transfer_account_id,omitempty"`
	CustomerBankDetails []*FundTransferRequestBankDetails `protobuf:"bytes,7,rep,name=customer_bank_details,json=customerBankDetails,proto3" json:"customer_bank_details,omitempty"`
	LoanId              *string                           `protobuf:"bytes,8,opt,name=loan_id,json=loanId,proto3,oneof" json:"loan_id,omitempty"`
	Mode                string                            `protobuf:"bytes,9,opt,name=mode,proto3" json:"mode,omitempty"`
	TransferDestination *string                           `protobuf:"bytes,10,opt,name=transfer_destination,json=transferDestination,proto3,oneof" json:"transfer_destination,omitempty"`
}

func (x *FundTransferRequest) Reset() {
	*x = FundTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_fundtransfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundTransferRequest) ProtoMessage() {}

func (x *FundTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payments_fundtransfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundTransferRequest.ProtoReflect.Descriptor instead.
func (*FundTransferRequest) Descriptor() ([]byte, []int) {
	return file_payments_fundtransfer_proto_rawDescGZIP(), []int{0}
}

func (x *FundTransferRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *FundTransferRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *FundTransferRequest) GetExternalOrderId() string {
	if x != nil {
		return x.ExternalOrderId
	}
	return ""
}

func (x *FundTransferRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FundTransferRequest) GetTransferSource() string {
	if x != nil {
		return x.TransferSource
	}
	return ""
}

func (x *FundTransferRequest) GetTransferAccountId() string {
	if x != nil {
		return x.TransferAccountId
	}
	return ""
}

func (x *FundTransferRequest) GetCustomerBankDetails() []*FundTransferRequestBankDetails {
	if x != nil {
		return x.CustomerBankDetails
	}
	return nil
}

func (x *FundTransferRequest) GetLoanId() string {
	if x != nil && x.LoanId != nil {
		return *x.LoanId
	}
	return ""
}

func (x *FundTransferRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *FundTransferRequest) GetTransferDestination() string {
	if x != nil && x.TransferDestination != nil {
		return *x.TransferDestination
	}
	return ""
}

type FundTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode      int32   `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	CustomerId      int32   `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ExternalOrderId string  `protobuf:"bytes,4,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	Amount          float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Message         *string `protobuf:"bytes,6,opt,name=message,proto3,oneof" json:"message,omitempty"`
	TxnDate         *string `protobuf:"bytes,7,opt,name=txn_date,json=txnDate,proto3,oneof" json:"txn_date,omitempty"`
}

func (x *FundTransferResponse) Reset() {
	*x = FundTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_fundtransfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundTransferResponse) ProtoMessage() {}

func (x *FundTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payments_fundtransfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundTransferResponse.ProtoReflect.Descriptor instead.
func (*FundTransferResponse) Descriptor() ([]byte, []int) {
	return file_payments_fundtransfer_proto_rawDescGZIP(), []int{1}
}

func (x *FundTransferResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FundTransferResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FundTransferResponse) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *FundTransferResponse) GetExternalOrderId() string {
	if x != nil {
		return x.ExternalOrderId
	}
	return ""
}

func (x *FundTransferResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FundTransferResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *FundTransferResponse) GetTxnDate() string {
	if x != nil && x.TxnDate != nil {
		return *x.TxnDate
	}
	return ""
}

type FundTransferRequestBankDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNo string `protobuf:"bytes,1,opt,name=account_no,json=accountNo,proto3" json:"account_no,omitempty"`
	Ifsc      string `protobuf:"bytes,2,opt,name=ifsc,proto3" json:"ifsc,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *FundTransferRequestBankDetails) Reset() {
	*x = FundTransferRequestBankDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_fundtransfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundTransferRequestBankDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundTransferRequestBankDetails) ProtoMessage() {}

func (x *FundTransferRequestBankDetails) ProtoReflect() protoreflect.Message {
	mi := &file_payments_fundtransfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundTransferRequestBankDetails.ProtoReflect.Descriptor instead.
func (*FundTransferRequestBankDetails) Descriptor() ([]byte, []int) {
	return file_payments_fundtransfer_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FundTransferRequestBankDetails) GetAccountNo() string {
	if x != nil {
		return x.AccountNo
	}
	return ""
}

func (x *FundTransferRequestBankDetails) GetIfsc() string {
	if x != nil {
		return x.Ifsc
	}
	return ""
}

func (x *FundTransferRequestBankDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FundTransferRequestBankDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FundTransferRequestBankDetails) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_payments_fundtransfer_proto protoreflect.FileDescriptor

var file_payments_fundtransfer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x22, 0xf0, 0x04, 0x0a, 0x13, 0x66, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x6b, 0x0a, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x66,
	0x75, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x1a, 0x81,
	0x01, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x66, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x66,
	0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x17,
	0x0a, 0x15, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x02, 0x0a, 0x14, 0x66, 0x75, 0x6e, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x74, 0x78, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x78,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payments_fundtransfer_proto_rawDescOnce sync.Once
	file_payments_fundtransfer_proto_rawDescData = file_payments_fundtransfer_proto_rawDesc
)

func file_payments_fundtransfer_proto_rawDescGZIP() []byte {
	file_payments_fundtransfer_proto_rawDescOnce.Do(func() {
		file_payments_fundtransfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_payments_fundtransfer_proto_rawDescData)
	})
	return file_payments_fundtransfer_proto_rawDescData
}

var file_payments_fundtransfer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payments_fundtransfer_proto_goTypes = []interface{}{
	(*FundTransferRequest)(nil),            // 0: payments.fundtransfer.fundTransferRequest
	(*FundTransferResponse)(nil),           // 1: payments.fundtransfer.fundTransferResponse
	(*FundTransferRequestBankDetails)(nil), // 2: payments.fundtransfer.fundTransferRequest.bank_details
}
var file_payments_fundtransfer_proto_depIdxs = []int32{
	2, // 0: payments.fundtransfer.fundTransferRequest.customer_bank_details:type_name -> payments.fundtransfer.fundTransferRequest.bank_details
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payments_fundtransfer_proto_init() }
func file_payments_fundtransfer_proto_init() {
	if File_payments_fundtransfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payments_fundtransfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundTransferRequest); i {
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
		file_payments_fundtransfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundTransferResponse); i {
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
		file_payments_fundtransfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundTransferRequestBankDetails); i {
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
	file_payments_fundtransfer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_payments_fundtransfer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payments_fundtransfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payments_fundtransfer_proto_goTypes,
		DependencyIndexes: file_payments_fundtransfer_proto_depIdxs,
		MessageInfos:      file_payments_fundtransfer_proto_msgTypes,
	}.Build()
	File_payments_fundtransfer_proto = out.File
	file_payments_fundtransfer_proto_rawDesc = nil
	file_payments_fundtransfer_proto_goTypes = nil
	file_payments_fundtransfer_proto_depIdxs = nil
}
