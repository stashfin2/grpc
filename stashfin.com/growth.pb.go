// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: growth.proto

package stashfin_com

import (
	growth "github.com/stashfin2/grpc/stashfin.com/growth"
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

var File_growth_proto protoreflect.FileDescriptor

var file_growth_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1d, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x63, 0x68, 0x72, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f, 0x67, 0x65, 0x74,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x70, 0x6c, 0x61, 0x6e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f,
	0x67, 0x65, 0x74, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2f, 0x67, 0x65, 0x74, 0x63,
	0x68, 0x72, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc8, 0x06, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68,
	0x12, 0x6b, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x63, 0x68, 0x72, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x63, 0x68, 0x72, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x63, 0x68, 0x72, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x63,
	0x68, 0x72, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x63, 0x68, 0x72, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a,
	0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2d, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6c, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61,
	0x6e, 0x73, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6c,
	0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61, 0x6e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x7f, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x70, 0x6c, 0x61,
	0x6e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x31, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74,
	0x68, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x70, 0x6c, 0x61, 0x6e, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x70, 0x6c, 0x61, 0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x72,
	0x6f, 0x77, 0x74, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x70, 0x6c,
	0x61, 0x6e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x70, 0x6c, 0x61, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5a, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61, 0x6e, 0x73,
	0x12, 0x23, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x68, 0x72,
	0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x67,
	0x65, 0x74, 0x63, 0x68, 0x72, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x70, 0x6c,
	0x61, 0x6e, 0x73, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01,
	0x0a, 0x15, 0x67, 0x65, 0x74, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x37, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68,
	0x2e, 0x67, 0x65, 0x74, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x0f,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x2e, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x7c, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x63, 0x68, 0x72, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x2e, 0x67, 0x72, 0x6f, 0x77,
	0x74, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x68, 0x72, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x72,
	0x6f, 0x77, 0x74, 0x68, 0x2e, 0x67, 0x65, 0x74, 0x63, 0x68, 0x72, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_growth_proto_goTypes = []interface{}{
	(*growth.Notifychrrequest)(nil),           // 0: growth.notifychrpayment.notifychrrequest
	(*growth.Chrplansstatusrequest)(nil),      // 1: growth.checkplanstatus.chrplansstatusrequest
	(*growth.Chrattachplanrequest)(nil),       // 2: growth.getattachplandetails.chrattachplanrequest
	(*growth.Chrplansrequest)(nil),            // 3: growth.getchrplans.chrplansrequest
	(*growth.Credithealthreportrequest)(nil),  // 4: growth.getcredithealthreport.credithealthreportrequest
	(*growth.Initiatepaymentrequest)(nil),     // 5: growth.initiatepayment.initiatepaymentrequest
	(*growth.Paymentstatusrequest)(nil),       // 6: growth.getchrpaymentstatus.paymentstatusrequest
	(*growth.Notifychrresponse)(nil),          // 7: growth.notifychrpayment.notifychrresponse
	(*growth.Chrplanstatusresponse)(nil),      // 8: growth.checkplanstatus.chrplanstatusresponse
	(*growth.Chrattachplanresponse)(nil),      // 9: growth.getattachplandetails.chrattachplanresponse
	(*growth.Chrplansresponse)(nil),           // 10: growth.getchrplans.chrplansresponse
	(*growth.Credithealthreportresponse)(nil), // 11: growth.getcredithealthreport.credithealthreportresponse
	(*growth.Initiatepaymentresponse)(nil),    // 12: growth.initiatepayment.initiatepaymentresponse
	(*growth.Paymentstatusresponse)(nil),      // 13: growth.getchrpaymentstatus.paymentstatusresponse
}
var file_growth_proto_depIdxs = []int32{
	0,  // 0: service.growth.notifychrpayment:input_type -> growth.notifychrpayment.notifychrrequest
	1,  // 1: service.growth.checkplanstatus:input_type -> growth.checkplanstatus.chrplansstatusrequest
	2,  // 2: service.growth.getattachplandetails:input_type -> growth.getattachplandetails.chrattachplanrequest
	3,  // 3: service.growth.getchrplans:input_type -> growth.getchrplans.chrplansrequest
	4,  // 4: service.growth.getcredithealthreport:input_type -> growth.getcredithealthreport.credithealthreportrequest
	5,  // 5: service.growth.initiatepayment:input_type -> growth.initiatepayment.initiatepaymentrequest
	6,  // 6: service.growth.getchrpaymentstatus:input_type -> growth.getchrpaymentstatus.paymentstatusrequest
	7,  // 7: service.growth.notifychrpayment:output_type -> growth.notifychrpayment.notifychrresponse
	8,  // 8: service.growth.checkplanstatus:output_type -> growth.checkplanstatus.chrplanstatusresponse
	9,  // 9: service.growth.getattachplandetails:output_type -> growth.getattachplandetails.chrattachplanresponse
	10, // 10: service.growth.getchrplans:output_type -> growth.getchrplans.chrplansresponse
	11, // 11: service.growth.getcredithealthreport:output_type -> growth.getcredithealthreport.credithealthreportresponse
	12, // 12: service.growth.initiatepayment:output_type -> growth.initiatepayment.initiatepaymentresponse
	13, // 13: service.growth.getchrpaymentstatus:output_type -> growth.getchrpaymentstatus.paymentstatusresponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_growth_proto_init() }
func file_growth_proto_init() {
	if File_growth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_growth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_growth_proto_goTypes,
		DependencyIndexes: file_growth_proto_depIdxs,
	}.Build()
	File_growth_proto = out.File
	file_growth_proto_rawDesc = nil
	file_growth_proto_goTypes = nil
	file_growth_proto_depIdxs = nil
}
