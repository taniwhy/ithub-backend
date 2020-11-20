package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/config"
	"github.com/taniwhy/ithub-backend/datastore"
	"github.com/taniwhy/ithub-backend/domain/service"
	"github.com/taniwhy/ithub-backend/handler"
	"github.com/taniwhy/ithub-backend/middleware/cors"
)

// Init : TODO
func Init(dbConn *gorm.DB) *gin.Engine {
	dbConn.LogMode(true)

	userDatastore := datastore.NewUserDatastore(dbConn)
	userService := service.NewUserService(userDatastore)
	authHandler := handler.NewAuthHandler(userDatastore, userService)
	userHandler := handler.NewUserHandler()

	r := gin.Default()
	store := cookie.NewStore([]byte(config.SignBytes))
	r.Use(sessions.Sessions("session", store))
	r.Use(cors.Write())

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	users := v1.Group("/users")
	{
		v1.GET("/", func(c *gin.Context) {
			session := sessions.Default(c)
			c.JSON(200, gin.H{"message": session.Get("state")})
		})
		auth.POST("/google/login", authHandler.Login)
		auth.DELETE("/logout", authHandler.Logout)
		auth.POST("/reflesh")

		v1.GET("/me", userHandler.GetMe)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("", userHandler.Update)
		users.DELETE("", userHandler.Delete)
	}
	return r
}
