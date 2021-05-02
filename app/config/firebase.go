package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// firebaseAuthClient firebase authのclientを作成
func firebaseAuthClient() *auth.Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/refreshToken.json")
	// firebase appの作成
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	firebaseAuthClient, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	return firebaseAuthClient
}

// 以下mock用コードなので削除すること
type MockFirebaseClient struct{}

func newMockFirebaseAuthClient() *MockFirebaseClient {
	return &MockFirebaseClient{}
}

func (m *MockFirebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return nil, nil
}
