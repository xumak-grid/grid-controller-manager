// Copyright  2018 TikalTechnologies.io
// Do not distribute.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/xumak-grid/grid-controller-manager/pkg/apis/hippo/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HippoLister helps list Hippos.
type HippoLister interface {
	// List lists all Hippos in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Hippo, err error)
	// Hippos returns an object that can list and get Hippos.
	Hippos(namespace string) HippoNamespaceLister
	HippoListerExpansion
}

// hippoLister implements the HippoLister interface.
type hippoLister struct {
	indexer cache.Indexer
}

// NewHippoLister returns a new HippoLister.
func NewHippoLister(indexer cache.Indexer) HippoLister {
	return &hippoLister{indexer: indexer}
}

// List lists all Hippos in the indexer.
func (s *hippoLister) List(selector labels.Selector) (ret []*v1alpha1.Hippo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Hippo))
	})
	return ret, err
}

// Hippos returns an object that can list and get Hippos.
func (s *hippoLister) Hippos(namespace string) HippoNamespaceLister {
	return hippoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HippoNamespaceLister helps list and get Hippos.
type HippoNamespaceLister interface {
	// List lists all Hippos in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Hippo, err error)
	// Get retrieves the Hippo from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Hippo, error)
	HippoNamespaceListerExpansion
}

// hippoNamespaceLister implements the HippoNamespaceLister
// interface.
type hippoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Hippos in the indexer for a given namespace.
func (s hippoNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Hippo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Hippo))
	})
	return ret, err
}

// Get retrieves the Hippo from the indexer for a given namespace and name.
func (s hippoNamespaceLister) Get(name string) (*v1alpha1.Hippo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hippo"), name)
	}
	return obj.(*v1alpha1.Hippo), nil
}
