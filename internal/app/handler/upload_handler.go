package handler

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	imgupload "github.com/olahol/go-imageupload"
)

// IUploadHandler : インターフェース
type IUploadHandler interface {
	UploadImage(c *gin.Context)
}

type uploadHandler struct {
	name string
}

// NewUploadHandler : ハンドラの生成
func NewUploadHandler() IUploadHandler {
	return &uploadHandler{}
}

func (h *uploadHandler) UploadImage(c *gin.Context) {

	/*
		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/" + filename + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	*/
	dstDir := "./web/images"
	img, err := imgupload.Process(c.Request, "file")
	if err != nil {
		panic(err)
	}

	thumb, err := imgupload.ThumbnailJPEG(img, 300, 300, 90)
	if err != nil {
		panic(err)
	}

	//s := sha1.Sum(thumb.Data)
	savepath := filepath.Join(dstDir, fmt.Sprintf("" /*, time.Now().Format("20060102150405"), s[:4]*/))
	thumb.Save(savepath)
}
