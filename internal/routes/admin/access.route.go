package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/internal/middlewares"
)

type AccessRoute struct{}

func (ac *AccessRoute) InitAccessRoute(Route *gin.RouterGroup, middleware middlewares.IMiddleware) {

	publicAccessRoute := Route.Group("/admin")
	{
		publicAccessRoute.POST("/login")
	}

	privateAccessRoute := Route.Group("/admin")
	privateAccessRoute.Use(middleware.Authentication())
	privateAccessRoute.Use(middleware.PermissionCheck())
	{
		privateAccessRoute.GET("/shops")              // Get all shop
		privateAccessRoute.PATCH("/shops/active/:id") // Active shop
		privateAccessRoute.POST("/logout")
	}
}
