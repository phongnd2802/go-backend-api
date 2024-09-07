package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/internal/middlewares"
	"github.com/phongnd2802/go-backend-api/internal/wire"
)

type AccessRoute struct{}

func (ac *AccessRoute) InitAccessRoute(Route *gin.RouterGroup, middleware middlewares.IMiddleware) {
	accessController, _ := wire.InitAccessRouterHandler()
	// Public Router
	accessPublicRouter := Route.Group("/shop")
	{
		accessPublicRouter.POST("/signup", accessController.SignUp)
		accessPublicRouter.POST("/signin", accessController.SignIn)
		accessPublicRouter.GET("/otp", accessController.VerifyOTP)
		accessPublicRouter.POST("/otp", accessController.SendOTP)
		accessPublicRouter.POST("/reset-password", accessController.ResetPassword)
	}

	// Private Route
	accessPrivateRouter := Route.Group("/shop")
	//Route.Use() //Authentication
	{
		accessPrivateRouter.GET("/profile")
	}
}
