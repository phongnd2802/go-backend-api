package middlewares

import "github.com/gin-gonic/gin"

type IMiddleware interface {
	ApiKey() gin.HandlerFunc
	Authentication() gin.HandlerFunc
	PermissionCheck() gin.HandlerFunc
	CorsMiddleware() gin.HandlerFunc
}

type middleware struct{}

func New() IMiddleware {
	return &middleware{}
}
