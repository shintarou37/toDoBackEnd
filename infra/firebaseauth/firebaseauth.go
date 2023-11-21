package firebaseauth

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"toDoBackEnd/domain/auth"
)

func New(ctx context.Context) (auth.Auth, error) {
	sa := option.WithCredentialsFile("./serviceAccount.json.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return app.Auth(ctx)
}
