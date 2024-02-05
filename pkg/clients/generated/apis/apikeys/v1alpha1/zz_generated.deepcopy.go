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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysKey) DeepCopyInto(out *APIKeysKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysKey.
func (in *APIKeysKey) DeepCopy() *APIKeysKey {
	if in == nil {
		return nil
	}
	out := new(APIKeysKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKeysKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysKeyList) DeepCopyInto(out *APIKeysKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIKeysKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysKeyList.
func (in *APIKeysKeyList) DeepCopy() *APIKeysKeyList {
	if in == nil {
		return nil
	}
	out := new(APIKeysKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKeysKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysKeySpec) DeepCopyInto(out *APIKeysKeySpec) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Restrictions != nil {
		in, out := &in.Restrictions, &out.Restrictions
		*out = new(KeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysKeySpec.
func (in *APIKeysKeySpec) DeepCopy() *APIKeysKeySpec {
	if in == nil {
		return nil
	}
	out := new(APIKeysKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysKeyStatus) DeepCopyInto(out *APIKeysKeyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.KeyString != nil {
		in, out := &in.KeyString, &out.KeyString
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysKeyStatus.
func (in *APIKeysKeyStatus) DeepCopy() *APIKeysKeyStatus {
	if in == nil {
		return nil
	}
	out := new(APIKeysKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyAllowedApplications) DeepCopyInto(out *KeyAllowedApplications) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyAllowedApplications.
func (in *KeyAllowedApplications) DeepCopy() *KeyAllowedApplications {
	if in == nil {
		return nil
	}
	out := new(KeyAllowedApplications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyAndroidKeyRestrictions) DeepCopyInto(out *KeyAndroidKeyRestrictions) {
	*out = *in
	if in.AllowedApplications != nil {
		in, out := &in.AllowedApplications, &out.AllowedApplications
		*out = make([]KeyAllowedApplications, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyAndroidKeyRestrictions.
func (in *KeyAndroidKeyRestrictions) DeepCopy() *KeyAndroidKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(KeyAndroidKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyApiTargets) DeepCopyInto(out *KeyApiTargets) {
	*out = *in
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyApiTargets.
func (in *KeyApiTargets) DeepCopy() *KeyApiTargets {
	if in == nil {
		return nil
	}
	out := new(KeyApiTargets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyBrowserKeyRestrictions) DeepCopyInto(out *KeyBrowserKeyRestrictions) {
	*out = *in
	if in.AllowedReferrers != nil {
		in, out := &in.AllowedReferrers, &out.AllowedReferrers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyBrowserKeyRestrictions.
func (in *KeyBrowserKeyRestrictions) DeepCopy() *KeyBrowserKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(KeyBrowserKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyIosKeyRestrictions) DeepCopyInto(out *KeyIosKeyRestrictions) {
	*out = *in
	if in.AllowedBundleIds != nil {
		in, out := &in.AllowedBundleIds, &out.AllowedBundleIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyIosKeyRestrictions.
func (in *KeyIosKeyRestrictions) DeepCopy() *KeyIosKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(KeyIosKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyRestrictions) DeepCopyInto(out *KeyRestrictions) {
	*out = *in
	if in.AndroidKeyRestrictions != nil {
		in, out := &in.AndroidKeyRestrictions, &out.AndroidKeyRestrictions
		*out = new(KeyAndroidKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiTargets != nil {
		in, out := &in.ApiTargets, &out.ApiTargets
		*out = make([]KeyApiTargets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BrowserKeyRestrictions != nil {
		in, out := &in.BrowserKeyRestrictions, &out.BrowserKeyRestrictions
		*out = new(KeyBrowserKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	if in.IosKeyRestrictions != nil {
		in, out := &in.IosKeyRestrictions, &out.IosKeyRestrictions
		*out = new(KeyIosKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerKeyRestrictions != nil {
		in, out := &in.ServerKeyRestrictions, &out.ServerKeyRestrictions
		*out = new(KeyServerKeyRestrictions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyRestrictions.
func (in *KeyRestrictions) DeepCopy() *KeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(KeyRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyServerKeyRestrictions) DeepCopyInto(out *KeyServerKeyRestrictions) {
	*out = *in
	if in.AllowedIps != nil {
		in, out := &in.AllowedIps, &out.AllowedIps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyServerKeyRestrictions.
func (in *KeyServerKeyRestrictions) DeepCopy() *KeyServerKeyRestrictions {
	if in == nil {
		return nil
	}
	out := new(KeyServerKeyRestrictions)
	in.DeepCopyInto(out)
	return out
}
