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

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/run/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRunJobs implements RunJobInterface
type FakeRunJobs struct {
	Fake *FakeRunV1beta1
	ns   string
}

var runjobsResource = schema.GroupVersionResource{Group: "run.cnrm.cloud.google.com", Version: "v1beta1", Resource: "runjobs"}

var runjobsKind = schema.GroupVersionKind{Group: "run.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RunJob"}

// Get takes name of the runJob, and returns the corresponding runJob object, and an error if there is any.
func (c *FakeRunJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.RunJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(runjobsResource, c.ns, name), &v1beta1.RunJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RunJob), err
}

// List takes label and field selectors, and returns the list of RunJobs that match those selectors.
func (c *FakeRunJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RunJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(runjobsResource, runjobsKind, c.ns, opts), &v1beta1.RunJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RunJobList{ListMeta: obj.(*v1beta1.RunJobList).ListMeta}
	for _, item := range obj.(*v1beta1.RunJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runJobs.
func (c *FakeRunJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(runjobsResource, c.ns, opts))

}

// Create takes the representation of a runJob and creates it.  Returns the server's representation of the runJob, and an error, if there is any.
func (c *FakeRunJobs) Create(ctx context.Context, runJob *v1beta1.RunJob, opts v1.CreateOptions) (result *v1beta1.RunJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(runjobsResource, c.ns, runJob), &v1beta1.RunJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RunJob), err
}

// Update takes the representation of a runJob and updates it. Returns the server's representation of the runJob, and an error, if there is any.
func (c *FakeRunJobs) Update(ctx context.Context, runJob *v1beta1.RunJob, opts v1.UpdateOptions) (result *v1beta1.RunJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(runjobsResource, c.ns, runJob), &v1beta1.RunJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RunJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRunJobs) UpdateStatus(ctx context.Context, runJob *v1beta1.RunJob, opts v1.UpdateOptions) (*v1beta1.RunJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(runjobsResource, "status", c.ns, runJob), &v1beta1.RunJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RunJob), err
}

// Delete takes name of the runJob and deletes it. Returns an error if one occurs.
func (c *FakeRunJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(runjobsResource, c.ns, name, opts), &v1beta1.RunJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRunJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(runjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.RunJobList{})
	return err
}

// Patch applies the patch and returns the patched runJob.
func (c *FakeRunJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.RunJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(runjobsResource, c.ns, name, pt, data, subresources...), &v1beta1.RunJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RunJob), err
}
