package handlers

import (
	"app_name/app/config"
	"app_name/app/handler"
	"app_name/app/repository"
	"app_name/app/service"
	"app_name/app/test"
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
	cfg := test.LoadConfig()

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

			handler.NewDefault(service.NewDefault(repository.NewDefault(db))).Ping(ginContext)
			asserts.Equal(td.outputStatus, recorder.Code)
			if td.outputJSON != "" {
				asserts.JSONEq(td.outputJSON, recorder.Body.String())
			} else {
				asserts.Equal(td.outputJSON, recorder.Body.String())
			}
		})
	}
}
