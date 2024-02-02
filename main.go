package main

import (
	"github.com/CirillaQL/kubepanopticon/pkg/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	router.InitAPI(engine)

	//k8sClient, err := k8s.NewK8SOutClusterClient()
	//if err != nil {
	//	logger.Log.Error("Init k8s clientSet failed", zap.Error(err))
	//	panic(err)
	//}
	//namespaces, err := k8sClient.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//for _, namespace := range namespaces.Items {
	//	logger.Log.Info("获取到namespace信息", zap.String("namespace_name", namespace.Name))
	//}

}
