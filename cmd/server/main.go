package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Tendo33/go-template/internal/config"
	"github.com/Tendo33/go-template/internal/httpserver"
	"github.com/Tendo33/go-template/internal/observability"
	"go.uber.org/zap"
)

const shutdownTimeout = 10 * time.Second

func main() {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	logger := observability.NewLogger(cfg)
	server := httpserver.NewServer(cfg, logger)
	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("starting server", zap.String("addr", server.Addr))

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(shutdownSignals)

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return
		}

		logger.Fatal("server stopped", zap.Error(err))
	case sig := <-shutdownSignals:
		logger.Info("shutdown signal received", zap.String("signal", sig.String()))

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Fatal("graceful shutdown failed", zap.Error(err))
		}

		logger.Info("server shutdown complete")
	}
}
