package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type successResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type successMsgResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Success : レスポンス生成とレスポンス実行
func Success(c *gin.Context, data interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, data)
		return
	}
	c.JSON(http.StatusOK, successMsgResponse{
		Status: http.StatusText(http.StatusOK),
	})
	return

}

// Error : レスポンス生成とレスポンス実行
func Error(c *gin.Context, httpCode, code int, msg string) {
	c.JSON(httpCode, errorResponse{
		Code:    code,
		Status:  http.StatusText(httpCode),
		Message: msg,
	})
	return
}
