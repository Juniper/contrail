package logic

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/apparentlymart/go-cidr/cidr"

	"github.com/Juniper/contrail/pkg/models"
)

// ReadAll will fetch all Subnets.
func (*Subnet) ReadAll(rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	var response []*SubnetResponse

	ctx := context.Background()
	virtualNetworks, err := listVirtualNetworks(ctx, rp, filters)
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

func subnetVncToNeutron(vn *models.VirtualNetwork, ipam *models.IpamSubnetType) *SubnetResponse {
	subnet := &SubnetResponse{
		ID:         ipam.GetSubnetUUID(),
		Name:       ipam.GetSubnetName(),
		TenantID:   strings.Replace(vn.GetParentUUID(), "-", "", -1),
		NetworkID:  vn.GetUUID(),
		EnableDHCP: ipam.GetEnableDHCP(),
		Shared:     subnetIsShared(vn),
		CreatedAt:  ipam.GetCreated(),
		UpdatedAt:  ipam.GetLastModified(),
	}

	subnetCIDRToNeutron(ipam.GetSubnet(), subnet)
	subnetGatewayToNeutron(ipam.GetDefaultGateway(), subnet)
	subnetHostRoutesToNeutron(ipam.GetHostRoutes(), subnet)

	subnetDNSNameServersToNeutron(ipam.GetDHCPOptionList(), subnet)
	subnetDNSServerAddressToNeutron(ipam.GetDNSServerAddress(), subnet)

	ipamHasSubnet := ipam.GetSubnet() != nil
	subnetAllocationPoolsToNeutron(ipam.GetAllocationPools(), ipamHasSubnet, subnet)

	return subnet
}

func subnetIsShared(vn *models.VirtualNetwork) bool {
	return vn.GetIsShared() || (vn.GetPerms2() != nil && len(vn.GetPerms2().GetShare()) > 0)
}

func subnetCIDRToNeutron(ipamType *models.SubnetType, subnet *SubnetResponse) {
	if ipamType == nil {
		subnet.Cidr = "0.0.0.0/0"
		subnet.IPVersion = ipV4
	} else {
		subnet.Cidr = fmt.Sprintf("%v/%v", ipamType.GetIPPrefix(), ipamType.GetIPPrefixLen())
		ipV, err := getIPVersion(ipamType.GetIPPrefix())
		if err == nil {
			subnet.IPVersion = int64(ipV)
		}
	}
}

func subnetGatewayToNeutron(gateway string, subnet *SubnetResponse) {
	if gateway == "0.0.0.0" {
		return
	}
	subnet.GatewayIP = gateway
}

func subnetAllocationPoolsToNeutron(aps []*models.AllocationPoolType, ipamHasSubnet bool, subnet *SubnetResponse) {
	for _, ap := range aps {
		subnet.AllocationPools = append(subnet.AllocationPools, &AllocationPool{
			FirstIP: ap.GetStart(),
			LastIP:  ap.GetEnd(),
		})
	}

	if !ipamHasSubnet {
		subnet.AllocationPools = append(subnet.AllocationPools, &AllocationPool{
			FirstIP: "0.0.0.0",
			LastIP:  "255.255.255.255",
		})
	} else if ipamHasSubnet && len(subnet.AllocationPools) == 0 {
		defaultAllocationPool := subnetDefaultAllocationPool(subnet.GatewayIP, subnet.Cidr)
		if defaultAllocationPool != nil {
			subnet.AllocationPools = append(subnet.AllocationPools, defaultAllocationPool)
		}
	}
}

func subnetDefaultAllocationPool(gateway, subnetCIDR string) *AllocationPool {
	gatewayIP := net.ParseIP(gateway)
	_, netIP, err := net.ParseCIDR(subnetCIDR)
	if gatewayIP == nil || err != nil {
		return nil
	}

	firstIP, lastIP := cidr.AddressRange(netIP)
	firstIP = cidr.Inc(firstIP)
	lastIP = cidr.Dec(lastIP)

	if gatewayIP.Equal(firstIP) {
		firstIP = cidr.Inc(firstIP)
	}

	return &AllocationPool{
		FirstIP: firstIP.String(),
		LastIP: lastIP.String(),
	}
}

func subnetDNSNameServersToNeutron(dhcpOptions *models.DhcpOptionsListType, subnet *SubnetResponse) {
	if dhcpOptions == nil {
		return
	}
	for _, opt := range dhcpOptions.GetDHCPOption() {
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

func subnetDNSServerAddressToNeutron(address string, subnet *SubnetResponse) {
	// TODO(pawel.zadrozny): Check if contrail_extensions_enabled is True neutron_plugin_db.py:1724
	subnet.DNSServerAddress = address
}

func subnetHostRoutesToNeutron(routeTable *models.RouteTableType, subnet *SubnetResponse) {
	if routeTable == nil {
		return
	}
	for _, route := range routeTable.GetRoute() {
		subnet.HostRoutes = append(subnet.HostRoutes, &RouteTableType{
			Destination: route.GetPrefix(),
			Nexthop:     route.GetNextHop(),
			SubnetID:    subnet.ID,
		})
	}
}
