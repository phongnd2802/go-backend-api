package user

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/internal/middlewares"
)

type AccessRoute struct{}

func (ac *AccessRoute) InitAccessRoute(Route *gin.RouterGroup, middleware middlewares.IMiddleware) {

}
