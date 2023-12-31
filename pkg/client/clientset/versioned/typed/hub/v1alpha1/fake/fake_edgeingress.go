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

package fake

import (
	"context"

	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEdgeIngresses implements EdgeIngressInterface
type FakeEdgeIngresses struct {
	Fake *FakeHubV1alpha1
	ns   string
}

var edgeingressesResource = v1alpha1.SchemeGroupVersion.WithResource("edgeingresses")

var edgeingressesKind = v1alpha1.SchemeGroupVersion.WithKind("EdgeIngress")

// Get takes name of the edgeIngress, and returns the corresponding edgeIngress object, and an error if there is any.
func (c *FakeEdgeIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EdgeIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(edgeingressesResource, c.ns, name), &v1alpha1.EdgeIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeIngress), err
}

// List takes label and field selectors, and returns the list of EdgeIngresses that match those selectors.
func (c *FakeEdgeIngresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EdgeIngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(edgeingressesResource, edgeingressesKind, c.ns, opts), &v1alpha1.EdgeIngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EdgeIngressList{ListMeta: obj.(*v1alpha1.EdgeIngressList).ListMeta}
	for _, item := range obj.(*v1alpha1.EdgeIngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested edgeIngresses.
func (c *FakeEdgeIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(edgeingressesResource, c.ns, opts))

}

// Create takes the representation of a edgeIngress and creates it.  Returns the server's representation of the edgeIngress, and an error, if there is any.
func (c *FakeEdgeIngresses) Create(ctx context.Context, edgeIngress *v1alpha1.EdgeIngress, opts v1.CreateOptions) (result *v1alpha1.EdgeIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(edgeingressesResource, c.ns, edgeIngress), &v1alpha1.EdgeIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeIngress), err
}

// Update takes the representation of a edgeIngress and updates it. Returns the server's representation of the edgeIngress, and an error, if there is any.
func (c *FakeEdgeIngresses) Update(ctx context.Context, edgeIngress *v1alpha1.EdgeIngress, opts v1.UpdateOptions) (result *v1alpha1.EdgeIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(edgeingressesResource, c.ns, edgeIngress), &v1alpha1.EdgeIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeIngress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEdgeIngresses) UpdateStatus(ctx context.Context, edgeIngress *v1alpha1.EdgeIngress, opts v1.UpdateOptions) (*v1alpha1.EdgeIngress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(edgeingressesResource, "status", c.ns, edgeIngress), &v1alpha1.EdgeIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeIngress), err
}

// Delete takes name of the edgeIngress and deletes it. Returns an error if one occurs.
func (c *FakeEdgeIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(edgeingressesResource, c.ns, name, opts), &v1alpha1.EdgeIngress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEdgeIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(edgeingressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EdgeIngressList{})
	return err
}

// Patch applies the patch and returns the patched edgeIngress.
func (c *FakeEdgeIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EdgeIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(edgeingressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EdgeIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeIngress), err
}
