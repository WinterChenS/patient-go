package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"winterchen.com/patient-go/patient/global"
	"go.uber.org/zap"
)

// global logger for gin requests
func LoggerForGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		global.Log.Info("access log",
			zap.Int("status", status),
			zap.String("latency", latency.String()),
			zap.String("clientIP", clientIP),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("raw", raw),
		)
	}

}
