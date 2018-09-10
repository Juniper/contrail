package logic

import (
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementing Intent Compiler's type-specific logic.
type Service struct {
	services.BaseService
	// WriteService is used to create/update/delete lower-level resources
	WriteService services.WriteService
	cache        *intent.Cache
}

// NewService creates a Service
func NewService(apiClient services.WriteService, cache *intent.Cache) *Service {
	return &Service{
		WriteService: apiClient,
		cache:        cache,
	}
}

func (s *Service) evaluateContext() *EvaluateContext {
	return &EvaluateContext{
		WriteService: s.WriteService,
		IntentLoader: s.cache,
	}
}
