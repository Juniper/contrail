package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	deviceOwnerRouterInterface = "network:router_interface"

	testPortID = "f7274d2b-bd18-438f-8bed-076d3de2f825"
)

func newRouterError(name errorType, format string, args ...interface{}) error {
	return newNeutronError(name, errorFields{
		"resource": "router",
		"msg":      fmt.Sprintf(format, args...),
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
	lr, err := r.neutronToVnc(ctx, rp)
	if err != nil {
		return nil, err
	}

	lrResponse, err := rp.WriteService.CreateLogicalRouter(ctx, &services.CreateLogicalRouterRequest{
		LogicalRouter: lr,
	})
	if err != nil {
		// TODO Wrap.
		return nil, err
	}

	// TODO: Update gateway.

	// TODO Wrap err.
	return r.vncToNeutron(lrResponse.GetLogicalRouter())
}

// Delete logic
func (r *Router) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	// TODO: Check VMI refs.

	if _, err := rp.WriteService.DeleteLogicalRouter(ctx, &services.DeleteLogicalRouterRequest{
		ID: id,
	}); err != nil {
		// TODO Wrap.
		return nil, err
	}

	return &RouterResponse{}, nil
}

// AddInterface logic
func (r *Router) AddInterface(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	//lrRes, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
	//	ID: id,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//lr := lrRes.GetLogicalRouter()

	/*
	check if port_id
		invoke _check_for_dup_router_subnet
	check if subnet id
		invoke _check_for_dup_router_subnet
	update logical router
	*/
	var portID, subnetTenantID string
	switch {
	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldPortID)):
		portID = r.PortID

		portRes, err := r.getPortResponse(ctx, rp, portID)
		if err != nil {
			return nil, err
		}

		if err = r.validatePort(ctx, rp, portRes); err != nil {
			return nil, err
		}
		subnetID := portRes.FixedIps[0].SubnetID
		subnetRes, err := r.getSubnetResponse(ctx, rp, subnetID)
		if err != nil {
			return nil, err
		}
		_ = r.checkForDupRouterSubnet(ctx, rp, id, subnetRes.NetworkID, subnetRes.ID, subnetRes.Cidr)
		subnetTenantID = subnetRes.TenantID

	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldSubnetID)):
		subnet, err := r.getSubnetResponse(ctx, rp, r.SubnetID)
		if err != nil {
			return nil, err
		}
		subnetTenantID = subnet.TenantID

		if err = r.validateSubnet(ctx, rp, subnet); err != nil {
			return nil, err
		}

		_ = r.checkForDupRouterSubnet(ctx, rp, id, subnet.NetworkID, subnet.ID, subnet.Cidr)

		portID, err = r.createPort(ctx, rp, id, subnet)
		if err != nil {
			return nil, err
		}

	default:
		return nil, newRouterError(badRequest, "Either port or subnet must be specified")
	}

	vmi, err := r.readVMIFromVnc(ctx, rp, portID)
	if err != nil {
		return nil, err
	}

	vmi.VirtualMachineInterfaceDeviceOwner = deviceOwnerRouterInterface
	_, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})
	if err != nil {
		return nil, err
	}

	//lr.AddVirtualMachineInterfaceRef(&models.LogicalRouterVirtualMachineInterfaceRef{
	//	To: vmi.FQName,
	//	UUID: vmi.UUID,
	//	Href: vmi.Href,
	//})

	// TODO update lr
	if err != nil {
		return nil, err
	}

	info := map[string]string{
		"id":        id,
		"tenant_id": subnetTenantID,
		"port_id":   portID,
		"subnet_id": r.SubnetID,
	}
	return info, nil
}

// DeleteInterface logic
func (r *Router) DeleteInterface(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	lrRes, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	lr := lrRes.GetLogicalRouter()

	var portID, subnetID string
	switch {
	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldPortID)):
		portID = r.PortID
		portRes, err := r.getPortResponse(ctx, rp, portID)
		if err != nil {
			return nil, err
		}
		if portRes.DeviceOwner != deviceOwnerRouterInterface || portRes.DeviceID != id {
			return nil, newRouterError(badRequest, "Router interface not found")
		}
		subnetID = portRes.FixedIps[0].SubnetID
		if r.SubnetID != "" && subnetID != r.SubnetID {
			return nil, newRouterError(subnetMismatchForPort, "Router interface not found")
		}

	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldSubnetID)):
		subnetID = r.SubnetID

		for _, intf := range lr.GetVirtualMachineInterfaceRefs() {
			portID = intf.UUID
			portRes, err := r.getPortResponse(ctx, rp, portID)
			if err != nil {
				return nil, err
			}
			if subnetID == portRes.FixedIps[0].SubnetID {
				break
			}
		}
		// TODO check if subnet is connected to router

	default:
		return nil, newRouterError(badRequest, "Either port or subnet must be specified")
	}

	//vmi, err := r.readVMIFromVnc(ctx, rp, portID)
	//if err != nil {
	//	return nil, err
	//}

	// TODO delete vmi from lr

	if _, err := (&Port{}).Delete(ctx, rp, portID); err != nil {
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
	// TODO implement DeleteInteface logic
	return nil, newNeutronError(badRequest, errorFields{
		"msg": "DeleteInterface is not implemented yet",
	})

}

func (r *Router) neutronToVnc(ctx context.Context, rp RequestParameters) (*models.LogicalRouter, error) {
	projectUUID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
	if err != nil {
		// TODO Wrap.
		return nil, err
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
	}, nil
}

func (r *Router) vncToNeutron(lr *models.LogicalRouter) (*RouterResponse, error) {
	response := &RouterResponse{
		ID:           lr.GetUUID(),
		Name:         lr.GetDisplayName(), // TODO or Name if it's empty.
		TenantID:     VncUUIDToNeutronID(lr.GetParentUUID()),
		AdminStateUp: lr.GetIDPerms().GetEnable(),
		Shared:       false,
		Status:       netStatusActive,
		GWPortID:     "",
		// TODO ExternalGatewayInfo.
		Description: lr.GetIDPerms().GetDescription(),
		CreatedAt:   lr.GetIDPerms().GetCreated(),
		UpdatedAt:   lr.GetIDPerms().GetLastModified(),
	}

	if contrailExtensionsEnabled {
		response.FQName = lr.GetFQName()
	}

	return response, nil
}

func (r *Router) checkForDupRouterSubnet(
	ctx context.Context, rp RequestParameters, routerID, networkID, subnetID, subnetCIDR string,
) error {
	return nil
	// TODO
	rportsRes, _ := (&Port{}).ReadAll(ctx, rp, Filters{deviceIDKey: []string{routerID}}, Fields{})
	rports := rportsRes.([]*PortResponse)
	//new_ipnet := netaddr.IPNetwork(subnetCIDR)

	for _, p := range rports {
		for _, ip := range p.FixedIps {
			if ip.SubnetID == subnetID {
				return newRouterError(badRequest, "Router %s already has a port on subnet %s", routerID, subnetID)
			}
			subID := ip.SubnetID
			subnetRes, err := r.getSubnetResponse(ctx, rp, subID)
			if err != nil {
				return err
			}
			//match1 = netaddr.all_matching_cidrs(new_ipnet, [cidr])
			//match2 = netaddr.all_matching_cidrs(ipnet, [subnet_cidr])
			cidr := subnetRes.Cidr
			_ = cidr

		}
	}
	// TODO implement function\
	return nil
}

func (r *Router) getSubnetResponse(ctx context.Context, rp RequestParameters, subnetID string) (*SubnetResponse, error) {
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

func (r *Router) validateSubnet(ctx context.Context, rp RequestParameters, subnet *SubnetResponse) error {
	if !rp.RequestContext.IsAdmin && sanitizeID(rp.RequestContext.TenantID) != subnet.TenantID {
		return newRouterError(routerInterfaceNotFoundForSubnet, "Router interface not found for subnet %s",
			r.SubnetID,
		)
	}
	if subnet.GatewayIP == "" {
		return newRouterError(badRequest, "Subnet four router interface must have a gateway IP")
	}
	return nil
}

func (r *Router) validatePort(ctx context.Context, rp RequestParameters, port *PortResponse) error {
	if !rp.RequestContext.IsAdmin && sanitizeID(rp.RequestContext.TenantID) != port.TenantID {
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

func (r *Router) readVMIFromVnc(ctx context.Context, rp RequestParameters, id string, ) (*models.VirtualMachineInterface, error) {
	vmiRes, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return vmiRes.GetVirtualMachineInterface(), nil
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

func (r *Router) createPort(
	ctx context.Context, rp RequestParameters, deviceId string, subnet *SubnetResponse,
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
		DeviceID:            deviceId,
		DeviceOwner:         deviceOwnerRouterInterface,
		Name:                "",
		PortSecurityEnabled: false,
	}
	portRes, err := (portObj).Create(ctx, rp)
	if err != nil {
		return "", err
	}
	port, ok := portRes.(*PortResponse)
	if !ok {
		return "", newRouterError(internalServerError, "Cannot convert response to port structure")
	}

	return port.ID, nil
}

func (r *Router) updateVMI(ctx context.Context, rp RequestParameters, id string, ) error {
	vmi, err := r.readVMIFromVnc(ctx, rp, id)
	if err != nil {
		return err
	}
	vmi.VirtualMachineInterfaceDeviceOwner = deviceOwnerRouterInterface

	_, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})

	return err
}

func sanitizeID(id string) string {
	return strings.Replace(id, "-", "", -1)
}
