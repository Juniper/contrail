package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
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

func (s *Service) evaluateContext() *intent.EvaluateContext {
	return &intent.EvaluateContext{
		WriteService: s.WriteService,
		Cache:        s.cache,
	}
}

// TODO use GetObject from Intent interface instead of passing r
func (s *Service) handleCreate(
	ctx context.Context,
	i intent.Intent,
	intentLogic func(ctx context.Context, ec *intent.EvaluateContext) error,
	r services.Resource,
) error {
	s.cache.Store(i)

	ec := s.evaluateContext()

	if intentLogic != nil {
		err := intentLogic(ctx, ec)
		if err != nil {
			return err
		}
	}

	if err := s.EvaluateDependencies(ctx, ec, r); err != nil {
		return errors.Wrap(err, "failed to evaluate Logical Router dependencies")
	}
	return nil
}
