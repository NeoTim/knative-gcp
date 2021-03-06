/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/google/knative-gcp/pkg/apis/events/v1alpha1"
	scheme "github.com/google/knative-gcp/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudStorageSourcesGetter has a method to return a CloudStorageSourceInterface.
// A group's client should implement this interface.
type CloudStorageSourcesGetter interface {
	CloudStorageSources(namespace string) CloudStorageSourceInterface
}

// CloudStorageSourceInterface has methods to work with CloudStorageSource resources.
type CloudStorageSourceInterface interface {
	Create(ctx context.Context, cloudStorageSource *v1alpha1.CloudStorageSource, opts v1.CreateOptions) (*v1alpha1.CloudStorageSource, error)
	Update(ctx context.Context, cloudStorageSource *v1alpha1.CloudStorageSource, opts v1.UpdateOptions) (*v1alpha1.CloudStorageSource, error)
	UpdateStatus(ctx context.Context, cloudStorageSource *v1alpha1.CloudStorageSource, opts v1.UpdateOptions) (*v1alpha1.CloudStorageSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CloudStorageSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CloudStorageSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudStorageSource, err error)
	CloudStorageSourceExpansion
}

// cloudStorageSources implements CloudStorageSourceInterface
type cloudStorageSources struct {
	client rest.Interface
	ns     string
}

// newCloudStorageSources returns a CloudStorageSources
func newCloudStorageSources(c *EventsV1alpha1Client, namespace string) *cloudStorageSources {
	return &cloudStorageSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudStorageSource, and returns the corresponding cloudStorageSource object, and an error if there is any.
func (c *cloudStorageSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudStorageSource, err error) {
	result = &v1alpha1.CloudStorageSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudStorageSources that match those selectors.
func (c *cloudStorageSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudStorageSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudStorageSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudStorageSources.
func (c *cloudStorageSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloudStorageSource and creates it.  Returns the server's representation of the cloudStorageSource, and an error, if there is any.
func (c *cloudStorageSources) Create(ctx context.Context, cloudStorageSource *v1alpha1.CloudStorageSource, opts v1.CreateOptions) (result *v1alpha1.CloudStorageSource, err error) {
	result = &v1alpha1.CloudStorageSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudStorageSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloudStorageSource and updates it. Returns the server's representation of the cloudStorageSource, and an error, if there is any.
func (c *cloudStorageSources) Update(ctx context.Context, cloudStorageSource *v1alpha1.CloudStorageSource, opts v1.UpdateOptions) (result *v1alpha1.CloudStorageSource, err error) {
	result = &v1alpha1.CloudStorageSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		Name(cloudStorageSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudStorageSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloudStorageSources) UpdateStatus(ctx context.Context, cloudStorageSource *v1alpha1.CloudStorageSource, opts v1.UpdateOptions) (result *v1alpha1.CloudStorageSource, err error) {
	result = &v1alpha1.CloudStorageSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		Name(cloudStorageSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudStorageSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloudStorageSource and deletes it. Returns an error if one occurs.
func (c *cloudStorageSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudStorageSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudstoragesources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloudStorageSource.
func (c *cloudStorageSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudStorageSource, err error) {
	result = &v1alpha1.CloudStorageSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudstoragesources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
