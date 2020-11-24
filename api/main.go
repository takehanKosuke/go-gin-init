package main

import (
	"app_name/api/config"
	"app_name/api/handlers"
	"app_name/api/repositories"
	"app_name/api/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type APIApplication struct {
	cfg    *config.Config
	router *gin.Engine
	db     *gorm.DB
}

func NewAPIApplication(cfg *config.Config, router *gin.Engine, db *gorm.DB) APIApplication {
	return APIApplication{
		cfg:    cfg,
		router: router,
		db:     db,
	}
}

func setupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	defo := handlers.NewDefault(services.NewDefault(repositories.NewDefault(cfg)))

	r.GET("/", defo.Ping)

	return r
}

func main() {
	api, err := InitializeApplication()
	if err != nil {
		panic(err)
	}

	err = api.router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
