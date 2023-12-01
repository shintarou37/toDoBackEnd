package grpc

import (
	"context"
	"toDoBackEnd/application"
	pb "toDoBackEnd/proto/proto-gen/pb"

	"go.uber.org/zap"
)

type UserServer struct {
	logger      *zap.Logger
	userService application.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserServer(
	userService application.UserService,
	logger *zap.Logger,
) pb.UserServiceServer {
	return &UserServer{
		userService: userService,
		logger:      logger.Named("HelloServer"),
	}
}

func (s *UserServer) Get(ctx context.Context, req *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	user, err := s.userService.Get(context.Background(), req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.UserGetResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) Create(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	user, err := s.userService.Create(context.Background(), req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.UserCreateResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) Update(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	user, err := s.userService.Update(context.Background(), req.GetId(), req.GetName(), req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &pb.UserUpdateResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	err := s.userService.Login(context.Background(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.UserLoginResponse{}, nil
}
