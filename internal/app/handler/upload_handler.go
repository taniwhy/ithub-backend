package handler

import (
	"crypto/sha1"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	imgupload "github.com/olahol/go-imageupload"
)

// IUploadHandler : インターフェース
type IUploadHandler interface {
	UploadImage(c *gin.Context)
}

type uploadHandler struct {
}

// NewUploadHandler : ハンドラの生成
func NewUploadHandler() IUploadHandler {
	return &uploadHandler{}
}

func (h *uploadHandler) UploadImage(c *gin.Context) {
	dstDir := "./web/images/"
	img, err := imgupload.Process(c.Request, "file")
	if err != nil {
		c.JSON(400, c.Request.Header.Get("Content-Type"))
	}

	thumb, err := imgupload.ThumbnailJPEG(img, 300, 300, 90)
	if err != nil {
		panic(err)
	}

	s := sha1.Sum(thumb.Data)
	savepath := filepath.Join(dstDir, fmt.Sprintf("%s_%x.jpg", time.Now().Format("20060102150405"), s[:4]))
	thumb.Save(savepath)
}
