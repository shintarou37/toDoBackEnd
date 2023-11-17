//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"toDoBackEnd/di/container"
	"toDoBackEnd/presentation/grpc"
	pb "toDoBackEnd/proto/proto-gen/pb"
)

var wireSet = wire.NewSet(
	container.WireSet,
	grpc.WireSet,
)

func NewHelloServer() pb.HelloServiceServer {
	wire.Build(wireSet)
	return nil
}
