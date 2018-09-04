package logic

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	virtualNetworkKey       = "VirtualNetwork"
	virtualNetworkIntentKey = "VirtualNetworkIntent"
)

// VirtualNetworkIntent contains state for VirtualNetwork.
type VirtualNetworkIntent struct {
	BaseIntent
	*models.VirtualNetwork
}

// CreateVirtualNetwork caches VirtualNetworkIntent and evaluates VirtualNetwork dependencies.
func (s *Service) CreateVirtualNetwork(
	ctx context.Context, request *services.CreateVirtualNetworkRequest,
) (*services.CreateVirtualNetworkResponse, error) {
	initializeVNIntentStore()

	vn := request.GetVirtualNetwork()
	storeVNIntent(vn)

	if err := EvaluateDependencies(
		ctx,
		&EvaluateContext{WriteService: s.WriteService},
		vn,
		virtualNetworkKey,
	); err != nil {
		return nil, errors.Wrap(err, "failed to evaluate virtual network dependencies")
	}

	return s.BaseService.CreateVirtualNetwork(ctx, request)
}

func initializeVNIntentStore() {
	if _, ok := compilationif.ObjsCache.Load(virtualNetworkIntentKey); !ok {
		compilationif.ObjsCache.Store(virtualNetworkIntentKey, &sync.Map{})
	}
}

func storeVNIntent(vn *models.VirtualNetwork) {
	objMap, ok := compilationif.ObjsCache.Load(virtualNetworkIntentKey)
	if ok {
		objMap.(*sync.Map).Store(vn.GetUUID(), &VirtualNetworkIntent{
			VirtualNetwork: vn,
		})
	}
}

// Evaluate evaluates VirtualNetworkIntent.
func (s *VirtualNetworkIntent) Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error {
	// TODO: get RI
	ri := &models.RoutingInstance{}

	// TODO: calculate RI refs based on vn.RI_lists

	// TODO: updateRouteTargetList (manage ri.staleRouteTargets and do ri.APIServerUpdate)
	_, err := evaluateCtx.WriteService.UpdateRoutingInstance(ctx, &services.UpdateRoutingInstanceRequest{
		RoutingInstance: ri,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update routing instance route target refs")
	}

	return nil
}

// code moved from APISRV (TODO: remove)

// MakeDefaultRoutingInstance returns the default routing instance for the network.
//func (m *VirtualNetwork) MakeDefaultRoutingInstance() *RoutingInstance {
//	return &RoutingInstance{
//      other fields...
//		RouteTargetRefs:           m.MakeImportExportRouteTargetRefs(),
//	}
//}

// MakeImportExportRouteTargetRefs returns refs to RouteTarget's from import and export lists.
//func (m *VirtualNetwork) MakeImportExportRouteTargetRefs() []*RoutingInstanceRouteTargetRef {
//	return append(
//		m.GetImportRouteTargetList().AsRefs(&InstanceTargetType{ImportExport: "import"}),
//		m.GetExportRouteTargetList().AsRefs(&InstanceTargetType{ImportExport: "export"})...,
//	)
//}

// AsRefs returns refs with instanceTargetType from a RoutingInstance to route targets in the list.
//func (m *RouteTargetList) AsRefs(instanceTargetType *InstanceTargetType) (refs []*RoutingInstanceRouteTargetRef) {
//	for _, rt := range m.GetRouteTarget() {
//		refs = append(refs, &RoutingInstanceRouteTargetRef{
//			To:   []string{rt},
//			Attr: instanceTargetType,
//		})
//	}
//	return refs
//}
