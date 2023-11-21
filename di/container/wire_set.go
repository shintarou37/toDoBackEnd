package container

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	FirestoreClient,
	Logger,
	Auth,
)
