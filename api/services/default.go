package services

import "app_name/api/repositories"

type Default interface {
	Ping()
}

type DefaultImpl struct {
	repository repositories.Default
}

func NewDefault(repository repositories.Default) Default {
	return &DefaultImpl{repository: repository}
}

func (d *DefaultImpl) Ping() {}
