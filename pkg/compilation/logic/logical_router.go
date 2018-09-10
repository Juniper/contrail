package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// LogicalRouterIntent contains Intent Compiler state for LogicalRouter.
type LogicalRouterIntent struct {
	intent.BaseIntent
	*models.LogicalRouter
	virtualNetworks map[string]*VirtualNetworkIntent
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

	ec := s.evaluateContext()

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

func (i *LogicalRouterIntent) checkVnDiff(
	vns map[string]*VirtualNetworkIntent,
) bool {
	for k := range vns {
		_, present := i.virtualNetworks[k]
		if !present {
			return false
		}
	}
	return true
}

func (i *LogicalRouterIntent) Evaluate(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
) error {
	i.updateVirtualNetworks(ctx, evaluateCtx)
	return nil
}

func (i *LogicalRouterIntent) updateVirtualNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
) {
	var vns map[string]*VirtualNetworkIntent
	for _, vmi := range i.VirtualMachineInterfaceRefs {
		vn, ok := LoadVirtualNetworkIntent(evaluateCtx.Cache, vmi.UUID)
		if ok {
			vns[vmi.UUID] = vn
		}
	}
	if i.checkVnDiff(vns) {
		return
	}
	// TODO implement logic for changed vns
}

func (i *LogicalRouterIntent) createDefaultRouteTarget(
	ctx context.Context,
	evaluateContext *intent.EvaluateContext,
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
