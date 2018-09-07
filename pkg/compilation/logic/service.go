package logic

import (
	"context"

	"github.com/pkg/errors"

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

// TODO use GetObject from Intent interface instead of passing r
func (s *Service) handleCreate(
	ctx context.Context,
	i intent.Intent,
	intentLogic func(ctx context.Context, ec *intent.EvaluateContext) error,
	r services.Resource,
) error {
	s.cache.Store(i)

	ec := &intent.EvaluateContext{
		WriteService: s.WriteService,
	}

	if intentLogic != nil {
		err := intentLogic(ctx, ec)
		if err != nil {
			return err
		}
	}

	if err := s.EvaluateDependencies(ctx, ec, r); err != nil {
		return errors.Wrapf(err, "failed to evaluate %s dependencies", i.Kind())
	}
	return nil
}
