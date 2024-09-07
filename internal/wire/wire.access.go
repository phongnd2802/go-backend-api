//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/phongnd2802/go-backend-api/internal/controllers"
	"github.com/phongnd2802/go-backend-api/internal/repositories"
	"github.com/phongnd2802/go-backend-api/internal/services"
)

func InitAccessRouterHandler() (*controllers.AccessController, error) {
	wire.Build(
		repositories.NewOTPRepository,
		repositories.NewShopRepository,
		repositories.NewTokenRepository,
		services.NewAccessService,
		controllers.NewAccessController,
	)
	return &controllers.AccessController{}, nil
}

