package logic

import "github.com/Juniper/contrail/pkg/services"

// Service implementing Intent Compiler's type-specific logic.
type Service struct {
	services.BaseService
	// WriteService is used to create/update/delete lower-level resources
	WriteService     services.WriteService
	IntPoolAllocator services.IntPoolAllocator
}

// NewService creates a Service
func NewService(apiClient services.WriteService, allocator services.IntPoolAllocator) *Service {
	return &Service{
		WriteService:     apiClient,
		IntPoolAllocator: allocator,
	}
}
