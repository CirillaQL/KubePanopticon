package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, code int, data any, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
