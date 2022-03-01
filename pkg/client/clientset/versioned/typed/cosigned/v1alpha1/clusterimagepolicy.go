// Copyright 2022 The Sigstore Authors.
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/sigstore/cosign/pkg/apis/cosigned/v1alpha1"
	scheme "github.com/sigstore/cosign/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterImagePoliciesGetter has a method to return a ClusterImagePolicyInterface.
// A group's client should implement this interface.
type ClusterImagePoliciesGetter interface {
	ClusterImagePolicies() ClusterImagePolicyInterface
}

// ClusterImagePolicyInterface has methods to work with ClusterImagePolicy resources.
type ClusterImagePolicyInterface interface {
	Create(ctx context.Context, clusterImagePolicy *v1alpha1.ClusterImagePolicy, opts v1.CreateOptions) (*v1alpha1.ClusterImagePolicy, error)
	Update(ctx context.Context, clusterImagePolicy *v1alpha1.ClusterImagePolicy, opts v1.UpdateOptions) (*v1alpha1.ClusterImagePolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterImagePolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterImagePolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterImagePolicy, err error)
	ClusterImagePolicyExpansion
}

// clusterImagePolicies implements ClusterImagePolicyInterface
type clusterImagePolicies struct {
	client rest.Interface
}

// newClusterImagePolicies returns a ClusterImagePolicies
func newClusterImagePolicies(c *CosignedV1alpha1Client) *clusterImagePolicies {
	return &clusterImagePolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterImagePolicy, and returns the corresponding clusterImagePolicy object, and an error if there is any.
func (c *clusterImagePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterImagePolicy, err error) {
	result = &v1alpha1.ClusterImagePolicy{}
	err = c.client.Get().
		Resource("clusterimagepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterImagePolicies that match those selectors.
func (c *clusterImagePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterImagePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterImagePolicyList{}
	err = c.client.Get().
		Resource("clusterimagepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterImagePolicies.
func (c *clusterImagePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterimagepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterImagePolicy and creates it.  Returns the server's representation of the clusterImagePolicy, and an error, if there is any.
func (c *clusterImagePolicies) Create(ctx context.Context, clusterImagePolicy *v1alpha1.ClusterImagePolicy, opts v1.CreateOptions) (result *v1alpha1.ClusterImagePolicy, err error) {
	result = &v1alpha1.ClusterImagePolicy{}
	err = c.client.Post().
		Resource("clusterimagepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterImagePolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterImagePolicy and updates it. Returns the server's representation of the clusterImagePolicy, and an error, if there is any.
func (c *clusterImagePolicies) Update(ctx context.Context, clusterImagePolicy *v1alpha1.ClusterImagePolicy, opts v1.UpdateOptions) (result *v1alpha1.ClusterImagePolicy, err error) {
	result = &v1alpha1.ClusterImagePolicy{}
	err = c.client.Put().
		Resource("clusterimagepolicies").
		Name(clusterImagePolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterImagePolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterImagePolicy and deletes it. Returns an error if one occurs.
func (c *clusterImagePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterimagepolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterImagePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterimagepolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterImagePolicy.
func (c *clusterImagePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterImagePolicy, err error) {
	result = &v1alpha1.ClusterImagePolicy{}
	err = c.client.Patch(pt).
		Resource("clusterimagepolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}