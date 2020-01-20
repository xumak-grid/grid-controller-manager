package databases

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	gridclientset "github.com/xumak-grid/grid-controller-manager/pkg/client/clientset/versioned"
	gridinformers "github.com/xumak-grid/grid-controller-manager/pkg/client/informers/externalversions"
	listers "github.com/xumak-grid/grid-controller-manager/pkg/client/listers/databases/v1alpha1"
	"github.com/xumak-grid/grid-controller-manager/pkg/cloudprovider"
)

// Controller is the controller implementation for Database resources.
type Controller struct {
	// kubeclienset is a standard kubernetes clientset
	kubeclienset kubernetes.Interface
	// gridclienset is a clientset for our APIs
	gridclienset gridclientset.Interface

	deploymentLister appslisters.DeploymentLister
	deploymentSynced cache.InformerSynced

	dbLister  listers.DatabaseLister
	dbSynced  cache.InformerSynced
	dbService cloudprovider.DatabaseService

	workqueue workqueue.RateLimitingInterface
}

// NewController returns a new ElasticPath controller.
func NewController(
	kubeclientset kubernetes.Interface,
	gridclientset gridclientset.Interface,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
	gridInformerFactory gridinformers.SharedInformerFactory,
	dbService cloudprovider.DatabaseService) *Controller {

	deploymentInformer := kubeInformerFactory.Apps().V1().Deployments()
	databaseInformer := gridInformerFactory.Databases().V1alpha1().Databases()

	controller := &Controller{
		kubeclienset:     kubeclientset,
		gridclienset:     gridclientset,
		deploymentLister: deploymentInformer.Lister(),
		deploymentSynced: deploymentInformer.Informer().HasSynced,
		dbLister:         databaseInformer.Lister(),
		dbSynced:         databaseInformer.Informer().HasSynced,
		dbService:        dbService,
		workqueue:        workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Databases"),
	}

	return controller
}

// Run setups handlers and waits for first synced cache.
func (c *Controller) Run(stopCh <-chan struct{}) {
	defer c.workqueue.ShutDown()
	if ok := cache.WaitForCacheSync(stopCh, c.deploymentSynced, c.dbSynced); !ok {
		return
	}
	wait.Until(c.runWoker, time.Second, stopCh)
	<-stopCh
}

func (c *Controller) runWoker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()
	// queue is shutdown phase
	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		return nil
	}(obj)
	if err != nil {
		return false
	}
	return true
}
