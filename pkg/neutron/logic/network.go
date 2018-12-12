package logic

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// Create logic
func (n *Network) Create(rp RequestParameters) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}

func listVirtualNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if len(filters) == 0 {
		return listVNWithoutFilters(ctx, rp)
	}

	if filtersHas(filters, "id") {
		return listVNByID(ctx, rp, filters["id"])
	}

	var routerExternal string
	if filtersHas(filters, "router:external") {
		routerExternal = filters["router:external"][0]
	}
	shared := filtersHas(filters, "shared") && filters["shared"][0] == "true"
	return listVNSharedOrExternal(ctx, rp, shared, routerExternal)
}

func listVNWithoutFilters(ctx context.Context, rp RequestParameters) ([]*models.VirtualNetwork, error) {
	var projectID string
	if rp.RequestContext.IsAdmin {
		projectID = rp.RequestContext.Tenant
	}

	var vNetworks []*models.VirtualNetwork
	vn, err := listVNInProject(ctx, rp, projectID)
	if err != nil {
		return nil, err
	}
	vNetworks = append(vNetworks, vn...)

	vn, err = listVNSharedOrExternal(ctx, rp, true, "")
	if err != nil {
		return nil, err
	}
	vNetworks = append(vNetworks, vn...)
	return vNetworks, err
}

func listVNByID(ctx context.Context, rp RequestParameters, IDs []string) ([]*models.VirtualNetwork, error) {
	kvs, err := retrieveValues(ctx, IDs)
	if err != nil {
		return nil, err
	}

	vnIDs := make([]string, 0, len(kvs))
	for _, kv := range kvs {
		vnIDs = append(vnIDs, strings.Split(kv, " ")[0])
	}

	vNetworkResponse, err := rp.ReadService.ListVirtualNetwork(
		ctx,
		&services.ListVirtualNetworkRequest{
			Spec: &baseservices.ListSpec{ObjectUUIDs: vnIDs, Detail: true},
		},
	)
	if err != nil {
		return nil, err
	}
	return vNetworkResponse.VirtualNetworks, err
}

func retrieveValues(ctx context.Context, keys []string) (vals []string, err error) {
	// TODO: blocked by useragentKV missing gRPC interface.
	return vals, err
}

func listVNInProject(
	ctx context.Context,
	rp RequestParameters,
	projectID string,
) ([]*models.VirtualNetwork, error) {
	// TODO: please do
	return []*models.VirtualNetwork{}, nil
}

func listVNSharedOrExternal(
	ctx context.Context,
	rp RequestParameters,
	shared bool,
	routerExternal string,
) ([]*models.VirtualNetwork, error) {
	// TODO: please do
	return []*models.VirtualNetwork{}, nil
}
