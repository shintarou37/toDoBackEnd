package ctx

import (
	"context"
	"toDoBackEnd/infra/clientinfo"
)

type clientInfoCtxKey struct{}

func WithClientInfo(parent context.Context, clientInfo *clientinfo.ClientInfo) context.Context {
	return context.WithValue(parent, clientInfoCtxKey{}, clientInfo)
}

func ClientInfo(ctx context.Context) (*clientinfo.ClientInfo, bool) {
	ver, ok := ctx.Value(clientInfoCtxKey{}).(*clientinfo.ClientInfo)
	return ver, ok
}
