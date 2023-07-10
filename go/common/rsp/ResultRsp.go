package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 正确状态处理
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

// Error 错误状态处理
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK,
		gin.H{
			"code": 400,
			"msg":  msg,
			"data": ' ',
		})
}
