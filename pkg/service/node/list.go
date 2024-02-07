package node

import (
	"github.com/CirillaQL/kubepanopticon/pkg/client/k8s"
	"github.com/CirillaQL/kubepanopticon/pkg/constant"
	"github.com/CirillaQL/kubepanopticon/pkg/service/cache/informer"
	"github.com/CirillaQL/kubepanopticon/pkg/service/response"
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
)

type NodeListItem struct {
	Node   *v1.Node `json:"node"`
	Status []string `json:"status"`
}

func List(c *gin.Context) {
	logger.Log.Info("get node list")
	if k8s.K8SClientset == nil {
		logger.Log.Error("Informer can't load k8sClientset, k8sClientset is empty")
	}

	nodeIndexInformer := informer.ResourceIndexInformerMap["node"]
	if nodeIndexInformer == nil {
		logger.Log.Error("can't get nodeIndexInformer, it is nil")
		response.Response(c, constant.OK, nil, "nodeIndexInformer doesn't exist")
		return
	}

	var nodeList []*NodeListItem

	nodeIndexLister := nodeIndexInformer.GetIndexer()
	for _, obj := range nodeIndexLister.List() {
		node, ok := obj.(*v1.Node)
		if !ok {
			logger.Log.Warn("failed to convert interface{} into node", zap.Any("node_info", obj))
		}
		conditions := handleNodeConditions(node)
		item := &NodeListItem{
			Node:   node,
			Status: conditions,
		}
		nodeList = append(nodeList, item)
	}

	response.Response(c, constant.OK, nodeList, "ok")
	logger.Log.Info("finish get node list")
}

func handleNodeConditions(node *v1.Node) []string {
	var conditions []string
	for _, condition := range node.Status.Conditions {
		if condition.Status == v1.ConditionTrue {
			conditions = append(conditions, string(condition.Type))
		}
	}
	return conditions
}
