package service

import "app_name/app/repository"

type Default interface {
	Ping()
}

type DefaultImpl struct {
	repository repository.Default
}

func NewDefault(repository repository.Default) Default {
	return &DefaultImpl{repository: repository}
}

func (d *DefaultImpl) Ping() {}
