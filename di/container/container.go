package container

import (
	"context"
	"go.uber.org/zap"
)

type C struct {
	logger *zap.Logger
}

var c C

func InitializeContainer(ctx context.Context, logger *zap.Logger) error {
	c = C{
		logger: logger,
	}
	return nil
}

func Logger() *zap.Logger {
	return c.logger
}
