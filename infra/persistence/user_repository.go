package persistence

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
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

func (r *userRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	ref := r.collection().NewDoc()
	u.CreatedAt = time.Now()
	u.ID = ref.ID
	_, err := ref.Create(ctx, u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create document from Firestore")
	}
	return u, nil
}

func (r *userRepository) FindOneByEmail(ctx context.Context, email string) (*user.User, error) {
	iter := r.collection().Where("email", "==", email).Limit(1).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query document from Firestore")
	}

	return r.toUser(doc)
}

func (r *userRepository) Update(ctx context.Context, u *user.User) (*user.User, error) {
	_, err := r.collection().Doc(u.ID).Update(ctx, []firestore.Update{
		{Path: "name", Value: u.Name},
		{Path: "email", Value: u.Email},
		{Path: "updated_at", Value: time.Now()},
	})
	return u, err
}

func (r *userRepository) toUser(d *firestore.DocumentSnapshot) (*user.User, error) {
	var u user.User
	if err := d.DataTo(&u); err != nil {
		return nil, errors.Wrap(err, "failed to convert document data to user")
	}
	u.ID = d.Ref.ID
	return &u, nil
}
