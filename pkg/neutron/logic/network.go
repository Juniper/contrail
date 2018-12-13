package logic

import (
	"context"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/services"
	"strconv"
)

const (
	id  = "id"
	name = "name"
	fqName = "fq_name"
	shared = "shared"
	routerExternal = "router:external"
	tenantID = "tenant_id"
)

type ListReq struct {
	ParentID string
	Filters []*baseservices.Filter
	Detail bool
	Count bool
}

// Create logic
func (n *Network) Create(rp RequestParameters) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}

// ReadAll logic
func (n *Network) ReadAll(rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	var resp Response
	ctx := context.Background()

	vncNets, err := n.collectVirtualNetworks(ctx, rp, filters)
	if err != nil{
		return nil, err
	}

	resp = pruneNetworks(ctx, rp, filters, vncNets)
	return resp, nil
}

func collectNonAdminNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if filtersHas(filters, name) {
		return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant)
	}

	var req ListReq
	if filtersHas(filters, shared) || filtersHas(filters, routerExternal) {
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, &req)
	}

	return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant) // filters nil??
}

func addDBFilter(req *ListReq, key string, values []string) {
	filter := baseservices.Filter{Key: key, Values: values}
	req.Filters = append(req.Filters, &filter)
}


func collectNetworkForTenant(ctx context.Context, rp RequestParameters, filters Filters, tenant string) ([]*models.VirtualNetwork, error) {
	var vns []*models.VirtualNetwork
	var req ListReq
	req.ParentID = tenant
	tenantVNs, err := listNetworksForProject(ctx, rp, &req)
	if err != nil {
		return nil, err
	}
	vns = append(vns, tenantVNs...)

	if len(filters) == 0 {
		return vns, nil
	}

	req.Filters = []*baseservices.Filter{}
	addDBFilter(&req,"is_shared", []string{"true"})
	sharedVNs, err :=  listNetworksForProject(ctx, rp, &req)
	if err != nil {
		return nil, err
	}
	vns = append(vns, sharedVNs...)

	req.Filters = []*baseservices.Filter{}
	addDBFilter(&req,"router_external", []string{"true"})
	rtExtVNs, err :=  listNetworksForProject(ctx, rp, &req)
	if err != nil {
		return nil, err
	}
	vns = append(vns, rtExtVNs...)

	return vns, nil
}

func collectFilteredAdminNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if filtersHas(filters, tenantID) {
		return collectUsingTenantID(ctx, rp, filters)
	}

	var req ListReq
	if filtersHas(filters, name) {
		return listNetworksForProject(ctx, rp, &req)
	}

	if filtersHas(filters, shared) || filtersHas(filters, routerExternal) {
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, &req)
	}
	return nil, nil
}

func collectUsingTenantID(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	projects := validateProjectByID(rp, filters)

	var collectedVNs, vns []*models.VirtualNetwork
	var err error
	var req ListReq

	for _, p := range projects {
		req.ParentID = p
		vns, err = listNetworksForProject(ctx, rp, &req)
		if err != nil {
			return nil, nil
		}
		collectedVNs = append(collectedVNs, vns...)
	}

	if filtersHas(filters, routerExternal) {
		req.ParentID = ""
		rtExtFilter := []*baseservices.Filter{{Key: "router_external", Values: filters[routerExternal]}} //TODO: change to bool after fix to filters
		req.Filters = rtExtFilter
		vns, err = listNetworksForProject(ctx, rp, &req)
		if err != nil {
			return nil, nil
		}
		collectedVNs = append(collectedVNs, vns...)
	}
	return collectedVNs, nil
}

func collectAllNetworksByProjects(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	var req ListReq
	req.Detail=true
	return listNetworksForProject(ctx, rp, &req)
}

func validateProjectByID(rp RequestParameters, filters Filters) []string {
	if !rp.RequestContext.IsAdmin {
		return []string {rp.RequestContext.Tenant}//TODO: handle tables in context
	}

	return filters[tenantID]
}

func (n *Network) collectVirtualNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if filtersHas(filters, id) {
		return collectWithoutPrune(ctx, rp, filters[id])
	}

	if !rp.RequestContext.IsAdmin {
		return collectNonAdminNetworks(ctx, rp, filters)
	}
	if filters != nil {
		return collectFilteredAdminNetworks(ctx, rp, filters)
	}
	return collectAllNetworksByProjects(ctx, rp, filters)
}


func pruneNetworks(ctx context.Context, rp RequestParameters, filters Filters, vns []*models.VirtualNetwork) Response {
	var nns []NetworkResponse
	for _, vn := range vns {
		if containsNetworkWithUUID(nns, vn.GetUUID()) {
			continue
		}
		//TODO: check fqName
		//if !checkFilterValue(filters, fqName, vn.GetFQName()) {
		//
		//}
		if !checkFilterValue(filters, name, vn.GetName()) && !checkFilterValue(filters, name, vn.GetDisplayName()) {
			continue
		}

		isShared := vn.GetIsShared()
		if vn.GetPerms2() != nil && isSharedWithTenant(&rp.RequestContext, vn.GetPerms2().GetShare()) {
			isShared = true
		}
		if !checkFilterValue(filters, shared, strconv.FormatBool(isShared)) {
			continue
		}
		nn := networkVNCToNeutron(ctx, rp, vn, "LIST")
		if nn == nil {
			continue
		}
		nns = append(nns, nn.(NetworkResponse)) //TODO: consider how to handle neutronObjects
	}
	return nns
}

func isSharedWithTenant(rc *RequestContext, sharedList []*models.ShareType) bool {
	if rc != nil && len(sharedList) > 0 {
		for _, t := range sharedList {
			if rc.Tenant == t.Tenant {
				return true
			}
		}
	}
	return false
}

func checkFilterValue(filters Filters, key string, value string) bool {
	if filtersHas(filters, key) && filters[key][0] == value {
		return true
	}
	return false
}

func containsNetworkWithUUID (nns []NetworkResponse, uuid string) bool {
	for _, nn := range nns {
		if nn.ID == uuid {
			return true
		}
	}
	return false
}

func networkVNCToNeutron(ctx context.Context, rp RequestParameters, vn * models.VirtualNetwork, operation string) Response {
	var nn NetworkResponse
	nn.ID = vn.GetUUID()
	if vn.GetDisplayName() == "" {
		nn.Name = vn.FQName[len(vn.FQName) -  1]
	}else {
		nn.Name = vn.GetDisplayName()
	}

	nn.TenantID = contrailUUIDToNeutronID(vn.GetParentUUID())
	nn.ProjectID = contrailUUIDToNeutronID(vn.GetParentUUID())
	nn.AdminStateUp = vn.GetIDPerms().GetEnable()

	if vn.GetIsShared() || (vn.GetPerms2() != nil && isSharedWithTenant(&rp.RequestContext, vn.GetPerms2().GetShare())) {
		nn.Shared = true
	}
	//Rest of convert logic

}

// consider better name
func collectWithoutPrune(ctx context.Context, rp RequestParameters, uuids []string) ([]*models.VirtualNetwork, error) {
	var vns []*models.VirtualNetwork
	for _, uuid := range uuids {
		vnResp, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: uuid})
		if err != nil {
			return nil, err
		}
		vns = append(vns, vnResp.GetVirtualNetwork())
	}
	return vns, nil //TODO: vnc network to neutron network
}

//TODO: filters should support bool values (check if method needed)
func listFilteredNetworks(ctx context.Context,
	rp RequestParameters,
	dbFilters []*baseservices.Filter,
) ([]*models.VirtualNetwork, error) {
	var req ListReq
	req.Filters = dbFilters
	return listNetworksForProject(ctx, rp, &req)
}

func listNetworksForProject(ctx context.Context, rp RequestParameters, req *ListReq) ([]*models.VirtualNetwork, error) {
	sp := prepareVirtualNetworkListSpec(req)
	vnResp, err := rp.ReadService.ListVirtualNetwork(
		ctx,
		&services.ListVirtualNetworkRequest{
			Spec: sp,
		},
	)
	if err != nil {
		return nil, err
	}

	return vnResp.GetVirtualNetworks(), nil
}

func prepareVirtualNetworkListSpec(req *ListReq) *baseservices.ListSpec {
	//TODO: add making table of Filter objects from filters from request
	return &baseservices.ListSpec{
		Filters:      req.Filters,
		Detail:       true,
		Count:        req.Count,
		Shared:       true,
		ParentUUIDs:  []string{req.ParentID},
	}
}

func collectSharedOrRouterExtNetworks(ctx context.Context, rp RequestParameters, filters Filters, req *ListReq) ([]*models.VirtualNetwork, error) {
	var dbFilters []*baseservices.Filter
	if filtersHas(filters, shared) {
		sharedFilter := &baseservices.Filter{Key: "is_shared", Values: filters[shared]} //TODO: change to bool after fix to filters
		dbFilters = append(dbFilters, sharedFilter)
	}

	if filtersHas(filters, routerExternal) {
		rtExtFilter := &baseservices.Filter{Key: "router_external", Values: filters[routerExternal]} //TODO: change to bool after fix to filters
		dbFilters = append(dbFilters, rtExtFilter)
	}

	req.Filters = dbFilters
	return listNetworksForProject(ctx, rp, req)
}

