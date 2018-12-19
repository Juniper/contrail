package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/pkg/errors"
)

// Create logic
func (port *Port) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	portRes := &PortResponse{
		Name:                "c9088d8c-ecd1-472a-9190-905c28a22f53",
		PortSecurityEnabled: true,
		Status:              "DOWN",
		BindingHostID:       "host_id",
	}

	if len(port.ID) == 0 {
		port.ID = uuid.NewV4().String()
	}

	// if mac-address is specified, check against the exisitng ports
	// to see if there exists a port with the same mac-address

	if err := port.checkMacAddress(ctx, rp); err != nil {
		return nil, err
	}

	vn, err := port.getVirtualNetwork(ctx, rp)
	if err != nil {
		return nil, err
	}

	if err = port.createVirtualMachineInterface(ctx, rp, vn); err != nil {
		return nil, err
	}

	// TODO implement create logic
	return portRes, nil
}

func (port *Port) getVirtualNetwork(
	ctx context.Context, rp RequestParameters,
) (*models.VirtualNetwork, error) {
	res, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: port.NetworkID,
	})
	return res.GetVirtualNetwork(), err
}

func (port *Port) getProjectID() string {
	uuid, err := uuid.Parse(port.TenantID)
	if err != nil {
		return ""
	}
	return uuid.String()
}

func (port *Port) ensureInstanceExists(
	ctx context.Context, rp RequestParameters,
) (*models.VirtualMachine, error) {
	vm := &models.VirtualMachine{
		Name: port.DeviceID,
		Perms2: &models.PermType2{
			Owner: port.getProjectID(),
		},
	}

	uuid, err := uuid.Parse(port.DeviceID)
	// if instance_id is not a valid uuid, let
	// virtual_machine_create generate uuid for the vm
	if err == nil {
		vm.UUID = uuid.String()
	}

	//TODO: Handle baremetal
	vm.ServerType = "virtual-server"

	createRes, err := rp.WriteService.CreateVirtualMachine(ctx, &services.CreateVirtualMachineRequest{
		VirtualMachine: vm,
	})

	if errutil.IsConflict(err) {
		// VM already exists try to read id
		readRes, err := rp.ReadService.GetVirtualMachine(ctx, &services.GetVirtualMachineRequest{
			ID: vm.GetUUID(),
		})

		if err != nil {
			return nil, errors.Wrapf(err, "couldn't get virtual machine uuid %s", vm.GetUUID())
		}
		//TODO: Handle baremetal
		vm = readRes.GetVirtualMachine()
	} else if err != nil {
		return nil, errors.Wrapf(err, "couldn't ensure vm instance (%s) existence", vm.GetUUID())
	} else {
		vm = createRes.GetVirtualMachine()
	}

	return vm, nil
}

func (port *Port) setVMInstance(ctx context.Context, rp RequestParameters,
	vmi *models.VirtualMachineInterface) error {
	//TODO: Delete old virtual machine object associated with the port

	if len(port.DeviceID) == 0 {
		vmi.VirtualMachineRefs = nil
		return nil
	}

	vm, err := port.ensureInstanceExists(ctx, rp)
	if err != nil {
		return err
	}

	vmi.VirtualMachineRefs = append(vmi.VirtualMachineRefs,
		&models.VirtualMachineInterfaceVirtualMachineRef{
			UUID: vm.GetUUID(),
		},
	)

	return nil
}

func (port *Port) createVirtualMachineInterface(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork,
) error {

	vmi := &models.VirtualMachineInterface{
		UUID:       port.ID,
		ParentType: models.KindProject,
		ParentUUID: port.ProjectID,
		IDPerms: &models.IdPermsType{
			Enable: true,
		},
		VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
			MacAddress: []string{port.MacAddress},
		},
	}

	if len(port.Name) == 0 {
		vmi.Name = port.ID
	} else {
		vmi.Name = port.Name
		vmi.DisplayName = port.Name
	}

	if len(port.MacAddress) != 0 {
		vmi.VirtualMachineInterfaceMacAddresses = &models.MacAddressesType{
			MacAddress: []string{port.MacAddress},
		}
	}

	vmi.VirtualNetworkRefs = append(vmi.VirtualNetworkRefs, &models.VirtualMachineInterfaceVirtualNetworkRef{
		UUID: vn.GetUUID(),
	})

	if port.DeviceOwner != "network:router_interface" &&
		port.DeviceOwner != "network:router_gateway" && len(port.DeviceID) != 0 {
		if err := port.setVMInstance(ctx, rp, vmi); err != nil {
			return err
		}
	}

	vmi.VirtualMachineInterfaceDeviceOwner = port.DeviceOwner
	if port.BindingVnicType != "" {
		kvps := &models.KeyValuePairs{}
		kvps.KeyValuePair = append(kvps.KeyValuePair, &models.KeyValuePair{
			Key:   "vnic_type",
			Value: port.BindingVnicType,
		})
	}

	vmi.PortSecurityEnabled = port.PortSecurityEnabled

	_, err := rp.WriteService.CreateVirtualMachineInterface(ctx, &services.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})
	if err != nil {
		return errors.Wrapf(err, "couldn't create virtual-machine-interface")
	}

	return nil
}

func (port *Port) checkMacAddress(ctx context.Context, rp RequestParameters) error {
	if len(port.MacAddress) == 0 {
		return nil
	}

	res, err := rp.ReadService.ListVirtualMachineInterface(ctx, &services.ListVirtualMachineInterfaceRequest{
		Spec: &baseservices.ListSpec{
			Filters: []*baseservices.Filter{
				{
					Key:    "virtual_machine_interface_mac_addresses",
					Values: []string{port.MacAddress},
				},
			},
		},
	})

	if err != nil {
		return nil
	}

	if res.GetVirtualMachineInterfaceCount() != 0 {
		errors.Errorf("MacAddressInUse: mac_address = %s", port.MacAddress)
	}

	return nil
}

// Read default implementation
func (port *Port) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	return PortResponse{}, nil
}

// ReadAll logic
func (port *Port) ReadAll(ctx context.Context, rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	// TODO implement ReadAll logic
	return []PortResponse{}, nil
}
