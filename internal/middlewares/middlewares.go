package middlewares

import (
	"github.com/phongnd2802/go-backend-api/internal/repositories"
)

type Middleware struct{
	authRepo repositories.IAuthRepository
}

func NewMiddleware(authRepo repositories.IAuthRepository) *Middleware {
	return &Middleware{
		authRepo: authRepo,
	}
}
