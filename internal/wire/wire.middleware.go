//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/phongnd2802/go-backend-api/internal/middlewares"
	"github.com/phongnd2802/go-backend-api/internal/repositories"
)

func InitMiddlewareHandler() (*middlewares.Middleware, error) {
	wire.Build(
		repositories.NewTokenRepository,
		repositories.NewAuthRepository,
		middlewares.NewMiddleware,
	)
	return &middlewares.Middleware{}, nil
}