package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// VirtualNetworkIntent contains Intent Compiler state for VirtualNetwork.
type VirtualNetworkIntent struct {
	BaseIntent
	*models.VirtualNetwork
}

// CreateVirtualNetwork creates resources that depend on VirtualNetwork.
func (s *Service) CreateVirtualNetwork(
	ctx context.Context, request *services.CreateVirtualNetworkRequest,
) (*services.CreateVirtualNetworkResponse, error) {
	// TODO: should it be noop?

	//obj := request.GetVirtualNetwork()
	//
	//intent := &VirtualNetworkIntent{
	//	VirtualNetwork: obj,
	//}
	//
	//if _, ok := compilationif.ObjsCache.Load("VirtualNetworkIntent"); !ok {
	//	compilationif.ObjsCache.Store("VirtualNetworkIntent", &sync.Map{})
	//}
	//
	//objMap, ok := compilationif.ObjsCache.Load("VirtualNetworkIntent")
	//if ok {
	//	objMap.(*sync.Map).Store(obj.GetUUID(), intent)
	//}
	//
	//ec := &EvaluateContext{
	//	WriteService: s.WriteService,
	//}
	//err := EvaluateDependencies(ctx, ec, obj, "VirtualNetwork")
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to evaluate Security Group dependencies")
	//}

	return s.BaseService.CreateVirtualNetwork(ctx, request)
}

// Evaluate evaluates VirtualNetwork dependencies.
func (s *VirtualNetworkIntent) Evaluate(ctx context.Context, evaluateContext *EvaluateContext) error {
	// TODO: get RI

	// TODO: calculate RI refs based on vn.RI_lists

	// TODO: updateRouteTargetList (manage ri.staleRouteTargets and do ri.APIServerUpdate)

	//ingressACL, egressACL := s.DefaultACLs()
	//_, err := evaluateContext.WriteService.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
	//	AccessControlList: ingressACL,
	//})
	//if err != nil {
	//	return errors.Wrap(err, "failed to create ingress access control list")
	//}

	//_, err = evaluateContext.WriteService.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
	//	AccessControlList: egressACL,
	//})
	//if err != nil {
	//	return errors.Wrap(err, "failed to create egress access control list")
	//}

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
