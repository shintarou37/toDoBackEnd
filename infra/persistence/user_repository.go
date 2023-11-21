package persistence

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"toDoBackEnd/domain/model/user"
)

const collectionNameUser = "User"

type userRepository struct {
	fsClient *firestore.Client
	logger   *zap.Logger
}

func NewUserRepository(
	fsClient *firestore.Client,
	logger *zap.Logger,
) user.Repository {
	return &userRepository{
		fsClient: fsClient,
		logger:   logger.Named("user/userRepository"),
	}
}

func (r *userRepository) collection() *firestore.CollectionRef {
	return r.fsClient.Collection(collectionNameUser)
}

func (r *userRepository) FindOne(ctx context.Context, id string) (*user.User, error) {
	d, err := r.collection().Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get document from Firestore")
	}
	return r.toUser(d)
}

func (r *userRepository) toUser(d *firestore.DocumentSnapshot) (*user.User, error) {
	var u user.User
	if err := d.DataTo(&u); err != nil {
		return nil, errors.Wrap(err, "failed to convert document data to user")
	}
	u.ID = d.Ref.ID
	return &u, nil
}
