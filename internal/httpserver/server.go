package httpserver

import (
	"net/http"

	"github.com/Tendo33/go-template/internal/config"
	"go.uber.org/zap"
)

func NewServer(cfg config.Config, logger *zap.Logger) *http.Server {
	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: NewRouter(cfg.ServiceName, logger),
	}
}
