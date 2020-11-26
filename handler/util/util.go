package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessDataResponser : サクセスデータレスポンス生成とレスポンス実行
func SuccessDataResponser(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   data,
	})
}

// SuccessMessageResponser : サクセスメッセージレスポンス生成とレスポンス実行
func SuccessMessageResponser(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  http.StatusText(http.StatusOK),
		"message": message,
	})
}

// ErrorResponser : エラーレスポンス生成とレスポンス実行
func ErrorResponser(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"status":  http.StatusText(code),
		"message": message,
	})
}
