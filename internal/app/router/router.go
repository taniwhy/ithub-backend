package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	uploadHandler := handler.NewUploadHandler()

	dbConn.LogMode(true)

	r := gin.Default()
	store := cookie.NewStore([]byte(configs.SecretKey))
	r.Use(sessions.Sessions("_session", store))
	r.Use(cors.Write())

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
		images.POST("/upload", uploadHandler.UploadImage)
	}
	return r
}
