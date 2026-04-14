package httpserver

import (
	"net/http"

	"github.com/Tendo33/go-template/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(serviceName string, logger *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	healthService := service.NewHealthService(serviceName)
	router := gin.New()
	router.Use(defaultMiddleware(logger)...)
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, healthService.Status())
	})

	return router
}
