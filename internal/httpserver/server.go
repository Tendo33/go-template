package httpserver

import (
	"net/http"
	"time"

	"github.com/Tendo33/go-template/internal/config"
	"go.uber.org/zap"
)

const (
	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 15 * time.Second
	defaultWriteTimeout      = 15 * time.Second
	defaultIdleTimeout       = 60 * time.Second
)

func NewServer(cfg config.Config, logger *zap.Logger) *http.Server {
	return &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           NewRouter(cfg.ServiceName, logger),
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		IdleTimeout:       defaultIdleTimeout,
	}
}
