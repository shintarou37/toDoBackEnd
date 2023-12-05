package application

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewUserService,
	NewTodoService,
)
