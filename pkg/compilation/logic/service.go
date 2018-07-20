package logic

import (
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementing Intent Compiler's type-specific logic.
type Service struct {
	services.BaseService
	// WriteService is used to create/update/delete lower-level resources
	WriteService *client.HTTP
}

// NewService creates a Service
func NewService(apiClient *client.HTTP) *Service {
	return &Service{
		WriteService: apiClient,
	}
}
