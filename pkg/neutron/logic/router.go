package logic

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	deviceOwnerRouterInterface = "network:router_interface"
)

func newRouterError(name errorType, format string, args ...interface{}) error {
	return newNeutronError(name, errorFields{
		"resource": "router",
		"msg":      fmt.Sprintf(format, args...),
	})
}

func newRouterNotFoundError(id string) *Error {
	return newNeutronError(routerNotFound, errorFields{
		"router_id": id,
	})
}

func newRouterInUseError(id string) *Error {
	return newNeutronError(routerInUse, errorFields{
		"router_id": id,
	})
}

// ReadAll logic
func (r *Router) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement ReadAll logic
	return []RouterResponse{}, nil
}

// Create logic
func (r *Router) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	lr, err := r.neutronToVnc(rp.RequestContext)
	if err != nil {
		return nil, err
	}

	lrResponse, err := rp.WriteService.CreateLogicalRouter(ctx, &services.CreateLogicalRouterRequest{
		LogicalRouter: lr,
	})
	if err != nil {
		return nil, err
	}

	return r.vncToNeutron(lrResponse.GetLogicalRouter()), nil
}

// Read logic
func (r *Router) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	// TODO: If fields == ["tenant_id"], return {"id": id, "tenant_id": None}

	lrResponse, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: id,
	})
	if errutil.IsNotFound(err) {
		return nil, newRouterNotFoundError(id)
	} else if err != nil {
		return nil, err
	}

	return r.vncToNeutron(lrResponse.GetLogicalRouter()), nil
}

// Update logic
func (r *Router) Update(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	lr, err := r.neutronToVnc(rp.RequestContext)
	if err != nil {
		return nil, err
	}
	lr.UUID = id

	// TODO: Check the referred VN's RouterExternal.

	_, err = rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{
			Paths: []string{models.LogicalRouterFieldVirtualNetworkRefs},
			// TODO: Update other fields.
		},
	})
	if errutil.IsNotFound(err) {
		return nil, newRouterNotFoundError(id)
	} else if err != nil {
		return nil, err
	}

	return r.Read(ctx, rp, id)
}

// Delete logic
func (r *Router) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	// TODO: Check VMI refs.

	if _, err := rp.WriteService.DeleteLogicalRouter(ctx, &services.DeleteLogicalRouterRequest{
		ID: id,
	}); errutil.IsNotFound(err) {
		return nil, newRouterNotFoundError(id)
	} else if errutil.IsConflict(err) {
		return nil, newRouterInUseError(id)
	} else if err != nil {
		return nil, err
	}

	return &RouterResponse{}, nil
}

func (r *Router) neutronToVnc(rc RequestContext) (*models.LogicalRouter, error) {
	projectUUID, err := neutronIDToVncUUID(rc.TenantID)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid tenant id: %s", rc.TenantID)
	}
	return &models.LogicalRouter{
		Name:        r.Name,
		DisplayName: r.Name,
		UUID:        r.ID,
		ParentUUID:  projectUUID,
		ParentType:  models.KindProject,
		IDPerms: &models.IdPermsType{
			Enable:      r.AdminStateUp,
			Description: r.Description,
		},
		VirtualNetworkRefs: r.makeVNRefs(),
	}, nil
}

func (r *Router) vncToNeutron(lr *models.LogicalRouter) *RouterResponse {
	response := &RouterResponse{
		ID:                  lr.GetUUID(),
		TenantID:            VncUUIDToNeutronID(lr.GetParentUUID()),
		AdminStateUp:        lr.GetIDPerms().GetEnable(),
		Shared:              false,
		Status:              netStatusActive,
		GWPortID:            "",
		ExternalGatewayInfo: r.makeExternalGatewayInfo(lr),
		Description:         lr.GetIDPerms().GetDescription(),
		CreatedAt:           lr.GetIDPerms().GetCreated(),
		UpdatedAt:           lr.GetIDPerms().GetLastModified(),
	}

	response.Name = lr.GetDisplayName()
	if response.Name == "" {
		response.Name = lr.GetName()
	}

	if contrailExtensionsEnabled {
		response.FQName = lr.GetFQName()
	}

	return response
}

func (r *Router) makeExternalGatewayInfo(lr *models.LogicalRouter) ExtGatewayInfo {
	vnRefs := lr.GetVirtualNetworkRefs()
	if len(vnRefs) == 0 {
		return ExtGatewayInfo{}
	}
	vnUUID := vnRefs[0].GetUUID()
	if vnUUID == "" {
		return ExtGatewayInfo{}
	}

	return ExtGatewayInfo{
		NetworkID:  vnUUID,
		EnableSnat: true,
	}
}

func (r *Router) makeVNRefs() (refs []*models.LogicalRouterVirtualNetworkRef) {
	// TODO: Make r.ExternalGatewayInfo a pointer and check if it is nil instead.
	if r.ExternalGatewayInfo.NetworkID != "" {
		refs = append(refs, &models.LogicalRouterVirtualNetworkRef{
			UUID: r.ExternalGatewayInfo.NetworkID,
		})
	}

	return refs
}

// AddInterface logic
func (r *Router) AddInterface(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	var portID, subnetTenantID, subnetID string

	lr, err := r.getLogicalRouter(ctx, rp, id)
	if err != nil {
		return nil, err
	}

	switch {
	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldPortID)):
		portID = r.PortID
		if subnetID, subnetTenantID, err = r.addInterfaceWithPort(ctx, rp, id); err != nil {
			return nil, err
		}

	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldSubnetID)):
		subnetID = r.SubnetID
		if portID, subnetTenantID, err = r.addInterfaceWithSubnet(ctx, rp, id); err != nil {
			return nil, err
		}

	default:
		return nil, newRouterError(badRequest, "Either port or subnet must be specified")
	}

	if err := r.updateLRForAddInterface(ctx, rp, lr, portID); err != nil {
		return nil, err
	}

	info := map[string]string{
		"id":        id,
		"tenant_id": subnetTenantID,
		"port_id":   portID,
		"subnet_id": subnetID,
	}
	return info, nil
}

func (r *Router) getLogicalRouter(
	ctx context.Context, rp RequestParameters, routerID string,
) (*models.LogicalRouter, error) {
	lrRes, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: routerID,
	})
	return lrRes.GetLogicalRouter(), err
}

func (r *Router) addInterfaceWithPort(
	ctx context.Context, rp RequestParameters, routerID string,
) (string, string, error) {
	portRes, err := r.getPortResponse(ctx, rp, r.PortID)
	if err != nil {
		return "", "", err
	}

	if err = r.validatePort(ctx, rp, portRes); err != nil {
		return "", "", err
	}

	subnet, err := r.getSubnetResponse(ctx, rp, portRes.FixedIps[0].SubnetID)
	if err != nil {
		return "", "", err
	}
	if err := r.checkForDupRouterSubnet(ctx, rp, routerID, subnet.NetworkID, subnet.ID, subnet.Cidr); err != nil {
		return "", "", err
	}

	return subnet.ID, subnet.TenantID, nil
}

func (r *Router) getPortResponse(ctx context.Context, rp RequestParameters, portID string) (*PortResponse, error) {
	portResNeutron, err := (&Port{}).Read(ctx, rp, portID)
	if err != nil {
		return nil, err
	}
	portRes, ok := portResNeutron.(*PortResponse)
	if !ok {
		return nil, newRouterError(internalServerError, "Cannot convert interface to port response structure")
	}
	return portRes, nil
}

func (r *Router) validatePort(ctx context.Context, rp RequestParameters, port *PortResponse) error {
	if !rp.RequestContext.IsAdmin && VncUUIDToNeutronID(rp.RequestContext.TenantID) != port.TenantID {
		return newRouterError(routerInterfaceNotFound, "Router interface not found")
	}

	if port.DeviceOwner == deviceOwnerRouterInterface && port.DeviceID != "" {
		return newRouterError(l3PortInUse, "network_id %s, port_id %s, device_id %s",
			port.NetworkID, port.ID, port.DeviceID,
		)
	}
	if len(port.FixedIps) != 1 {
		return newRouterError(badRequest, "Router port must have exactly one fixed IP")
	}

	return nil
}

func (r *Router) getSubnetResponse(
	ctx context.Context, rp RequestParameters, subnetID string,
) (*SubnetResponse, error) {
	subnetResNeutron, err := (&Subnet{}).Read(ctx, rp, subnetID)
	if err != nil {
		return nil, err
	}
	subnetRes, ok := subnetResNeutron.(*SubnetResponse)
	if !ok {
		return nil, newRouterError(internalServerError, "Cannot convert interface to subnet response structure")
	}
	return subnetRes, nil
}

func (r *Router) checkForDupRouterSubnet(
	ctx context.Context, rp RequestParameters, routerID, networkID, subnetID, subnetCIDR string,
) error {
	rportsRes, err := (&Port{}).ReadAll(ctx, rp, Filters{deviceIDKey: []string{routerID}}, Fields{})
	if err != nil {
		return err
	}

	rports, ok := rportsRes.([]*PortResponse)
	if !ok {
		return newRouterError(internalServerError, "Cannot convert interface to port response structure")
	}

	for _, p := range rports {
		for _, ip := range p.FixedIps {
			if ip.SubnetID == subnetID {
				return newRouterError(badRequest, "Router %s already has a port on subnet %s", routerID, subnetID)
			}
			subnetRes, err := r.getSubnetResponse(ctx, rp, ip.SubnetID)
			if err != nil {
				return err
			}

			if overlap, err := models.CheckSubnetsOverlap(subnetCIDR, subnetRes.Cidr); err != nil {
				return err
			} else if overlap == true {
				return newRouterError(badRequest,
					"Cidr %s of subnet %s overlaps with cidr %s of subnet %s",
					subnetCIDR, subnetID, subnetRes.Cidr, subnetRes.ID,
				)
			}
		}
	}
	return nil
}

func (r *Router) addInterfaceWithSubnet(
	ctx context.Context, rp RequestParameters, routerID string,
) (string, string, error) {
	subnet, err := r.getSubnetResponse(ctx, rp, r.SubnetID)
	if err != nil {
		return "", "", err
	}

	if err = r.validateSubnet(ctx, rp, subnet); err != nil {
		return "", "", err
	}

	if err = r.checkForDupRouterSubnet(ctx, rp, routerID, subnet.NetworkID, subnet.ID, subnet.Cidr); err != nil {
		return "", "", err
	}

	portID, err := r.createPort(ctx, rp, routerID, subnet)
	if err != nil {
		return "", "", err
	}

	return portID, subnet.TenantID, nil
}

func (r *Router) validateSubnet(ctx context.Context, rp RequestParameters, subnet *SubnetResponse) error {
	if !rp.RequestContext.IsAdmin && VncUUIDToNeutronID(rp.RequestContext.TenantID) != subnet.TenantID {
		return newRouterError(routerInterfaceNotFoundForSubnet, "Router interface not found for subnet %s",
			r.SubnetID,
		)
	}
	if subnet.GatewayIP == "" {
		return newRouterError(badRequest, "Subnet four router interface must have a gateway IP")
	}
	return nil
}

func (r *Router) createPort(
	ctx context.Context, rp RequestParameters, deviceID string, subnet *SubnetResponse,
) (string, error) {
	fixedIps := []*FixedIp{
		{
			SubnetID:  subnet.ID,
			IPAddress: subnet.GatewayIP,
		},
	}
	portObj := &Port{
		TenantID:            subnet.TenantID,
		NetworkID:           subnet.NetworkID,
		FixedIps:            fixedIps,
		AdminStateUp:        true,
		DeviceID:            deviceID,
		DeviceOwner:         deviceOwnerRouterInterface,
		Name:                "",
		PortSecurityEnabled: false,
	}
	portRes, err := portObj.Create(ctx, rp)
	if err != nil {
		return "", err
	}
	port, ok := portRes.(*PortResponse)
	if !ok {
		return "", newRouterError(internalServerError, "Cannot convert response to port structure")
	}

	return port.ID, nil
}

func (r *Router) updateLRForAddInterface(
	ctx context.Context, rp RequestParameters, lr *models.LogicalRouter, portID string,
) error {
	_, err := rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: &models.VirtualMachineInterface{
			UUID:                               portID,
			VirtualMachineInterfaceDeviceOwner: deviceOwnerRouterInterface,
		},
		FieldMask: types.FieldMask{Paths: []string{
			models.VirtualMachineInterfaceFieldVirtualMachineInterfaceDeviceOwner,
		}},
	})
	if err != nil {
		return err
	}

	lr.AddVirtualMachineInterfaceRef(&models.LogicalRouterVirtualMachineInterfaceRef{
		UUID: portID,
	})

	_, err = rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{Paths: []string{
			models.LogicalRouterFieldVirtualMachineInterfaceRefs,
		}},
	})

	return err
}

func (r *Router) readVMIFromVnc(
	ctx context.Context, rp RequestParameters, id string,
) (*models.VirtualMachineInterface, error) {
	vmiRes, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return vmiRes.GetVirtualMachineInterface(), nil
}

// DeleteInterface logic
func (r *Router) DeleteInterface(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	var portID, subnetID string

	lr, err := r.getLogicalRouter(ctx, rp, id)
	if err != nil {
		return nil, err
	}

	switch {
	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldPortID)):
		if portID, subnetID, err = r.deleteInterfaceWithPort(ctx, rp, id); err != nil {
			return nil, err
		}

	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldSubnetID)):
		if portID, subnetID, err = r.deleteInterfaceWithSubnet(ctx, rp, id, lr); err != nil {
			return nil, err
		}

	default:
		return nil, newRouterError(badRequest, "Either port or subnet must be specified")
	}

	if err = r.updateLRForDelInterface(ctx, rp, lr, portID); err != nil {
		return nil, err
	}

	subnetRes, err := r.getSubnetResponse(ctx, rp, subnetID)
	if err != nil {
		return nil, err
	}

	info := map[string]string{
		"id":        id,
		"tenant_id": subnetRes.TenantID,
		"port_id":   portID,
		"subnet_id": subnetID,
	}
	return info, nil
}

func (r *Router) deleteInterfaceWithPort(
	ctx context.Context, rp RequestParameters, routerID string,
) (string, string, error) {
	portRes, err := r.getPortResponse(ctx, rp, r.PortID)
	if err != nil {
		return "", "", err
	}
	if portRes.DeviceOwner != deviceOwnerRouterInterface || portRes.DeviceID != routerID {
		return "", "", newRouterError(badRequest, "Router interface not found")
	}
	subnetID := portRes.FixedIps[0].SubnetID
	if r.SubnetID != "" && subnetID != r.SubnetID {
		return "", "", newRouterError(subnetMismatchForPort, "Router interface not found")
	}

	return r.PortID, subnetID, nil
}

func (r *Router) deleteInterfaceWithSubnet(
	ctx context.Context, rp RequestParameters, routerID string, lr *models.LogicalRouter,
) (string, string, error) {
	for _, intf := range lr.GetVirtualMachineInterfaceRefs() {
		portRes, err := r.getPortResponse(ctx, rp, intf.UUID)
		if err != nil {
			return "", "", err
		}
		if r.SubnetID == portRes.FixedIps[0].SubnetID {
			return intf.UUID, r.SubnetID, nil
		}
	}
	return "", "", newRouterError(badRequest, "Subnet %s is not connected to router %s", r.SubnetID, routerID)

}

func (r *Router) updateLRForDelInterface(
	ctx context.Context, rp RequestParameters, lr *models.LogicalRouter, portID string,
) error {

	lr.RemoveVirtualMachineInterfaceRefWithID(portID)

	_, err := rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{Paths: []string{
			models.LogicalRouterFieldVirtualMachineInterfaceRefs,
		}},
	})
	if err != nil {
		return errors.Wrapf(err, "Cannot update logical router %s", lr.UUID)
	}

	_, err = (&Port{}).Delete(ctx, rp, portID)
	return err
}
