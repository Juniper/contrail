package logic

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

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

	return rtResponse.GetRouteTarget(), err
}
