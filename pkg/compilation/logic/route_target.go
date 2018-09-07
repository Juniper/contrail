package logic

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// TODO: get_autonomous_system method
const (
	defaultAutonomousSystem = 64512
	routeTargetIntPoolID    = "route_target_number"
)

func createDefaultRouteTarget(
	ctx context.Context, evaluateContext *EvaluateContext,
) (*models.RouteTarget, error) {
	fmt.Println(evaluateContext)
	target, err := evaluateContext.IntPoolAllocator.AllocateInt(ctx, routeTargetIntPoolID)
	if err != nil {
		return nil, err
	}

	rtKey := models.RouteTargetString(defaultAutonomousSystem, target)

	rtResponse, err := evaluateContext.WriteService.CreateRouteTarget(
		ctx,
		&services.CreateRouteTargetRequest{
			RouteTarget: &models.RouteTarget{
				FQName:      []string{rtKey},
				DisplayName: rtKey,
			},
		},
	)

	return rtResponse.GetRouteTarget(), err
}
