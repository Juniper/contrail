package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// VirtualMachineInterfaceIntent intent
type VirtualMachineInterfaceIntent struct {
	intent.BaseIntent
	*models.VirtualMachineInterface
}

// GetObject returns embedded resource object
func (i *VirtualMachineInterfaceIntent) GetObject() basemodels.Object {
	return i.VirtualMachineInterface
}

// NewVirtualMachineInterfaceIntent returns a new virtual machine interface intent.
func NewVirtualMachineInterfaceIntent(
	ctx context.Context,
	ReadService services.ReadService,
	request *services.CreateVirtualMachineInterfaceRequest,
) *VirtualMachineInterfaceIntent {
	return &VirtualMachineInterfaceIntent{
		VirtualMachineInterface: request.GetVirtualMachineInterface(),
	}
}

// CreateVirtualMachineInterface evaluates VirtualMachineInterface dependencies.
func (s *Service) CreateVirtualMachineInterface(
	ctx context.Context,
	request *services.CreateVirtualMachineInterfaceRequest,
) (*services.CreateVirtualMachineInterfaceResponse, error) {

	i := NewVirtualMachineInterfaceIntent(ctx, s.ReadService, request)

	err := s.handleCreate(ctx, i)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateVirtualMachineInterface(ctx, request)
}

// LoadVirtualMachineInterfaceIntent loads a virtual machine interface intent from cache.
func LoadVirtualMachineInterfaceIntent(
	loader intent.Loader,
	uuid string,
) *VirtualMachineInterfaceIntent {
	i := loader.Load(models.KindVirtualMachineInterface, intent.ByUUID(uuid))
	actual, _ := i.(*VirtualMachineInterfaceIntent)
	return actual
}
