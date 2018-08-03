package logic

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// TODO: get_autonomous_system method and int pool allocator endpoint
const (
	defaultAutonomousSystem    = 64512
	defaultRoutingTargetNumber = 8000002
	// This number should be generated from int pool allocator.
)

// CreateRoutingInstance may create default Route Target.
func (s *Service) CreateRoutingInstance(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) (*services.CreateRoutingInstanceResponse, error) {

	ri := request.GetRoutingInstance()

	if ri.GetRoutingInstanceIsDefault() {
		fqName := ri.GetFQName()
		if ri.IsIPFabric() || ri.IsLinkLocal() {
			return &services.CreateRoutingInstanceResponse{RoutingInstance: ri}, nil
		}
		if err := s.createDefaultRouteTarget(ctx, request); err != nil {
			return nil, err
		}
	} else {
		// TODO: handle the situation in case if it's not default Routing Instance
		// and creating non default route targets
	}

	return &services.CreateRoutingInstanceResponse{RoutingInstance: ri}, nil
}

func (s *Service) createDefaultRouteTarget(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) error {
	rtKey := fmt.Sprintf("target:%v:%v", autonomousSystem, genFromIntPoolAllocator)

	rtResponse, err := s.WriteService.CreateRouteTarget(ctx, &services.CreateRouteTargetRequest{
		RouteTarget: &models.RouteTarget{
			FQName:      []string{rtKey},
			DisplayName: rtKey,
		},
	})

	if err != nil {
		return err
	}

	_, err = s.WriteService.CreateRoutingInstanceRouteTargetRef(ctx, &services.CreateRoutingInstanceRouteTargetRefRequest{
		ID: request.RoutingInstance.GetUUID(),
		RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
			UUID: rtResponse.RouteTarget.GetUUID(),
		},
	})

	return err
}
