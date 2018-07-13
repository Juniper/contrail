package logic

import "github.com/Juniper/contrail/pkg/services"

// Service implementing Intent Compiler's type-specific logic.
type Service struct {
	services.BaseService
	// api is used to create/update/delete lower-level resources
	api services.Service
}

// NewService creates a Service
func NewService(apiClient services.Service) *Service {
	return &Service{
		api: apiClient,
	}
}
