package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
)

// CreateQosConfig does pre-check for QoS Config create.
func (sv *ContrailTypeLogicService) CreateQosConfig(
	ctx context.Context,
	request *services.CreateQosConfigRequest,
) (response *services.CreateQosConfigResponse, err error) {
	qos := request.GetQosConfig()
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var defaultGSCUUID string
			defaultGSCUUID, err = sv.getDefaultGlobalSystemConfigUUID(ctx)
			if err != nil {
				return err
			}

			qos.AddDefaultGlobalSystemConfigRef(defaultGSCUUID, sv.getDefaultGlobalSystemConfigFqName())

			if err = qos.CheckQoSValues(); err != nil {
				return err
			}

			response, err = sv.BaseService.CreateQosConfig(ctx, request)
			return err
		})
	return response, err
}

// UpdateQosConfig does pre-check for QoS Config update.
func (sv *ContrailTypeLogicService) UpdateQosConfig(
	ctx context.Context,
	request *services.UpdateQosConfigRequest,
) (response *services.UpdateQosConfigResponse, err error) {
	qos := request.GetQosConfig()

	if err = qos.CheckQoSValues(); err != nil {
		return response, err
	}

	return sv.BaseService.UpdateQosConfig(ctx, request)
}
