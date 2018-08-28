package types

import (
	"context"

	"github.com/gogo/protobuf/types"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

//CreateLogicalRouter validates logical-router create request
func (sv *ContrailTypeLogicService) CreateLogicalRouter(
	ctx context.Context,
	request *services.CreateLogicalRouterRequest) (*services.CreateLogicalRouterResponse, error) {

	var response *services.CreateLogicalRouterResponse
	logicalRouter := request.GetLogicalRouter()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			err = sv.checkForExternalGateway(ctx, logicalRouter, nil)
			if err != nil {
				return err
			}

			err = sv.checkPortGatewayNetworks(ctx, logicalRouter, nil)
			if err != nil {
				return err
			}

			err = sv.checkPortAvailability(ctx, logicalRouter)
			if err != nil {
				return err
			}

			logicalRouter.VxlanNetworkIdentifier = logicalRouter.GetVXLanIDInLogicaRouter()

			//TODO check router supports vpn type
			response, err = sv.BaseService.CreateLogicalRouter(ctx, request)
			return err

			//TODO post-create creating internal virtual network
		})

	return response, err
}

//UpdateLogicalRouter validates logical-router update request
func (sv *ContrailTypeLogicService) UpdateLogicalRouter(
	ctx context.Context,
	request *services.UpdateLogicalRouterRequest) (*services.UpdateLogicalRouterResponse, error) {

	var response *services.UpdateLogicalRouterResponse
	logicalRouter := request.GetLogicalRouter()
	fieldMask := request.GetFieldMask()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			err = sv.checkForExternalGateway(ctx, logicalRouter, &fieldMask)
			if err != nil {
				return err
			}

			dbLogicalRouter, err := sv.getLogicalRouter(ctx, logicalRouter.GetUUID())
			if err != nil {
				return err
			}

			err = sv.checkPortGatewayNetworks(ctx, logicalRouter, dbLogicalRouter)
			if err != nil {
				return err
			}

			err = sv.checkPortAvailability(ctx, logicalRouter)
			if err != nil {
				return err
			}

			logicalRouter.VxlanNetworkIdentifier = logicalRouter.GetVXLanIDInLogicaRouter()

			//TODO check BGP VPNs
			response, err = sv.BaseService.UpdateLogicalRouter(ctx, request)
			return err

			//TODO post update changes
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getLogicalRouter(
	ctx context.Context, id string,
) (*models.LogicalRouter, error) {

	logicalRouterResponse, err := sv.ReadService.GetLogicalRouter(
		ctx,
		&services.GetLogicalRouterRequest{
			ID: id,
		},
	)

	return logicalRouterResponse.GetLogicalRouter(), err
}

func (sv *ContrailTypeLogicService) checkForExternalGateway(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
	fm *types.FieldMask) error {

	if fm != nil && !basemodels.FieldMaskContains(fm, models.LogicalRouterFieldParentUUID) &&
		!basemodels.FieldMaskContains(fm, models.LogicalRouterFieldVirtualNetworkRefs) {
		return nil
	}

	enabled, err := sv.isVxlanRoutingEnabled(ctx, logicalRouter)
	if err != nil {
		return err
	}

	if enabled {
		for _, vn := range logicalRouter.GetVirtualNetworkRefs() {
			if vn.GetAttr().GetLogicalRouterVirtualNetworkType() != "InternalVirtualNetwork" {
				return common.ErrorBadRequest("external gateway not supported with VxLAN")
			}
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) isVxlanRoutingEnabled(
	ctx context.Context,
	logicalRouter *models.LogicalRouter) (bool, error) {

	projectUUID := logicalRouter.GetParentUUID()
	if projectUUID == "" {

		uuid := logicalRouter.GetUUID()
		if uuid == "" {
			return false, common.ErrorBadRequest("no input to derive parent for Logical Router")
		}

		logicalRouterResponse, err := sv.ReadService.GetLogicalRouter(
			ctx,
			&services.GetLogicalRouterRequest{
				ID: uuid,
			},
		)
		if err != nil {
			return false, err
		}

		projectUUID = logicalRouterResponse.GetLogicalRouter().GetParentUUID()
	}

	projectResponse, err := sv.ReadService.GetProject(
		ctx,
		&services.GetProjectRequest{
			ID:     projectUUID,
			Fields: []string{models.ProjectFieldVxlanRouting},
		},
	)

	return projectResponse.GetProject().GetVxlanRouting(), err
}

func (sv *ContrailTypeLogicService) checkPortGatewayNetworks(
	ctx context.Context,
	requestLR *models.LogicalRouter,
	databaseLR *models.LogicalRouter,
) error {

	err := sv.checkPortGateway(ctx, requestLR.GetVirtualMachineInterfaceRefs(), requestLR.GetVirtualNetworkRefs())
	if err != nil {
		return err
	}

	if databaseLR == nil {
		return nil
	}

	err = sv.checkPortGateway(ctx, requestLR.GetVirtualMachineInterfaceRefs(), databaseLR.GetVirtualNetworkRefs())
	if err != nil {
		return err
	}

	return sv.checkPortGateway(ctx, databaseLR.GetVirtualMachineInterfaceRefs(), requestLR.GetVirtualNetworkRefs())
}

func (sv *ContrailTypeLogicService) checkPortGateway(
	ctx context.Context,
	vmiRefs []*models.LogicalRouterVirtualMachineInterfaceRef,
	vnRefs []*models.LogicalRouterVirtualNetworkRef,
) error {

	var interfaceNetworkUUIDs []string
	for _, vmi := range vmiRefs {
		vmiResponse, err := sv.ReadService.GetVirtualMachineInterface(
			ctx,
			&services.GetVirtualMachineInterfaceRequest{
				ID:     vmi.GetUUID(),
				Fields: []string{models.VirtualMachineInterfaceFieldVirtualNetworkRefs},
			},
		)
		if err != nil {
			return err
		}

		vnRefs := vmiResponse.GetVirtualMachineInterface().GetVirtualNetworkRefs()
		if len(vnRefs) > 0 {
			interfaceNetworkUUIDs = append(
				interfaceNetworkUUIDs,
				vnRefs[0].GetUUID(),
			)
		}
	}

	for _, vn := range vnRefs {
		if common.ContainsString(interfaceNetworkUUIDs, vn.GetUUID()) {
			return common.ErrorBadRequestf(
				"logical router interface and gateway cannot be in the same virtual network(%s)", vn.GetUUID(),
			)
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkPortAvailability(
	ctx context.Context,
	logicalRouter *models.LogicalRouter) error {

	for _, vmiRef := range logicalRouter.GetVirtualMachineInterfaceRefs() {
		vmiResponse, err := sv.ReadService.GetVirtualMachineInterface(
			ctx,
			&services.GetVirtualMachineInterfaceRequest{
				ID: vmiRef.GetUUID(),
			},
		)
		if err != nil {
			return err
		}

		vmi := vmiResponse.GetVirtualMachineInterface()
		if vmi.GetParentType() == models.KindVirtualMachine || len(vmi.GetVirtualMachineRefs()) > 0 {
			return common.ErrorConflictf("port(%s) already in use by virtual-machine", vmi.GetUUID())
		}
	}

	return nil
}
