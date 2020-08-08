package main

import (
	"app_name/api/config"
	"app_name/api/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	defo := handlers.NewDefault()

	r.GET("/", defo.Ping)

	return r
}

func main() {
	// db接続
	config.ConnectDB()

	r := setupRouter()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
