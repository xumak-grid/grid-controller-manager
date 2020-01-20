package controller

import (
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"

	gridclientset "github.com/xumak-grid/grid-controller-manager/pkg/client/clientset/versioned"
	gridinformers "github.com/xumak-grid/grid-controller-manager/pkg/client/informers/externalversions"
	listers "github.com/xumak-grid/grid-controller-manager/pkg/client/listers/elasticpath/v1alpha1"
)

// Controller is the controller implementation for ElasticPath resources.
type Controller struct {
	// kubeclienset is a standard kubernetes clientset
	kubeclienset kubernetes.Interface
	// gridclienset is a clientset for our APIs
	gridclienset gridclientset.Interface

	deploymentLister appslisters.DeploymentLister

	epLister listers.DeploymentLister
}

// NewController returns a new ElasticPath controller.
func NewController(
	kubeclientset kubernetes.Interface,
	gridclientset gridclientset.Interface,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
	gridInformerFactory gridinformers.SharedInformerFactory) *Controller {

	deploymentInformer := kubeInformerFactory.Apps().V1().Deployments()
	elasticPathInformer := gridInformerFactory.Elasticpath().V1alpha1().Deployments()

	controller := &Controller{
		kubeclienset:     kubeclientset,
		gridclienset:     gridclientset,
		deploymentLister: deploymentInformer.Lister(),
		epLister:         elasticPathInformer.Lister(),
	}

	return controller
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	return
}
