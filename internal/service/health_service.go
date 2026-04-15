package service

import (
	"context"

	"github.com/Tendo33/go-template/internal/model"
	"github.com/Tendo33/go-template/internal/observability"
)

type HealthService struct {
	serviceName string
}

func NewHealthService(serviceName string) HealthService {
	return HealthService{serviceName: serviceName}
}

func (s HealthService) Status(ctx context.Context) model.HealthResponse {
	observability.FromContext(ctx).Info("health status requested")

	return model.HealthResponse{
		Status:  "ok",
		Service: s.serviceName,
	}
}
