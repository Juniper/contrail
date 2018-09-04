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
	initializeVNIntentCache()

	vn := request.GetVirtualNetwork()
	cacheVNIntent(vn)

	if err := EvaluateDependencies(
		ctx,
		&EvaluateContext{WriteService: s.WriteService},
		vn,
		virtualNetworkKey,
	); err != nil {
		return nil, errors.Wrapf(err, "failed to evaluate dependencies of virtual network"+
			"with UUID %q and FQName %q", vn.UUID, vn.FQName)
	}

	return s.BaseService.CreateVirtualNetwork(ctx, request)
}

func initializeVNIntentCache() {
	if _, ok := compilationif.ObjsCache.Load(virtualNetworkIntentKey); !ok {
		compilationif.ObjsCache.Store(virtualNetworkIntentKey, &sync.Map{})
	}
}

func cacheVNIntent(vn *models.VirtualNetwork) {
	syncMap, ok := compilationif.ObjsCache.Load(virtualNetworkIntentKey)
	if ok {
		syncMap.(*sync.Map).Store(vn.GetUUID(), &VirtualNetworkIntent{
			VirtualNetwork: vn,
		})
	}
}

// Evaluate evaluates VirtualNetworkIntent.
func (intent *VirtualNetworkIntent) Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error {
	if intent.routeTargetListsArePresent() {
		riList := intent.RoutingInstances
		if len(riList) == 0 {
			return errors.Errorf("there is no child routing instances for virtual network"+
				"with UUID %q and FQName %q", intent.UUID, intent.FQName)
		}

		ri := riList[0] // TODO: handle all RIs (check if necessary)

		ri.RouteTargetRefs = makeImportExportRouteTargetRefs(
			intent.GetImportRouteTargetList().GetRouteTarget(),
			intent.GetExportRouteTargetList().GetRouteTarget(),
		)

		// TODO: manage ri.staleRouteTargets (check if necessary)

		_, err := evaluateCtx.WriteService.UpdateRoutingInstance(ctx, &services.UpdateRoutingInstanceRequest{
			RoutingInstance: ri,
		})
		if err != nil {
			return errors.Wrapf(err, "failed to update route target refs of routing instance"+
				"with UUID %q and FQName %q", ri.UUID, ri.FQName)
		}
	}

	return nil
}

func (intent *VirtualNetworkIntent) routeTargetListsArePresent() bool {
	if rtList := intent.VirtualNetwork.GetImportRouteTargetList(); rtList != nil {
		if len(rtList.RouteTarget) > 0 {
			return true
		}
	}
	if rtList := intent.VirtualNetwork.GetExportRouteTargetList(); rtList != nil {
		if len(rtList.RouteTarget) > 0 {
			return true
		}
	}
	return false
}

func makeImportExportRouteTargetRefs(
	importRTList, exportRTList []string,
) []*models.RoutingInstanceRouteTargetRef {
	return append(
		rtRefsFromRTList(importRTList, &models.InstanceTargetType{ImportExport: "import"}),
		rtRefsFromRTList(exportRTList, &models.InstanceTargetType{ImportExport: "export"})...,
	)
}

func rtRefsFromRTList(
	rtList []string, instanceTargetType *models.InstanceTargetType,
) (refs []*models.RoutingInstanceRouteTargetRef) {
	for _, rt := range rtList {
		refs = append(refs, &models.RoutingInstanceRouteTargetRef{
			To:   []string{rt},
			Attr: instanceTargetType,
		})
	}
	return refs
}
