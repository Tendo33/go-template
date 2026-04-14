package service

import "github.com/Tendo33/go-template/internal/model"

type HealthService struct {
	serviceName string
}

func NewHealthService(serviceName string) HealthService {
	return HealthService{serviceName: serviceName}
}

func (s HealthService) Status() model.HealthResponse {
	return model.HealthResponse{
		Status:  "ok",
		Service: s.serviceName,
	}
}
