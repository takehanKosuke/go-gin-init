//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package middleware

import (
	"app_name/app/response"
	"context"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type FirebaseAuth interface {
	Authentication(c *gin.Context)
}

type FirebaseAuthClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

type firebaseAuthImpl struct {
	authClient FirebaseAuthClient
}

func NewFirebaseAuth(authClient FirebaseAuthClient) FirebaseAuth {
	return &firebaseAuthImpl{
		authClient: authClient,
	}
}

func (f *firebaseAuthImpl) Authentication(c *gin.Context) {
	token := c.GetHeader("Authorization")
	// 認証トークンを検証
	idToken, err := f.authClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		response.HTTPResponseInternalServerError(c, "error")
		return
	}

	c.Set("firebaseUUID", idToken.UID)
}
