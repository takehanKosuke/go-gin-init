package handlers

import (
	"app_name/app/services"

	"github.com/gin-gonic/gin"
)

// Default Interface
type Default interface {
	Ping(c *gin.Context)
}

// DefaultImpl DefaultImpl struct
type DefaultImpl struct {
	service services.Default
}

// NewDefault 新規Default structを作成
func NewDefault(service services.Default) Default {
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
