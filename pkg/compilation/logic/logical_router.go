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

// GetObject returns embedded resource object
func (i *LogicalRouterIntent) GetObject() basemodels.Object {
	return i.LogicalRouter
}

// CreateLogicalRouter evaluates LogicalRouter dependencies.
func (s *Service) CreateLogicalRouter(
	ctx context.Context, request *services.CreateLogicalRouterRequest,
) (*services.CreateLogicalRouterResponse, error) {

	i := &LogicalRouterIntent{
		LogicalRouter: request.GetLogicalRouter(),
	}

	c := func(ctx context.Context, ec *intent.EvaluateContext) error {
		if len(i.LogicalRouter.GetRouteTargetRefs()) == 0 {
			if err := i.createDefaultRouteTarget(ctx, ec); err != nil {
				return errors.Wrap(err, "failed to create Logical Router's default Route Target")
			}
		}
		return nil
	}

	if err := s.handleCreate(ctx, i, c, i.LogicalRouter); err != nil {
		return nil, err
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
