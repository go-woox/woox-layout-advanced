//go:build wireinject
// +build wireinject

package assembly

import (
	hServer "github.com/cloudwego/hertz/pkg/app/server"
	ahServer "github.com/go-woox/woox-layout-advanced/internal/app/server"
	"github.com/go-woox/woox-layout-advanced/internal/server"
	"github.com/go-woox/woox-layout-advanced/internal/service"
	"github.com/google/wire"
)

func NewHttpServer() (*hServer.Hertz, func(), error) {
	panic(wire.Build(
		ahServer.ProviderHandlerSet,
		service.ProviderServiceSet,
		server.ProviderHttpServerSet,
	))
}
