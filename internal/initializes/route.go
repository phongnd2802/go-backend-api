package initializes

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/internal/routes"
	"github.com/phongnd2802/go-backend-api/internal/wire"
	"github.com/phongnd2802/go-backend-api/pkg/response"
)

func initRoute() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		r = gin.New()
	}

	// Middlewares

	middleware, _ := wire.InitMiddlewareHandler()
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.LoggerWrite())
	r.Use(middleware.ApiKey())
	// Routes
	adminRoute := routes.RouteApp.Admin
	shopRoute := routes.RouteApp.Shop
	userRoute := routes.RouteApp.User
	MainGroup := r.Group("api/v1")
	{
		MainGroup.GET("/checkStatus", func(c *gin.Context) {
			clientIP := c.Request.RemoteAddr
			response.SuccessResponse(c, response.SuccessOK, clientIP)
		})
	}
	{
		adminRoute.InitAccessRoute(MainGroup, middleware)
	}
	{
		shopRoute.InitAccessRoute(MainGroup, middleware)
	}
	{
		userRoute.InitAccessRoute(MainGroup, middleware)
	}

	return r
}
