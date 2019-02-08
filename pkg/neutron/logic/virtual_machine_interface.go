package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// VirtualMachineInterfaceRead read logic.
func VirtualMachineInterfaceRead(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	vmi, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: id,
		Fields: []string{
			models.VirtualMachineInterfaceFieldLogicalRouterBackRefs,
			models.VirtualMachineInterfaceFieldInstanceIPBackRefs,
			models.VirtualMachineInterfaceFieldFloatingIPBackRefs,
		},
	})
	return vmi, err
}
