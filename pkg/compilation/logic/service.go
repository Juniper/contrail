package logic

import (
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementing Intent Compiler's type-specific logic.
type Service struct {
	services.BaseService
	// WriteService is used to create/update/delete lower-level resources
	WriteService services.WriteService
}

// NewService creates a Service
func NewService(apiClient services.WriteService) *Service {
	return &Service{
		WriteService: apiClient,
	}
}
