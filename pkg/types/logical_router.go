package types

import (
	"context"

	"github.com/gogo/protobuf/types"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	//"github.com/Juniper/contrail/pkg/models/basemodels"
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

			if err = sv.checkRouterSupportsVpnType(ctx, logicalRouter); err != nil {
				return err
			}

			if err = sv.checkRouterHasBgpvpnAssocViaNetwork(ctx, logicalRouter, nil); err != nil {
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

			if err = sv.checkRouterSupportsVpnType(ctx, logicalRouter); err != nil {
				return err
			}

			if err = sv.checkRouterHasBgpvpnAssocViaNetwork(ctx, logicalRouter, dbLogicalRouter); err != nil {
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

	if fm != nil && !basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldParentUUID) &&
		!basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldVirtualNetworkRefs) {
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

func (sv *ContrailTypeLogicService) checkRouterSupportsVpnType(
	ctx context.Context,
	logicalRouter *models.LogicalRouter) error {

	bgpvpnRefs := logicalRouter.GetBGPVPNRefs()
	bgpvpnUUIDs := make([]string, 0, len(bgpvpnRefs))
	for _, bgpvpnRef := range bgpvpnRefs {
		bgpvpnUUIDs = append(bgpvpnUUIDs, bgpvpnRef.GetUUID())
	}

	if len(bgpvpnUUIDs) == 0 {
		return nil
	}

	bgpvpns, err := sv.ReadService.ListBGPVPN(ctx,
		&services.ListBGPVPNRequest{
			Spec: &baseservices.ListSpec{
				Filters: []*baseservices.Filter{{
					Key:    models.BGPVPNFieldUUID,
					Values: bgpvpnUUIDs,
				}},
				Fields: []string{models.BGPVPNFieldBGPVPNType},
			},
		},
	)
	if err != nil {
		return err
	}

	for _, bgpvpn := range bgpvpns.GetBGPVPNs() {
		if bgpvpn.GetBGPVPNType() != models.L3VPNType {
			return common.ErrorBadRequestf("Can only associate '%s' type BGPVPNs to a logical router", models.L3VPNType)
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkRouterHasBgpvpnAssocViaNetwork( //nolint: gocyclo
	ctx context.Context,
	logicalRouter, dbLogicalRouter *models.LogicalRouter) error {

	bgpvpnRefs := logicalRouter.GetBGPVPNRefs()
	if len(bgpvpnRefs) == 0 {
		bgpvpnRefs = dbLogicalRouter.GetBGPVPNRefs()
	}
	if len(bgpvpnRefs) == 0 {
		return nil
	}

	vmiRefs := logicalRouter.GetVirtualMachineInterfaceRefs()
	if len(vmiRefs) == 0 {
		vmiRefs = dbLogicalRouter.GetVirtualMachineInterfaceRefs()
	}
	if len(vmiRefs) == 0 {
		return nil
	}

	vmiUUIDs := make([]string, 0, len(vmiRefs))
	for _, vmiRef := range vmiRefs {
		vmiUUIDs = append(vmiUUIDs, vmiRef.GetUUID())
	}

	vmisResp, err := sv.ReadService.ListVirtualMachineInterface(ctx,
		&services.ListVirtualMachineInterfaceRequest{
			Spec: &baseservices.ListSpec{
				Filters: []*baseservices.Filter{{
					Key:    models.VirtualMachineInterfaceFieldUUID,
					Values: vmiUUIDs,
				}},
				Detail: true,
				Fields: []string{models.VirtualMachineInterfaceFieldVirtualNetworkRefs},
			},
		},
	)
	if err != nil {
		return err
	}

	var vnUUIDs []string
	for _, vmi := range vmisResp.GetVirtualMachineInterfaces() {
		vnRefs := vmi.GetVirtualNetworkRefs()
		for _, vnRef := range vnRefs {
			vnUUIDs = append(vnUUIDs, vnRef.GetUUID())
		}
	}
	if len(vnUUIDs) == 0 {
		return nil
	}

	vnResp, err := sv.ReadService.ListVirtualNetwork(ctx,
		&services.ListVirtualNetworkRequest{
			Spec: &baseservices.ListSpec{
				Filters: []*baseservices.Filter{{
					Key:    models.VirtualNetworkFieldUUID,
					Values: vnUUIDs,
				}},
				Detail: true,
				Fields: []string{models.VirtualNetworkFieldBGPVPNRefs},
			},
		},
	)
	if err != nil {
		return err
	}

	for _, vn := range vnResp.GetVirtualNetworks() {
		if len(vn.GetBGPVPNRefs()) != 0 {
			return common.ErrorBadRequestf(
				"Can not associate BGPVPN to router which is linked to a network(%s) "+
					"which already has BGPVPN associated", vn.GetUUID())
		}
	}

	return nil
}
