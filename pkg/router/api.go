package router

import (
	"github.com/CirillaQL/kubepanopticon/pkg/service/node"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAPI(e *gin.Engine) {
	e.Use(Cors())

	baseAPI := e.Group("/api")
	baseAPI.GET("/healthz", healthz)

	v1 := baseAPI.Group("/v1")
	NodeRouter(v1)

	e.Run(":7000")
}

func NodeRouter(group *gin.RouterGroup) {
	nodeGroup := group.Group("/nodes")
	nodeGroup.GET("/", node.List)
}

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
