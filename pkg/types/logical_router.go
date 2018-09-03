package types

import (
	"context"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateLogicalRouter validates logical-router create request
func (sv *ContrailTypeLogicService) CreateLogicalRouter(
	ctx context.Context,
	request *services.CreateLogicalRouterRequest) (*services.CreateLogicalRouterResponse, error) {

	var response *services.CreateLogicalRouterResponse
	logicalRouter := request.GetLogicalRouter()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			vxLanRouting, err := sv.checkForExternalGateway(ctx, logicalRouter, nil)
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

			logicalRouter.VxlanNetworkIdentifier, err = logicalRouter.GetVXLanIDInLogicaRouter()
			if err != nil {
				return err
			}

			if vxLanRouting {
				var internalVN *models.VirtualNetwork
				internalVN, err = sv.createInternalVirtualNetwork(ctx, logicalRouter)
				if err != nil {
					return err
				}

				//return common.ErrorBadRequestf("ZLEEEEEEEEEEEEEEEeeee: %s", internalVN.ParentUUID)

				vnRef := &models.LogicalRouterVirtualNetworkRef{
					UUID: internalVN.GetUUID(),
					To:   internalVN.GetFQName(),
					Attr: &models.LogicalRouterVirtualNetworkType{
						LogicalRouterVirtualNetworkType: "InternalVirtualNetwork",
					},
				}
				request.LogicalRouter.AddVirtualNetworkRef(vnRef)
			}

			//TODO check router supports vpn type
			response, err = sv.BaseService.CreateLogicalRouter(ctx, request)
			return err
		})

	return response, err
}

// UpdateLogicalRouter validates logical-router update request
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

			vxLanRouting, err := sv.checkForExternalGateway(ctx, logicalRouter, &fieldMask)
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

			logicalRouter.VxlanNetworkIdentifier, err = logicalRouter.GetVXLanIDInLogicaRouter()
			if err != nil {
				return err
			}

			if vxLanRouting {
				err = sv.updateInternalVirtualNetwork(ctx, logicalRouter, dbLogicalRouter, &fieldMask)
				if err != nil {
					return err
				}
			}

			//TODO check BGP VPNs
			response, err = sv.BaseService.UpdateLogicalRouter(ctx, request)
			return err
		})

	return response, err
}

// DeleteLogicalRouter deletes internal vn if xvlan is enabled
func (sv *ContrailTypeLogicService) DeleteLogicalRouter(
	ctx context.Context,
	request *services.DeleteLogicalRouterRequest) (*services.DeleteLogicalRouterResponse, error) {

	var response *services.DeleteLogicalRouterResponse
	uuid := request.GetID()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			logicalRouterResponse, err := sv.ReadService.GetLogicalRouter(
				ctx,
				&services.GetLogicalRouterRequest{
					ID: uuid,
				},
			)
			if err != nil {
				return err
			}

			logicalRouter := logicalRouterResponse.GetLogicalRouter()
			vxLanRouting, err := sv.isVxlanRoutingEnabled(ctx, logicalRouter)
			if err != nil {
				return err
			}

			if vxLanRouting {
				err = sv.deleteInternalVirtualNetwork(ctx, logicalRouter)
				if err != nil {
					return err
				}
			}

			response, err = sv.BaseService.DeleteLogicalRouter(ctx, request)
			return err
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
	fm *types.FieldMask,
) (bool, error) {

	if fm != nil && !basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldParentUUID) &&
		!basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldVirtualNetworkRefs) {
		return false, nil
	}

	enabled, err := sv.isVxlanRoutingEnabled(ctx, logicalRouter)
	if err != nil {
		return enabled, err
	}

	if enabled {
		for _, vn := range logicalRouter.GetVirtualNetworkRefs() {
			if vn.GetAttr().GetLogicalRouterVirtualNetworkType() != "InternalVirtualNetwork" {
				return false, common.ErrorBadRequest("external gateway not supported with VxLAN")
			}
		}
	}

	return enabled, nil
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

func (sv *ContrailTypeLogicService) createInternalVirtualNetwork(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
) (*models.VirtualNetwork, error) {

	internalVN := models.MakeVirtualNetwork()
	internalVN.UUID = uuid.NewV4().String()
	internalVN.ParentUUID = logicalRouter.GetParentUUID()
	internalVN.ParentType = models.KindProject
	internalVN.Name = logicalRouter.GetInternalVNName()
	internalVN.RouteTargetList = logicalRouter.GetConfiguredRouteTargetList()
	internalVN.IDPerms = &models.IdPermsType{
		Enable:      true,
		UserVisible: false,
	}

	vxlanNetworkID, _ := strconv.ParseInt(logicalRouter.GetVxlanNetworkIdentifier(), 10, 64)
	internalVN.VirtualNetworkProperties = &models.VirtualNetworkType{
		VxlanNetworkIdentifier: vxlanNetworkID,
		ForwardingMode:         models.L3Mode,
	}

	request := &services.CreateVirtualNetworkRequest{
		VirtualNetwork: internalVN,
	}

	response, err := sv.WriteService.CreateVirtualNetwork(ctx, request)
	return response.GetVirtualNetwork(), err
}

func (sv *ContrailTypeLogicService) updateInternalVirtualNetwork(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
	dbLogicalRouter *models.LogicalRouter,
	fm *types.FieldMask,
) error {

	if !basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldConfiguredRouteTargetList) &&
		!basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldVxlanNetworkIdentifier) {
		return nil
	}

	var uuid string
	for _, vn := range dbLogicalRouter.GetVirtualNetworkRefs() {
		if vn.GetAttr().GetLogicalRouterVirtualNetworkType() == "InternalVirtualNetwork" {
			uuid = vn.GetUUID()
			break
		}
	}

	if uuid == "" {
		return nil
	}

	updateVN := models.MakeVirtualNetwork()
	updateVN.UUID = uuid
	var updatePaths []string

	if basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldConfiguredRouteTargetList) {
		updateVN.RouteTargetList = logicalRouter.GetConfiguredRouteTargetList()
		updatePaths = append(updatePaths, models.VirtualNetworkFieldRouteTargetList)
	}

	if basemodels.FieldMaskContains(*fm, models.LogicalRouterFieldVxlanNetworkIdentifier) {
		vxlanNetworkID, _ := strconv.ParseInt(logicalRouter.GetVxlanNetworkIdentifier(), 10, 64)
		updateVN.VirtualNetworkProperties = &models.VirtualNetworkType{
			VxlanNetworkIdentifier: vxlanNetworkID,
		}

		path := []string{
			models.VirtualNetworkFieldVirtualNetworkProperties,
			models.VirtualNetworkTypeFieldVxlanNetworkIdentifier,
		}
		updatePaths = append(updatePaths, strings.Join(path, "."))
	}

	_, err := sv.WriteService.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
		VirtualNetwork: updateVN,
		FieldMask:      types.FieldMask{Paths: updatePaths},
	})

	return err
}

func (sv *ContrailTypeLogicService) deleteInternalVirtualNetwork(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
) error {

	m, err := sv.MetaDataGetter.GetMetadata(ctx, "", fqName)
	if err != nil {
		return err
	}

	_, err = sv.WriteService.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{
		ID: m.GetMetadata().UUID,
	})

	return err
}
