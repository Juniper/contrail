package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// LogicalRouterIntent contains Intent Compiler state for LogicalRouter.
type LogicalRouterIntent struct {
	intent.BaseIntent
	*models.LogicalRouter
}

func (i *LogicalRouterIntent) GetObject() basemodels.Object {
	return i.LogicalRouter
}

// CreateLogicalRouter evaluates LogicalRouter dependencies.
func (s *Service) CreateLogicalRouter(
	ctx context.Context, request *services.CreateLogicalRouterRequest,
) (*services.CreateLogicalRouterResponse, error) {

	obj := request.GetLogicalRouter()

	i := &LogicalRouterIntent{
		LogicalRouter: obj,
	}

	s.cache.Store(i)

	ec := &intent.EvaluateContext{
		WriteService: s.WriteService,
	}

	if len(obj.GetRouteTargetRefs()) == 0 {
		if err := i.createDefaultRouteTarget(ctx, ec); err != nil {
			return nil, errors.Wrap(err, "failed to create Logical Router's default Route Target")
		}
	}

	if err := s.EvaluateDependencies(ctx, ec, obj, "LogicalRouter"); err != nil {
		return nil, errors.Wrap(err, "failed to evaluate Logical Router dependencies")
	}

	return s.BaseService.CreateLogicalRouter(ctx, request)
}

func (i *LogicalRouterIntent) createDefaultRouteTarget(
	ctx context.Context, evaluateContext *intent.EvaluateContext,
) error {
	rt, err := createDefaultRouteTarget(ctx, evaluateContext)
	if err != nil {
		return err
	}

	_, err = evaluateContext.WriteService.CreateLogicalRouterRouteTargetRef(
		ctx,
		&services.CreateLogicalRouterRouteTargetRefRequest{
			ID: i.GetUUID(),
			LogicalRouterRouteTargetRef: &models.LogicalRouterRouteTargetRef{
				UUID: rt.GetUUID(),
			},
		},
	)

	return err
}
