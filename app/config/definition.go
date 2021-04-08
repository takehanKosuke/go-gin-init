package config

import (
	"fmt"
	"strings"

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
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$PWD/../../config/")
	v.AddConfigPath("app/config/")

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("read error: %s", err)
	}

	var cfg Config

	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("server error: %s", err)
	}

	return &cfg, nil
}
