package logic

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"net"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
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
	lr, err := r.getLogicalRouter(ctx, rp, id)
	if err != nil {
		return nil, err
	}

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
		FieldMask: types.FieldMask{Paths:[]string {
			models.VirtualMachineInterfaceFieldVirtualMachineInterfaceDeviceOwner,
		}},
	})
	if err != nil {
		return nil, err
	}

	lr.AddVirtualMachineInterfaceRef(&models.LogicalRouterVirtualMachineInterfaceRef{
		To: vmi.FQName,
		UUID: vmi.UUID,
		Href: vmi.Href,
	})

	_, err = rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{Paths:[]string {
			models.LogicalRouterFieldVirtualMachineInterfaceRefs,
		}},
	})
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
		if portID == "" {
			return nil, newRouterError(badRequest, "Subnet %s is not connected to router %s", subnetID, id)
		}

	default:
		return nil, newRouterError(badRequest, "Either port or subnet must be specified")
	}

	vmi, err := r.readVMIFromVnc(ctx, rp, portID)
	if err != nil {
		return nil, err
	}

	lr.RemoveVirtualMachineInterfaceRef(&models.LogicalRouterVirtualMachineInterfaceRef{
		To: vmi.FQName,
		UUID: vmi.UUID,
		Href: vmi.Href,
	})

	_, err = rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{Paths:[]string {
			models.LogicalRouterFieldVirtualMachineInterfaceRefs,
		}},
	})
	if err != nil {
		return nil, err
	}

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
		Name:         lr.GetDisplayName(),
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
	// TODO check error
	rportsRes, _ := (&Port{}).ReadAll(ctx, rp, Filters{deviceIDKey: []string{routerID}}, Fields{})
	rports := rportsRes.([]*PortResponse)

	newIP, newSub, err := net.ParseCIDR(subnetCIDR)
	if err != nil{
		return err
	}

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
			ip2, sub2, err := net.ParseCIDR(subnetRes.Cidr)
			if newSub.Contains(ip2) || sub2.Contains(newIP) {
				return newRouterError(badRequest, "Cidr %s of subnet %s overlaps with cidr %s of subnet %s", subnetCIDR, subnetID, subnetRes.Cidr, subnetRes.ID)
			}
		}
	}
	return nil
}

func (r *Router) getLogicalRouter(ctx context.Context, rp RequestParameters, routerID string) (*models.LogicalRouter, error) {
	lrRes, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: routerID,
	})
	if err != nil {
		return nil, err
	}
	return lrRes.GetLogicalRouter(), nil
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

func sanitizeID(id string) string {
	return strings.Replace(id, "-", "", -1)
}
