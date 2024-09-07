package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/pkg/response"
)

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (m *Middleware) ApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		if apiKey == "" {
			response.ErrorMiddlewareReponse(c, response.ErrCodeForbidden, "error")
			return
		}

		key, _ := m.authRepo.GetAPIKey(apiKey)
		if key == nil {
			response.ErrorMiddlewareReponse(c, response.ErrCodeForbidden, "apiKey not found!")
			return
		}
		
		c.Next()
	}
}

func (m *Middleware) PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
