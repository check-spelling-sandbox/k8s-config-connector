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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConnectionAccessRole struct {
	/* The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection. */
	IamRoleID string `json:"iamRoleID"`
}

type ConnectionAws struct {
	/* Authentication using Google owned service account to assume into customer's AWS IAM Role. */
	AccessRole ConnectionAccessRole `json:"accessRole"`
}

type ConnectionAzure struct {
	/* The id of customer's directory that host the data. */
	CustomerTenantID string `json:"customerTenantID"`

	/* The client ID of the user's Azure Active Directory Application used for a federated connection. */
	// +optional
	FederatedApplicationClientID *string `json:"federatedApplicationClientID,omitempty"`
}

type ConnectionCloudResource struct {
}

type ConnectionCloudSQL struct {
	/* Cloud SQL credential. */
	Credential ConnectionCredential `json:"credential"`

	/* Database name. */
	Database string `json:"database"`

	/* Reference to the Cloud SQL instance ID. */
	InstanceRef v1alpha1.ResourceRef `json:"instanceRef"`

	/* Type of the Cloud SQL database. */
	Type string `json:"type"`
}

type ConnectionCloudSpanner struct {
	/* Reference to a spanner database ID. */
	DatabaseRef v1alpha1.ResourceRef `json:"databaseRef"`

	/* Optional. Cloud Spanner database role for fine-grained access control.
	The Cloud Spanner admin should have provisioned the database role with
	appropriate permissions, such as `SELECT` and `INSERT`. Other users should
	only use roles provided by their Cloud Spanner admins.

	For more details, see [About fine-grained access control]
	(https://cloud.google.com/spanner/docs/fgac-about).

	REQUIRES: The database role name must start with a letter, and can only
	contain letters, numbers, and underscores. */
	// +optional
	DatabaseRole *string `json:"databaseRole,omitempty"`

	/* Allows setting max parallelism per query when executing on Spanner
	independent compute resources. If unspecified, default values of
	parallelism are chosen that are dependent on the Cloud Spanner instance
	configuration.

	REQUIRES: `use_parallelism` must be set.
	REQUIRES: Either `use_data_boost` or `use_serverless_analytics` must be
	set. */
	// +optional
	MaxParallelism *int32 `json:"maxParallelism,omitempty"`

	/* If set, the request will be executed via Spanner independent compute
	resources.
	REQUIRES: `use_parallelism` must be set.

	NOTE: `use_serverless_analytics` will be deprecated. Prefer
	`use_data_boost` over `use_serverless_analytics`. */
	// +optional
	UseDataBoost *bool `json:"useDataBoost,omitempty"`

	/* If parallelism should be used when reading from Cloud Spanner */
	// +optional
	UseParallelism *bool `json:"useParallelism,omitempty"`

	/* If the serverless analytics service should be used to read data from Cloud Spanner. Note: `use_parallelism` must be set when using serverless analytics. */
	// +optional
	UseServerlessAnalytics *bool `json:"useServerlessAnalytics,omitempty"`
}

type ConnectionCredential struct {
	/* The password for the credential. */
	// +optional
	Password *string `json:"password,omitempty"`

	/* The username for the credential. */
	// +optional
	Username *string `json:"username,omitempty"`
}

type ConnectionMetastoreService struct {
	/* Optional. Resource name of an existing Dataproc Metastore service.

	Example:

	* `projects/[project_id]/locations/[region]/services/[service_id]` */
	// +optional
	MetastoreServiceRef *v1alpha1.ResourceRef `json:"metastoreServiceRef,omitempty"`
}

type ConnectionSpark struct {
	/* Optional. Dataproc Metastore Service configuration for the connection. */
	// +optional
	MetastoreService *ConnectionMetastoreService `json:"metastoreService,omitempty"`

	/* Optional. Spark History Server configuration for the connection. */
	// +optional
	SparkHistoryServer *ConnectionSparkHistoryServer `json:"sparkHistoryServer,omitempty"`
}

type ConnectionSparkHistoryServer struct {
	/* Optional. Resource name of an existing Dataproc Cluster to act as a Spark
	History Server for the connection.

	Example:

	* `projects/[project_id]/regions/[region]/clusters/[cluster_name]` */
	// +optional
	DataprocClusterRef *v1alpha1.ResourceRef `json:"dataprocClusterRef,omitempty"`
}

type BigQueryConnectionConnectionSpec struct {
	/* Amazon Web Services (AWS) properties. */
	// +optional
	Aws *ConnectionAws `json:"aws,omitempty"`

	/* Azure properties. */
	// +optional
	Azure *ConnectionAzure `json:"azure,omitempty"`

	/* Use Cloud Resource properties. */
	// +optional
	CloudResource *ConnectionCloudResource `json:"cloudResource,omitempty"`

	/* Cloud SQL properties. */
	// +optional
	CloudSQL *ConnectionCloudSQL `json:"cloudSQL,omitempty"`

	/* Cloud Spanner properties. */
	// +optional
	CloudSpanner *ConnectionCloudSpanner `json:"cloudSpanner,omitempty"`

	/* User provided description. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* User provided display name for the connection. */
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty"`

	/* Immutable. */
	Location string `json:"location"`

	/* The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* The BigQuery ConnectionID. This is a server-generated ID in the UUID format. If not provided, ConfigConnector will create a new Connection and store the UUID in `status.serviceGeneratedID` field. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Spark properties. */
	// +optional
	Spark *ConnectionSpark `json:"spark,omitempty"`
}

type ConnectionAccessRoleStatus struct {
	/* A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's AWS IAM Role. */
	// +optional
	Identity *string `json:"identity,omitempty"`
}

type ConnectionAwsStatus struct {
	// +optional
	AccessRole *ConnectionAccessRoleStatus `json:"accessRole,omitempty"`
}

type ConnectionAzureStatus struct {
	/* The name of the Azure Active Directory Application. */
	// +optional
	Application *string `json:"application,omitempty"`

	/* The client id of the Azure Active Directory Application. */
	// +optional
	ClientID *string `json:"clientID,omitempty"`

	/* A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's Azure Active Directory Application. */
	// +optional
	Identity *string `json:"identity,omitempty"`

	/* The object id of the Azure Active Directory Application. */
	// +optional
	ObjectID *string `json:"objectID,omitempty"`

	/* The URL user will be redirected to after granting consent during connection setup. */
	// +optional
	RedirectUri *string `json:"redirectUri,omitempty"`
}

type ConnectionCloudResourceStatus struct {
	/* The account ID of the service created for the purpose of this
	connection.

	The service account does not have any permissions associated with it
	when it is created. After creation, customers delegate permissions
	to the service account. When the connection is used in the context of an
	operation in BigQuery, the service account will be used to connect to the
	desired resources in GCP.

	The account ID is in the form of:
	<service-1234>@gcp-sa-bigquery-cloudresource.iam.gserviceaccount.com */
	// +optional
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

type ConnectionCloudSQLStatus struct {
	/* The account ID of the service used for the purpose of this connection.

	When the connection is used in the context of an operation in
	BigQuery, this service account will serve as the identity being used for
	connecting to the CloudSQL instance specified in this connection. */
	// +optional
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

type ConnectionObservedStateStatus struct {
	// +optional
	Aws *ConnectionAwsStatus `json:"aws,omitempty"`

	// +optional
	Azure *ConnectionAzureStatus `json:"azure,omitempty"`

	// +optional
	CloudResource *ConnectionCloudResourceStatus `json:"cloudResource,omitempty"`

	// +optional
	CloudSQL *ConnectionCloudSQLStatus `json:"cloudSQL,omitempty"`

	/* The description for the connection. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The display name for the connection. */
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty"`

	/* Output only. True, if credential is configured for this connection. */
	// +optional
	HasCredential *bool `json:"hasCredential,omitempty"`

	// +optional
	Spark *ConnectionSparkStatus `json:"spark,omitempty"`
}

type ConnectionSparkStatus struct {
	/* The account ID of the service created for the purpose of this
	connection.

	The service account does not have any permissions associated with it when
	it is created. After creation, customers delegate permissions to the
	service account. When the connection is used in the context of a stored
	procedure for Apache Spark in BigQuery, the service account is used to
	connect to the desired resources in Google Cloud.

	The account ID is in the form of:
	bqcx-<projectnumber>-<uniqueid>@gcp-sa-bigquery-consp.iam.gserviceaccount.com */
	// +optional
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

type BigQueryConnectionConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryConnectionConnection's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* A unique specifier for the BigQueryConnectionConnection resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *ConnectionObservedStateStatus `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryconnectionconnection;gcpbigqueryconnectionconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryConnectionConnection is the Schema for the bigqueryconnection API
// +k8s:openapi-gen=true
type BigQueryConnectionConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryConnectionConnectionSpec   `json:"spec,omitempty"`
	Status BigQueryConnectionConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigQueryConnectionConnectionList contains a list of BigQueryConnectionConnection
type BigQueryConnectionConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryConnectionConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryConnectionConnection{}, &BigQueryConnectionConnectionList{})
}
