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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/securesourcemanager/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecureSourceManagerInstancesGetter has a method to return a SecureSourceManagerInstanceInterface.
// A group's client should implement this interface.
type SecureSourceManagerInstancesGetter interface {
	SecureSourceManagerInstances(namespace string) SecureSourceManagerInstanceInterface
}

// SecureSourceManagerInstanceInterface has methods to work with SecureSourceManagerInstance resources.
type SecureSourceManagerInstanceInterface interface {
	Create(ctx context.Context, secureSourceManagerInstance *v1alpha1.SecureSourceManagerInstance, opts v1.CreateOptions) (*v1alpha1.SecureSourceManagerInstance, error)
	Update(ctx context.Context, secureSourceManagerInstance *v1alpha1.SecureSourceManagerInstance, opts v1.UpdateOptions) (*v1alpha1.SecureSourceManagerInstance, error)
	UpdateStatus(ctx context.Context, secureSourceManagerInstance *v1alpha1.SecureSourceManagerInstance, opts v1.UpdateOptions) (*v1alpha1.SecureSourceManagerInstance, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SecureSourceManagerInstance, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SecureSourceManagerInstanceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecureSourceManagerInstance, err error)
	SecureSourceManagerInstanceExpansion
}

// secureSourceManagerInstances implements SecureSourceManagerInstanceInterface
type secureSourceManagerInstances struct {
	client rest.Interface
	ns     string
}

// newSecureSourceManagerInstances returns a SecureSourceManagerInstances
func newSecureSourceManagerInstances(c *SecuresourcemanagerV1alpha1Client, namespace string) *secureSourceManagerInstances {
	return &secureSourceManagerInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secureSourceManagerInstance, and returns the corresponding secureSourceManagerInstance object, and an error if there is any.
func (c *secureSourceManagerInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecureSourceManagerInstance, err error) {
	result = &v1alpha1.SecureSourceManagerInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecureSourceManagerInstances that match those selectors.
func (c *secureSourceManagerInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecureSourceManagerInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecureSourceManagerInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secureSourceManagerInstances.
func (c *secureSourceManagerInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a secureSourceManagerInstance and creates it.  Returns the server's representation of the secureSourceManagerInstance, and an error, if there is any.
func (c *secureSourceManagerInstances) Create(ctx context.Context, secureSourceManagerInstance *v1alpha1.SecureSourceManagerInstance, opts v1.CreateOptions) (result *v1alpha1.SecureSourceManagerInstance, err error) {
	result = &v1alpha1.SecureSourceManagerInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secureSourceManagerInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a secureSourceManagerInstance and updates it. Returns the server's representation of the secureSourceManagerInstance, and an error, if there is any.
func (c *secureSourceManagerInstances) Update(ctx context.Context, secureSourceManagerInstance *v1alpha1.SecureSourceManagerInstance, opts v1.UpdateOptions) (result *v1alpha1.SecureSourceManagerInstance, err error) {
	result = &v1alpha1.SecureSourceManagerInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		Name(secureSourceManagerInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secureSourceManagerInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *secureSourceManagerInstances) UpdateStatus(ctx context.Context, secureSourceManagerInstance *v1alpha1.SecureSourceManagerInstance, opts v1.UpdateOptions) (result *v1alpha1.SecureSourceManagerInstance, err error) {
	result = &v1alpha1.SecureSourceManagerInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		Name(secureSourceManagerInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secureSourceManagerInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the secureSourceManagerInstance and deletes it. Returns an error if one occurs.
func (c *secureSourceManagerInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secureSourceManagerInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched secureSourceManagerInstance.
func (c *secureSourceManagerInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecureSourceManagerInstance, err error) {
	result = &v1alpha1.SecureSourceManagerInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("securesourcemanagerinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
