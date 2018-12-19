package logic

import (
	"context"

	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/pkg/errors"
)

// Create logic
func (port *Port) Create(rp RequestParameters) (Response, error) {
	portRes := &PortResponse{
		Name:                "c9088d8c-ecd1-472a-9190-905c28a22f53",
		PortSecurityEnabled: true,
		Status:              "DOWN",
		BindingHostID:       "host_id",
	}

	ctx := context.Background()
	// if mac-address is specified, check against the exisitng ports
	// to see if there exists a port with the same mac-address

	vn, err := port.getVirtualNetwork(ctx, rp)
	if err != nil {
		return nil, err
	}

	if err = port.checkMacAddress(ctx, rp); err != nil {
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

func (port *Port) createVirtualMachineInterface(ctx context.Context, rp RequestParameters) error {

	vmi := &models.VirtualMachineInterface{
		IDPerms: &models.IdPermsType{
			Enable: true,
		},
	}

	portUUID := uuid.NewV4().String()

	if len(port.Name) == 0 {
		vmi.Name = portUUID
	} else {
		vmi.Name = port.Name
	}

	_, err := rp.WriteService.CreateVirtualMachineInterface(ctx, &services.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})
	if err != nil {
		return err
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
func (port *Port) Read(rp RequestParameters, id string) (Response, error) {
	return PortResponse{}, nil
}

// ReadAll logic
func (port *Port) ReadAll(rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	// TODO implement ReadAll logic
	return []PortResponse{}, nil
}
