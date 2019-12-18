package types

import (
	"context"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateBGPAsAService creates BGP as a service instance and ensures that shared one have IP address
func (sv *ContrailTypeLogicService) CreateBGPAsAService(
	ctx context.Context, request *services.CreateBGPAsAServiceRequest,
) (response *services.CreateBGPAsAServiceResponse, err error) {
	if err = sv.checkSharedAddress(request.GetBGPAsAService()); err != nil {
		return nil, errutil.ErrorBadRequest("BGPaaS IP Address needs to be configured if BGPaaS is shared")
	}
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			response, err = sv.BaseService.CreateBGPAsAService(ctx, request)
			return err
		})
	return response, err
}

// UpdateBGPAsAService updates BGP as a service after validation
func (sv *ContrailTypeLogicService) UpdateBGPAsAService(
	ctx context.Context, request *services.UpdateBGPAsAServiceRequest,
) (response *services.UpdateBGPAsAServiceResponse, err error) {
	id := request.GetBGPAsAService().GetUUID()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			fm := request.GetFieldMask()
			if basemodels.FieldMaskContains(&fm, models.BGPAsAServiceFieldBgpaasShared) {
				bgpaas, terr := sv.ReadService.GetBGPAsAService(ctx, &services.GetBGPAsAServiceRequest{ID: id})
				if terr != nil {
					return terr
				}
				if bgpaas.GetBGPAsAService().GetBgpaasShared() != request.GetBGPAsAService().GetBgpaasShared() {
					return errutil.ErrorBadRequest("BGPaaS sharing cannot be modified")
				}
			}
			response, err = sv.BaseService.UpdateBGPAsAService(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) checkSharedAddress(bgpaas *models.BGPAsAService) error {
	if bgpaas.GetBgpaasShared() && len(bgpaas.GetBgpaasIPAddress()) == 0 {
		return errutil.ErrorBadRequest("BGPaaS IP Address needs to be configured if BGPaaS is shared")
	}
	return nil
}
