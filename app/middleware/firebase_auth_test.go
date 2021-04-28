package middleware

import (
	"app_name/app/middleware/mock_middleware"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	dummyIDToken = "asdfghj12345"
)

func TestAuthentication(t *testing.T) {
	t.Parallel()
	asserts := assert.New(t)

	tests := []struct {
		name        string
		insertToken string
		outputToken *auth.Token
		err         error
	}{
		{
			name:        "適切にcontextにidTokenが取得",
			insertToken: "bearer hogehoge",
			outputToken: &auth.Token{
				UID: dummyIDToken,
			},
			err:       nil,
		},
		{
			name:        "tokenがinvlidする",
			insertToken: "bearer hogehoge",
			outputToken: nil,
			err:         fmt.Errorf("auth error"),
		},
	}
	for _, td := range tests {
		td := td
		t.Run(fmt.Sprintf("Authentication: %s", td.name), func(t *testing.T) {
			// gomockの設定
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockFirebaseAuthClient := mock_middleware.NewMockFirebaseAuthClient(ctrl)
			mockFirebaseAuthClient.EXPECT().VerifyIDToken(gomock.Any(), gomock.Any()).Return(td.outputToken, td.err).AnyTimes()

			mockAuthMiddleware := NewFirebaseAuth(mockFirebaseAuthClient)

			ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

			// リクエストの生成
			req, _ := http.NewRequest("GET", "/", nil)

			// ヘッダー情報の追加
			req.Header.Add("Authorization", td.insertToken)

			// リクエスト情報をコンテキストに入れる
			ginContext.Request = req

			mockAuthMiddleware.Authentication(ginContext)
			returnToken, _ := ginContext.Get("firebaseUUID")
			if td.outputToken == nil {
				asserts.Equal(nil, returnToken)
			} else {
				asserts.Equal(td.outputToken.UID, returnToken)
			}
		})
	}
}
