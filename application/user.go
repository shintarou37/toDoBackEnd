package application

import (
	"context"

	"go.uber.org/zap"
	"toDoBackEnd/domain/model/user"
)

type UserService interface {
	Get(ctx context.Context, id string) (*user.User, error)
}

type userService struct {
	logger *zap.Logger
}

func NewUserService(
	logger *zap.Logger,
) UserService {
	return &userService{
		logger: logger.Named("UserService"),
	}
}

func (s *userService) Get(ctx context.Context, id string) (*user.User, error) {
	return &user.User{
		Name: "test user",
	}, nil
}
