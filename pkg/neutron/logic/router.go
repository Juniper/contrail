package logic

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/gogo/protobuf/types"

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
	var portID, subnetTenantID string

	lr, err := r.getLogicalRouter(ctx, rp, id)
	if err != nil {
		return nil, err
	}

	switch {
	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldPortID)):
		portID, subnetTenantID, err = r.addInterfaceWithPort(ctx, rp, id)
		if err != nil {
			return nil, err
		}

	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldSubnetID)):
		portID, subnetTenantID, err = r.addInterfaceWithSubnet(ctx, rp, id)
		if err != nil {
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
		"subnet_id": r.SubnetID,
	}
	return info, nil
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
		portID, subnetID, err = r.deleteInterfaceWithPort(ctx, rp, id)
		if err != nil {
			return nil, err
		}

	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldSubnetID)):
		portID, subnetID, err = r.deleteInterfaceWithSubnet(ctx, rp, id, lr)
		if err != nil {
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

func (r *Router) addInterfaceWithPort(
	ctx context.Context, rp RequestParameters, routerID string,
) (string, string, error) {
	portID := r.PortID

	portRes, err := r.getPortResponse(ctx, rp, portID)
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

	return portID, subnet.TenantID, nil
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

func (r *Router) deleteInterfaceWithPort(
	ctx context.Context, rp RequestParameters, routerID string,
) (string, string, error) {
	portID := r.PortID
	portRes, err := r.getPortResponse(ctx, rp, portID)
	if err != nil {
		return "", "", err
	}
	if portRes.DeviceOwner != deviceOwnerRouterInterface || portRes.DeviceID != r.ID {
		return "", "", newRouterError(badRequest, "Router interface not found")
	}
	subnetID := portRes.FixedIps[0].SubnetID
	if r.SubnetID != "" && subnetID != r.SubnetID {
		return "", "", newRouterError(subnetMismatchForPort, "Router interface not found")
	}

	return portID, subnetID, nil
}

func (r *Router) deleteInterfaceWithSubnet(
	ctx context.Context, rp RequestParameters, routerID string, lr *models.LogicalRouter,
) (string, string, error) {
	var portID, subnetID string
	subnetID = r.SubnetID

	for _, intf := range lr.GetVirtualMachineInterfaceRefs() {
		portID = intf.UUID
		portRes, err := r.getPortResponse(ctx, rp, portID)
		if err != nil {
			return "", "", err
		}
		if subnetID == portRes.FixedIps[0].SubnetID {
			break
		}
	}
	if portID == "" {
		return "", "", newRouterError(badRequest, "Subnet %s is not connected to router %s", subnetID, routerID)
	}
	return portID, subnetID, nil
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
			subID := ip.SubnetID
			subnetRes, err := r.getSubnetResponse(ctx, rp, subID)
			if err != nil {
				return err
			}

			if err = doesOverlap(subnetCIDR, subnetRes.Cidr); err != nil {
				return nil
			}

		}
	}
	return nil
}

func doesOverlap(cidr1, cidr2 string) error {
	ip1, sub1, err := net.ParseCIDR(cidr1)
	if err != nil { return err }

	ip2, sub2, err := net.ParseCIDR(cidr2)
	if err != nil {
		return err
	}

	if sub1.Contains(ip2) || sub2.Contains(ip1) {
		return newRouterError(badRequest,
			"Cidr %s of subnet %s overlaps with cidr %s of subnet %s",
			cidr1, cidr2,
		)
	}

	return nil
}

func (r *Router) updateLRForAddInterface(
	ctx context.Context, rp RequestParameters, lr *models.LogicalRouter, portID string,
) error {
	vmi, err := r.readVMIFromVnc(ctx, rp, portID)
	if err != nil {
		return err
	}

	vmi.VirtualMachineInterfaceDeviceOwner = deviceOwnerRouterInterface
	_, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
		FieldMask: types.FieldMask{Paths: []string{
			models.VirtualMachineInterfaceFieldVirtualMachineInterfaceDeviceOwner,
		}},
	})
	if err != nil {
		return err
	}

	lr.AddVirtualMachineInterfaceRef(&models.LogicalRouterVirtualMachineInterfaceRef{
		To:   vmi.FQName,
		UUID: vmi.UUID,
		Href: vmi.Href,
	})

	_, err = rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{Paths: []string{
			models.LogicalRouterFieldVirtualMachineInterfaceRefs,
		}},
	})

	return err
}

func (r *Router) updateLRForDelInterface(
	ctx context.Context, rp RequestParameters, lr *models.LogicalRouter, portID string,
) error {
	vmi, err := r.readVMIFromVnc(ctx, rp, portID)
	if err != nil {
		return err
	}

	lr.RemoveVirtualMachineInterfaceRef(&models.LogicalRouterVirtualMachineInterfaceRef{
		To:   vmi.FQName,
		UUID: vmi.UUID,
		Href: vmi.Href,
	})

	_, err = rp.WriteService.UpdateLogicalRouter(ctx, &services.UpdateLogicalRouterRequest{
		LogicalRouter: lr,
		FieldMask: types.FieldMask{Paths: []string{
			models.LogicalRouterFieldVirtualMachineInterfaceRefs,
		}},
	})
	if err != nil {
		return err
	}

	_, err = (&Port{}).Delete(ctx, rp, portID)
	return err
}

func (r *Router) getLogicalRouter(
	ctx context.Context, rp RequestParameters, routerID string,
) (*models.LogicalRouter, error) {
	lrRes, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: routerID,
	})
	if err != nil {
		return nil, err
	}
	return lrRes.GetLogicalRouter(), nil
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
