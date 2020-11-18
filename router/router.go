package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/handler"
	"github.com/taniwhy/ithub-backend/middleware/cors"
)

// Init : TODO
func Init(dbConn *gorm.DB) *gin.Engine {
	authHandler := handler.NewAuthHandler()
	fmt.Print(dbConn)
	r := gin.Default()
	r.Use(cors.Write())

	v1 := r.Group("/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})
	auth := v1.Group("/auth")
	auth.POST("/google/login", authHandler.Login)
	auth.POST("/reflesh")
	return r
}
