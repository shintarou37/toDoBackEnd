//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"toDoBackEnd/application"
	"toDoBackEnd/di/container"
	"toDoBackEnd/infra/persistence"
	"toDoBackEnd/presentation/grpc"
	pb "toDoBackEnd/proto/proto-gen/pb"
)

var wireSet = wire.NewSet(
	container.WireSet,
	grpc.WireSet,
	application.WireSet,
	persistence.WireSet,
)

func NewUserServer() pb.UserServiceServer {
	wire.Build(wireSet)
	return nil
}
