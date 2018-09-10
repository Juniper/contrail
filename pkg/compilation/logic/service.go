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
	ReadService  services.ReadService
	cache        *intent.Cache
}

// NewService creates a Service
func NewService(
	writeService services.WriteService,
	readService services.ReadService,
	cache *intent.Cache,
) *Service {
	return &Service{
		WriteService: writeService,
		ReadService:  readService,
		cache:        cache,
	}
}

func (s *Service) evaluateContext() *intent.EvaluateContext {
	return &intent.EvaluateContext{
		WriteService: s.WriteService,
		ReadService:  s.ReadService,
		IntentLoader: s.cache,
	}
}

// TODO use GetObject from Intent interface instead of passing r
func (s *Service) handleCreate(
	ctx context.Context,
	i intent.Intent,
	r services.Resource,
) error {
	s.cache.Store(i)

	ec := s.evaluateContext()

	if err := s.EvaluateDependencies(ctx, ec, r); err != nil {
		return errors.Wrapf(err, "failed to evaluate %s dependencies", i.Kind())
	}
	return nil
}
