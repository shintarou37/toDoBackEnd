package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"
)

func New(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsFile("/Users/shint/go/study/toDoBackEnd/serviceAccount.json")
	config := &firebase.Config{ProjectID: "test-3eaa4"}
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		return nil, err
	}

	client, _ := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil

}
