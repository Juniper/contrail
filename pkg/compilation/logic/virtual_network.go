package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// VirtualNetworkIntent intent
type VirtualNetworkIntent struct {
	intent.BaseIntent
	*models.VirtualNetwork
}

func NewVirtualNetworkIntent(
	ctx context.Context,
	ReadService services.ReadService,
	request *services.CreateVirtualNetworkRequest,
) *VirtualNetworkIntent {
	vn := &VirtualNetworkIntent{
		VirtualNetwork: request.GetVirtualNetwork(),
	}

	return vn
}

// CreateVirtualNetwork evaluates VirtualNetwork dependencies.
func (s *Service) CreateVirtualNetwork(
	ctx context.Context,
	request *services.CreateVirtualNetworkRequest,
) (*services.CreateVirtualNetworkResponse, error) {

	i := NewVirtualNetworkIntent(ctx, s.ReadService, request)

	err := s.handleCreate(ctx, i, i.VirtualNetwork)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateVirtualNetwork(ctx, request)
}

func LoadVirtualNetworkIntent(
	c intent.Loader,
	uuid string,
) (*VirtualNetworkIntent, bool) {
	i, ok := c.Load(models.KindVirtualNetwork, uuid)
	if !ok {
		return nil, false
	}
	actual, ok := i.(*VirtualNetworkIntent)
	if !ok {
		return nil, false
	}
	return actual, true
}

func (i *VirtualNetworkIntent) GetPrimaryRoutingInstanceIntent(
	ctx context.Context,
	ec *intent.EvaluateContext,
) *RoutingInstanceIntent {
	fqName := i.DefaultRoutingInstanceFQName()
	ri, _ := ec.IntentLoader.LoadByFQName(models.TypeNameRoutingInstance, fqName)
	actualIntent, _ := ri.(*RoutingInstanceIntent)
	return actualIntent
}
