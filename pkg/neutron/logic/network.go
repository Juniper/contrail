package logic

import (
	"context"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	id  = "id"
	name = "name"
	shared = "shared"
	routerExternal = "router:external"
)

type ListReq struct {
	parentID string
	objUUIDs []string
	fields Fields
	filters []*baseservices.Filter
	detail bool
	count bool
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
	var err error
	ctx := context.Background()

	resp, err = n.collectVirtualNetworks(ctx, rp, filters)
	if err != nil{
		return nil, err
	}
	//TODO: prune networks

	return resp, nil
}

func collectNonAdminNetworks(ctx context.Context, rp RequestParameters, filters Filters) (Response, error) {
	if filters == nil {
		return nil, nil //TODO: collect shared and then collect router:external
	}

	if filtersHas(filters, id) {
		return collectWithoutPrune(filters[id])
	}

	if filtersHas(filters, name) {
		return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant)
	}

	if filtersHas(filters, shared) || filtersHas(filters, routerExternal) {
		return collectShared()
	}


	return nil, nil
}

func collectShared() (Response, error) {
	return nil
}

func collectNetworkForTenant(ctx context.Context, rp RequestParameters, filters Filters, tenant string) (Response, error) {
	var vns []*models.VirtualNetwork
	var req ListReq
	req.parentID = tenant
	tenantVNs, err := listNetworksForProject(ctx, rp, &req, false)
	if err != nil {
		return nil, err
	}
	vns = append(vns, tenantVNs...)

	sharedFilter := []*baseservices.Filter{{Key: "is_shared", Values: []string{"true"}}} //TODO: change to bool after fix to filters
	req.filters = sharedFilter
	sharedVNs, err :=  listNetworksForProject(ctx, rp, &req, false)
	if err != nil {
		return nil, err
	}
	vns = append(vns, sharedVNs...)

	rtExtFilter := []*baseservices.Filter{{Key: "router_external", Values: []string{"true"}}} //TODO: change to bool after fix to filters
	req.filters = rtExtFilter
	rtExtVNs, err :=  listNetworksForProject(ctx, rp, &req, false)
	if err != nil {
		return nil, err
	}
	vns = append(vns, rtExtVNs...)

	return vns, nil
}

func collectFilteredAdminNetworks(ctx context.Context, rp RequestParameters, filters Filters) (Response, error) {
	if filtersHas(filters, id) {
		return collectWithoutPrune(filters[id])
	}

	return nil, nil
}

func collectAllNetworksByProjects(ctx context.Context, rp RequestParameters, filters Filters) (Response, error) {
	return nil, nil
}

func (n *Network) collectVirtualNetworks(ctx context.Context, rp RequestParameters, filters Filters) (Response, error) {
	if !rp.RequestContext.IsAdmin {
		return collectNonAdminNetworks(ctx, rp, filters)
	}
	if filters != nil {
		return collectFilteredAdminNetworks(ctx, rp, filters)
	}
	return collectAllNetworksByProjects(ctx, rp, filters)
}

func networkVNCToNeutron(ctx context.Context, rp RequestParameters, n * models.VirtualNetwork, operation string) Response {
	return nil
}

// consider better name
func collectWithoutPrune(uuids []string) (Response, error) {
	return nil, nil
}

//TODO: filters should support bool values
func listFilteredNetworks(ctx context.Context,
	rp RequestParameters,
	dbFilters []*baseservices.Filter,
) ([]*models.VirtualNetwork, error) {
	var req ListReq
	req.filters = dbFilters
	return listNetworksForProject(ctx, rp, &req, false)
}

func listNetworksForProject(
	ctx context.Context,
	rp RequestParameters,
	req *ListReq, count bool,
) ([]*models.VirtualNetwork, error) {

	if count {
		req.count = true
	} else {
		req.detail = true
	}

	return listVirtualNetworks(ctx, rp, req)
}

func listVirtualNetworks(ctx context.Context, rp RequestParameters, req *ListReq) ([]*models.VirtualNetwork, error) {
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
		Filters:      nil, //TODO: change that
		Detail:       req.detail,
		Count:        req.count,
		Shared:       true,
		ParentUUIDs:  []string{req.parentID},
		ObjectUUIDs:  req.objUUIDs,
		Fields:       req.fields,
	}
}

