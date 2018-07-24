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

// CreateRoutingInstance creates default Route Target for already created Routing Instance.
// It is based on __init__ method from class RoutingInstanceST in Python
func (s *Service) CreateRoutingInstance(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) (*services.CreateRoutingInstanceResponse, error) {

	ri := request.GetRoutingInstance()

	if ri == nil {
		return nil, errors.Errorf("couldn't find Routing Instance")
	}

	if FQname := request.RoutingInstance.GetFQName(); reflect.DeepEqual(FQname, ipFabricRiFqName) || reflect.DeepEqual(FQname, linkLocalRiFqName) {
		return &services.CreateRoutingInstanceResponse{RoutingInstance: request.RoutingInstance}, nil
	}

	if err := locateRouteTarget(ctx, request, s); err != nil {
		return nil, err
	}

	if request.RoutingInstance.GetRoutingInstanceIsDefault() {

	}

	// TODO: Add connections from routing_instance_refs

	if request.RoutingInstance.GetRoutingInstanceIsDefault() {
		// TODO: load request.VirtualNetwork from DB Cache
		// vn := ...
		// if vn == nil {
		//	 return &services.CreateRoutingInstanceResponse{RoutingInstance: request.RoutingInstance}, nil
		// }

		// TODO: add routing instance connections from VNs and if primary RI is connected
		// to another primary RI, we also have to create connection between VNs
	}

	// TODO: if this routing instance is default update static routes

	return &services.CreateRoutingInstanceResponse{RoutingInstance: request.RoutingInstance}, nil
}

func locateRouteTarget(
	ctx context.Context, request *services.CreateRoutingInstanceRequest, s *Service,
) error {
	// TODO: get_autonomous_system method and int pool allocator endpoint
	var autonomousSystem = 64512
	var genFromIntPoolAllocator = 8000002

	rtRefs := request.RoutingInstance.GetRouteTargetRefs()
	if len(rtRefs) == 0 {
		return errors.Errorf("couldn't find Route Target Reference")
	}
	rtuuid := rtRefs[0].GetUUID()

	// TODO: Autonomous system number and create int pool allocator endpoint
	rtKey := "target:" + strconv.Itoa(autonomousSystem) + ":" + strconv.Itoa(genFromIntPoolAllocator)

	s.api.CreateRouteTarget(ctx, &services.CreateRouteTargetRequest{
		RouteTarget: &models.RouteTarget{
			UUID:        rtuuid,
			FQName:      []string{rtKey},
			DisplayName: rtKey,
		},
	})

	//request.RoutingInstance.ParentUUID
	return nil
}
