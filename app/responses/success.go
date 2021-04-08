package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Success struct {
	Type    int
	Message string
}

func HTTPResponseStatusOK(c *gin.Context, message string) {
	e := Error{http.StatusOK, message}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"type":    e.Type,
		"message": e.Message,
	})
}
