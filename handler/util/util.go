package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponser : サクセスレスポンス生成とレスポンス実行
func SuccessResponser(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   data,
	})
}

// ErrorResponser : エラーレスポンス生成とレスポンス実行
func ErrorResponser(c *gin.Context, code int, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"status":  http.StatusText(code),
		"message": message,
	})
}
