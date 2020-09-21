package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("api/config/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("read error: %s \n", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("server error: %s \n", err)
	}

	return &cfg, nil
}
