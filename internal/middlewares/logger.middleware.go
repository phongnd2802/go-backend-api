package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phongnd2802/go-backend-api/global"
	"go.uber.org/zap"
)

// LoggerWrite implements IMiddleware.
func (m *Middleware) LoggerWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		clientIP := c.ClientIP()

		c.Next()

		// Wrire Log
		duration := time.Since(startTime)
		global.Logger.Info(fmt.Sprintf("From IP: %s", clientIP),
		 zap.String("Request Method: ", c.Request.Method),
		 zap.String("Request URI: ", c.Request.RequestURI),
		 zap.Int("Status: ", c.Writer.Status()),
		 zap.Duration("Duration: ", duration),
		)

	}
}