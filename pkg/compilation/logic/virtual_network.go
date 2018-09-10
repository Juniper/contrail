package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// VirtualNetworkIntent intent
type VirtualNetworkIntent struct {
	intent.BaseIntent
	*models.VirtualNetwork
}

// GetObject returns embedded resource object
func (i *VirtualNetworkIntent) GetObject() basemodels.Object {
	return i.VirtualNetwork
}

// NewVirtualNetworkIntent returns a new virtual network intent.
func NewVirtualNetworkIntent(
	ctx context.Context,
	ReadService services.ReadService,
	request *services.CreateVirtualNetworkRequest,
) *VirtualNetworkIntent {
	return &VirtualNetworkIntent{
		VirtualNetwork: request.GetVirtualNetwork(),
	}
}

// CreateVirtualNetwork evaluates VirtualNetwork dependencies.
func (s *Service) CreateVirtualNetwork(
	ctx context.Context,
	request *services.CreateVirtualNetworkRequest,
) (*services.CreateVirtualNetworkResponse, error) {

	i := NewVirtualNetworkIntent(ctx, s.ReadService, request)

	err := s.handleCreate(ctx, i)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateVirtualNetwork(ctx, request)
}

// LoadVirtualNetworkIntent loads a virtual network intent from cache.
func LoadVirtualNetworkIntent(
	c intent.Loader,
	uuid string,
) *VirtualNetworkIntent {
	i := c.Load(models.KindVirtualNetwork, intent.ByUUID(uuid))
	actual, _ := i.(*VirtualNetworkIntent)
	return actual
}

// GetPrimaryRoutingInstanceIntent returns the virtual network's default routing instance intent.
func (i *VirtualNetworkIntent) GetPrimaryRoutingInstanceIntent(
	ctx context.Context,
	ec *intent.EvaluateContext,
) *RoutingInstanceIntent {
	fqName := i.DefaultRoutingInstanceFQName()
	ri := ec.IntentLoader.Load(models.TypeNameRoutingInstance, intent.ByFQName(fqName))
	actualIntent, _ := ri.(*RoutingInstanceIntent)
	return actualIntent
}
