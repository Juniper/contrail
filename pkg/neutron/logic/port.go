package logic

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// Create handles port create request.
func (port *Port) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	if port.ID == "" {
		port.ID = uuid.NewV4().String()
	}

	if err := port.checkIfMacAddressIsAvailable(ctx, rp); err != nil {
		return nil, err
	}

	vn, err := port.getVirtualNetwork(ctx, rp)
	if err != nil {
		return nil, err
	}

	vmi, err := port.createVirtualMachineInterface(ctx, rp, vn)
	if err != nil {
		return nil, err
	}

	iip, err := port.allocateIPAddress(ctx, rp, vn, vmi)
	if err != nil {
		return nil, err
	}

	// TODO:
	// create interface route table for the port if
	// subnet has a host route for this port ip.

	return makePortResponse(vn, vmi, []*models.InstanceIP{iip}), nil
}

// Update handles port update request.
func (port *Port) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vmi, vn, err := port.readVNCPort(ctx, rp, id)
	if err != nil {
		return nil, newNeutronError(portNotFound, errorFields{
			"port_id": id,
		})
	}

	fm := types.FieldMask{}

	if port.Name != "" {
		vmi.DisplayName = port.Name
		basemodels.FieldMaskAppend(&fm, models.VirtualMachineInterfaceFieldDisplayName)
	}

	if port.DeviceOwner != "network:router_interface" &&
		port.DeviceOwner != "network:router_gateway" && port.DeviceID != "" {
		if err = port.setVMInstance(ctx, rp, vmi); err != nil {
			return nil, err
		}
		basemodels.FieldMaskAppend(&fm, models.VirtualMachineInterfaceFieldVirtualMachineRefs)
	}

	if port.DeviceOwner != "" {
		vmi.VirtualMachineInterfaceDeviceOwner = port.DeviceOwner
		basemodels.FieldMaskAppend(&fm, models.VirtualMachineInterfaceFieldVirtualMachineInterfaceDeviceOwner)
	}

	port.setBindings(vmi)
	basemodels.FieldMaskAppend(&fm, models.VirtualMachineInterfaceFieldVirtualMachineInterfaceBindings)
	//TODO handle mac address change
	//TODO port security enabled update
	//TODO id perms update
	//TODO allowed_address_pairs update
	//TODO fixed_ips update

	if _, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
		FieldMask:               fm,
	}); err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	return makePortResponse(vn, vmi, vmi.GetInstanceIPBackRefs()), nil
}

// Delete handles port delete requests.
func (port *Port) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vmi, _, err := port.readVNCPort(ctx, rp, id)
	if err != nil {
		return nil, newNeutronError(portNotFound, errorFields{
			"port_id": id,
		})
	}

	if len(vmi.GetLogicalRouterBackRefs()) > 0 {
		return nil, newNeutronError(l3PortInUse, errorFields{
			"port_id":      id,
			"device_owner": "network:router_interface",
		})
	}

	// release instance IP address
	for _, iip := range vmi.GetInstanceIPBackRefs() {
		// TODO handle shared ip case
		if _, err = rp.WriteService.DeleteInstanceIP(ctx, &services.DeleteInstanceIPRequest{
			ID: iip.GetUUID(),
		}); err != nil {
			// instance ip could be deleted by svc monitor if it is
			// a shared ip. Ignore this error
		}
	}

	//TODO disassociate any floating IP used by instance

	if _, err = rp.WriteService.DeleteVirtualMachineInterface(ctx, &services.DeleteVirtualMachineInterfaceRequest{
		ID: vmi.GetUUID(),
	}); err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	if vmID := port.getAsssociatedVirtualMachineID(vmi); vmID != "" {
		_, err = rp.WriteService.DeleteVirtualMachine(ctx, &services.DeleteVirtualMachineRequest{
			ID: vmID,
		})
		// delete instance if this was the last port
		if err != nil && !errutil.IsNotFound(err) && !errutil.IsConflict(err) {
			return nil, newNeutronError(badRequest, errorFields{
				"resource": "port",
				"msg":      err.Error(),
			})
		}
	}

	//TODO delete any interface route table associated with the port to handle
	// subnet host route Neutron extension, un-reference others

	//TODO make correct delete response
	return &PortResponse{}, nil
}

// Read handles port read requests.
func (port *Port) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vmi, vn, err := port.readVNCPort(ctx, rp, id)
	if err != nil {
		return nil, newNeutronError(portNotFound, errorFields{
			"port_id": id,
		})
	}

	return makePortResponse(vn, vmi, vmi.GetInstanceIPBackRefs()), nil
}

// ReadAll handles port read all requests.
func (port *Port) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement rest of ReadAll logic
	ps := []*PortResponse{}
	if filters.HaveKeys("device_id") {
		deviceUUIDs := filters["device_id"]
		if len(deviceUUIDs) != 1 {
			return ps, nil
		}

		idToFQNameRes, err := rp.IDToFQNameService.IDToFQName(ctx, &services.IDToFQNameRequest{
			UUID: deviceUUIDs[0],
		})

		if errutil.IsNotFound(err) {
			return ps, nil
		}

		if err != nil {
			return nil, newNeutronError(badRequest, errorFields{
				"resource": "port",
				"msg":      err.Error(),
			})
		}

		// TODO handle another resources associated with port using device_id field in filters
		if idToFQNameRes.GetType() == models.KindVirtualMachine {
			return port.readPortsAssociatedWithVM(ctx, rp, filters, deviceUUIDs[0])
		}
	}
	return ps, nil
}

func (port *Port) getAsssociatedVirtualMachineID(vmi *models.VirtualMachineInterface) string {
	if vmi.GetParentType() == models.KindVirtualMachine {
		return vmi.GetParentUUID()
	}

	if len(vmi.GetVirtualMachineRefs()) > 0 {
		return vmi.VirtualMachineRefs[0].GetUUID()
	}

	return ""
}

func (port *Port) readPortsAssociatedWithVM(
	ctx context.Context, rp RequestParameters, filters Filters, vmUUID string,
) ([]*PortResponse, error) {
	vmRes, err := rp.ReadService.GetVirtualMachine(ctx, &services.GetVirtualMachineRequest{
		ID: vmUUID,
	})
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	ps := []*PortResponse{}
	vmiBackRefs := vmRes.GetVirtualMachine().GetVirtualMachineInterfaceBackRefs()
	if len(vmiBackRefs) > 0 {
		var vmi *models.VirtualMachineInterface
		var vn *models.VirtualNetwork
		vmi, vn, err = port.readVNCPort(ctx, rp, vmiBackRefs[0].GetUUID())
		if err != nil {
			return nil, newNeutronError(portNotFound, errorFields{
				"port_id": vmiBackRefs[0].GetUUID(),
			})
		}

		ps = append(ps, makePortResponse(vn, vmi, vmi.GetInstanceIPBackRefs()))
	}
	return ps, nil
}

func (port *Port) getAssociatedVirtualNetwork(ctx context.Context, rp RequestParameters,
	vmi *models.VirtualMachineInterface,
) (*models.VirtualNetwork, error) {
	vnRefs := vmi.GetVirtualNetworkRefs()
	if len(vnRefs) == 0 {
		return nil, nil
	}
	vnRes, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: vnRefs[0].GetUUID(),
	})
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}
	return vnRes.GetVirtualNetwork(), nil
}

func (port *Port) readVNCPort(
	ctx context.Context, rp RequestParameters, id string,
) (*models.VirtualMachineInterface, *models.VirtualNetwork, error) {
	vmiRes, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: id,
	})

	if err != nil {
		return nil, nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	vmi := vmiRes.GetVirtualMachineInterface()
	vn, err := port.getAssociatedVirtualNetwork(ctx, rp, vmi)
	if errutil.IsNotFound(err) {
		return nil, nil, newNeutronError(networkNotFound, errorFields{
			"net_id": id,
		})
	} else if err != nil {
		return nil, nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	return vmi, vn, nil
}

func (port *Port) allocateIPAddress(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork, vmi *models.VirtualMachineInterface,
) (*models.InstanceIP, error) {

	//TODO handle fixed_ips
	if len(vn.NetworkIpamRefs) == 0 {
		return nil, nil
	}

	return port.createInstanceIP(ctx, rp, vn, vmi, "", "")
}

func (port *Port) createInstanceIP(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork, vmi *models.VirtualMachineInterface,
	subnetUUID string, ipAddress string,
) (*models.InstanceIP, error) {

	ipUUID := uuid.NewV4().String()
	iip := &models.InstanceIP{
		Name:              ipUUID,
		UUID:              ipUUID,
		SubnetUUID:        subnetUUID,
		InstanceIPAddress: ipAddress,
		Perms2: &models.PermType2{
			Owner: vmi.GetPerms2().GetOwner(),
		},
	}

	iip.AddVirtualMachineInterfaceRef(&models.InstanceIPVirtualMachineInterfaceRef{
		UUID: vmi.UUID,
	})

	iip.AddVirtualNetworkRef(&models.InstanceIPVirtualNetworkRef{
		UUID: vn.UUID,
	})

	iipRes, err := rp.WriteService.CreateInstanceIP(ctx, &services.CreateInstanceIPRequest{
		InstanceIP: iip,
	})

	if err != nil {
		errorName := ipAddressGenerationFailure
		if errutil.IsBadRequest(err) {
			errorName = badRequest
		}
		return nil, newNeutronError(errorName, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	return iipRes.InstanceIP, nil
}

func (port *Port) getVirtualNetwork(
	ctx context.Context, rp RequestParameters,
) (*models.VirtualNetwork, error) {
	res, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: port.NetworkID,
	})

	if err != nil {
		return nil, newNeutronError(networkNotFound, errorFields{
			"net_id": port.NetworkID,
		})
	}

	return res.GetVirtualNetwork(), err
}

func (port *Port) getProjectID() string {
	uuid, err := uuid.Parse(port.TenantID)
	if err != nil {
		return ""
	}
	return uuid.String()
}

func (port *Port) parseDeviceID() string {
	uuid, err := uuid.Parse(port.DeviceID)
	if err != nil {
		return ""
	}
	return uuid.String()

}

func (port *Port) ensureInstanceExists(
	ctx context.Context, rp RequestParameters, tenantID string,
) (*models.VirtualMachine, error) {
	vm := &models.VirtualMachine{
		Name: port.DeviceID,
		UUID: port.parseDeviceID(),
		Perms2: &models.PermType2{
			Owner: tenantID,
		},
	}

	//TODO: Handle baremetal
	vm.ServerType = "virtual-server"

	createRes, err := rp.WriteService.CreateVirtualMachine(ctx, &services.CreateVirtualMachineRequest{
		VirtualMachine: vm,
	})

	if errutil.IsConflict(err) {
		var vmRes *services.GetVirtualMachineResponse
		vmRes, err = rp.ReadService.GetVirtualMachine(ctx, &services.GetVirtualMachineRequest{
			ID: vm.GetUUID(),
		})

		if err != nil {
			return nil, newNeutronError(badRequest, errorFields{
				"resource": "port",
				"msg":      err.Error(),
			})
		}
		//TODO: Handle baremetal
		vm = vmRes.GetVirtualMachine()
	} else if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	} else {
		vm = createRes.GetVirtualMachine()
	}

	return vm, nil
}

func (port *Port) setVMInstance(ctx context.Context, rp RequestParameters,
	vmi *models.VirtualMachineInterface) error {
	//TODO: Delete old virtual machine object associated with the port

	if port.DeviceID == "" {
		vmi.VirtualMachineRefs = nil
		return nil
	}

	vm, err := port.ensureInstanceExists(ctx, rp, vmi.GetPerms2().GetOwner())
	if err != nil {
		return newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	vmi.AddVirtualMachineRef(&models.VirtualMachineInterfaceVirtualMachineRef{
		UUID: vm.GetUUID(),
		To:   vm.GetFQName(),
	})

	return nil
}

func (port *Port) setPortSecurity(
	ctx context.Context, rp RequestParameters, vmi *models.VirtualMachineInterface, vn *models.VirtualNetwork,
) error {
	vmi.PortSecurityEnabled = port.PortSecurityEnabled
	if !vmi.PortSecurityEnabled {
		vmi.PortSecurityEnabled = vn.PortSecurityEnabled
	}

	res, err := rp.ReadService.ListSecurityGroup(ctx, &services.ListSecurityGroupRequest{
		Spec: &baseservices.ListSpec{
			ObjectUUIDs: port.SecurityGroups,
			Fields:      []string{"uuid", "fqname"},
		},
	})

	if errutil.IsNotFound(err) {
		// TODO add information which group is missing
		return newNeutronError(securityGroupNotFound, errorFields{
			"device_owner": "network:router_interface",
		})
	} else if err != nil {
		return newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	securityGroups := res.GetSecurityGroups()
	for _, sc := range securityGroups {
		vmi.AddSecurityGroupRef(&models.VirtualMachineInterfaceSecurityGroupRef{
			UUID: sc.GetUUID(),
		})
	}

	if len(vmi.SecurityGroupRefs) == 0 && vmi.PortSecurityEnabled {
		vmi.AddSecurityGroupRef(&models.VirtualMachineInterfaceSecurityGroupRef{
			To: sgNoRuleFQName,
		})
	}

	//TODO Handle default security group

	return nil
}

func (port *Port) createVirtualMachineInterface(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork,
) (*models.VirtualMachineInterface, error) {

	vmi := &models.VirtualMachineInterface{
		UUID:       port.ID,
		ParentType: models.KindProject,
		ParentUUID: port.getProjectID(),
		IDPerms: &models.IdPermsType{
			Enable: true,
		},
		Perms2: &models.PermType2{
			Owner: port.TenantID,
		},
	}

	if port.Name == "" {
		vmi.Name = port.ID
	} else {
		vmi.Name = port.Name
		vmi.DisplayName = port.Name
	}

	if port.MacAddress != "" {
		vmi.VirtualMachineInterfaceMacAddresses = &models.MacAddressesType{
			MacAddress: []string{port.MacAddress},
		}
	}

	vmi.AddVirtualNetworkRef(&models.VirtualMachineInterfaceVirtualNetworkRef{
		UUID: vn.GetUUID(),
	})

	if port.DeviceOwner != "network:router_interface" &&
		port.DeviceOwner != "network:router_gateway" && port.DeviceID != "" {
		if err := port.setVMInstance(ctx, rp, vmi); err != nil {
			return nil, err
		}
	}

	vmi.VirtualMachineInterfaceDeviceOwner = port.DeviceOwner
	port.setBindings(vmi)

	if err := port.setPortSecurity(ctx, rp, vmi, vn); err != nil {
		return nil, err
	}

	//TODO Handle allowed address pair
	//TODO Handle fixed ips

	vmiRes, err := rp.WriteService.CreateVirtualMachineInterface(ctx, &services.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	return vmiRes.GetVirtualMachineInterface(), nil
}

func (port *Port) setBinding(vmi *models.VirtualMachineInterface, key string, value string) {
	if vmi.VirtualMachineInterfaceBindings == nil {
		vmi.VirtualMachineInterfaceBindings = models.MakeKeyValuePairs()
	}

	//TODO SetInMap shouldn't return an error
	_ = vmi.GetVirtualMachineInterfaceBindings().SetInMap(&models.KeyValuePair{ // nolint: errcheck
		Key:   key,
		Value: value,
	})
}

func (port *Port) setBindings(vmi *models.VirtualMachineInterface) {
	if port.BindingVnicType != "" {
		port.setBinding(vmi, "vnic_type", port.BindingVnicType)
	}
	if port.BindingVifType != "" {
		port.setBinding(vmi, "vif_type", port.BindingVifType)
	}
	if port.BindingHostID != "" {
		port.setBinding(vmi, "host_id", port.BindingHostID)
	}
}

func (port *Port) checkIfMacAddressIsAvailable(ctx context.Context, rp RequestParameters) error {
	if port.MacAddress == "" {
		return nil
	}

	vmisRes, err := rp.ReadService.ListVirtualMachineInterface(ctx, &services.ListVirtualMachineInterfaceRequest{
		Spec: &baseservices.ListSpec{
			Filters: []*baseservices.Filter{
				{
					Key:    models.VirtualMachineInterfaceFieldVirtualMachineInterfaceMacAddresses,
					Values: []string{port.MacAddress},
				},
			},
		},
	})

	if err != nil {
		return newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	if vmisRes.GetVirtualMachineInterfaceCount() != 0 {
		return newNeutronError(macAddressInUse, errorFields{
			"net_id": port.NetworkID,
			"mac":    port.MacAddress,
		})
	}

	return nil
}
