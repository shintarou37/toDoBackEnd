package persistence

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
	"toDoBackEnd/domain/model/todo"
)

const collectionNameTodo = "Todo"

type todoRepository struct {
	fsClient *firestore.Client
	logger   *zap.Logger
}

func NewTodoRepository(
	fsClient *firestore.Client,
	logger *zap.Logger,
) todo.Repository {
	return &todoRepository{
		fsClient: fsClient,
		logger:   logger.Named("todo/todoRepository"),
	}
}

func (r *todoRepository) collection() *firestore.CollectionRef {
	return r.fsClient.Collection(collectionNameTodo)
}

func (r *todoRepository) FindOne(ctx context.Context, id string) (*todo.Todo, error) {
	d, err := r.collection().Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get document from Firestore")
	}
	return r.toTodo(d)
}

func (r *todoRepository) List(ctx context.Context, userID string) ([]*todo.Todo, error) {
	iter := r.collection().Where("user_id", "==", userID).OrderBy("priority", firestore.Asc).Documents(ctx)
	return r.documentsToTodoList(ctx, iter)
}

func (r *todoRepository) Create(ctx context.Context, u *todo.Todo) (*todo.Todo, error) {
	ref := r.collection().NewDoc()
	u.CreatedAt = time.Now()
	_, err := ref.Create(ctx, u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create document from Firestore")
	}
	return u, nil
}

func (r *todoRepository) Update(ctx context.Context, u *todo.Todo) (*todo.Todo, error) {
	_, err := r.collection().Doc(u.ID).Update(ctx, []firestore.Update{
		{Path: "title", Value: u.Title},
		{Path: "context", Value: u.Context},
		{Path: "priority", Value: u.Priority},
		{Path: "updated_at", Value: time.Now()},
	})
	return u, err
}

func (r *todoRepository) toTodo(d *firestore.DocumentSnapshot) (*todo.Todo, error) {
	var u todo.Todo
	if err := d.DataTo(&u); err != nil {
		return nil, errors.Wrap(err, "failed to convert document data to todo")
	}
	return &u, nil
}

func (r *todoRepository) documentsToTodoList(ctx context.Context, iter *firestore.DocumentIterator) ([]*todo.Todo, error) {
	var invs []*todo.Todo
	for {
		ds, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var inv todo.Todo
		if err := ds.DataTo(&inv); err != nil {
			return nil, err
		}
		inv.ID = ds.Ref.ID
		invs = append(invs, &inv)
	}
	return invs, nil
}
