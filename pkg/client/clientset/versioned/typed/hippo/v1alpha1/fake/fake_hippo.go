// Copyright  2018 TikalTechnologies.io
// Do not distribute.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/xumak-grid/grid-controller-manager/pkg/apis/hippo/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHippos implements HippoInterface
type FakeHippos struct {
	Fake *FakeHippoV1alpha1
	ns   string
}

var hipposResource = schema.GroupVersionResource{Group: "hippo.xumak.io", Version: "v1alpha1", Resource: "hippos"}

var hipposKind = schema.GroupVersionKind{Group: "hippo.xumak.io", Version: "v1alpha1", Kind: "Hippo"}

// Get takes name of the hippo, and returns the corresponding hippo object, and an error if there is any.
func (c *FakeHippos) Get(name string, options v1.GetOptions) (result *v1alpha1.Hippo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hipposResource, c.ns, name), &v1alpha1.Hippo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hippo), err
}

// List takes label and field selectors, and returns the list of Hippos that match those selectors.
func (c *FakeHippos) List(opts v1.ListOptions) (result *v1alpha1.HippoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hipposResource, hipposKind, c.ns, opts), &v1alpha1.HippoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HippoList{}
	for _, item := range obj.(*v1alpha1.HippoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hippos.
func (c *FakeHippos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hipposResource, c.ns, opts))

}

// Create takes the representation of a hippo and creates it.  Returns the server's representation of the hippo, and an error, if there is any.
func (c *FakeHippos) Create(hippo *v1alpha1.Hippo) (result *v1alpha1.Hippo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hipposResource, c.ns, hippo), &v1alpha1.Hippo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hippo), err
}

// Update takes the representation of a hippo and updates it. Returns the server's representation of the hippo, and an error, if there is any.
func (c *FakeHippos) Update(hippo *v1alpha1.Hippo) (result *v1alpha1.Hippo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hipposResource, c.ns, hippo), &v1alpha1.Hippo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hippo), err
}

// Delete takes name of the hippo and deletes it. Returns an error if one occurs.
func (c *FakeHippos) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hipposResource, c.ns, name), &v1alpha1.Hippo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHippos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hipposResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HippoList{})
	return err
}

// Patch applies the patch and returns the patched hippo.
func (c *FakeHippos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Hippo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hipposResource, c.ns, name, data, subresources...), &v1alpha1.Hippo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hippo), err
}
