package application

import (
	"context"

	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	userModel "toDoBackEnd/domain/model/user"
)

type UserService interface {
	Get(ctx context.Context, id string) (*userModel.User, error)
	Create(ctx context.Context, name, email, password string) (*userModel.User, error)
	Update(ctx context.Context, id, name, email string) (*userModel.User, error)
	Login(ctx context.Context, email, password string) error
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
	a := bcrypt.CompareHashAndPassword(user.Password, []byte("Password"))
	fmt.Println(a)
	return user, nil
}

func (u *userService) Create(ctx context.Context, name, email, password string) (*userModel.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepository.Create(ctx, &userModel.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Update(ctx context.Context, id, name, email string) (*userModel.User, error) {
	user, err := u.userRepository.Update(ctx, &userModel.User{
		ID:    id,
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Login(ctx context.Context, email, password string) error {
	user, err := u.userRepository.FindOneByEmail(ctx, email)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		// パスワードが一致しない場合のエラーメッセージ
		return errors.New("invalid password")
	} else if err != nil {
		return errors.Wrap(err, "error comparing password")
	}
	return nil
}
