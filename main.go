package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/taniwhy/ithub-backend/db/dao"
	"github.com/taniwhy/ithub-backend/router"
)

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
		log.Fatal(err)
		panic(err)
	}
}
