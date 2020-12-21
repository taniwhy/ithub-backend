package router

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	imgupload "github.com/olahol/go-imageupload"
	"github.com/taniwhy/ithub-backend/configs"
	"github.com/taniwhy/ithub-backend/internal/app/datastore"
	"github.com/taniwhy/ithub-backend/internal/app/domain/service"
	"github.com/taniwhy/ithub-backend/internal/app/handler"
	"github.com/taniwhy/ithub-backend/internal/app/middleware/cors"
)

// Init : Init関数は依存性の注入とURLパスルーティングを行います
func Init(dbConn *gorm.DB) *gin.Engine {
	userDatastore := datastore.NewUserDatastore(dbConn)
	tagDatastore := datastore.NewTagDatastore(dbConn)
	userService := service.NewUserService(userDatastore)
	authHandler := handler.NewAuthHandler(userDatastore, userService)
	userHandler := handler.NewUserHandler(userDatastore)
	tagHandler := handler.NewTagHandler(tagDatastore)

	dbConn.LogMode(true)

	r := gin.Default()
	store := cookie.NewStore([]byte(configs.SecretKey))
	r.Use(sessions.Sessions("_session", store))
	r.Use(cors.Write())

	var test string

	type body struct {
		Mark string `json:"mark" validate:"required"`
	}

	r.GET("", func(c *gin.Context) {
        c.JSON(200, test)
    })

	r.POST("", func(c *gin.Context) {
		body := body{}
		_ = c.Bind(&body)
		test = body.Mark
        c.JSON(200, gin.H{
            "message": test,
        })
    })

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	users := v1.Group("/users")
	tags := v1.Group("/tags")
	static := r.Group("static")
	images := static.Group("/images")
	{
		auth.POST("/google/login", authHandler.Login)
		auth.DELETE("/logout", authHandler.Logout)

		v1.GET("/me", userHandler.GetMe)
		users.GET("/:name", userHandler.GetByName)
		users.PUT("/", userHandler.Update)
		users.DELETE("/", userHandler.Delete)

		tags.GET("/", tagHandler.GetList)
		tags.POST("/", tagHandler.Create)

		images.Static("/", "./web/images/")
		images.POST("/upload",  func(c *gin.Context) {
			dstDir := "./web/images/"
            img, err := imgupload.Process(c.Request, "file")
            if err != nil {
                c.JSON(400, err.Error())
            }

            thumb, err := imgupload.ThumbnailPNG(img, 96, 96)
            if err != nil {
                panic(err)
            }

            h := sha1.Sum(thumb.Data)
            filename := fmt.Sprintf("%s_%x.png", time.Now().Format("20060102150405"), h[:4])
            savepath := filepath.Join(dstDir, filename)
            thumb.Save(savepath)
            c.JSON(200, gin.H{"link": os.Getenv("HOST") + "/static/images/" + filename})
        })
	}
	return r
}
