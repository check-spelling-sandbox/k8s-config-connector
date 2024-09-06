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
// source: mockgcp/firestore/admin/v1/schedule.proto

package adminpb

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
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

// A backup schedule for a Cloud Firestore Database.
//
// This resource is owned by the database it is backing up, and is deleted along
// with the database. The actual backups are not though.
type BackupSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique backup schedule identifier across all locations and
	// databases for the given project.
	//
	// This will be auto-assigned.
	//
	// Format is
	// `projects/{project}/databases/{database}/backupSchedules/{backup_schedule}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The timestamp at which this backup schedule was created and
	// effective since.
	//
	// No backups will be created for this schedule before this time.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The timestamp at which this backup schedule was most recently
	// updated. When a backup schedule is first created, this is the same as
	// create_time.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// At what relative time in the future, compared to its creation time,
	// the backup should be deleted, e.g. keep backups for 7 days.
	//
	// The maximum supported retention period is 14 weeks.
	Retention *duration.Duration `protobuf:"bytes,6,opt,name=retention,proto3" json:"retention,omitempty"`
	// A oneof field to represent when backups will be taken.
	//
	// Types that are assignable to Recurrence:
	//
	//	*BackupSchedule_DailyRecurrence
	//	*BackupSchedule_WeeklyRecurrence
	Recurrence isBackupSchedule_Recurrence `protobuf_oneof:"recurrence"`
}

func (x *BackupSchedule) Reset() {
	*x = BackupSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupSchedule) ProtoMessage() {}

func (x *BackupSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupSchedule.ProtoReflect.Descriptor instead.
func (*BackupSchedule) Descriptor() ([]byte, []int) {
	return file_mockgcp_firestore_admin_v1_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *BackupSchedule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BackupSchedule) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BackupSchedule) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *BackupSchedule) GetRetention() *duration.Duration {
	if x != nil {
		return x.Retention
	}
	return nil
}

func (m *BackupSchedule) GetRecurrence() isBackupSchedule_Recurrence {
	if m != nil {
		return m.Recurrence
	}
	return nil
}

func (x *BackupSchedule) GetDailyRecurrence() *DailyRecurrence {
	if x, ok := x.GetRecurrence().(*BackupSchedule_DailyRecurrence); ok {
		return x.DailyRecurrence
	}
	return nil
}

func (x *BackupSchedule) GetWeeklyRecurrence() *WeeklyRecurrence {
	if x, ok := x.GetRecurrence().(*BackupSchedule_WeeklyRecurrence); ok {
		return x.WeeklyRecurrence
	}
	return nil
}

type isBackupSchedule_Recurrence interface {
	isBackupSchedule_Recurrence()
}

type BackupSchedule_DailyRecurrence struct {
	// For a schedule that runs daily.
	DailyRecurrence *DailyRecurrence `protobuf:"bytes,7,opt,name=daily_recurrence,json=dailyRecurrence,proto3,oneof"`
}

type BackupSchedule_WeeklyRecurrence struct {
	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence *WeeklyRecurrence `protobuf:"bytes,8,opt,name=weekly_recurrence,json=weeklyRecurrence,proto3,oneof"`
}

func (*BackupSchedule_DailyRecurrence) isBackupSchedule_Recurrence() {}

func (*BackupSchedule_WeeklyRecurrence) isBackupSchedule_Recurrence() {}

// Represents a recurring schedule that runs every day.
//
// The time zone is UTC.
type DailyRecurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DailyRecurrence) Reset() {
	*x = DailyRecurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyRecurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyRecurrence) ProtoMessage() {}

func (x *DailyRecurrence) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyRecurrence.ProtoReflect.Descriptor instead.
func (*DailyRecurrence) Descriptor() ([]byte, []int) {
	return file_mockgcp_firestore_admin_v1_schedule_proto_rawDescGZIP(), []int{1}
}

// Represents a recurring schedule that runs on a specified day of the week.
//
// The time zone is UTC.
type WeeklyRecurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The day of week to run.
	//
	// DAY_OF_WEEK_UNSPECIFIED is not allowed.
	Day dayofweek.DayOfWeek `protobuf:"varint,2,opt,name=day,proto3,enum=google.type.DayOfWeek" json:"day,omitempty"`
}

func (x *WeeklyRecurrence) Reset() {
	*x = WeeklyRecurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeeklyRecurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklyRecurrence) ProtoMessage() {}

func (x *WeeklyRecurrence) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklyRecurrence.ProtoReflect.Descriptor instead.
func (*WeeklyRecurrence) Descriptor() ([]byte, []int) {
	return file_mockgcp_firestore_admin_v1_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *WeeklyRecurrence) GetDay() dayofweek.DayOfWeek {
	if x != nil {
		return x.Day
	}
	return dayofweek.DayOfWeek(0)
}

var File_mockgcp_firestore_admin_v1_schedule_proto protoreflect.FileDescriptor

var file_mockgcp_firestore_admin_v1_schedule_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x64, 0x61, 0x79, 0x6f, 0x66, 0x77, 0x65, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa4, 0x04, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x10, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x66,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x11, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x72,
	0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x10, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x3a, 0x77, 0xea, 0x41, 0x74, 0x0a, 0x27, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x49, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x7d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x7d, 0x42, 0x0c, 0x0a, 0x0a, 0x72, 0x65,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x57,
	0x65, 0x65, 0x6b, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x79, 0x4f, 0x66,
	0x57, 0x65, 0x65, 0x6b, 0x52, 0x03, 0x64, 0x61, 0x79, 0x42, 0xdd, 0x01, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62,
	0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x46, 0x53, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x46,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mockgcp_firestore_admin_v1_schedule_proto_rawDescOnce sync.Once
	file_mockgcp_firestore_admin_v1_schedule_proto_rawDescData = file_mockgcp_firestore_admin_v1_schedule_proto_rawDesc
)

func file_mockgcp_firestore_admin_v1_schedule_proto_rawDescGZIP() []byte {
	file_mockgcp_firestore_admin_v1_schedule_proto_rawDescOnce.Do(func() {
		file_mockgcp_firestore_admin_v1_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_firestore_admin_v1_schedule_proto_rawDescData)
	})
	return file_mockgcp_firestore_admin_v1_schedule_proto_rawDescData
}

var file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mockgcp_firestore_admin_v1_schedule_proto_goTypes = []interface{}{
	(*BackupSchedule)(nil),      // 0: mockgcp.firestore.admin.v1.BackupSchedule
	(*DailyRecurrence)(nil),     // 1: mockgcp.firestore.admin.v1.DailyRecurrence
	(*WeeklyRecurrence)(nil),    // 2: mockgcp.firestore.admin.v1.WeeklyRecurrence
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 4: google.protobuf.Duration
	(dayofweek.DayOfWeek)(0),    // 5: google.type.DayOfWeek
}
var file_mockgcp_firestore_admin_v1_schedule_proto_depIdxs = []int32{
	3, // 0: mockgcp.firestore.admin.v1.BackupSchedule.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: mockgcp.firestore.admin.v1.BackupSchedule.update_time:type_name -> google.protobuf.Timestamp
	4, // 2: mockgcp.firestore.admin.v1.BackupSchedule.retention:type_name -> google.protobuf.Duration
	1, // 3: mockgcp.firestore.admin.v1.BackupSchedule.daily_recurrence:type_name -> mockgcp.firestore.admin.v1.DailyRecurrence
	2, // 4: mockgcp.firestore.admin.v1.BackupSchedule.weekly_recurrence:type_name -> mockgcp.firestore.admin.v1.WeeklyRecurrence
	5, // 5: mockgcp.firestore.admin.v1.WeeklyRecurrence.day:type_name -> google.type.DayOfWeek
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mockgcp_firestore_admin_v1_schedule_proto_init() }
func file_mockgcp_firestore_admin_v1_schedule_proto_init() {
	if File_mockgcp_firestore_admin_v1_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupSchedule); i {
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
		file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyRecurrence); i {
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
		file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeeklyRecurrence); i {
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
	file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BackupSchedule_DailyRecurrence)(nil),
		(*BackupSchedule_WeeklyRecurrence)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mockgcp_firestore_admin_v1_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_firestore_admin_v1_schedule_proto_goTypes,
		DependencyIndexes: file_mockgcp_firestore_admin_v1_schedule_proto_depIdxs,
		MessageInfos:      file_mockgcp_firestore_admin_v1_schedule_proto_msgTypes,
	}.Build()
	File_mockgcp_firestore_admin_v1_schedule_proto = out.File
	file_mockgcp_firestore_admin_v1_schedule_proto_rawDesc = nil
	file_mockgcp_firestore_admin_v1_schedule_proto_goTypes = nil
	file_mockgcp_firestore_admin_v1_schedule_proto_depIdxs = nil
}
