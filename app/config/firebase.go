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
