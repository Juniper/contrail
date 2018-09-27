package logic

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
)

// EvaluateDependencies evaluates the dependencies upon object change
func (s *Service) EvaluateDependencies(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	i intent.Intent,
) error {

	log.Printf("EvaluateDependencies called for (%s): \n", i.TypeName())
	dependencies := s.cache.GetDependencies(i, "Self")

	for _, dependency := range dependencies {
		log.WithFields(log.Fields{"type-name": i.Kind(), "uuid": i.GetUUID()}).Printf("Processing intent")
		err := dependency.Evaluate(ctx, evaluateCtx)
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to evaluate intent of type %s with uuid %s",
				i.Kind(),
				i.GetUUID(),
			)
		}
	}
	return nil
}
