package repository

import (
	"github.com/jinzhu/gorm"
)

type Default interface {
	Ping()
}

type DefaultImpl struct {
	db *gorm.DB
}

func NewDefault(db *gorm.DB) Default {
	return &DefaultImpl{db: db}
}

func (d *DefaultImpl) Ping() {}
