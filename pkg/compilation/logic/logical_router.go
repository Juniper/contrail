package logic

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// LogicalRouterIntent contains Intent Compiler state for LogicalRouter
type LogicalRouterIntent struct {
	BaseIntent
	*models.LogicalRouter
}

// CreateLogicalRouter evaluates LogicalRouter dependencies
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
			return nil, err
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

func createDefaultRouteTarget(
	ctx context.Context, evaluateContext *EvaluateContext,
) (*models.RouteTarget, error) {
	rtKey := fmt.Sprintf("target:%v:%v", defaultAutonomousSystem, generateRandomRouteTargetNumber())

	rtResponse, err := evaluateContext.WriteService.CreateRouteTarget(
		ctx,
		&services.CreateRouteTargetRequest{
			RouteTarget: &models.RouteTarget{
				FQName:      []string{rtKey},
				DisplayName: rtKey,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return rtResponse.GetRouteTarget(), err
}
