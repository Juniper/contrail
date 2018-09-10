package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// VirtualMachineInterfaceIntent intent
type VirtualMachineInterfaceIntent struct {
	intent.BaseIntent
	*models.VirtualMachineInterface
}

func NewVirtualMachineInterfaceIntent(
	ctx context.Context,
	ReadService services.ReadService,
	request *services.CreateVirtualMachineInterfaceRequest,
) *VirtualMachineInterfaceIntent {
	i := &VirtualMachineInterfaceIntent{
		VirtualMachineInterface: request.GetVirtualMachineInterface(),
	}
	return i
}

// CreateVirtualMachineInterface evaluates VirtualMachineInterface dependencies.
func (s *Service) CreateVirtualMachineInterface(
	ctx context.Context,
	request *services.CreateVirtualMachineInterfaceRequest,
) (*services.CreateVirtualMachineInterfaceResponse, error) {

	i := NewVirtualMachineInterfaceIntent(ctx, s.ReadService, request)

	err := s.handleCreate(ctx, i, i.VirtualMachineInterface)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateVirtualMachineInterface(ctx, request)
}

func LoadVirtualMachineInterfaceIntent(
	loader intent.Loader,
	uuid string,
) *VirtualMachineInterfaceIntent {
	i := loader.Load(models.KindVirtualMachineInterface, intent.UUID(uuid))
	actual, _ := i.(*VirtualMachineInterfaceIntent)
	return actual
}
