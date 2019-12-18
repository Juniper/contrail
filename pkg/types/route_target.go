package types

import (
	"context"

	"github.com/Juniper/asf/pkg/errutil"
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
			err = routeTarget.Validate()
			if err != nil {
				return errutil.ErrorBadRequestf("Validation of route target name failed with error: %v", err)
			}

			response, err = sv.BaseService.CreateRouteTarget(ctx, request)
			return err
		})
	return response, err
}
