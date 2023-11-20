// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"toDoBackEnd/application"
	"toDoBackEnd/di/container"
	"toDoBackEnd/infra/persistence"
	"toDoBackEnd/presentation/grpc"
	"toDoBackEnd/proto/proto-gen/pb"
)

// Injectors from wire.go:

func NewHelloServer() pb.UserServiceServer {
	logger := container.Logger()
	userService := application.NewUserService(logger)
	userServiceServer := grpc.NewHelloServer(userService, logger)
	return userServiceServer
}

// wire.go:

var wireSet = wire.NewSet(container.WireSet, grpc.WireSet, application.WireSet, persistence.WireSet)
