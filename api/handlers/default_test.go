package handlers

import (
	"app_name/api/config"
	"app_name/api/repositories"
	"app_name/api/services"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	t.Parallel()
	asserts := assert.New(t)
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db := config.ConnectDB(cfg)

	tests := []struct {
		name         string
		outputStatus int
		outputJSON   string
	}{
		{
			name:         "適切にレスポンスが帰ってくる",
			outputStatus: 200,
			outputJSON:   `{"message":"pong"}`,
		},
	}
	for _, td := range tests {
		td := td
		t.Run(fmt.Sprintf("Ping: %s", td.name), func(t *testing.T) {
			// test用のgin contextの生成
			recorder := httptest.NewRecorder()
			ginContext, _ := gin.CreateTestContext(recorder)
			req, _ := http.NewRequest("GET", "/", nil)
			ginContext.Request = req

			NewDefault(services.NewDefault(repositories.NewDefault(db))).Ping(ginContext)
			asserts.Equal(td.outputStatus, recorder.Code)
			if td.outputJSON != "" {
				asserts.JSONEq(td.outputJSON, recorder.Body.String())
			} else {
				asserts.Equal(td.outputJSON, recorder.Body.String())
			}
		})
	}
}
