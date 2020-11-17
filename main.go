package main

import (
	"net/http"
	"fmt"

	"github.com/taniwhy/ithub-backend/db/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := dao.NewDatabase()
	defer dbConn.Close()
	
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":8000")
}