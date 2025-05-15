package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func RequsetLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		logAttrs := []slog.Attr{
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int("status", statusCode),
			slog.String("ip", c.ClientIP()),
			slog.Duration("latency", latency),
		}

		if len(c.Errors) > 0 {
			logAttrs = append(logAttrs, slog.String("errors", c.Errors.String()))
		}

		logger.LogAttrs(c.Request.Context(), slog.LevelInfo, "HTTP Request", logAttrs...)
		//logger.Info("HTTP Request", logAttrs...)
	}
}
