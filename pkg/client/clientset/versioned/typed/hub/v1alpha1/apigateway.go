/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2024 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	scheme "github.com/traefik/hub-crds/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// APIGatewaysGetter has a method to return a APIGatewayInterface.
// A group's client should implement this interface.
type APIGatewaysGetter interface {
	APIGateways() APIGatewayInterface
}

// APIGatewayInterface has methods to work with APIGateway resources.
type APIGatewayInterface interface {
	Create(ctx context.Context, aPIGateway *v1alpha1.APIGateway, opts v1.CreateOptions) (*v1alpha1.APIGateway, error)
	Update(ctx context.Context, aPIGateway *v1alpha1.APIGateway, opts v1.UpdateOptions) (*v1alpha1.APIGateway, error)
	UpdateStatus(ctx context.Context, aPIGateway *v1alpha1.APIGateway, opts v1.UpdateOptions) (*v1alpha1.APIGateway, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIGateway, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIGatewayList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIGateway, err error)
	APIGatewayExpansion
}

// aPIGateways implements APIGatewayInterface
type aPIGateways struct {
	client rest.Interface
}

// newAPIGateways returns a APIGateways
func newAPIGateways(c *HubV1alpha1Client) *aPIGateways {
	return &aPIGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPIGateway, and returns the corresponding aPIGateway object, and an error if there is any.
func (c *aPIGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIGateway, err error) {
	result = &v1alpha1.APIGateway{}
	err = c.client.Get().
		Resource("apigateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIGateways that match those selectors.
func (c *aPIGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIGatewayList{}
	err = c.client.Get().
		Resource("apigateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIGateways.
func (c *aPIGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apigateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIGateway and creates it.  Returns the server's representation of the aPIGateway, and an error, if there is any.
func (c *aPIGateways) Create(ctx context.Context, aPIGateway *v1alpha1.APIGateway, opts v1.CreateOptions) (result *v1alpha1.APIGateway, err error) {
	result = &v1alpha1.APIGateway{}
	err = c.client.Post().
		Resource("apigateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIGateway).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIGateway and updates it. Returns the server's representation of the aPIGateway, and an error, if there is any.
func (c *aPIGateways) Update(ctx context.Context, aPIGateway *v1alpha1.APIGateway, opts v1.UpdateOptions) (result *v1alpha1.APIGateway, err error) {
	result = &v1alpha1.APIGateway{}
	err = c.client.Put().
		Resource("apigateways").
		Name(aPIGateway.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIGateway).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPIGateways) UpdateStatus(ctx context.Context, aPIGateway *v1alpha1.APIGateway, opts v1.UpdateOptions) (result *v1alpha1.APIGateway, err error) {
	result = &v1alpha1.APIGateway{}
	err = c.client.Put().
		Resource("apigateways").
		Name(aPIGateway.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIGateway).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIGateway and deletes it. Returns an error if one occurs.
func (c *aPIGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apigateways").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apigateways").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIGateway.
func (c *aPIGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIGateway, err error) {
	result = &v1alpha1.APIGateway{}
	err = c.client.Patch(pt).
		Resource("apigateways").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
