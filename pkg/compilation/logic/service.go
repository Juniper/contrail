package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/plugins/contrail/dependencies"
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementing Intent Compiler's type-specific logic.
type Service struct {
	services.BaseService
	// WriteService is used to create/update/delete lower-level resources
	WriteService        services.WriteService
	IntPoolAllocator    services.IntPoolAllocator
	ReadService         services.ReadService
	cache               *intent.Cache
	dependencyProcessor *dependencies.DependencyProcessor
}

// NewService creates a Service
func NewService(
	apiClient services.WriteService,
	readService services.ReadService,
	allocator services.IntPoolAllocator,
	cache *intent.Cache,
	dependencyProcessor *dependencies.DependencyProcessor,
) *Service {
	return &Service{
		WriteService:        apiClient,
		ReadService:         readService,
		IntPoolAllocator:    allocator,
		cache:               cache,
		dependencyProcessor: dependencyProcessor,
	}
}

func (s *Service) evaluateContext() *intent.EvaluateContext {
	return &intent.EvaluateContext{
		WriteService:     s.WriteService,
		ReadService:      s.ReadService,
		IntPoolAllocator: s.IntPoolAllocator,
		IntentLoader:     s.cache,
	}
}

// TODO use GetObject from Intent interface instead of passing r
func (s *Service) handleCreate(
	ctx context.Context,
	i intent.Intent,
) error {
	s.cache.Store(i)

	ec := s.evaluateContext()

	if err := s.EvaluateDependencies(ctx, ec, i); err != nil {
		return errors.Wrapf(err, "failed to evaluate %s dependencies", i.Kind())
	}
	return nil
}
