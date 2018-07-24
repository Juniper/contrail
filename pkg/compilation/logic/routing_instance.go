package logic

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateRoutingInstance creates default Route Target.
//nolint :lll
// https://github.com/Juniper/contrail-controller/blob/335a2184c9c3d5093e74fc3151c34bfbd40508d9/src/config/schema-transformer/config_db.py#L2075-L2129
func (s *Service) CreateRoutingInstance(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) (*services.CreateRoutingInstanceResponse, error) {

	ri := request.GetRoutingInstance()

	if ri.GetRoutingInstanceIsDefault() {
		fqName := ri.GetFQName()
		if isIPFabricRiFqName(fqName) || isLinkLocalRiFqName(fqName) {
			return &services.CreateRoutingInstanceResponse{RoutingInstance: ri}, nil
		}
		if err := s.createDefaultRouteTarget(ctx, request); err != nil {
			return nil, err
		}
	} else {
		// TODO: handle the situation in case if it's not default Routing Instance
		// and creating non default route targets
		//nolint: lll
		// https://github.com/Juniper/contrail-controller/blob/335a2184c9c3d5093e74fc3151c34bfbd40508d9/src/config/schema-transformer/config_db.py#L2246-L2336
	}

	return &services.CreateRoutingInstanceResponse{RoutingInstance: request.RoutingInstance}, nil
}

// This method is a small part of
//nolint: lll
// https://github.com/Juniper/contrail-controller/blob/335a2184c9c3d5093e74fc3151c34bfbd40508d9/src/config/schema-transformer/config_db.py#L2246-L2336
func (s *Service) createDefaultRouteTarget(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) error {
	// TODO: get_autonomous_system method and int pool allocator endpoint
	//nolint: lll
	// https://github.com/Juniper/contrail-controller/blob/335a2184c9c3d5093e74fc3151c34bfbd40508d9/src/config/schema-transformer/config_db.py#L2248-L2251
	//nolint: lll
	// https://github.com/Juniper/contrail-controller/blob/335a2184c9c3d5093e74fc3151c34bfbd40508d9/src/config/schema-transformer/db.py#L197-L216
	autonomousSystem := 64512
	genFromIntPoolAllocator := 8000002

	rtKey := "target:" + strconv.Itoa(autonomousSystem) + ":" + strconv.Itoa(genFromIntPoolAllocator)

	rtResponse, err := s.WriteService.CreateRouteTarget(ctx, &services.CreateRouteTargetRequest{
		RouteTarget: &models.RouteTarget{
			FQName:      []string{rtKey},
			DisplayName: rtKey,
		},
	})

	_, err = s.WriteService.CreateRoutingInstanceRouteTargetRef(ctx, &services.CreateRoutingInstanceRouteTargetRefRequest{
		ID: request.RoutingInstance.GetUUID(),
		RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
			UUID: rtResponse.RouteTarget.GetUUID(),
		},
	})

	return err
}

func isIPFabricRiFqName(fqName []string) bool {
	fq := []string{"default-domain", "default-project", "ip-fabric", "__default__"}
	return models.FQNameToString(fq) == models.FQNameToString(fqName)
}

func isLinkLocalRiFqName(fqName []string) bool {
	fq := []string{"default-domain", "default-project", "__link_local__", "__link_local__"}
	return models.FQNameToString(fq) == models.FQNameToString(fqName)
}
