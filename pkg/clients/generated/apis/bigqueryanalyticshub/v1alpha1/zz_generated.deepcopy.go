//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchange) DeepCopyInto(out *BigQueryAnalyticsHubDataExchange) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
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
	return
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
func (in *BigQueryAnalyticsHubDataExchangeSpec) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
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
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(DataexchangeObservedStateStatus)
		(*in).DeepCopyInto(*out)
	}
	return
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
	return
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
	return
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
func (in *BigQueryAnalyticsHubListingSpec) DeepCopyInto(out *BigQueryAnalyticsHubListingSpec) {
	*out = *in
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DataProvider != nil {
		in, out := &in.DataProvider, &out.DataProvider
		*out = new(ListingDataProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(ListingPublisher)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestAccess != nil {
		in, out := &in.RequestAccess, &out.RequestAccess
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
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
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(ListingObservedStateStatus)
		(*in).DeepCopyInto(*out)
	}
	return
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
func (in *DataexchangeObservedStateStatus) DeepCopyInto(out *DataexchangeObservedStateStatus) {
	*out = *in
	if in.ListingCount != nil {
		in, out := &in.ListingCount, &out.ListingCount
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataexchangeObservedStateStatus.
func (in *DataexchangeObservedStateStatus) DeepCopy() *DataexchangeObservedStateStatus {
	if in == nil {
		return nil
	}
	out := new(DataexchangeObservedStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingDataProvider) DeepCopyInto(out *ListingDataProvider) {
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingDataProvider.
func (in *ListingDataProvider) DeepCopy() *ListingDataProvider {
	if in == nil {
		return nil
	}
	out := new(ListingDataProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingObservedStateStatus) DeepCopyInto(out *ListingObservedStateStatus) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingObservedStateStatus.
func (in *ListingObservedStateStatus) DeepCopy() *ListingObservedStateStatus {
	if in == nil {
		return nil
	}
	out := new(ListingObservedStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingPublisher) DeepCopyInto(out *ListingPublisher) {
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingPublisher.
func (in *ListingPublisher) DeepCopy() *ListingPublisher {
	if in == nil {
		return nil
	}
	out := new(ListingPublisher)
	in.DeepCopyInto(out)
	return out
}
