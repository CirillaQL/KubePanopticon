package cluster

import (
	"github.com/CirillaQL/kubepanopticon/pkg/client/prometheus"
	"github.com/CirillaQL/kubepanopticon/pkg/constant"
	"github.com/CirillaQL/kubepanopticon/pkg/constant/errorcode"
	"github.com/CirillaQL/kubepanopticon/pkg/service/response"
	"github.com/CirillaQL/kubepanopticon/utils/config"
	"github.com/CirillaQL/kubepanopticon/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func PrometheusClusterAllPodCoreUsage(c *gin.Context) {
	logger.Log.Info("check cluster core")
	prometheusClient, err := prometheus.NewPrometheusClient(config.GlobalConfig.PrometheusURL)
	if err != nil {
		logger.Log.Error("failed init prometheus client when try to get cluster all core", zap.Error(err))
		response.Response(c, errorcode.ErrPrometheusClientInit, nil, "failed init prometheus client when try to get cluster all core")
		return
	}

	startTime := time.Now().Add(-6 * time.Hour)

	result := prometheus.QueryWithTimeRange(prometheusClient, startTime, time.Now(), time.Minute, prometheus.Cluster_All_Pods_Cpu_Usage)

	response.Response(c, constant.OK, result, "ok")

}
