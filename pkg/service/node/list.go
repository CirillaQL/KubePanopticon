package node

import (
	"context"
	"github.com/CirillaQL/kubepanopticon/pkg/client/k8s"
	"github.com/CirillaQL/kubepanopticon/pkg/service/response"
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NodeList(c *gin.Context) {
	if k8s.K8SClientset == nil {
		logger.Log.Error("Informer can't load k8sClientset, k8sClientset is empty")
	}

	nodeList, err := k8s.K8SClientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		logger.Log.Error("")
		response.Response(c, 200, nil, err.Error())
	}

	response.Response(c, 200, nodeList.Items, "ok")
}
