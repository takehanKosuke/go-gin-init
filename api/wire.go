//+build wireinject

package main

import (
	"app_name/api/config"

	"github.com/google/wire"
)

func InitializeApplication() (APIApplication, error) {
	wire.Build(
		NewAPIApplication,
		config.Load,
		config.SetupRouter,
		config.ConnectDB,
	)
	return APIApplication{}, nil
}
