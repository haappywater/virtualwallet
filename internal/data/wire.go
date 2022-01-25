package data

import "github.com/google/wire"

var Provider = wire.NewSet(NewData, NewGormDB)
