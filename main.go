package main

import (
	"flag"
	"github.com/CirillaQL/kubepanopticon/pkg/router"
	"github.com/CirillaQL/kubepanopticon/pkg/service/cache/informer"
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"k8s.io/klog/v2"
)

func main() {
	flag.Parse()

	// Handle klog output
	klog.SetLogger(zapr.NewLogger(logger.Log.Logger))

	err := informer.InitInformer()
	if err != nil {
		logger.Log.Error("failed init informer", zap.Error(err))
		panic(err)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())

	router.InitAPI(engine)

}
