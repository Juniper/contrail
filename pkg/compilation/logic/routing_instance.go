package logic

import (
	"context"
	"reflect"
	"strconv"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// TODO: This is default autonomous system value in Python.
var autonomousSystem = 64512

// TODO: This value exists here only because there is no int pool allocator endpoint yet
var genFromIntPoolAllocator = 8000002

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

	FQname := request.RoutingInstance.GetFQName()

	if reflect.DeepEqual(FQname, ipFabricRiFqName) || reflect.DeepEqual(FQname, linkLocalRiFqName) {
		return &services.CreateRoutingInstanceResponse{RoutingInstance: request.RoutingInstance}, nil
	}

	RTrefs := request.RoutingInstance.GetRouteTargetRefs()

	if len(RTrefs) == 0 {
		return nil, errors.Errorf("couldn't find Route Target Reference")
	}

	rtuuid := RTrefs[0].GetUUID()

	// TODO: Autonomous system number and create int pool allocator endpoint
	rtname := "target:" + strconv.Itoa(autonomousSystem) + ":" + strconv.Itoa(genFromIntPoolAllocator)

	// TODO: Create route target
	s.api.CreateRouteTarget(ctx, &services.CreateRouteTargetRequest{
		RouteTarget: &models.RouteTarget{
			UUID:        rtuuid,
			FQName:      []string{rtname},
			DisplayName: rtname,
		},
	})

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
