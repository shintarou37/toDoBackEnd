package application

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	todoModel "toDoBackEnd/domain/model/todo"
	"unicode/utf8"
)

type TodoService interface {
	Get(ctx context.Context, id string) (*todoModel.Todo, error)
	List(ctx context.Context, userID string) ([]*todoModel.Todo, error)
	Create(ctx context.Context, title, context, userID string, priority int32) (*todoModel.Todo, error)
	Update(ctx context.Context, title, context, todoID string, priority int32) (*todoModel.Todo, error)
}

type todoService struct {
	todoRepository todoModel.Repository
	logger         *zap.Logger
}

func NewTodoService(
	todoRepository todoModel.Repository,
	logger *zap.Logger,
) TodoService {
	return &todoService{
		todoRepository: todoRepository,
		logger:         logger.Named("TodoService"),
	}
}

func (t *todoService) Get(ctx context.Context, id string) (*todoModel.Todo, error) {
	todo, err := t.todoRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todoService) List(ctx context.Context, userID string) ([]*todoModel.Todo, error) {
	todoList, err := t.todoRepository.List(ctx, userID)
	if err != nil {
		return nil, err
	}
	return todoList, nil
}

func (t *todoService) Create(ctx context.Context, title, context, userID string, priority int32) (*todoModel.Todo, error) {
	if err := t.validateCreateAndUpdate(title, context); err != nil {
		return nil, err
	}
	todo, err := t.todoRepository.Create(ctx, &todoModel.Todo{
		Title:    title,
		Context:  context,
		Priority: priority,
		UserID:   userID,
	})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todoService) Update(ctx context.Context, title, context, todoID string, priority int32) (*todoModel.Todo, error) {
	if err := t.validateCreateAndUpdate(title, context); err != nil {
		return nil, err
	}
	todo, err := t.todoRepository.Update(ctx, &todoModel.Todo{
		ID:       todoID,
		Title:    title,
		Context:  context,
		Priority: priority,
	})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todoService) validateCreateAndUpdate(title, context string) error {
	if title == "" || utf8.RuneCountInString(title) > 20 {
		return errors.New("invalid title")
	}
	if context == "" || utf8.RuneCountInString(context) > 1000 {
		return errors.New("invalid context")
	}
	return nil
}
