// Copyright  2018 TikalTechnologies.io
// Do not distribute.

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	glog "github.com/golang/glog"
	databasesv1alpha1 "github.com/xumak-grid/grid-controller-manager/pkg/client/clientset/versioned/typed/databases/v1alpha1"
	elasticpathv1alpha1 "github.com/xumak-grid/grid-controller-manager/pkg/client/clientset/versioned/typed/elasticpath/v1alpha1"
	hippov1alpha1 "github.com/xumak-grid/grid-controller-manager/pkg/client/clientset/versioned/typed/hippo/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	DatabasesV1alpha1() databasesv1alpha1.DatabasesV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Databases() databasesv1alpha1.DatabasesV1alpha1Interface
	ElasticpathV1alpha1() elasticpathv1alpha1.ElasticpathV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Elasticpath() elasticpathv1alpha1.ElasticpathV1alpha1Interface
	HippoV1alpha1() hippov1alpha1.HippoV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Hippo() hippov1alpha1.HippoV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	databasesV1alpha1   *databasesv1alpha1.DatabasesV1alpha1Client
	elasticpathV1alpha1 *elasticpathv1alpha1.ElasticpathV1alpha1Client
	hippoV1alpha1       *hippov1alpha1.HippoV1alpha1Client
}

// DatabasesV1alpha1 retrieves the DatabasesV1alpha1Client
func (c *Clientset) DatabasesV1alpha1() databasesv1alpha1.DatabasesV1alpha1Interface {
	return c.databasesV1alpha1
}

// Deprecated: Databases retrieves the default version of DatabasesClient.
// Please explicitly pick a version.
func (c *Clientset) Databases() databasesv1alpha1.DatabasesV1alpha1Interface {
	return c.databasesV1alpha1
}

// ElasticpathV1alpha1 retrieves the ElasticpathV1alpha1Client
func (c *Clientset) ElasticpathV1alpha1() elasticpathv1alpha1.ElasticpathV1alpha1Interface {
	return c.elasticpathV1alpha1
}

// Deprecated: Elasticpath retrieves the default version of ElasticpathClient.
// Please explicitly pick a version.
func (c *Clientset) Elasticpath() elasticpathv1alpha1.ElasticpathV1alpha1Interface {
	return c.elasticpathV1alpha1
}

// HippoV1alpha1 retrieves the HippoV1alpha1Client
func (c *Clientset) HippoV1alpha1() hippov1alpha1.HippoV1alpha1Interface {
	return c.hippoV1alpha1
}

// Deprecated: Hippo retrieves the default version of HippoClient.
// Please explicitly pick a version.
func (c *Clientset) Hippo() hippov1alpha1.HippoV1alpha1Interface {
	return c.hippoV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.databasesV1alpha1, err = databasesv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.elasticpathV1alpha1, err = elasticpathv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.hippoV1alpha1, err = hippov1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.databasesV1alpha1 = databasesv1alpha1.NewForConfigOrDie(c)
	cs.elasticpathV1alpha1 = elasticpathv1alpha1.NewForConfigOrDie(c)
	cs.hippoV1alpha1 = hippov1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.databasesV1alpha1 = databasesv1alpha1.New(c)
	cs.elasticpathV1alpha1 = elasticpathv1alpha1.New(c)
	cs.hippoV1alpha1 = hippov1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
