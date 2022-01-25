//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/haappywater/virtualwallet/internal/biz"
	"github.com/haappywater/virtualwallet/internal/conf"
	"github.com/haappywater/virtualwallet/internal/data"
	"github.com/haappywater/virtualwallet/internal/server"
	"github.com/haappywater/virtualwallet/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
