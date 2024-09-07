package middlewares

import (
	"github.com/phongnd2802/go-backend-api/internal/repositories"
)

type Middleware struct{
	tokenRepo repositories.ITokenRepository
	authRepo repositories.IAuthRepository
}

func NewMiddleware(authRepo repositories.IAuthRepository, tokenRepo repositories.ITokenRepository) *Middleware {
	return &Middleware{
		authRepo: authRepo,
		tokenRepo: tokenRepo,
	}
}
