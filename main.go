package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"toDoBackEnd/di"
	"toDoBackEnd/di/container"
	l "toDoBackEnd/infra/log"
	"toDoBackEnd/proto/proto-gen/pb"
)

const port = 9000
func main() {
	ctx := context.Background()
	logger, err := l.New("DEBUG")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup logger: %s\n", err)
		os.Exit(1)
	}
	if err := container.InitializeContainer(ctx, logger); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to initialize container: %s\n", err)
		os.Exit(1)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	pb.RegisterHelloServiceServer(s, di.NewHelloServer())

	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] failed to listen: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Starting gRPC server on %d\n", port)
	if err := s.Serve(lis); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] failed to serve: %v", err)
		os.Exit(1)
	}
}
