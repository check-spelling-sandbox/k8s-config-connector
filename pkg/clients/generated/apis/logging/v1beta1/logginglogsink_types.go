// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LogsinkBigqueryOptions struct {
	/* Whether to use BigQuery's partition tables. By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned tables the date suffix is no longer present and special query syntax has to be used instead. In both cases, tables are sharded based on UTC timezone. */
	UsePartitionedTables bool `json:"usePartitionedTables"`
}

type LogsinkDestination struct {
	// +optional
	BigQueryDatasetRef *v1alpha1.ResourceRef `json:"bigQueryDatasetRef,omitempty"`

	/* Only `external` field is supported to configure the reference. */
	// +optional
	LoggingLogBucketRef *v1alpha1.ResourceRef `json:"loggingLogBucketRef,omitempty"`

	// +optional
	PubSubTopicRef *v1alpha1.ResourceRef `json:"pubSubTopicRef,omitempty"`

	// +optional
	StorageBucketRef *v1alpha1.ResourceRef `json:"storageBucketRef,omitempty"`
}

type LogsinkExclusions struct {
	/* A description of this exclusion. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* If set to True, then this exclusion is disabled and it does not exclude any log entries. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. */
	Filter string `json:"filter"`

	/* A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric. */
	Name string `json:"name"`
}

type LoggingLogSinkSpec struct {
	/* Options that affect sinks exporting data to BigQuery. */
	// +optional
	BigqueryOptions *LogsinkBigqueryOptions `json:"bigqueryOptions,omitempty"`

	/* A description of this sink. The maximum length of the description is 8000 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	Destination LogsinkDestination `json:"destination"`

	/* If set to True, then this sink is disabled and it does not export any log entries. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion's filters, it will not be exported. */
	// +optional
	Exclusions []LogsinkExclusions `json:"exclusions,omitempty"`

	/* The filter to apply when exporting logs. Only log entries that match the filter are exported. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* The folder in which to create the sink. Only one of projectRef,
	folderRef, or organizationRef may be specified. */
	// +optional
	FolderRef *v1alpha1.ResourceRef `json:"folderRef,omitempty"`

	/* Immutable. Whether or not to include children organizations in the sink export. If true, logs associated with child projects are also exported; otherwise only logs relating to the provided organization are included. */
	// +optional
	IncludeChildren *bool `json:"includeChildren,omitempty"`

	/* The organization in which to create the sink. Only one of projectRef,
	folderRef, or organizationRef may be specified. */
	// +optional
	OrganizationRef *v1alpha1.ResourceRef `json:"organizationRef,omitempty"`

	/* The project in which to create the sink. Only one of projectRef,
	folderRef, or organizationRef may be specified. */
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Whether or not to create a unique identity associated with this sink. If false (the default), then the writer_identity used is serviceAccount:cloud-logs@system.gserviceaccount.com. If true, then a unique service account is created and used for this sink. If you wish to publish logs across projects, you must set unique_writer_identity to true. */
	// +optional
	UniqueWriterIdentity *bool `json:"uniqueWriterIdentity,omitempty"`
}

type LoggingLogSinkStatus struct {
	/* Conditions represent the latest available observations of the
	   LoggingLogSink's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The identity associated with this sink. This identity must be granted write access to the configured destination. */
	// +optional
	WriterIdentity *string `json:"writerIdentity,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoggingLogSink is the Schema for the logging API
// +k8s:openapi-gen=true
type LoggingLogSink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingLogSinkSpec   `json:"spec,omitempty"`
	Status LoggingLogSinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoggingLogSinkList contains a list of LoggingLogSink
type LoggingLogSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingLogSink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingLogSink{}, &LoggingLogSinkList{})
}
