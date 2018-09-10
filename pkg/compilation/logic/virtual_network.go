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
	i := &VirtualNetworkIntent{
		VirtualNetwork: request.GetVirtualNetwork(),
	}
	return i
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
) *VirtualNetworkIntent {
	i := c.Load(models.KindVirtualNetwork, intent.UUID(uuid))
	actual, _ := i.(*VirtualNetworkIntent)
	return actual
}

func (i *VirtualNetworkIntent) GetPrimaryRoutingInstanceIntent(
	ctx context.Context,
	ec *intent.EvaluateContext,
) *RoutingInstanceIntent {
	fqName := i.DefaultRoutingInstanceFQName()
	ri := ec.IntentLoader.Load(models.TypeNameRoutingInstance, intent.FQName(fqName))
	actualIntent, _ := ri.(*RoutingInstanceIntent)
	return actualIntent
}
