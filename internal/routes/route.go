package routes

import (
	"github.com/phongnd2802/go-backend-api/internal/routes/admin"
	"github.com/phongnd2802/go-backend-api/internal/routes/shop"
	"github.com/phongnd2802/go-backend-api/internal/routes/user"
)

type RouteGroup struct {
	Admin admin.AdminRouteGroup
	Shop  shop.ShopRouteGroup
	User  user.UserRouteGroup
}

var RouteApp = new(RouteGroup)
