package logic

import (
	"context"
	"reflect"
	"strconv"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

var ipFabricRiFqName = [...]string{"default-domain", "default-project", "ip-fabric", "__default__"}
var linkLocalRiFqName = [...]string{"default-domain", "default-project", "__link_local__", "__link_local__"}

// CreateRoutingInstance creates default Route Target.
// It is a small part of __init__ method for RoutingInstanceST in config_db.py.
func (s *Service) CreateRoutingInstance(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) (*services.CreateRoutingInstanceResponse, error) {

	ri := request.GetRoutingInstance()

	if ri == nil {
		return nil, errors.Errorf("couldn't find Routing Instance")
	}

	if request.GetRoutingInstance().GetRoutingInstanceIsDefault() {
		fqName := request.RoutingInstance.GetFQName()
		if reflect.DeepEqual(fqName, ipFabricRiFqName) || reflect.DeepEqual(fqName, linkLocalRiFqName) {
			return &services.CreateRoutingInstanceResponse{RoutingInstance: request.RoutingInstance}, nil
		}
		if err := s.createDefaultRouteTarget(ctx, request); err != nil {
			return nil, err
		}
	} else {
		// TODO: handle the situation in case if it's not default Routing Instance
		// and creating non default route targets (notice locate_route_target
		// method from class RoutingInstanceST in config_db.py)
	}

	return &services.CreateRoutingInstanceResponse{RoutingInstance: request.RoutingInstance}, nil
}

// This method is a small part of locate_route_target method
// for class RoutingInstanceST from config_db.py
func (s *Service) createDefaultRouteTarget(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) error {
	// TODO: get_autonomous_system method and int pool allocator endpoint
	var autonomousSystem = 64512
	var genFromIntPoolAllocator = 8000002

	// TODO: Autonomous system number and create int pool allocator endpoint
	rtKey := "target:" + strconv.Itoa(autonomousSystem) + ":" + strconv.Itoa(genFromIntPoolAllocator)

	rtResponse, err := s.api.CreateRouteTarget(ctx, &services.CreateRouteTargetRequest{
		RouteTarget: &models.RouteTarget{
			FQName:      []string{rtKey},
			DisplayName: rtKey,
		},
	})

	s.api.CreateRoutingInstanceRouteTargetRef(ctx, &services.CreateRoutingInstanceRouteTargetRefRequest{
		ID: request.RoutingInstance.GetUUID(),
		RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
			UUID: rtResponse.RouteTarget.GetUUID(),
		},
	})

	return err
}
