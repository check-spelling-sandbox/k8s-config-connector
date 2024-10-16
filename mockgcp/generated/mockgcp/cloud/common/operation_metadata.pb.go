// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/cloud/common/operation_metadata.proto

package common

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Represents the metadata of the long-running operation.
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time the operation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time the operation finished running.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Output only. Server-defined resource path for the target of the operation.
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// Output only. Name of the verb executed by the operation.
	Verb string `protobuf:"bytes,4,opt,name=verb,proto3" json:"verb,omitempty"`
	// Output only. Human-readable status of the operation, if any.
	StatusDetail string `protobuf:"bytes,5,opt,name=status_detail,json=statusDetail,proto3" json:"status_detail,omitempty"`
	// Output only. Identifies whether the user has requested cancellation
	// of the operation. Operations that have successfully been cancelled
	// have [Operation.error][] value with a [google.rpc.Status.code][] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelRequested bool `protobuf:"varint,6,opt,name=cancel_requested,json=cancelRequested,proto3" json:"cancel_requested,omitempty"`
	// Output only. API version used to start the operation.
	ApiVersion string `protobuf:"bytes,7,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_common_operation_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_common_operation_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_common_operation_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *OperationMetadata) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *OperationMetadata) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *OperationMetadata) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *OperationMetadata) GetStatusDetail() string {
	if x != nil {
		return x.StatusDetail
	}
	return ""
}

func (x *OperationMetadata) GetCancelRequested() bool {
	if x != nil {
		return x.CancelRequested
	}
	return false
}

func (x *OperationMetadata) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

var File_mockgcp_cloud_common_operation_metadata_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_common_operation_metadata_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x02, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x76, 0x65, 0x72,
	0x62, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x10, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x57, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a,
	0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mockgcp_cloud_common_operation_metadata_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_common_operation_metadata_proto_rawDescData = file_mockgcp_cloud_common_operation_metadata_proto_rawDesc
)

func file_mockgcp_cloud_common_operation_metadata_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_common_operation_metadata_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_common_operation_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_common_operation_metadata_proto_rawDescData)
	})
	return file_mockgcp_cloud_common_operation_metadata_proto_rawDescData
}

var file_mockgcp_cloud_common_operation_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mockgcp_cloud_common_operation_metadata_proto_goTypes = []interface{}{
	(*OperationMetadata)(nil),   // 0: mockgcp.cloud.common.OperationMetadata
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_mockgcp_cloud_common_operation_metadata_proto_depIdxs = []int32{
	1, // 0: mockgcp.cloud.common.OperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: mockgcp.cloud.common.OperationMetadata.end_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_common_operation_metadata_proto_init() }
func file_mockgcp_cloud_common_operation_metadata_proto_init() {
	if File_mockgcp_cloud_common_operation_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_common_operation_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
			RawDescriptor: file_mockgcp_cloud_common_operation_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_cloud_common_operation_metadata_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_common_operation_metadata_proto_depIdxs,
		MessageInfos:      file_mockgcp_cloud_common_operation_metadata_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_common_operation_metadata_proto = out.File
	file_mockgcp_cloud_common_operation_metadata_proto_rawDesc = nil
	file_mockgcp_cloud_common_operation_metadata_proto_goTypes = nil
	file_mockgcp_cloud_common_operation_metadata_proto_depIdxs = nil
}
