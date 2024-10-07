//go:build !ignore_autogenerated

// Copyright 2020 Google LLC
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchange) DeepCopyInto(out *BigQueryAnalyticsHubDataExchange) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchange.
func (in *BigQueryAnalyticsHubDataExchange) DeepCopy() *BigQueryAnalyticsHubDataExchange {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryAnalyticsHubDataExchange) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeList) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryAnalyticsHubDataExchange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeList.
func (in *BigQueryAnalyticsHubDataExchangeList) DeepCopy() *BigQueryAnalyticsHubDataExchangeList {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryAnalyticsHubDataExchangeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeObservedState) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeObservedState) {
	*out = *in
	if in.ListingCount != nil {
		in, out := &in.ListingCount, &out.ListingCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeObservedState.
func (in *BigQueryAnalyticsHubDataExchangeObservedState) DeepCopy() *BigQueryAnalyticsHubDataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeSpec) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeSpec) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeSpec.
func (in *BigQueryAnalyticsHubDataExchangeSpec) DeepCopy() *BigQueryAnalyticsHubDataExchangeSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeStatus) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeStatus) {
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
		*out = new(BigQueryAnalyticsHubDataExchangeObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeStatus.
func (in *BigQueryAnalyticsHubDataExchangeStatus) DeepCopy() *BigQueryAnalyticsHubDataExchangeStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubListing) DeepCopyInto(out *BigQueryAnalyticsHubListing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubListing.
func (in *BigQueryAnalyticsHubListing) DeepCopy() *BigQueryAnalyticsHubListing {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubListing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryAnalyticsHubListing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubListingList) DeepCopyInto(out *BigQueryAnalyticsHubListingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryAnalyticsHubListing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubListingList.
func (in *BigQueryAnalyticsHubListingList) DeepCopy() *BigQueryAnalyticsHubListingList {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubListingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryAnalyticsHubListingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubListingObservedState) DeepCopyInto(out *BigQueryAnalyticsHubListingObservedState) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubListingObservedState.
func (in *BigQueryAnalyticsHubListingObservedState) DeepCopy() *BigQueryAnalyticsHubListingObservedState {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubListingObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubListingParent) DeepCopyInto(out *BigQueryAnalyticsHubListingParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubListingParent.
func (in *BigQueryAnalyticsHubListingParent) DeepCopy() *BigQueryAnalyticsHubListingParent {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubListingParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubListingRef) DeepCopyInto(out *BigQueryAnalyticsHubListingRef) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(BigQueryAnalyticsHubListingParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubListingRef.
func (in *BigQueryAnalyticsHubListingRef) DeepCopy() *BigQueryAnalyticsHubListingRef {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubListingRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubListingSpec) DeepCopyInto(out *BigQueryAnalyticsHubListingSpec) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.DataProvider != nil {
		in, out := &in.DataProvider, &out.DataProvider
		*out = new(DataProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(Publisher)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestAccess != nil {
		in, out := &in.RequestAccess, &out.RequestAccess
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubListingSpec.
func (in *BigQueryAnalyticsHubListingSpec) DeepCopy() *BigQueryAnalyticsHubListingSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubListingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubListingStatus) DeepCopyInto(out *BigQueryAnalyticsHubListingStatus) {
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
		*out = new(BigQueryAnalyticsHubListingObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubListingStatus.
func (in *BigQueryAnalyticsHubListingStatus) DeepCopy() *BigQueryAnalyticsHubListingStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubListingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoolValue) DeepCopyInto(out *BoolValue) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoolValue.
func (in *BoolValue) DeepCopy() *BoolValue {
	if in == nil {
		return nil
	}
	out := new(BoolValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExchange) DeepCopyInto(out *DataExchange) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SharingEnvironmentConfig != nil {
		in, out := &in.SharingEnvironmentConfig, &out.SharingEnvironmentConfig
		*out = new(SharingEnvironmentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExchange.
func (in *DataExchange) DeepCopy() *DataExchange {
	if in == nil {
		return nil
	}
	out := new(DataExchange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProvider) DeepCopyInto(out *DataProvider) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProvider.
func (in *DataProvider) DeepCopy() *DataProvider {
	if in == nil {
		return nil
	}
	out := new(DataProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listing) DeepCopyInto(out *Listing) {
	*out = *in
	if in.BigqueryDataset != nil {
		in, out := &in.BigqueryDataset, &out.BigqueryDataset
		*out = new(Listing_BigQueryDatasetSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.DataProvider != nil {
		in, out := &in.DataProvider, &out.DataProvider
		*out = new(DataProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(Publisher)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestAccess != nil {
		in, out := &in.RequestAccess, &out.RequestAccess
		*out = new(string)
		**out = **in
	}
	if in.RestrictedExportConfig != nil {
		in, out := &in.RestrictedExportConfig, &out.RestrictedExportConfig
		*out = new(Listing_RestrictedExportConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listing.
func (in *Listing) DeepCopy() *Listing {
	if in == nil {
		return nil
	}
	out := new(Listing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listing_BigQueryDatasetSource) DeepCopyInto(out *Listing_BigQueryDatasetSource) {
	*out = *in
	if in.Dataset != nil {
		in, out := &in.Dataset, &out.Dataset
		*out = new(string)
		**out = **in
	}
	if in.SelectedResources != nil {
		in, out := &in.SelectedResources, &out.SelectedResources
		*out = make([]Listing_BigQueryDatasetSource_SelectedResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestrictedExportPolicy != nil {
		in, out := &in.RestrictedExportPolicy, &out.RestrictedExportPolicy
		*out = new(Listing_BigQueryDatasetSource_RestrictedExportPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listing_BigQueryDatasetSource.
func (in *Listing_BigQueryDatasetSource) DeepCopy() *Listing_BigQueryDatasetSource {
	if in == nil {
		return nil
	}
	out := new(Listing_BigQueryDatasetSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listing_BigQueryDatasetSource_RestrictedExportPolicy) DeepCopyInto(out *Listing_BigQueryDatasetSource_RestrictedExportPolicy) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(BoolValue)
		(*in).DeepCopyInto(*out)
	}
	if in.RestrictDirectTableAccess != nil {
		in, out := &in.RestrictDirectTableAccess, &out.RestrictDirectTableAccess
		*out = new(BoolValue)
		(*in).DeepCopyInto(*out)
	}
	if in.RestrictQueryResult != nil {
		in, out := &in.RestrictQueryResult, &out.RestrictQueryResult
		*out = new(BoolValue)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listing_BigQueryDatasetSource_RestrictedExportPolicy.
func (in *Listing_BigQueryDatasetSource_RestrictedExportPolicy) DeepCopy() *Listing_BigQueryDatasetSource_RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := new(Listing_BigQueryDatasetSource_RestrictedExportPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listing_BigQueryDatasetSource_SelectedResource) DeepCopyInto(out *Listing_BigQueryDatasetSource_SelectedResource) {
	*out = *in
	if in.Table != nil {
		in, out := &in.Table, &out.Table
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listing_BigQueryDatasetSource_SelectedResource.
func (in *Listing_BigQueryDatasetSource_SelectedResource) DeepCopy() *Listing_BigQueryDatasetSource_SelectedResource {
	if in == nil {
		return nil
	}
	out := new(Listing_BigQueryDatasetSource_SelectedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listing_RestrictedExportConfig) DeepCopyInto(out *Listing_RestrictedExportConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RestrictDirectTableAccess != nil {
		in, out := &in.RestrictDirectTableAccess, &out.RestrictDirectTableAccess
		*out = new(bool)
		**out = **in
	}
	if in.RestrictQueryResult != nil {
		in, out := &in.RestrictQueryResult, &out.RestrictQueryResult
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listing_RestrictedExportConfig.
func (in *Listing_RestrictedExportConfig) DeepCopy() *Listing_RestrictedExportConfig {
	if in == nil {
		return nil
	}
	out := new(Listing_RestrictedExportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Publisher) DeepCopyInto(out *Publisher) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Publisher.
func (in *Publisher) DeepCopy() *Publisher {
	if in == nil {
		return nil
	}
	out := new(Publisher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharingEnvironmentConfig) DeepCopyInto(out *SharingEnvironmentConfig) {
	*out = *in
	if in.DefaultExchangeConfig != nil {
		in, out := &in.DefaultExchangeConfig, &out.DefaultExchangeConfig
		*out = new(SharingEnvironmentConfig_DefaultExchangeConfig)
		**out = **in
	}
	if in.DcrExchangeConfig != nil {
		in, out := &in.DcrExchangeConfig, &out.DcrExchangeConfig
		*out = new(SharingEnvironmentConfig_DcrExchangeConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharingEnvironmentConfig.
func (in *SharingEnvironmentConfig) DeepCopy() *SharingEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := new(SharingEnvironmentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharingEnvironmentConfig_DcrExchangeConfig) DeepCopyInto(out *SharingEnvironmentConfig_DcrExchangeConfig) {
	*out = *in
	if in.SingleSelectedResourceSharingRestriction != nil {
		in, out := &in.SingleSelectedResourceSharingRestriction, &out.SingleSelectedResourceSharingRestriction
		*out = new(bool)
		**out = **in
	}
	if in.SingleLinkedDatasetPerCleanroom != nil {
		in, out := &in.SingleLinkedDatasetPerCleanroom, &out.SingleLinkedDatasetPerCleanroom
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharingEnvironmentConfig_DcrExchangeConfig.
func (in *SharingEnvironmentConfig_DcrExchangeConfig) DeepCopy() *SharingEnvironmentConfig_DcrExchangeConfig {
	if in == nil {
		return nil
	}
	out := new(SharingEnvironmentConfig_DcrExchangeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharingEnvironmentConfig_DefaultExchangeConfig) DeepCopyInto(out *SharingEnvironmentConfig_DefaultExchangeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharingEnvironmentConfig_DefaultExchangeConfig.
func (in *SharingEnvironmentConfig_DefaultExchangeConfig) DeepCopy() *SharingEnvironmentConfig_DefaultExchangeConfig {
	if in == nil {
		return nil
	}
	out := new(SharingEnvironmentConfig_DefaultExchangeConfig)
	in.DeepCopyInto(out)
	return out
}
