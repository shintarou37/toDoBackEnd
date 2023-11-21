package container

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"toDoBackEnd/domain/auth"
	"toDoBackEnd/infra/firebaseauth"
)

type C struct {
	firestoreClient *firestore.Client
	logger          *zap.Logger
	auth            auth.Auth
}

var c C

func InitializeContainer(ctx context.Context, logger *zap.Logger) error {
	fsClient, err := firestore.NewClient(ctx, "test")
	if err != nil {
		return errors.Wrap(err, "failed Firestore setting")
	}

	authCli, err := firebaseauth.New(ctx)
	if err != nil {
		return err
	}
	c = C{
		firestoreClient: fsClient,
		logger:          logger,
		auth:            authCli,
	}

	return nil
}

func FirestoreClient() *firestore.Client {
	return c.firestoreClient
}

func Logger() *zap.Logger {
	return c.logger
}

func Auth() auth.Auth {
	return c.auth
}
