package main

import (
	"github.com/Tendo33/go-template/internal/config"
	"github.com/Tendo33/go-template/internal/httpserver"
	"github.com/Tendo33/go-template/internal/observability"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	logger := observability.NewLogger(cfg.LogLevel)
	server := httpserver.NewServer(cfg, logger)
	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("starting server", zap.String("addr", server.Addr))

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("server stopped", zap.Error(err))
	}
}
