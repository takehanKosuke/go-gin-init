//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
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
