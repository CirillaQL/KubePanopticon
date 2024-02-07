package informer

import (
	"errors"
	"github.com/CirillaQL/kubepanopticon/pkg/client/k8s"
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"log"
	"time"
)

var ResourceIndexInformerMap map[string]cache.SharedIndexInformer

func InitInformer() error {
	if k8s.K8SClientset == nil {
		logger.Log.Error("Informer can't load k8sClientset, k8sClientset is empty")
		return errors.New("can't init k8s informer")
	}

	ResourceIndexInformerMap = make(map[string]cache.SharedIndexInformer)

	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformersFactory := informers.NewSharedInformerFactory(k8s.K8SClientset, 3*time.Minute)
	nodeInformer := sharedInformersFactory.Core().V1().Nodes().Informer()

	ResourceIndexInformerMap["node"] = nodeInformer

	go nodeInformer.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, nodeInformer.HasSynced) {
		log.Fatal("Error syncing cache")
	}

	return nil
}
