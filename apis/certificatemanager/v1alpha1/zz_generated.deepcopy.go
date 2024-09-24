//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateManagerDNSAuthorization) DeepCopyInto(out *CertificateManagerDNSAuthorization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateManagerDNSAuthorization.
func (in *CertificateManagerDNSAuthorization) DeepCopy() *CertificateManagerDNSAuthorization {
	if in == nil {
		return nil
	}
	out := new(CertificateManagerDNSAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateManagerDNSAuthorization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateManagerDNSAuthorizationList) DeepCopyInto(out *CertificateManagerDNSAuthorizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateManagerDNSAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateManagerDNSAuthorizationList.
func (in *CertificateManagerDNSAuthorizationList) DeepCopy() *CertificateManagerDNSAuthorizationList {
	if in == nil {
		return nil
	}
	out := new(CertificateManagerDNSAuthorizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateManagerDNSAuthorizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateManagerDNSAuthorizationSpec) DeepCopyInto(out *CertificateManagerDNSAuthorizationSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateManagerDNSAuthorizationSpec.
func (in *CertificateManagerDNSAuthorizationSpec) DeepCopy() *CertificateManagerDNSAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateManagerDNSAuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateManagerDNSAuthorizationStatus) DeepCopyInto(out *CertificateManagerDNSAuthorizationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.DnsResourceRecord != nil {
		in, out := &in.DnsResourceRecord, &out.DnsResourceRecord
		*out = make([]DnsauthorizationDnsResourceRecordStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateManagerDNSAuthorizationStatus.
func (in *CertificateManagerDNSAuthorizationStatus) DeepCopy() *CertificateManagerDNSAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateManagerDNSAuthorizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnsAuthorization_DnsResourceRecord) DeepCopyInto(out *DnsAuthorization_DnsResourceRecord) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnsAuthorization_DnsResourceRecord.
func (in *DnsAuthorization_DnsResourceRecord) DeepCopy() *DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := new(DnsAuthorization_DnsResourceRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnsauthorizationDnsResourceRecordStatus) DeepCopyInto(out *DnsauthorizationDnsResourceRecordStatus) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnsauthorizationDnsResourceRecordStatus.
func (in *DnsauthorizationDnsResourceRecordStatus) DeepCopy() *DnsauthorizationDnsResourceRecordStatus {
	if in == nil {
		return nil
	}
	out := new(DnsauthorizationDnsResourceRecordStatus)
	in.DeepCopyInto(out)
	return out
}
