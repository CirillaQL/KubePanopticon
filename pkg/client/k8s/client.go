package k8s

import (
	"flag"
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var K8SClientset *kubernetes.Clientset

func NewK8SClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Log.Error("Kubernetes init cluster config failed", zap.Error(err))
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Log.Error("Kubernetes init clientSet failed.", zap.Error(err))
		return nil, err
	}

	return clientSet, nil
}

func NewK8SOutClusterClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logger.Log.Error("Kubernetes init cluster config out of cluster failed", zap.Error(err))
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Log.Error("Kubernetes init clientSet out of cluster failed.", zap.Error(err))
		return nil, err
	}

	return clientset, nil
}

func init() {
	clientset, err := NewK8SOutClusterClient()
	if err != nil {
		logger.Log.Error("k8s ClientSet init failed", zap.Error(err))
		panic(err)
	}
	K8SClientset = clientset
}
