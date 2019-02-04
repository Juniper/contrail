package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

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
	_ context.Context,
	_ services.ReadService,
	request *services.CreateVirtualMachineInterfaceRequest,
) *VirtualMachineInterfaceIntent {
	return &VirtualMachineInterfaceIntent{
		VirtualMachineInterface: request.GetVirtualMachineInterface(),
	}
}

// LoadVirtualMachineInterfaceIntent loads a virtual machine interface intent from cache.
func LoadVirtualMachineInterfaceIntent(loader intent.Loader, query intent.Query) *VirtualMachineInterfaceIntent {
	intent := loader.Load(models.KindVirtualMachineInterface, query)
	vmiIntent, ok := intent.(*VirtualMachineInterfaceIntent)
	if ok == false {
		logrus.Warning("Cannot cast intent to Virtual Machine Interface Intent")
	}
	return vmiIntent
}

// CreateVirtualMachineInterface evaluates VirtualMachineInterface dependencies.
func (s *Service) CreateVirtualMachineInterface(
	ctx context.Context,
	request *services.CreateVirtualMachineInterfaceRequest,
) (*services.CreateVirtualMachineInterfaceResponse, error) {

	i := NewVirtualMachineInterfaceIntent(ctx, s.ReadService, request)

	err := s.storeAndEvaluate(ctx, i)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateVirtualMachineInterface(ctx, request)
}

// UpdateVirtualMachineInterface evaluates UpdateMachineInterface dependencies.
func (s *Service) UpdateVirtualMachineInterface(
	ctx context.Context,
	request *services.UpdateVirtualMachineInterfaceRequest,
) (*services.UpdateVirtualMachineInterfaceResponse, error) {
	vmi := request.GetVirtualMachineInterface()
	if vmi == nil {
		return nil, errors.New("failed to update Virtual Machine Interface." +
			" Virtual Machine Interface Request needs to contain resource!")
	}

	// TODO: Handle update

	i := LoadVirtualMachineInterfaceIntent(s.cache, intent.ByUUID(vmi.GetUUID()))
	if i == nil {
		return nil, errors.Errorf("cannot load intent for Virtual Machine Interface %v", vmi.GetUUID())
	}

	i.VirtualMachineInterface = vmi
	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, errors.Wrapf(err, "failed to update intent for Virtual Machine Interface :%v", vmi.GetUUID())
	}

	return s.BaseService.UpdateVirtualMachineInterface(ctx, request)
}
