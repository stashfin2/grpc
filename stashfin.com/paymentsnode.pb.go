// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: paymentsnode.proto

package stashfin_com

import (
	payments "github.com/stashfin2/grpc/stashfin.com/payments"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_paymentsnode_proto protoreflect.FileDescriptor

var file_paymentsnode_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x20, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6e, 0x61, 0x63,
	0x68, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6e, 0x61,
	0x63, 0x68, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66,
	0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x80, 0x06, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x6e, 0x6f,
	0x64, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x87, 0x01, 0x0a, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x10, 0x6e,
	0x61, 0x63, 0x68, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6e, 0x61, 0x63, 0x68, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x61, 0x63, 0x68,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6e,
	0x61, 0x63, 0x68, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6e, 0x61, 0x63, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x10, 0x6e, 0x61,
	0x63, 0x68, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6e, 0x61, 0x63, 0x68, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x61, 0x63, 0x68, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6e, 0x61,
	0x63, 0x68, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e,
	0x61, 0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0c, 0x66, 0x75, 0x6e,
	0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x66, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x66, 0x75,
	0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_paymentsnode_proto_goTypes = []interface{}{
	(*payments.GetPaymentOptionsRequest)(nil),   // 0: payments.getpaymentoptions.getPaymentOptionsRequest
	(*payments.InitiatePaymentRequest)(nil),     // 1: payments.initiatepayment.initiatePaymentRequest
	(*payments.CheckPaymentStatusRequest)(nil),  // 2: payments.checkpaymentstatus.checkPaymentStatusRequest
	(*payments.NachRegistrationRequest)(nil),    // 3: payments.nachregistration.nachRegistrationRequest
	(*payments.NachPresentationRequest)(nil),    // 4: payments.nachpresentation.nachPresentationRequest
	(*payments.FundTransferRequest)(nil),        // 5: payments.fundtransfer.fundTransferRequest
	(*payments.GetPaymentOptionsResponse)(nil),  // 6: payments.getpaymentoptions.getPaymentOptionsResponse
	(*payments.InitiatePaymentResponse)(nil),    // 7: payments.initiatepayment.initiatePaymentResponse
	(*payments.CheckPaymentStatusResponse)(nil), // 8: payments.checkpaymentstatus.checkPaymentStatusResponse
	(*payments.NachRegistrationResponse)(nil),   // 9: payments.nachregistration.nachRegistrationResponse
	(*payments.NachPresentationResponse)(nil),   // 10: payments.nachpresentation.nachPresentationResponse
	(*payments.FundTransferResponse)(nil),       // 11: payments.fundtransfer.fundTransferResponse
}
var file_paymentsnode_proto_depIdxs = []int32{
	0,  // 0: service.paymentsnode.getpaymentoptions:input_type -> payments.getpaymentoptions.getPaymentOptionsRequest
	1,  // 1: service.paymentsnode.initiatepayment:input_type -> payments.initiatepayment.initiatePaymentRequest
	2,  // 2: service.paymentsnode.checkpaymentstatus:input_type -> payments.checkpaymentstatus.checkPaymentStatusRequest
	3,  // 3: service.paymentsnode.nachregistration:input_type -> payments.nachregistration.nachRegistrationRequest
	4,  // 4: service.paymentsnode.nachpresentation:input_type -> payments.nachpresentation.nachPresentationRequest
	5,  // 5: service.paymentsnode.fundtransfer:input_type -> payments.fundtransfer.fundTransferRequest
	6,  // 6: service.paymentsnode.getpaymentoptions:output_type -> payments.getpaymentoptions.getPaymentOptionsResponse
	7,  // 7: service.paymentsnode.initiatepayment:output_type -> payments.initiatepayment.initiatePaymentResponse
	8,  // 8: service.paymentsnode.checkpaymentstatus:output_type -> payments.checkpaymentstatus.checkPaymentStatusResponse
	9,  // 9: service.paymentsnode.nachregistration:output_type -> payments.nachregistration.nachRegistrationResponse
	10, // 10: service.paymentsnode.nachpresentation:output_type -> payments.nachpresentation.nachPresentationResponse
	11, // 11: service.paymentsnode.fundtransfer:output_type -> payments.fundtransfer.fundTransferResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_paymentsnode_proto_init() }
func file_paymentsnode_proto_init() {
	if File_paymentsnode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_paymentsnode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paymentsnode_proto_goTypes,
		DependencyIndexes: file_paymentsnode_proto_depIdxs,
	}.Build()
	File_paymentsnode_proto = out.File
	file_paymentsnode_proto_rawDesc = nil
	file_paymentsnode_proto_goTypes = nil
	file_paymentsnode_proto_depIdxs = nil
}
