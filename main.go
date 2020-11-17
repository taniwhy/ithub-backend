package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/db/dao"
)

func main() {
	dbConn := dao.NewDatabase()
	defer dbConn.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world3",
		})
	})
	r.Run(":" + os.Getenv("PORT"))
}