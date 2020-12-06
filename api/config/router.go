package config

import (
	"app_name/api/handlers"
	"app_name/api/repositories"
	"app_name/api/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// NewHandler
	defo := handlers.NewDefault(services.NewDefault(repositories.NewDefault(db)))

	// Routing
	r.GET("/", defo.Ping)

	return r
}
