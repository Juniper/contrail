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

	visited := make(map[string]bool, len(virtualNetworks))
	for _, vn := range virtualNetworks {
		if _, ok := visited[vn.UUID]; ok {
			continue
		}
		visited[vn.UUID] = true
		for _, ipamRef := range vn.GetNetworkIpamRefs() {
			for _, subnetVnc := range ipamRef.GetAttr().GetIpamSubnets() {
				response = append(response, subnetVncToNeutron(vn, subnetVnc))
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

func subnetVncToNeutron(vn *models.VirtualNetwork, ipam *models.IpamSubnetType) SubnetResponse {
	const (
		defaultIPAddr = "0.0.0.0"
		defaultCIDR   = "0.0.0.0/0"
	)

	subnet := SubnetResponse{
		ID: ipam.GetSubnetUUID(),
		Name: ipam.GetSubnetName(),
		TenantID: strings.Replace(vn.GetParentUUID(), "-", "", -1),
		NetworkID: vn.GetUUID(),
		EnableDHCP: ipam.GetEnableDHCP(),
		Shared: subnetIsShared(vn),
		CreatedAt: ipam.GetCreated(),
		UpdatedAt: ipam.GetLastModified(),
	}

	if ipamSubnet := ipam.GetSubnet(); ipamSubnet != nil {
		subnet.Cidr = fmt.Sprintf("%v/%v", ipamSubnet.GetIPPrefix(), ipamSubnet.GetIPPrefixLen())
		ipV, err := getIPVersion(ipamSubnet.GetIPPrefix())
		if err == nil {
			subnet.IPVersion = int64(ipV)
		}
	} else {
		subnet.Cidr = defaultCIDR
		subnet.IPVersion = ipV4
	}

	if gateway := ipam.GetDefaultGateway(); gateway != defaultIPAddr {
		subnet.GatewayIP = gateway
	}

	for _, ap := range ipam.GetAllocationPools() {
		subnet.AllocationPools = append(subnet.AllocationPools, &AllocationPool{
			FirstIP: ap.GetStart(),
			LastIP: ap.GetEnd(),
		})
	}

	// TODO(pawel.zadrozny): offset network alloc pool by one: neutron_plugin_db.py:1688

	if dhcpOpts := ipam.GetDHCPOptionList(); dhcpOpts != nil {
		for _, opt := range dhcpOpts.GetDHCPOption() {
			if opt.GetDHCPOptionName() == "6" {
				dnsServers := strings.Split(opt.GetDHCPOptionValue(), "")
				for _, dnsServer := range dnsServers {
					subnet.DNSNameservers = append(subnet.DNSNameservers, &DnsNameserver{
						Address:  dnsServer,
						SubnetID: subnet.ID,
					})
				}
			}
		}
	}

	// TODO(pawel.zadrozny): Check if contrail_extensions_enabled is True neutron_plugin_db.py:1724
	subnet.DNSServerAddress = ipam.GetDNSServerAddress()

	if hostRoutes := ipam.GetHostRoutes(); hostRoutes != nil {
		for _, route := range hostRoutes.GetRoute() {
			subnet.HostRoutes = append(subnet.HostRoutes, &RouteTableType{
				Destination: route.GetPrefix(),
				Nexthop: route.GetNextHop(),
				SubnetID: subnet.ID,
			})
		}
	}

	return subnet
}

func subnetIsShared(vn *models.VirtualNetwork) bool {
	return vn.GetIsShared() || (vn.GetPerms2() != nil && len(vn.GetPerms2().GetShare())  > 0)
}
