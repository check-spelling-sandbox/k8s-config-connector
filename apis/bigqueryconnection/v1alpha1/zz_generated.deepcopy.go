//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsAccessRole) DeepCopyInto(out *AwsAccessRole) {
	*out = *in
	if in.IamRoleID != nil {
		in, out := &in.IamRoleID, &out.IamRoleID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsAccessRole.
func (in *AwsAccessRole) DeepCopy() *AwsAccessRole {
	if in == nil {
		return nil
	}
	out := new(AwsAccessRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsCrossAccountRole) DeepCopyInto(out *AwsCrossAccountRole) {
	*out = *in
	if in.IamRoleID != nil {
		in, out := &in.IamRoleID, &out.IamRoleID
		*out = new(string)
		**out = **in
	}
	if in.IamUserID != nil {
		in, out := &in.IamUserID, &out.IamUserID
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsCrossAccountRole.
func (in *AwsCrossAccountRole) DeepCopy() *AwsCrossAccountRole {
	if in == nil {
		return nil
	}
	out := new(AwsCrossAccountRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsProperties) DeepCopyInto(out *AwsProperties) {
	*out = *in
	if in.CrossAccountRole != nil {
		in, out := &in.CrossAccountRole, &out.CrossAccountRole
		*out = new(AwsCrossAccountRole)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessRole != nil {
		in, out := &in.AccessRole, &out.AccessRole
		*out = new(AwsAccessRole)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsProperties.
func (in *AwsProperties) DeepCopy() *AwsProperties {
	if in == nil {
		return nil
	}
	out := new(AwsProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureProperties) DeepCopyInto(out *AzureProperties) {
	*out = *in
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(string)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
	if in.CustomerTenantID != nil {
		in, out := &in.CustomerTenantID, &out.CustomerTenantID
		*out = new(string)
		**out = **in
	}
	if in.RedirectUri != nil {
		in, out := &in.RedirectUri, &out.RedirectUri
		*out = new(string)
		**out = **in
	}
	if in.FederatedApplicationClientID != nil {
		in, out := &in.FederatedApplicationClientID, &out.FederatedApplicationClientID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureProperties.
func (in *AzureProperties) DeepCopy() *AzureProperties {
	if in == nil {
		return nil
	}
	out := new(AzureProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryConnectionConnection) DeepCopyInto(out *BigQueryConnectionConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryConnectionConnection.
func (in *BigQueryConnectionConnection) DeepCopy() *BigQueryConnectionConnection {
	if in == nil {
		return nil
	}
	out := new(BigQueryConnectionConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryConnectionConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryConnectionConnectionList) DeepCopyInto(out *BigQueryConnectionConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryConnectionConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryConnectionConnectionList.
func (in *BigQueryConnectionConnectionList) DeepCopy() *BigQueryConnectionConnectionList {
	if in == nil {
		return nil
	}
	out := new(BigQueryConnectionConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryConnectionConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryConnectionConnectionObservedState) DeepCopyInto(out *BigQueryConnectionConnectionObservedState) {
	*out = *in
	if in.CloudResource != nil {
		in, out := &in.CloudResource, &out.CloudResource
		*out = new(CloudResourcePropertiesStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.FriendlyName != nil {
		in, out := &in.FriendlyName, &out.FriendlyName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HasCredential != nil {
		in, out := &in.HasCredential, &out.HasCredential
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryConnectionConnectionObservedState.
func (in *BigQueryConnectionConnectionObservedState) DeepCopy() *BigQueryConnectionConnectionObservedState {
	if in == nil {
		return nil
	}
	out := new(BigQueryConnectionConnectionObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryConnectionConnectionSpec) DeepCopyInto(out *BigQueryConnectionConnectionSpec) {
	*out = *in
	in.Parent.DeepCopyInto(&out.Parent)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.FriendlyName != nil {
		in, out := &in.FriendlyName, &out.FriendlyName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.CloudResourceSpec != nil {
		in, out := &in.CloudResourceSpec, &out.CloudResourceSpec
		*out = new(CloudResourcePropertiesSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryConnectionConnectionSpec.
func (in *BigQueryConnectionConnectionSpec) DeepCopy() *BigQueryConnectionConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryConnectionConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryConnectionConnectionStatus) DeepCopyInto(out *BigQueryConnectionConnectionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(BigQueryConnectionConnectionObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryConnectionConnectionStatus.
func (in *BigQueryConnectionConnectionStatus) DeepCopy() *BigQueryConnectionConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryConnectionConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResourceProperties) DeepCopyInto(out *CloudResourceProperties) {
	*out = *in
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourceProperties.
func (in *CloudResourceProperties) DeepCopy() *CloudResourceProperties {
	if in == nil {
		return nil
	}
	out := new(CloudResourceProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResourcePropertiesSpec) DeepCopyInto(out *CloudResourcePropertiesSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourcePropertiesSpec.
func (in *CloudResourcePropertiesSpec) DeepCopy() *CloudResourcePropertiesSpec {
	if in == nil {
		return nil
	}
	out := new(CloudResourcePropertiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResourcePropertiesStatus) DeepCopyInto(out *CloudResourcePropertiesStatus) {
	*out = *in
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourcePropertiesStatus.
func (in *CloudResourcePropertiesStatus) DeepCopy() *CloudResourcePropertiesStatus {
	if in == nil {
		return nil
	}
	out := new(CloudResourcePropertiesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSpannerProperties) DeepCopyInto(out *CloudSpannerProperties) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.UseParallelism != nil {
		in, out := &in.UseParallelism, &out.UseParallelism
		*out = new(bool)
		**out = **in
	}
	if in.MaxParallelism != nil {
		in, out := &in.MaxParallelism, &out.MaxParallelism
		*out = new(int32)
		**out = **in
	}
	if in.UseServerlessAnalytics != nil {
		in, out := &in.UseServerlessAnalytics, &out.UseServerlessAnalytics
		*out = new(bool)
		**out = **in
	}
	if in.UseDataBoost != nil {
		in, out := &in.UseDataBoost, &out.UseDataBoost
		*out = new(bool)
		**out = **in
	}
	if in.DatabaseRole != nil {
		in, out := &in.DatabaseRole, &out.DatabaseRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSpannerProperties.
func (in *CloudSpannerProperties) DeepCopy() *CloudSpannerProperties {
	if in == nil {
		return nil
	}
	out := new(CloudSpannerProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSqlCredential) DeepCopyInto(out *CloudSqlCredential) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSqlCredential.
func (in *CloudSqlCredential) DeepCopy() *CloudSqlCredential {
	if in == nil {
		return nil
	}
	out := new(CloudSqlCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSqlProperties) DeepCopyInto(out *CloudSqlProperties) {
	*out = *in
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(CloudSqlCredential)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSqlProperties.
func (in *CloudSqlProperties) DeepCopy() *CloudSqlProperties {
	if in == nil {
		return nil
	}
	out := new(CloudSqlProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.FriendlyName != nil {
		in, out := &in.FriendlyName, &out.FriendlyName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.CloudSql != nil {
		in, out := &in.CloudSql, &out.CloudSql
		*out = new(CloudSqlProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Aws != nil {
		in, out := &in.Aws, &out.Aws
		*out = new(AwsProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudSpanner != nil {
		in, out := &in.CloudSpanner, &out.CloudSpanner
		*out = new(CloudSpannerProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudResource != nil {
		in, out := &in.CloudResource, &out.CloudResource
		*out = new(CloudResourceProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Spark != nil {
		in, out := &in.Spark, &out.Spark
		*out = new(SparkProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.SalesforceDataCloud != nil {
		in, out := &in.SalesforceDataCloud, &out.SalesforceDataCloud
		*out = new(SalesforceDataCloudProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
	if in.LastModifiedTime != nil {
		in, out := &in.LastModifiedTime, &out.LastModifiedTime
		*out = new(int64)
		**out = **in
	}
	if in.HasCredential != nil {
		in, out := &in.HasCredential, &out.HasCredential
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreServiceConfig) DeepCopyInto(out *MetastoreServiceConfig) {
	*out = *in
	if in.MetastoreService != nil {
		in, out := &in.MetastoreService, &out.MetastoreService
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreServiceConfig.
func (in *MetastoreServiceConfig) DeepCopy() *MetastoreServiceConfig {
	if in == nil {
		return nil
	}
	out := new(MetastoreServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SalesforceDataCloudProperties) DeepCopyInto(out *SalesforceDataCloudProperties) {
	*out = *in
	if in.InstanceUri != nil {
		in, out := &in.InstanceUri, &out.InstanceUri
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SalesforceDataCloudProperties.
func (in *SalesforceDataCloudProperties) DeepCopy() *SalesforceDataCloudProperties {
	if in == nil {
		return nil
	}
	out := new(SalesforceDataCloudProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkHistoryServerConfig) DeepCopyInto(out *SparkHistoryServerConfig) {
	*out = *in
	if in.DataprocCluster != nil {
		in, out := &in.DataprocCluster, &out.DataprocCluster
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkHistoryServerConfig.
func (in *SparkHistoryServerConfig) DeepCopy() *SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := new(SparkHistoryServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkProperties) DeepCopyInto(out *SparkProperties) {
	*out = *in
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.MetastoreServiceConfig != nil {
		in, out := &in.MetastoreServiceConfig, &out.MetastoreServiceConfig
		*out = new(MetastoreServiceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkHistoryServerConfig != nil {
		in, out := &in.SparkHistoryServerConfig, &out.SparkHistoryServerConfig
		*out = new(SparkHistoryServerConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkProperties.
func (in *SparkProperties) DeepCopy() *SparkProperties {
	if in == nil {
		return nil
	}
	out := new(SparkProperties)
	in.DeepCopyInto(out)
	return out
}
