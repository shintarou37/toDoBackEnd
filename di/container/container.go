package container

import (
	"cloud.google.com/go/firestore"
	"context"
	"go.uber.org/zap"
	"toDoBackEnd/infra/firebase"
)

type C struct {
	firestoreClient *firestore.Client
	logger          *zap.Logger
	auth            *firestore.Client
}

var c C

func InitializeContainer(ctx context.Context, logger *zap.Logger) error {
	firestoreClient, err := firebase.New(ctx)
	if err != nil {
		return err
	}
	c = C{
		logger:          logger,
		firestoreClient: firestoreClient,
	}

	return nil
}

func FirestoreClient() *firestore.Client {
	return c.firestoreClient
}

func Logger() *zap.Logger {
	return c.logger
}
