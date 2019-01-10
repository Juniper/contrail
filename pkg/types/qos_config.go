package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const defaultGlobalSystemConfigFqName = "default-global-system-config"

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

			qos.AddDefaultGlobalSystemConfigRef(defaultGSCUUID, []string{defaultGlobalSystemConfigFqName})

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

func (sv *ContrailTypeLogicService) getDefaultGlobalSystemConfigUUID(ctx context.Context) (string, error) {
	defaultGSCUUID, err := sv.MetadataGetter.GetMetadata(
		ctx,
		basemodels.Metadata{
			FQName: []string{defaultGlobalSystemConfigFqName},
			Type:   models.KindGlobalSystemConfig,
		},
	)
	if err != nil {
		return "", errutil.ErrorBadRequestf("Cannot resolve Global System Config with FQName %v: %v",
			defaultGlobalSystemConfigFqName, err)
	}
	return defaultGSCUUID.UUID, nil
}
