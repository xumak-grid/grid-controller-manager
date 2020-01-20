package main

import (
	"flag"
	"time"

	"github.com/sirupsen/logrus"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	gridclientset "github.com/xumak-grid/grid-controller-manager/pkg/client/clientset/versioned"
	gridinformers "github.com/xumak-grid/grid-controller-manager/pkg/client/informers/externalversions"
	epcontroller "github.com/xumak-grid/grid-controller-manager/pkg/controllers/elasticpath"
	"github.com/xumak-grid/grid-controller-manager/pkg/util/signals"
)

var (
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to kubeconfig file, required for out of cluster e.g: ~/.kube/config")
}

func main() {
	flag.Parse()
	log := logrus.StandardLogger()
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error building kubernetes config", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("error building kubernetes clientset", err.Error())
	}
	gridClient, err := gridclientset.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("error building grid clientset", err.Error())
	}
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	gridInformerFactory := gridinformers.NewSharedInformerFactory(gridClient, time.Second*30)
	elasticPathController := epcontroller.NewController(kubeClient, gridClient, kubeInformerFactory, gridInformerFactory)
	stopCh := signals.SetupSignalHandler()
	go kubeInformerFactory.Start(stopCh)
	go gridInformerFactory.Start(stopCh)
	elasticPathController.Run(stopCh)
}
