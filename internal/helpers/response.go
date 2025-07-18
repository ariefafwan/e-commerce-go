package helpers

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, code int, data any, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

func Error(c *gin.Context, code int, data any, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}
