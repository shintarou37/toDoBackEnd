package grpc

import (
	"context"
	pb "toDoBackEnd/proto/proto-gen/pb"

	"go.uber.org/zap"
)

type HelloService struct {
	logger                             *zap.Logger
	pb.UnimplementedHelloServiceServer
}

func NewHelloServer(
	logger *zap.Logger,
) pb.HelloServiceServer {
	return &HelloService{
		logger: logger.Named("HelloServer"),
	}
}

func (s *HelloService) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{}, nil
}

func (s *HelloService) SayHello2(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{}, nil
}
