package middlewares

import "github.com/gin-gonic/gin"

func (m *middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (m *middleware) ApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (m *middleware) PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
