package handler

import (
	"app_name/app/service"

	"github.com/gin-gonic/gin"
)

// Default Interface
type Default interface {
	Ping(c *gin.Context)
}

// DefaultImpl DefaultImpl struct
type DefaultImpl struct {
	service service.Default
}

// NewDefault 新規Default structを作成
func NewDefault(service service.Default) Default {
	return &DefaultImpl{
		service: service,
	}
}

// Ping 初期デフォルトエンドポイント
func (d *DefaultImpl) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
