package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

//CreateLogicalRouter TODO
func (sv *ContrailTypeLogicService) CreateLogicalRouter(
	ctx context.Context,
	request *services.CreateLogicalRouterRequest) (*services.CreateLogicalRouterResponse, error) {

	var response *services.CreateLogicalRouterResponse
	logicalRouter := request.GetLogicalRouter()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			err = sv.checkForExternalGateway(ctx, logicalRouter)
			if err != nil {
				return err
			}

			err = sv.checkPortGatewayNotInSameNetwork(ctx, logicalRouter)
			if err != nil {
				return err
			}

			err = sv.checkPortInUseByVirtualMachine(ctx, logicalRouter)
			if err != nil {
				return err
			}

			vxLanID := logicalRouter.GetVXLanIDInLogicaRouter()

			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) checkForExternalGateway(
	ctx context.Context,
	logicalRouter *models.LogicalRouter) error {

	enabled, err := sv.isVxlanRoutingEnabled(ctx, logicalRouter)
	if err != nil {
		return err
	}

	if enabled {
		for _, vn := range logicalRouter.GetVirtualNetworkRefs() {
			if vn.GetAttr() == nil || vn.GetAttr().GetLogicalRouterVirtualNetworkType() != "InternalVirtualNetwork" {
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

		logicalRouterResponse, err := sv.ReadService.GetLogicalRouter(ctx,
			&services.GetLogicalRouterRequest{
				ID: uuid,
			})
		if err != nil {
			return false, err
		}

		projectUUID = logicalRouterResponse.GetLogicalRouter().GetParentUUID()
	}

	projectResponse, err := sv.ReadService.GetProject(ctx,
		&services.GetProjectRequest{
			ID: projectUUID,
		})
	if err != nil {
		return false, err
	}

	return projectResponse.GetProject().GetVxlanRouting(), nil
}

func (sv *ContrailTypeLogicService) getLogicalRouterVMIRefs(
	ctx context.Context,
	logicalRouter *models.LogicalRouter) ([]*models.VirtualMachineInterface, error) {

	var vmiRefs []*models.VirtualMachineInterface
	for _, vmi := range logicalRouter.GetVirtualMachineInterfaceRefs() {
		vmiResponse, err := sv.ReadService.GetVirtualMachineInterface(ctx,
			&services.GetVirtualMachineInterfaceRequest{
				ID: vmi.GetUUID(),
			})
		if err != nil {
			return nil, err
		}

		vmiRefs = append(vmiRefs, vmiResponse.GetVirtualMachineInterface())
	}

	return vmiRefs, nil
}

func (sv *ContrailTypeLogicService) checkPortGatewayNotInSameNetwork(
	ctx context.Context,
	logicalRouter *models.LogicalRouter) error {

	vmiRefs, err := sv.getLogicalRouterVMIRefs(ctx, logicalRouter)
	if err != nil {
		return err
	}

	var interfaceNetworkUUIDs []string
	for _, vmi := range vmiRefs {
		interfaceNetworkUUIDs = append(
			interfaceNetworkUUIDs,
			vmi.GetVirtualNetworkRefs()[0].GetUUID(),
		)
	}

	for _, vn := range logicalRouter.GetVirtualNetworkRefs() {
		if common.ContainsString(interfaceNetworkUUIDs, vn.GetUUID()) {
			return common.ErrorBadRequestf(
				"logical router interface and gateway cannot be in the same virtual network(%s)", vn.GetUUID(),
			)
		}
	}

	//TODO update part
	return nil
}

func (sv *ContrailTypeLogicService) checkPortInUseByVirtualMachine(
	ctx context.Context,
	logicalRouter *models.LogicalRouter) error {

	vmiRefs, err := sv.getLogicalRouterVMIRefs(ctx, logicalRouter)
	if err != nil {
		return err
	}

	for _, vmi := range vmiRefs {
		if vmi.GetParentType() == "virtual-machine" || len(vmi.GetVirtualMachineRefs()) > 0 {
			return common.ErrorConflictf("port(%s) already in use by virtual-machine", vmi.GetUUID())
		}
	}

	return nil
}
