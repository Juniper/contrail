package logic

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	floatingIPResourceName              = "floating_ip"
	virtualMachineInterfaceResourceName = "virtual_machine_interface"
)

// Read floating_ip by UUID
func (fip *Floatingip) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	uuid, err := neutronIDToContrailUUID(id)
	if err != nil {
		return nil, newFloatingIPError(fmt.Sprintf("invalid uuid format %v for READ FloatingIp", id), err)
	}
	resp, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{ID: uuid})
	if errutil.IsNotFound(err) {
		return nil, newNeutronError(floatingIPNotFound, errorFields{
			"floatingip_id": id,
			"msg":           err,
		})
	} else if err != nil {
		return nil, err
	}

	portRefs := resp.GetFloatingIP().GetVirtualMachineInterfaceRefs()

	var ports []*models.VirtualMachineInterface
	if portRefs != nil && len(portRefs) > 0 {
		ports, err = listVirtualMachineInterfaces(ctx, rp, portRefs)
		if err != nil {
			return nil, newFloatingIPError("can't get virtual machine interfaces.", err)
		}
	}

	return makeFloatingipResponse(ctx, rp, resp.GetFloatingIP(), ports, nil)
}

// ReadAll logic
func (fip *Floatingip) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement ReadAll logic
	return []FloatingipResponse{}, nil
}

func newFloatingIPError(message string, err error) *Error {
	if err != nil {
		message += fmt.Sprintf(" Error details: %+v", err)
	}
	return newNeutronError(badRequest, errorFields{
		"resource": floatingIPResourceName,
		"msg":      message,
	})
}

func listVirtualMachineInterfaces(ctx context.Context, rp RequestParameters,
	vmiRefs []*models.FloatingIPVirtualMachineInterfaceRef) ([]*models.VirtualMachineInterface, error) {
	vmiUUIDs := make([]string, 0, len(vmiRefs))
	for _, vmiRef := range vmiRefs {
		vmiUUIDs = append(vmiUUIDs, vmiRef.GetUUID())
	}

	vmisResp, err := rp.ReadService.ListVirtualMachineInterface(
		ctx, &services.ListVirtualMachineInterfaceRequest{
			Spec: baseservices.SimpleListSpec(vmiUUIDs, models.VirtualMachineInterfaceFieldVirtualNetworkRefs),
		},
	)

	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": virtualMachineInterfaceResourceName,
			"msg": fmt.Sprintf("error while reading virtual machine interfaces of ids %v from database. "+
				"Error details: %+v", vmiUUIDs, err),
		})
	}
	return vmisResp.GetVirtualMachineInterfaces(), nil
}
