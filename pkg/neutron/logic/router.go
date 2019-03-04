package logic

import (
	"context"
	"fmt"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
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

// AddInterface logic
func (r *Router) AddInterface(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	lrRes, err := rp.ReadService.GetLogicalRouter(ctx, &services.GetLogicalRouterRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	lr := lrRes.GetLogicalRouter()

	/*
	check if port_id
		read port
		check is admin
		reformat tenant_id from kebab
		check if port is used
		check if port has one fixed ip
		invoke _check_for_dup_router_subnet
		
	check if subnet id 
		read subnet
		check is admin
		check if subnet has gateway ip
		invoke _check_for_dup_router_subnet
		create port

	add vmi to router
	update logical router
	*/
	var portID, subnetTenantID string
	switch {
	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldPortID)):
		// TODO port_id logic
		return nil, newRouterError(badRequest, "Add interface with port id not implemented yet")
		portID = r.PortID
	case basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(RouterFieldSubnetID)):
		subnet, err := r.getSubnetResponse(ctx, rp)
		if err != nil {
			return nil, err
		}
		subnetTenantID = subnet.TenantID

		if err = r.validateSubnet(ctx, rp, subnet); err != nil{
			return nil, err
		}

		checkForDupRouterSubnet()

		portID, err = r.createPort(ctx, rp, id, subnet)
		if err != nil {
			return nil, err
		}

	default:
		return nil, newRouterError(badRequest, "Either port or subnet must be specified")
	}

	if err = updateVMI(ctx, rp, portID); err != nil {
		return nil, err
	}

	// TODO update lr

	//lr.AddVirtualMachineInterfaceRef(vmi.ref)
	_ = lr // silencing an error

	if err != nil {
		return nil, err
	}

	info := map[string]string{
		"id": id,
		"tenant_id": subnetTenantID,
		"port_id": portID,
		"subnet_id": r.SubnetID,
	}
	return info, nil
}

// DeleteInterface logic
func (r *Router) DeleteInterface(
	ctx context.Context, rp RequestParameters,
) (Response, error) {
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

func checkForDupRouterSubnet(){
	// TODO implement function
}

func (r *Router) getSubnetResponse(ctx context.Context, rp RequestParameters) (*SubnetResponse, error){
	subnetResNeutron, err := (&Subnet{}).Read(ctx, rp, r.SubnetID)
	if err != nil {
		return nil, err
	}
	subnetRes, ok := subnetResNeutron.(*SubnetResponse)
	if !ok {
		return nil, newRouterError(internalServerError, "Cannot convert interface to subnet response structure")
	}
	return subnetRes, nil
}

func (r *Router) validateSubnet(ctx context.Context, rp RequestParameters, subnet *SubnetResponse) error{
	if !rp.RequestContext.IsAdmin && sanitizeID(rp.RequestContext.TenantID) != subnet.TenantID {
		return newRouterError(badRequest, "Router interface not found for subnet %s",
			r.SubnetID,
		)
	}
	if subnet.GatewayIP == "" {
		return newRouterError(badRequest, "Subnet four router interface must have a gateway IP")
	}
	return nil
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
		TenantID: subnet.TenantID,
		NetworkID: subnet.NetworkID,
		FixedIps: fixedIps,
		AdminStateUp: true,
		DeviceID: deviceId,
		DeviceOwner: deviceOwnerRouterInterface,
		Name: "",
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

func updateVMI(ctx context.Context, rp RequestParameters, id string,) (error){
	vmiRes, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: id,
	})
	if err != nil {
		return err
	}
	vmi := vmiRes.GetVirtualMachineInterface()
	vmi.VirtualMachineInterfaceDeviceOwner = deviceOwnerRouterInterface

	_, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})

	return err

}

func sanitizeID(id string) string{
	return strings.Replace(id, "-", "", -1)
}
