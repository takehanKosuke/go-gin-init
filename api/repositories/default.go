package repositories

import (
	"app_name/api/config"

	"github.com/jinzhu/gorm"
)

type Default interface {
	Ping()
}

type DefaultImpl struct {
	db *gorm.DB
}

func NewDefault(cfg *config.Config) Default {
	return &DefaultImpl{db: cfg.Mysql.DB}
}

func (d *DefaultImpl) Ping() {}
