package middlewares

import (

	"github.com/gin-gonic/gin"
)

// CorsMiddleware implements IMiddleware.
func (m *Middleware) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}