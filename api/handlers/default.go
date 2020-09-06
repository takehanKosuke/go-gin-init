package handlers

import (
	"github.com/gin-gonic/gin"
)

// Default Interface
type Default interface {
	Ping(c *gin.Context)
}

// DefaultImpl DefaultImpl struct
type DefaultImpl struct{}

// NewDefault 新規Default structを作成
func NewDefault() Default {
	return &DefaultImpl{}
}

// Ping 初期デフォルトエンドポイント
func (d *DefaultImpl) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
