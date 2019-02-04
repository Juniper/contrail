package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
)

// EvaluateDependencies evaluates the dependencies upon object change
func (s *Service) EvaluateDependencies(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	i intent.Intent,
) error {

	logrus.WithFields(logrus.Fields{
		"kind": i.Kind(),
		"uuid": i.GetUUID(),
	}).Debug("Resolving dependencies.")
	dependencies := s.dependencyProcessor.GetDependencies(s.cache, i, "self")

	for _, dependency := range dependencies {
		logrus.WithFields(logrus.Fields{
			"kind": dependency.Kind(),
			"uuid": dependency.GetUUID(),
		}).Debug("Evaluating intent.")

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
