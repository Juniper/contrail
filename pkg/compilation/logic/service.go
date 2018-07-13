package logic

import "github.com/Juniper/contrail/pkg/services"

type Service struct {
	services.BaseService
	api services.Service
}

func NewService(apiClient services.Service) *Service {
	return &Service{
		api: apiClient,
	}
}
