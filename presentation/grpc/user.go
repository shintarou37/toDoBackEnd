package grpc

import (
	"context"
	"toDoBackEnd/application"
	pb "toDoBackEnd/proto/proto-gen/pb"

	"go.uber.org/zap"
)

type HelloService struct {
	logger      *zap.Logger
	userService application.UserService
	pb.UnimplementedUserServiceServer
}

func NewHelloServer(
	userService application.UserService,
	logger *zap.Logger,
) pb.UserServiceServer {
	return &HelloService{
		userService: userService,
		logger:      logger.Named("HelloServer"),
	}
}

func (s *HelloService) Get(context.Context, *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	user, err := s.userService.Get(context.Background(), "test")
	if err != nil {
		return nil, err
	}
	return &pb.UserGetResponse{
		Name: user.Name,
	}, nil
}
