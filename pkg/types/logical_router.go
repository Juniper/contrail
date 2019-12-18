package types

import (
	"context"
	"strconv"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/gogo/protobuf/types"

	uuid "github.com/satori/go.uuid"
)

// CreateLogicalRouter creates a logical router and if vxlan routing is enabled
// also creates internal virtual network connected to this logical router
func (sv *ContrailTypeLogicService) CreateLogicalRouter(
	ctx context.Context,
	request *services.CreateLogicalRouterRequest) (*services.CreateLogicalRouterResponse, error) {

	var response *services.CreateLogicalRouterResponse
	logicalRouter := request.GetLogicalRouter()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			project, err := sv.getLogicalRouterParentProject(ctx, logicalRouter, nil, nil)
			if err != nil {
				return err
			}

			vxLanRouting := project.GetVxlanRouting()

			err = sv.checkForExternalGateway(logicalRouter, nil, vxLanRouting)
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

			if vxLanRouting {

				logicalRouter.VxlanNetworkIdentifier = logicalRouter.GetVXLanIDInLogicaRouter()
				logicalRouter.VxlanNetworkIdentifier, err = sv.allocateVxlanNetworkID(ctx, logicalRouter, project)
				if err != nil {
					return err
				}

				var internalVN *models.VirtualNetwork
				internalVN, err = sv.createInternalVirtualNetwork(ctx, logicalRouter)
				if err != nil {
					return err
				}

				vnRef := &models.LogicalRouterVirtualNetworkRef{
					UUID: internalVN.GetUUID(),
					To:   internalVN.GetFQName(),
					Attr: &models.LogicalRouterVirtualNetworkType{
						LogicalRouterVirtualNetworkType: "InternalVirtualNetwork",
					},
				}
				request.LogicalRouter.AddVirtualNetworkRef(vnRef)
			}

			if err = sv.checkRouterSupportsVPNType(ctx, logicalRouter); err != nil {
				return err
			}

			if err = sv.checkRouterHasBGPVPNAssocViaNetwork(ctx, logicalRouter, nil, nil); err != nil {
				return err
			}

			response, err = sv.BaseService.CreateLogicalRouter(ctx, request)
			return err
		})

	return response, err
}

// UpdateLogicalRouter validates logical-router update request and if vxlan routing is
// enabled also updates internal virtual network associated with this logical router
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

			dbLogicalRouter, err := sv.getLogicalRouter(ctx, logicalRouter.GetUUID())
			if err != nil {
				return err
			}

			project, err := sv.getLogicalRouterParentProject(ctx, logicalRouter, dbLogicalRouter, &fieldMask)
			if err != nil {
				return err
			}

			vxLanRouting := project.GetVxlanRouting()

			err = sv.checkForExternalGateway(
				logicalRouter, &fieldMask, vxLanRouting,
			)
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

			if vxLanRouting {
				err = sv.updateInternalVirtualNetwork(ctx, logicalRouter, dbLogicalRouter, &fieldMask, project)
				if err != nil {
					return err
				}
			}

			if err = sv.checkRouterSupportsVPNType(ctx, logicalRouter); err != nil {
				return err
			}

			if err = sv.checkRouterHasBGPVPNAssocViaNetwork(ctx, logicalRouter, dbLogicalRouter, &fieldMask); err != nil {
				return err
			}

			response, err = sv.BaseService.UpdateLogicalRouter(ctx, request)
			return err
		})

	return response, err
}

// DeleteLogicalRouter deletes internal virtual network associated
// with this logical router if xvlan is enabled
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

			project, err := sv.getLogicalRouterParentProject(ctx, logicalRouter, nil, nil)
			if err != nil {
				return err
			}

			vxLanRouting := project.GetVxlanRouting()

			if vxLanRouting {
				var id int64
				id, err = logicalRouter.ConvertVXLanIDToInt()
				if err != nil {
					return err
				}

				err = sv.IntPoolAllocator.DeallocateInt(ctx, VirtualNetworkIDPoolKey, id)
				if err != nil {
					return err
				}

				err = sv.deleteInternalVirtualNetwork(ctx, logicalRouter, project)
				if err != nil {
					return err
				}
			}

			response, err = sv.BaseService.DeleteLogicalRouter(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) allocateVxlanNetworkID(
	ctx context.Context, logicalRouter *models.LogicalRouter, project *models.Project,
) (string, error) {

	dummyVN := &models.VirtualNetwork{FQName: logicalRouter.GetInternalVNFQName(project)}
	intOwner := dummyVN.VxLANIntOwner()

	vxlanNetworkID := logicalRouter.GetVxlanNetworkIdentifier()
	if vxlanNetworkID == "" {
		id, err := sv.IntPoolAllocator.AllocateInt(ctx, VirtualNetworkIDPoolKey, intOwner)
		return strconv.FormatInt(id, 10), err
	}

	id, err := logicalRouter.ConvertVXLanIDToInt()
	if err != nil {
		return "", err
	}

	err = sv.IntPoolAllocator.SetInt(ctx, VirtualNetworkIDPoolKey, id, intOwner)
	if err != nil {
		return "", errutil.ErrorBadRequestf("cannot allocate provided vxlan identifier(%s): %v", vxlanNetworkID, err)
	}

	return vxlanNetworkID, nil
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
	logicalRouter *models.LogicalRouter,
	fm *types.FieldMask,
	vxLanRouting bool,
) error {

	if fm != nil && !basemodels.FieldMaskContains(fm, models.LogicalRouterFieldVirtualNetworkRefs) {
		return nil
	}

	if vxLanRouting {
		for _, vn := range logicalRouter.GetVirtualNetworkRefs() {
			if vn.GetAttr().GetLogicalRouterVirtualNetworkType() != "InternalVirtualNetwork" {
				return errutil.ErrorBadRequest("external gateway not supported with VxLAN")
			}
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) isVxlanRoutingEnabled(ctx context.Context, project *models.Project) bool {
	return project.GetVxlanRouting()
}

func (sv *ContrailTypeLogicService) getLogicalRouterParentProject(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
	dbLogicalRouter *models.LogicalRouter,
	fm *types.FieldMask,
) (*models.Project, error) {

	projectUUID := ""
	if fm != nil && !basemodels.FieldMaskContains(fm, models.LogicalRouterFieldParentUUID) {
		projectUUID = dbLogicalRouter.GetParentUUID()
	} else {
		projectUUID = logicalRouter.GetParentUUID()
	}

	if projectUUID == "" {

		uuid := logicalRouter.GetUUID()
		if uuid == "" {
			return nil, errutil.ErrorBadRequest("no input to derive parent for Logical Router")
		}

		logicalRouterResponse, err := sv.ReadService.GetLogicalRouter(
			ctx,
			&services.GetLogicalRouterRequest{
				ID: uuid,
			},
		)
		if err != nil {
			return nil, err
		}

		projectUUID = logicalRouterResponse.GetLogicalRouter().GetParentUUID()
	}

	projectResponse, err := sv.ReadService.GetProject(ctx, &services.GetProjectRequest{
		ID: projectUUID,
		Fields: []string{
			models.ProjectFieldFQName,
			models.ProjectFieldVxlanRouting,
		},
	})
	return projectResponse.GetProject(), err
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
		if format.ContainsString(interfaceNetworkUUIDs, vn.GetUUID()) {
			return errutil.ErrorBadRequestf(
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
			return errutil.ErrorConflictf("port(%s) already in use by virtual-machine", vmi.GetUUID())
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

	vxlanNetworkID, err := logicalRouter.ConvertVXLanIDToInt()
	if err != nil {
		return nil, err
	}

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

func (sv *ContrailTypeLogicService) updateVxlanID(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
	dbLogicalRouter *models.LogicalRouter,
	fm *types.FieldMask,
	project *models.Project,
) error {

	if !basemodels.FieldMaskContains(fm, models.LogicalRouterFieldVxlanNetworkIdentifier) {
		return nil
	}

	oldVxlanID := dbLogicalRouter.GetVxlanNetworkIdentifier()
	newVxlanID := logicalRouter.GetVxlanNetworkIdentifier()
	if oldVxlanID == newVxlanID {
		return nil
	}

	idToDelete, err := dbLogicalRouter.ConvertVXLanIDToInt()
	if err != nil {
		return err
	}

	err = sv.IntPoolAllocator.DeallocateInt(ctx, VirtualNetworkIDPoolKey, idToDelete)
	if err != nil {
		return err
	}

	logicalRouter.VxlanNetworkIdentifier, err = sv.allocateVxlanNetworkID(ctx, logicalRouter, project)
	return err
}

func (sv *ContrailTypeLogicService) updateInternalVirtualNetwork(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
	dbLogicalRouter *models.LogicalRouter,
	fm *types.FieldMask,
	project *models.Project,
) error {

	logicalRouter.VxlanNetworkIdentifier = logicalRouter.GetVXLanIDInLogicaRouter()
	err := sv.updateVxlanID(ctx, logicalRouter, dbLogicalRouter, fm, project)
	if err != nil {
		return err
	}

	if !basemodels.FieldMaskContains(fm, models.LogicalRouterFieldConfiguredRouteTargetList) &&
		!basemodels.FieldMaskContains(fm, models.LogicalRouterFieldVxlanNetworkIdentifier) {
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

	if basemodels.FieldMaskContains(fm, models.LogicalRouterFieldConfiguredRouteTargetList) {
		updateVN.RouteTargetList = logicalRouter.GetConfiguredRouteTargetList()
		updatePaths = append(updatePaths, models.VirtualNetworkFieldRouteTargetList)
	}

	var vxlanNetworkID int64
	if basemodels.FieldMaskContains(fm, models.LogicalRouterFieldVxlanNetworkIdentifier) {
		vxlanNetworkID, err = logicalRouter.ConvertVXLanIDToInt()
		if err != nil {
			return err
		}

		updateVN.VirtualNetworkProperties = &models.VirtualNetworkType{
			VxlanNetworkIdentifier: vxlanNetworkID,
		}

		updatePaths = append(
			updatePaths,
			basemodels.JoinPath(
				models.VirtualNetworkFieldVirtualNetworkProperties,
				models.VirtualNetworkTypeFieldVxlanNetworkIdentifier,
			),
		)
	}

	_, err = sv.WriteService.UpdateVirtualNetwork(
		ctx, &services.UpdateVirtualNetworkRequest{
			VirtualNetwork: updateVN,
			FieldMask:      types.FieldMask{Paths: updatePaths},
		},
	)

	return err
}

func (sv *ContrailTypeLogicService) deleteInternalVirtualNetwork(
	ctx context.Context,
	logicalRouter *models.LogicalRouter,
	project *models.Project,
) error {

	fqName := logicalRouter.GetInternalVNFQName(project)
	uuid, err := sv.FQNameToUUID(ctx, fqName, models.KindVirtualNetwork)
	if err != nil {
		return errutil.ErrorNotFoundf("cannot find UUID of the internal virtual network: %v", err)
	}

	_, err = sv.WriteService.DeleteLogicalRouterVirtualNetworkRef(
		ctx, &services.DeleteLogicalRouterVirtualNetworkRefRequest{
			ID: logicalRouter.GetUUID(),
			LogicalRouterVirtualNetworkRef: &models.LogicalRouterVirtualNetworkRef{
				UUID: uuid,
			},
		},
	)
	if err != nil {
		return err
	}

	_, err = sv.WriteService.DeleteVirtualNetwork(
		ctx, &services.DeleteVirtualNetworkRequest{
			ID: uuid,
		},
	)
	return err
}

func (sv *ContrailTypeLogicService) checkRouterSupportsVPNType(ctx context.Context, lr *models.LogicalRouter) error {
	bgpvpnRefs := lr.GetBGPVPNRefs()
	if len(bgpvpnRefs) == 0 {
		return nil
	}

	bgpvpnUUIDs := make([]string, 0, len(bgpvpnRefs))
	for _, bgpvpnRef := range bgpvpnRefs {
		bgpvpnUUIDs = append(bgpvpnUUIDs, bgpvpnRef.GetUUID())
	}

	bgpvpns, err := sv.ReadService.ListBGPVPN(
		ctx, &services.ListBGPVPNRequest{Spec: baseservices.SimpleListSpec(bgpvpnUUIDs, models.BGPVPNFieldBGPVPNType)},
	)
	if err != nil {
		return err
	}

	for _, bgpvpn := range bgpvpns.GetBGPVPNs() {
		if bgpvpn.GetBGPVPNType() != models.L3VPNType {
			return errutil.ErrorBadRequestf("can only associate '%s' type BGPVPNs to a logical router", models.L3VPNType)
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) checkRouterHasBGPVPNAssocViaNetwork(
	ctx context.Context,
	lr, dbLR *models.LogicalRouter,
	fm *types.FieldMask,
) error {

	if !isBGPVPNOrVMIChangeRequested(fm) {
		return nil
	}

	if !hasBGPVPNRefs(lr, dbLR, fm) {
		return nil
	}

	vmiRefs := getVMIRefs(lr, dbLR, fm)
	if len(vmiRefs) == 0 {
		return nil
	}

	vnUUIDs, err := sv.getLinkedVnUUIDs(ctx, vmiRefs)
	if err != nil || len(vnUUIDs) == 0 {
		return err
	}

	vnsResp, err := sv.ReadService.ListVirtualNetwork(
		ctx, &services.ListVirtualNetworkRequest{
			Spec: baseservices.SimpleListSpec(vnUUIDs, models.VirtualNetworkFieldBGPVPNRefs),
		},
	)
	if err != nil {
		return err
	}

	for _, vn := range vnsResp.GetVirtualNetworks() {
		if len(vn.GetBGPVPNRefs()) != 0 {
			return errutil.ErrorBadRequestf(
				"can not associate BGPVPN to router which is linked to a network(%s) "+
					"which already has BGPVPN associated",
				vn.GetUUID())
		}
	}

	return nil
}

func isBGPVPNOrVMIChangeRequested(fm *types.FieldMask) bool {
	return fm == nil ||
		basemodels.FieldMaskContains(fm, models.LogicalRouterFieldBGPVPNRefs) ||
		basemodels.FieldMaskContains(fm, models.LogicalRouterFieldVirtualMachineInterfaceRefs)
}

func hasBGPVPNRefs(lr, dbLR *models.LogicalRouter, fm *types.FieldMask) bool {
	if fm == nil || basemodels.FieldMaskContains(fm, models.LogicalRouterFieldBGPVPNRefs) {
		return len(lr.GetBGPVPNRefs()) != 0
	}
	return len(dbLR.GetBGPVPNRefs()) != 0
}

func getVMIRefs(lr, dbLR *models.LogicalRouter, fm *types.FieldMask) []*models.LogicalRouterVirtualMachineInterfaceRef {
	if fm == nil || basemodels.FieldMaskContains(fm, models.LogicalRouterFieldVirtualMachineInterfaceRefs) {
		return lr.GetVirtualMachineInterfaceRefs()
	}
	return dbLR.GetVirtualMachineInterfaceRefs()
}

func (sv *ContrailTypeLogicService) getLinkedVnUUIDs(
	ctx context.Context,
	vmiRefs []*models.LogicalRouterVirtualMachineInterfaceRef,
) ([]string, error) {

	vmiUUIDs := make([]string, 0, len(vmiRefs))
	for _, vmiRef := range vmiRefs {
		vmiUUIDs = append(vmiUUIDs, vmiRef.GetUUID())
	}

	vmisResp, err := sv.ReadService.ListVirtualMachineInterface(
		ctx, &services.ListVirtualMachineInterfaceRequest{
			Spec: baseservices.SimpleListSpec(vmiUUIDs, models.VirtualMachineInterfaceFieldVirtualNetworkRefs),
		},
	)
	if err != nil {
		return nil, err
	}

	var vnUUIDs []string
	for _, vmi := range vmisResp.GetVirtualMachineInterfaces() {
		for _, vnRef := range vmi.GetVirtualNetworkRefs() {
			vnUUIDs = append(vnUUIDs, vnRef.GetUUID())
		}
	}
	return vnUUIDs, nil
}
