//+build wireinject

package main

import (
	"app_name/app/config"
	"app_name/app/handler"
	"app_name/app/repository"
	"app_name/app/service"

	"github.com/google/wire"
)

func InitializeApplication() (APIApplication, error) {
	wire.Build(
		NewAPIApplication,
		config.Load,
		config.SetupRouter,
		config.ConnectDB,

		// handler
		handler.NewDefault,

		// service
		service.NewDefault,

		// repository
		repository.NewDefault,
	)
	return APIApplication{}, nil
}
