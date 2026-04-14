package httpserver

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func defaultMiddleware(logger *zap.Logger) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Recovery(),
		requestLogger(logger),
	}
}

func requestLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		path := ctx.FullPath()
		if path == "" {
			path = ctx.Request.URL.Path
		}

		logger.Info(
			"request completed",
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.Int("status", ctx.Writer.Status()),
			zap.Int64("latency_ms", time.Since(start).Milliseconds()),
		)
	}
}
