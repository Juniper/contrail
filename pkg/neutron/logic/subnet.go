package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// ReadAll will fetch all Subnets.
func (*Subnet) ReadAll(rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	var response []SubnetResponse

	ctx := context.Background()
	virtualNetworks, err := listVirtualNetwork(ctx, rp, filters)
	if err != nil || len(virtualNetworks) == 0 {
		return response, err
	}

	processed := make(map[string]bool, len(virtualNetworks))
	for _, vn := range virtualNetworks {
		if _, ok := processed[vn.UUID]; ok {
			continue
		}
		processed[vn.UUID] = true
		for _, ipamRef := range vn.GetNetworkIpamRefs() {
			for _, subnetVnc := range ipamRef.GetAttr().GetIpamSubnets() {
				response = append(response, createResponseSubnet(vn, subnetVnc))
			}
		}
	}

	return response, err
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
		kvs, err := retrieveValues(ctx, filters["id"])
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
	routerExternal string,
) ([]*models.VirtualNetwork, error) {
	// TODO: please do
	return []*models.VirtualNetwork{}, nil
}

func retrieveValues(ctx context.Context, keys []string) (vals []string, err error) {
	// TODO: blocked by useragentKV missing gRPC interface.
	return vals, err
}

func createResponseSubnet(vn *models.VirtualNetwork, ipam *models.IpamSubnetType) SubnetResponse {
	// TODO: translate vn and ipam to SubnetResponse
	subnet := SubnetResponse{
		Name: ipam.GetSubnetName(),
		TenantID: vn.GetParentUUID(),
		NetworkID: vn.GetUUID(),
	}

	if ipam.GetSubnet() != nil {
		ipamSubnet := ipam.GetSubnet()
		subnet.Cidr = fmt.Sprintf("%v/%v", ipamSubnet.GetIPPrefix(), ipamSubnet.GetIPPrefixLen())
		ipV, err := getIPVersion(ipamSubnet.GetIPPrefix())
		if err == nil {
			subnet.IPVersion = int64(ipV)
		}
	} else {
		subnet.Cidr = "0.0.0.0/0"
		subnet.IPVersion = ipV4
	}

	return subnet
}
