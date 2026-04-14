package main

import (
	"log"

	"github.com/Tendo33/go-template/internal/config"
	"github.com/Tendo33/go-template/internal/httpserver"
	"github.com/Tendo33/go-template/internal/observability"
)

func main() {
	cfg := config.Load()
	logger := observability.NewLogger(cfg.LogLevel)
	server := httpserver.NewServer(cfg, logger)

	logger.Info("starting server", "addr", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
