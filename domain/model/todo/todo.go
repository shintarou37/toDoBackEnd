package todo

import (
	"context"
	"time"
	pb "toDoBackEnd/proto/proto-gen/pb"
)

type Repository interface {
	FindOne(ctx context.Context, id string) (*Todo, error)
	List(ctx context.Context, user_id string) ([]*Todo, error)
	Create(ctx context.Context, u *Todo) (*Todo, error)
	Update(ctx context.Context, u *Todo) (*Todo, error)
}

type Todo struct {
	ID        string    `firestore:"-" json:"id"`
	UserID    string    `firestore:"user_id" json:"user_id"`
	Title     string    `firestore:"title" json:"title"`
	Context   string    `firestore:"context" json:"context"`
	Priority  int32     `firestore:"priority" json:"priority"`
	CreatedAt time.Time `firestore:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `firestore:"updated_at" json:"updated_at,omitempty"`
}

func (t *Todo) ToPB() *pb.Todo {
	return &pb.Todo{
		Title:   t.Title,
		Context: t.Context,
	}
}

func ToPBList(todoList []*Todo) []*pb.Todo {
	var pbTodoList []*pb.Todo
	for _, todo := range todoList {
		pbTodoList = append(pbTodoList, todo.ToPB())
	}
	return pbTodoList
}
