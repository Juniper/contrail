package logic

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// LogicalRouterIntent contains Intent Compiler state for LogicalRouter.
type LogicalRouterIntent struct {
	BaseIntent
	*models.LogicalRouter
}

// CreateLogicalRouter evaluates LogicalRouter dependencies.
func (s *Service) CreateLogicalRouter(
	ctx context.Context, request *services.CreateLogicalRouterRequest,
) (*services.CreateLogicalRouterResponse, error) {

	obj := request.GetLogicalRouter()

	intent := &LogicalRouterIntent{
		LogicalRouter: obj,
	}

	if _, ok := compilationif.ObjsCache.Load("LogicalRouterIntent"); !ok {
		compilationif.ObjsCache.Store("LogicalRouterIntent", &sync.Map{})
	}

	objMap, ok := compilationif.ObjsCache.Load("LogicalRouterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intent)
	}

	ec := &EvaluateContext{
		WriteService: s.WriteService,
	}

	if len(obj.GetRouteTargetRefs()) == 0 {
		if err := intent.createDefaultRouteTarget(ctx, ec); err != nil {
			return nil, errors.Wrap(err, "failed to create Logical Router's default Route Target")
		}
	}

	if err := EvaluateDependencies(ctx, ec, obj, "LogicalRouter"); err != nil {
		return nil, errors.Wrap(err, "failed to evaluate Logical Router dependencies")
	}

	return s.BaseService.CreateLogicalRouter(ctx, request)
}

func (intent *LogicalRouterIntent) createDefaultRouteTarget(
	ctx context.Context, evaluateContext *EvaluateContext,
) error {
	rt, err := createDefaultRouteTarget(ctx, evaluateContext)
	if err != nil {
		return err
	}

	_, err = evaluateContext.WriteService.CreateLogicalRouterRouteTargetRef(
		ctx,
		&services.CreateLogicalRouterRouteTargetRefRequest{
			ID: intent.GetUUID(),
			LogicalRouterRouteTargetRef: &models.LogicalRouterRouteTargetRef{
				UUID: rt.GetUUID(),
			},
		},
	)

	return err
}
