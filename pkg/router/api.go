package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAPI(e *gin.Engine) {
	baseAPI := e.Group("/api")
	baseAPI.GET("/healthz", healthz)

	v1 := baseAPI.Group("/v1")
	NodeRouter(v1)

	e.Run(":8080")
}

func NodeRouter(group *gin.RouterGroup) {
	nodeGroup := group.Group("/nodes")
	nodeGroup.GET("/")
}

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
