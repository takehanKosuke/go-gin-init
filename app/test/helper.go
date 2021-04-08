package test

import "app_name/app/config"

// LoadConfig
func LoadConfig() *config.Config {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}
	return config
}
