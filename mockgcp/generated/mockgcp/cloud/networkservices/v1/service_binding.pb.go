// Copyright 2023 Google LLC
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
// source: mockgcp/cloud/networkservices/v1/service_binding.proto

package networkservicespb

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

// ServiceBinding is the resource that defines a Service Directory Service to
// be used in a BackendService resource.
type ServiceBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the ServiceBinding resource. It matches pattern
	// `projects/*/locations/global/serviceBindings/service_binding_name`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. A free-text description of the resource. Max length 1024
	// characters.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The timestamp when the resource was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The timestamp when the resource was updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. The full service directory service name of the format
	// /projects/*/locations/*/namespaces/*/services/*
	Service string `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	// Optional. Set of label tags associated with the ServiceBinding resource.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceBinding) Reset() {
	*x = ServiceBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceBinding) ProtoMessage() {}

func (x *ServiceBinding) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceBinding.ProtoReflect.Descriptor instead.
func (*ServiceBinding) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceBinding) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceBinding) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceBinding) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ServiceBinding) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ServiceBinding) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ServiceBinding) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Request used with the ListServiceBindings method.
type ListServiceBindingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The project and location from which the ServiceBindings should be
	// listed, specified in the format `projects/*/locations/global`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Maximum number of ServiceBindings to return per call.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value returned by the last `ListServiceBindingsResponse`
	// Indicates that this is a continuation of a prior `ListRouters` call,
	// and that the system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListServiceBindingsRequest) Reset() {
	*x = ListServiceBindingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceBindingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceBindingsRequest) ProtoMessage() {}

func (x *ListServiceBindingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceBindingsRequest.ProtoReflect.Descriptor instead.
func (*ListServiceBindingsRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescGZIP(), []int{1}
}

func (x *ListServiceBindingsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListServiceBindingsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListServiceBindingsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response returned by the ListServiceBindings method.
type ListServiceBindingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of ServiceBinding resources.
	ServiceBindings []*ServiceBinding `protobuf:"bytes,1,rep,name=service_bindings,json=serviceBindings,proto3" json:"service_bindings,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `next_page_token` is included. To get the next set of results, call this
	// method again using the value of `next_page_token` as `page_token`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListServiceBindingsResponse) Reset() {
	*x = ListServiceBindingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceBindingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceBindingsResponse) ProtoMessage() {}

func (x *ListServiceBindingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceBindingsResponse.ProtoReflect.Descriptor instead.
func (*ListServiceBindingsResponse) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescGZIP(), []int{2}
}

func (x *ListServiceBindingsResponse) GetServiceBindings() []*ServiceBinding {
	if x != nil {
		return x.ServiceBindings
	}
	return nil
}

func (x *ListServiceBindingsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request used by the GetServiceBinding method.
type GetServiceBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A name of the ServiceBinding to get. Must be in the format
	// `projects/*/locations/global/serviceBindings/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetServiceBindingRequest) Reset() {
	*x = GetServiceBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceBindingRequest) ProtoMessage() {}

func (x *GetServiceBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceBindingRequest.ProtoReflect.Descriptor instead.
func (*GetServiceBindingRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescGZIP(), []int{3}
}

func (x *GetServiceBindingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request used by the ServiceBinding method.
type CreateServiceBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource of the ServiceBinding. Must be in the
	// format `projects/*/locations/global`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Short name of the ServiceBinding resource to be created.
	ServiceBindingId string `protobuf:"bytes,2,opt,name=service_binding_id,json=serviceBindingId,proto3" json:"service_binding_id,omitempty"`
	// Required. ServiceBinding resource to be created.
	ServiceBinding *ServiceBinding `protobuf:"bytes,3,opt,name=service_binding,json=serviceBinding,proto3" json:"service_binding,omitempty"`
}

func (x *CreateServiceBindingRequest) Reset() {
	*x = CreateServiceBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceBindingRequest) ProtoMessage() {}

func (x *CreateServiceBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceBindingRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceBindingRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescGZIP(), []int{4}
}

func (x *CreateServiceBindingRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateServiceBindingRequest) GetServiceBindingId() string {
	if x != nil {
		return x.ServiceBindingId
	}
	return ""
}

func (x *CreateServiceBindingRequest) GetServiceBinding() *ServiceBinding {
	if x != nil {
		return x.ServiceBinding
	}
	return nil
}

// Request used by the DeleteServiceBinding method.
type DeleteServiceBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A name of the ServiceBinding to delete. Must be in the format
	// `projects/*/locations/global/serviceBindings/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteServiceBindingRequest) Reset() {
	*x = DeleteServiceBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceBindingRequest) ProtoMessage() {}

func (x *DeleteServiceBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceBindingRequest.ProtoReflect.Descriptor instead.
func (*DeleteServiceBindingRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteServiceBindingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_mockgcp_cloud_networkservices_v1_service_binding_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63,
	0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x04, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x7d, 0xea, 0x41, 0x7a, 0x0a, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x49, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x7d, 0x22, 0xa7, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x35, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2f, 0x12, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa2, 0x01, 0x0a,
	0x1b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x65, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2f,
	0x12, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x68, 0x0a, 0x1b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2f, 0x0a,
	0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0xf5, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70,
	0x62, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescData = file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDesc
)

func file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescData)
	})
	return file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDescData
}

var file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mockgcp_cloud_networkservices_v1_service_binding_proto_goTypes = []interface{}{
	(*ServiceBinding)(nil),              // 0: mockgcp.cloud.networkservices.v1.ServiceBinding
	(*ListServiceBindingsRequest)(nil),  // 1: mockgcp.cloud.networkservices.v1.ListServiceBindingsRequest
	(*ListServiceBindingsResponse)(nil), // 2: mockgcp.cloud.networkservices.v1.ListServiceBindingsResponse
	(*GetServiceBindingRequest)(nil),    // 3: mockgcp.cloud.networkservices.v1.GetServiceBindingRequest
	(*CreateServiceBindingRequest)(nil), // 4: mockgcp.cloud.networkservices.v1.CreateServiceBindingRequest
	(*DeleteServiceBindingRequest)(nil), // 5: mockgcp.cloud.networkservices.v1.DeleteServiceBindingRequest
	nil,                                 // 6: mockgcp.cloud.networkservices.v1.ServiceBinding.LabelsEntry
	(*timestamp.Timestamp)(nil),         // 7: google.protobuf.Timestamp
}
var file_mockgcp_cloud_networkservices_v1_service_binding_proto_depIdxs = []int32{
	7, // 0: mockgcp.cloud.networkservices.v1.ServiceBinding.create_time:type_name -> google.protobuf.Timestamp
	7, // 1: mockgcp.cloud.networkservices.v1.ServiceBinding.update_time:type_name -> google.protobuf.Timestamp
	6, // 2: mockgcp.cloud.networkservices.v1.ServiceBinding.labels:type_name -> mockgcp.cloud.networkservices.v1.ServiceBinding.LabelsEntry
	0, // 3: mockgcp.cloud.networkservices.v1.ListServiceBindingsResponse.service_bindings:type_name -> mockgcp.cloud.networkservices.v1.ServiceBinding
	0, // 4: mockgcp.cloud.networkservices.v1.CreateServiceBindingRequest.service_binding:type_name -> mockgcp.cloud.networkservices.v1.ServiceBinding
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_networkservices_v1_service_binding_proto_init() }
func file_mockgcp_cloud_networkservices_v1_service_binding_proto_init() {
	if File_mockgcp_cloud_networkservices_v1_service_binding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceBinding); i {
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
		file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceBindingsRequest); i {
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
		file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceBindingsResponse); i {
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
		file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceBindingRequest); i {
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
		file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceBindingRequest); i {
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
		file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceBindingRequest); i {
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
			RawDescriptor: file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_cloud_networkservices_v1_service_binding_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_networkservices_v1_service_binding_proto_depIdxs,
		MessageInfos:      file_mockgcp_cloud_networkservices_v1_service_binding_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_networkservices_v1_service_binding_proto = out.File
	file_mockgcp_cloud_networkservices_v1_service_binding_proto_rawDesc = nil
	file_mockgcp_cloud_networkservices_v1_service_binding_proto_goTypes = nil
	file_mockgcp_cloud_networkservices_v1_service_binding_proto_depIdxs = nil
}
