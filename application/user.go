package application

import (
	"context"

	"go.uber.org/zap"
	userModel "toDoBackEnd/domain/model/user"
)

type UserService interface {
	Get(ctx context.Context, id string) (*userModel.User, error)
}

type userService struct {
	userRepository userModel.Repository
	logger         *zap.Logger
}

func NewUserService(
	userRepository userModel.Repository,
	logger *zap.Logger,
) UserService {
	return &userService{
		userRepository: userRepository,
		logger:         logger.Named("UserService"),
	}
}

func (u *userService) Get(ctx context.Context, id string) (*userModel.User, error) {
	user, err := u.userRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return &userModel.User{
		Name: user.Name,
	}, nil
}
