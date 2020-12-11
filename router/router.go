package router

import (
	"crypto/sha1"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	imgupload "github.com/olahol/go-imageupload"
	"github.com/taniwhy/ithub-backend/config"
	"github.com/taniwhy/ithub-backend/datastore"
	"github.com/taniwhy/ithub-backend/domain/service"
	"github.com/taniwhy/ithub-backend/handler"
	"github.com/taniwhy/ithub-backend/middleware/cors"
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
	store := cookie.NewStore([]byte(config.SecretKey))
	r.Use(sessions.Sessions("_session", store))
	r.Use(cors.Write())

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	users := v1.Group("/users")
	tags := v1.Group("/tags")
	images := r.Group("/images")

	dstDir := "./public/images"

	{
		v1.GET("/", func(c *gin.Context) {
			session := sessions.Default(c)
			c.JSON(200, gin.H{"message": session.Get("_session")})
		})
		auth.POST("/google/login", authHandler.Login)
		auth.DELETE("/logout", authHandler.Logout)

		v1.GET("/me", userHandler.GetMe)
		users.GET("/:name", userHandler.GetByName)
		users.PUT("/", userHandler.Update)
		users.DELETE("/", userHandler.Delete)

		tags.GET("/", tagHandler.GetList)
		tags.POST("/", tagHandler.Create)

		images.POST("/upload", func(c *gin.Context) {
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

	}
	return r
}
