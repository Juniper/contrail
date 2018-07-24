package types

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateRouteTarget validates Route Target name before creation
func (sv *ContrailTypeLogicService) CreateRouteTarget(
	ctx context.Context,
	request *services.CreateRouteTargetRequest,
) (response *services.CreateRouteTargetResponse, err error) {
	routeTarget := request.GetRouteTarget()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			name := routeTarget.GetName()
			_, _, _, err = models.ParseRouteTarget(strings.Split(name, ":"))
			if err != nil {
				return common.ErrorBadRequestf("Validation of route target name failed with error: %v", err)
			}
			response, err = sv.BaseService.CreateRouteTarget(ctx, request)
			return err
		})
	return response, err
}
