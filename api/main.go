package main

import (
	"app_name/api/config"
	"app_name/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	defo := handlers.NewDefault()

	r.GET("/", defo.Ping)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
