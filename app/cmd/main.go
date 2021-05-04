package main

import (
	"app_name/app/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
