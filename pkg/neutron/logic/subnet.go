package logic

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// ReadAll will fetch all Subnets.
func (*Subnet) ReadAll(rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	ctx := context.Background()
	virtualNetworks, err := listVirtualNetwork(ctx, rp, filters)
	if err != nil || len(virtualNetworks) == 0 {
		return []SubnetResponse{}, err
	}

	processed := make(map[string]bool, len(virtualNetworks))
	for _, vn := range virtualNetworks {
		if _, ok := processed[vn.UUID]; ok {
			continue
		}
		processed[vn.UUID] = true


	}
	// TODO: prepare response
	// ListSubnet(RequestParameters.RequestParameters, *ListSubnetRequest) (*ListSubnetResponse, error)
	return []SubnetResponse{}, err
}

func listVirtualNetwork(
	ctx context.Context,
	rp RequestParameters,
	filters Filters,
) (vNetworks []*models.VirtualNetwork, err error) {
	if len(filters) == 0 {
		var projectID string
		if rp.RequestContext.IsAdmin {
			projectID = rp.RequestContext.Tenant
		}

		vn, err := listVirtualNetworkByProject(ctx, rp, projectID)
		if err != nil {
			return nil, err
		}
		vNetworks = append(vNetworks, vn...)

		vn, err = listVirtualNetworkFiltered(ctx, rp, true, "")
		if err != nil {
			return nil, err
		}
		vNetworks = append(vNetworks, vn...)
		return vNetworks, err
	}

	if filtersHas(filters, "id") {
		kvs, err := rp.DBService.RetrieveValues(ctx, filters["id"])
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

	if filtersHas(filters, "shared") || filtersHas(filters, "router:external") {
		var shared bool
		if filtersHas(filters, "shared") && filters["shared"][0] == "true" {
			shared = true
		}

		var routerExternal string
		if filtersHas(filters, "router:external") {
			routerExternal = filters["router:external"][0]
		}
		return listVirtualNetworkFiltered(ctx, rp, shared, routerExternal)
	}

	return vNetworks, err
}

func listVirtualNetworkByProject(
	ctx context.Context,
	rp RequestParameters,
	projectID string,
) ([]*models.VirtualNetwork, error) {
	// TODO: please do
	return []*models.VirtualNetwork{}, nil
}

func listVirtualNetworkFiltered(
	ctx context.Context,
	rp RequestParameters,
	shared bool,
	router_external string,
) ([]*models.VirtualNetwork, error) {
	// TODO: please do
	return []*models.VirtualNetwork{}, nil
}
