package logic

import (
	"context"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/gogo/protobuf/types"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// ReadAll will fetch all Subnets.
func (*Subnet) ReadAll(ctx context.Context, rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	var response []*SubnetResponse

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

func (s *Subnet) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	// neutron_plugin_db.py:3174
	virtualNetwork, err := s.getVirtualNetwork(ctx, rp)
	if err != nil {
		return nil, err
	}

	networkIpam, err := s.getNetworkIpam(ctx, rp, virtualNetwork)
	if err != nil {
		return nil, err
	}

	// check if subnet exists and does not overlap
	var networkIpamRef *models.VirtualNetworkNetworkIpamRef
	for _, ipamRef := range virtualNetwork.GetNetworkIpamRefs() {
		if strings.Join(ipamRef.GetTo(), "-") == strings.Join(networkIpam.GetFQName(), "-") {
			networkIpamRef = ipamRef
			break
		}
	}

	subnetVnc := subnetNeutronToVnc(s)
	if networkIpamRef != nil {
		for _, ipamSubnet := range networkIpamRef.GetAttr().GetIpamSubnets() {
			if subnetsOverlaps(subnetVnc, ipamSubnet) {
				return nil, errors.New(fmt.Sprintf(
					"Cidr %s overlaps with another subnet: %s",
					s.Cidr,
					ipamSubnet.GetSubnetUUID(),
				))
			}
		}
	}

	err = updateVirtualNetwork(ctx, rp, virtualNetwork, subnetVnc)
	if err != nil {
		return nil, err
	}

	// TODO(pawel.zadrozny) fetch newly created subnet, and make a proper response
	return &SubnetResponse{}, nil
}

func (s *Subnet) getVirtualNetwork(ctx context.Context, rp RequestParameters) (*models.VirtualNetwork, error) {
	virtualNetworkRequest := &services.GetVirtualNetworkRequest{ID: s.NetworkID}
	virtualNetworkResponse, err := rp.ReadService.GetVirtualNetwork(ctx, virtualNetworkRequest)
	if err != nil {
		return nil, err
	}
	return virtualNetworkResponse.GetVirtualNetwork(), nil
}

func (s *Subnet) getNetworkIpam(
	ctx context.Context,
	rp RequestParameters,
	network *models.VirtualNetwork,
) (*models.NetworkIpam, error) {
	// if requested ipam FQName has length of 3, create new NetworkIpam from it.
	if len(s.IpamFQName) == 3 {
		return &models.NetworkIpam{Name: s.IpamFQName[2], ParentType: "project", FQName: s.IpamFQName}, nil
	}

	// try to link with project's default ipam or global default ipam
	fqName := network.GetFQName()
	netIpamRes, err := rp.ReadService.ListNetworkIpam(ctx, &services.ListNetworkIpamRequest{
		Spec: &baseservices.ListSpec{
			Filters: []*baseservices.Filter{
				{
					Key:    "FQName",
					Values: []string{fqName[len(fqName)-1], "default-network-ipam"},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// if default global subnet does not exist create new empty NetworkIpam
	if netIpamRes.NetworkIpamCount == 0 {
		return &models.NetworkIpam{
			Name:       s.IpamFQName[len(s.IpamFQName)-1],
			ParentType: "project",
			FQName:     s.IpamFQName,
		}, nil
	}

	return netIpamRes.NetworkIpams[0], nil
}

func updateVirtualNetwork(
	ctx context.Context,
	rp RequestParameters,
	vn *models.VirtualNetwork,
	subnet *models.IpamSubnetType,
) error {
	_, err := rp.WriteService.UpdateVirtualNetwork(
		ctx,
		&services.UpdateVirtualNetworkRequest{
			VirtualNetwork: &models.VirtualNetwork{
				UUID: vn.GetUUID(),
				NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						To: []string{vn.GetUUID()},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{subnet},
						},
					},
				},
			},
			FieldMask: types.FieldMask{
				Paths: []string{
					models.VirtualNetworkFieldNetworkIpamRefs,
				},
			},
		},
	)
	return err
}

func subnetsOverlaps(ipamA, ipamB *models.IpamSubnetType) bool {
	// TODO(pawel.zadrozny) make it the right way
	subnetA := ipamA.GetSubnet()
	subnetB := ipamB.GetSubnet()
	if subnetA == nil || subnetB == nil {
		return false
	}

	cidrA := fmt.Sprintf("%s/%d", subnetA.GetIPPrefix(), subnetA.GetIPPrefixLen())
	cidrB := fmt.Sprintf("%s/%d", subnetB.GetIPPrefix(), subnetB.GetIPPrefixLen())

	return cidrA == cidrB
}

func subnetNeutronToVnc(subnet *Subnet) *models.IpamSubnetType {
	// TODO(pawel.zadrozny) please do
	return &models.IpamSubnetType{}
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

	subnet.CIDRFromVnc(ipam.GetSubnet())
	subnet.GatewayFromVnc(ipam.GetDefaultGateway())
	subnet.HostRoutesFromVnc(ipam.GetHostRoutes())

	subnet.DNSNameServersFromVnc(ipam.GetDHCPOptionList())
	subnet.DNSServerAddressFromVnc(ipam.GetDNSServerAddress())

	ipamHasSubnet := ipam.GetSubnet() != nil
	subnet.AllocationPoolsFromVnc(ipam.GetAllocationPools(), ipamHasSubnet)

	return subnet
}

func subnetIsShared(vn *models.VirtualNetwork) bool {
	return vn.GetIsShared() || (vn.GetPerms2() != nil && len(vn.GetPerms2().GetShare()) > 0)
}

// CIDRFromVnc converts VNC Subnet Type CIDR to neutron CIDR and IPVersion format.
func (s *SubnetResponse) CIDRFromVnc(ipamType *models.SubnetType) {
	if ipamType == nil {
		s.Cidr = "0.0.0.0/0"
		s.IPVersion = ipV4
	} else {
		s.Cidr = fmt.Sprintf("%v/%v", ipamType.GetIPPrefix(), ipamType.GetIPPrefixLen())
		ipV, err := getIPVersion(ipamType.GetIPPrefix())
		if err == nil {
			s.IPVersion = int64(ipV)
		}
	}
}

// GatewayFromVnc converts vnc Gateway to neutron Gateway.
func (s *SubnetResponse) GatewayFromVnc(gateway string) {
	if gateway == "0.0.0.0" {
		return
	}
	s.GatewayIP = gateway
}

// AllocationPoolsFromVnc converts VNC Allocation Pool Type to Neutron Allocation Pool format.
func (s *SubnetResponse) AllocationPoolsFromVnc(aps []*models.AllocationPoolType, ipamHasSubnet bool) {
	for _, ap := range aps {
		s.AllocationPools = append(s.AllocationPools, &AllocationPool{
			FirstIP: ap.GetStart(),
			LastIP:  ap.GetEnd(),
		})
	}

	if !ipamHasSubnet {
		s.AllocationPools = append(s.AllocationPools, &AllocationPool{
			FirstIP: "0.0.0.0",
			LastIP:  "255.255.255.255",
		})
	} else if ipamHasSubnet && len(s.AllocationPools) == 0 {
		defaultAllocationPool := subnetDefaultAllocationPool(s.GatewayIP, s.Cidr)
		if defaultAllocationPool != nil {
			s.AllocationPools = append(s.AllocationPools, defaultAllocationPool)
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
		LastIP:  lastIP.String(),
	}
}

// DNSNameServersFromVnc converts VNC DHCP Option List Type to Neutron DNS Nameservers format.
func (s *SubnetResponse) DNSNameServersFromVnc(dhcpOptions *models.DhcpOptionsListType) {
	if dhcpOptions == nil {
		return
	}
	splitter := regexp.MustCompile("[^\\s]+")
	for _, opt := range dhcpOptions.GetDHCPOption() {
		if opt.GetDHCPOptionName() == "6" {
			dnsServers := splitter.FindAllString(opt.GetDHCPOptionValue(), -1)
			for _, dnsServer := range dnsServers {
				s.DNSNameservers = append(s.DNSNameservers, &DnsNameserver{
					Address:  dnsServer,
					SubnetID: s.ID,
				})
			}
		}
	}
}

// DNSServerAddressFromVnc reassign DNS Address Server if contrail extensions are enabled.
func (s *SubnetResponse) DNSServerAddressFromVnc(address string) {
	// TODO(pawel.zadrozny): Check if contrail_extensions_enabled is True neutron_plugin_db.py:1724
	contrailExtensionsEnabled := true
	if contrailExtensionsEnabled {
		s.DNSServerAddress = address
	}
}

// HostRoutesFromVnc converts VNC Route Table Type to Neutron Host Routes format.
func (s *SubnetResponse) HostRoutesFromVnc(routeTable *models.RouteTableType) {
	if routeTable == nil {
		return
	}
	for _, route := range routeTable.GetRoute() {
		s.HostRoutes = append(s.HostRoutes, &RouteTableType{
			Destination: route.GetPrefix(),
			Nexthop:     route.GetNextHop(),
			SubnetID:    s.ID,
		})
	}
}
