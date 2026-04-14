package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/Tendo33/go-template/internal/config"
)

func NewServer(cfg config.Config, logger *slog.Logger) *http.Server {
	_ = logger

	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: NewRouter(cfg.ServiceName),
	}
}
