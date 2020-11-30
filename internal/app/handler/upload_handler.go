package handler

import (
	"github.com/gin-gonic/gin"
)

// IUploadHandler : インターフェース
type IUploadHandler interface {
	UploadImage(c *gin.Context)
}

type uploadHandler struct{}

// NewUploadHandler : ハンドラの生成
func NewUploadHandler() IUploadHandler {
	return &uploadHandler{}
}

func (h *uploadHandler) UploadImage(c *gin.Context) {

}
