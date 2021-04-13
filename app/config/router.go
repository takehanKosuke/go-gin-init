package config

import (
	"app_name/app/handler"
	"app_name/app/repository"
	"app_name/app/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// NewHandler
	defo := handler.NewDefault(service.NewDefault(repository.NewDefault(db)))

	// Routing
	r.GET("/", defo.Ping)

	return r
}
