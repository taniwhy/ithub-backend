package main

import (
	"crypto/sha1"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	imgupload "github.com/olahol/go-imageupload"
)

/*
func main() {
	dbConn := dao.NewDatabase()
	defer dbConn.Close()

	routers := router.Init(dbConn)

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        routers,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Serve failed")
		panic(err)
	}

	// Upload先のディレクトリ
	//dstDir := "./public/images"

	//router := gin.Default()

	//routers.MaxMultipartMemory = 8 << 20

	//routers.Static("/", "./views")
	/*
		routers.POST("/upload", func(c *gin.Context) {
			img, err := imgupload.Process(c.Request, "file")
			if err != nil {
				panic(err)
			}

			thumb, err := imgupload.ThumbnailJPEG(img, 300, 300, 90)
			if err != nil {
				panic(err)
			}

			h := sha1.Sum(thumb.Data)
			savepath := filepath.Join(dstDir, fmt.Sprintf("%s_%x.jpg", time.Now().Format("20060102150405"), h[:4]))
			thumb.Save(savepath)
		})

}*/
func main() {
	// Upload先のディレクトリ
	dstDir := "./public/images"

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20

	router.Static("/", "./views")

	router.POST("/upload", func(c *gin.Context) {
		img, err := imgupload.Process(c.Request, "file")
		if err != nil {
			panic(err)
		}

		thumb, err := imgupload.ThumbnailJPEG(img, 300, 300, 90)
		if err != nil {
			panic(err)
		}

		h := sha1.Sum(thumb.Data)
		savepath := filepath.Join(dstDir, fmt.Sprintf("%s_%x.jpg", time.Now().Format("20060102150405"), h[:4]))
		thumb.Save(savepath)
	})

	router.Run(":8000")
}
