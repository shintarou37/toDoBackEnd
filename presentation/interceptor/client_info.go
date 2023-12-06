package interceptor

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"toDoBackEnd/infra/clientinfo"
	infractx "toDoBackEnd/infra/ctx"
)

const (
	AppVersionKey = "x-my-app-version"
	AppOSKey      = "x-my-app-os"
)

func ClientInfoInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Error("failed to get incoming metadata", zap.Error(err))
			return handler(ctx, req)
		}

		ci := &clientinfo.ClientInfo{}

		mdAppVersion := md.Get(AppVersionKey)
		if len(mdAppVersion) > 0 {
			ci.Version = mdAppVersion[0]
		}

		mdOs := md.Get(AppOSKey)
		if len(mdOs) > 0 {
			ci.OS = clientinfo.ToOS(mdOs[0])
		}

		newCtx := infractx.WithClientInfo(ctx, ci)
		return handler(newCtx, req)
	}
}
