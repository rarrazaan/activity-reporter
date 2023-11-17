package middleware

import (
	"activity-reporter/logger"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Logger(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Info(map[string]interface{}{
			"id":      requestid.Get(c),
			"method":  reqMethod,
			"latency": latencyTime,
			"uri":     reqUri,
			"status":  statusCode,
			"ip":      clientIP,
		})

		c.Next()
	}
}
