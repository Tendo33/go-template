package httpserver

import (
	"github.com/Tendo33/go-template/internal/observability"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func defaultMiddleware(logger *zap.Logger) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		requestid.New(),
		contextLogger(logger),
		ginzap.GinzapWithConfig(logger, &ginzap.Config{
			TimeFormat:   "2006-01-02T15:04:05.000Z07:00",
			UTC:          false,
			DefaultLevel: zapcore.InfoLevel,
			Context: func(ctx *gin.Context) []zapcore.Field {
				fields := []zapcore.Field{
					zap.String("request_id", requestid.Get(ctx)),
				}

				if route := ctx.FullPath(); route != "" {
					fields = append(fields, zap.String("route", route))
				}

				return fields
			},
		}),
		ginzap.RecoveryWithZap(logger, true),
	}
}

func contextLogger(baseLogger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := requestid.Get(ctx)

		requestContext := observability.WithLogger(ctx.Request.Context(), baseLogger)
		requestContext = observability.WithRequestID(requestContext, requestID)
		ctx.Request = ctx.Request.WithContext(requestContext)

		ctx.Next()
	}
}
