// Copyright  2018 TikalTechnologies.io
// Do not distribute.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/xumak-grid/grid-controller-manager/pkg/apis/elasticpath/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeployments implements DeploymentInterface
type FakeDeployments struct {
	Fake *FakeElasticpathV1alpha1
	ns   string
}

var deploymentsResource = schema.GroupVersionResource{Group: "elasticpath.xumak.io", Version: "v1alpha1", Resource: "deployments"}

var deploymentsKind = schema.GroupVersionKind{Group: "elasticpath.xumak.io", Version: "v1alpha1", Kind: "Deployment"}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *FakeDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentsResource, c.ns, name), &v1alpha1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *FakeDeployments) List(opts v1.ListOptions) (result *v1alpha1.DeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentsResource, deploymentsKind, c.ns, opts), &v1alpha1.DeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeploymentList{}
	for _, item := range obj.(*v1alpha1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *FakeDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentsResource, c.ns, opts))

}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Create(deployment *v1alpha1.Deployment) (result *v1alpha1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentsResource, c.ns, deployment), &v1alpha1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployment), err
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Update(deployment *v1alpha1.Deployment) (result *v1alpha1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentsResource, c.ns, deployment), &v1alpha1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployment), err
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *FakeDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deploymentsResource, c.ns, name), &v1alpha1.Deployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeploymentList{})
	return err
}

// Patch applies the patch and returns the patched deployment.
func (c *FakeDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentsResource, c.ns, name, data, subresources...), &v1alpha1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Deployment), err
}
