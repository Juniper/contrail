package logic

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/gogo/protobuf/types"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	neutronNetworkIDKey = "network_id"
	neutronCIDRKey      = "cidr"

	// TODO(pawel.zadrozny) check if this config is still required or can be removed
	strictCompliance = false
)

// ReadAll will fetch all subnets.
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
				neutronSN := subnetVncToNeutron(vn, subnetVnc)
				if shouldSkipSubnet(filters, vn, neutronSN) {
					continue
				}
				response = append(response, neutronSN)
			}
		}
	}

	return response, err
}

func shouldSkipSubnet(filters Filters, vn *models.VirtualNetwork, neutronSN *SubnetResponse) bool {
	if len(filters) == 0 {
		return false
	}

	if filters.haveKeys(neutronSharedKey) && filters.checkValue(neutronSharedKey, "true") && !vn.GetIsShared() {
		return true
	}

	if !filters.checkValue(neutronIDKey, neutronSN.ID) {
		return true
	}

	if !filters.checkValue(neutronTenantIDKey, neutronSN.TenantID) {
		return true
	}

	if !filters.checkValue(neutronNetworkIDKey, neutronSN.NetworkID) {
		return true
	}

	if !filters.checkValue(neutronNameKey, neutronSN.Name) {
		return true
	}

	if !filters.checkValue(neutronCIDRKey, neutronSN.Cidr) {
		return true
	}

	return false
}

// Create new subnet for given network
func (s *Subnet) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	// TODO(pawel.zadrozny) validate if CIDR version is equal to ip_version neutron_plugin_db.py:1585
	virtualNetwork, err := s.getVirtualNetwork(ctx, rp)
	if err != nil {
		return nil, err
	}
	// TODO(pawel.zadrozny) check if subnet exists and does not overlap neutron_plugin_db.py:3217

	subnetVnc, err := s.toVnc()
	if err != nil {
		return nil, err
	}

	err = updateVirtualNetwork(ctx, rp, virtualNetwork, subnetVnc)
	if err != nil {
		return nil, err
	}

	virtualNetwork, err = s.getVirtualNetwork(ctx, rp)
	if err != nil {
		return nil, err
	}
	for _, ipamRefs := range virtualNetwork.GetNetworkIpamRefs() {
		for _, ipamType := range ipamRefs.GetAttr().GetIpamSubnets() {
			if netPrefix(ipamType.GetSubnet()) == netPrefix(subnetVnc.GetSubnet()) {
				return subnetVncToNeutron(virtualNetwork, ipamType), nil
			}
		}
	}

	return &SubnetResponse{}, nil
}

func (s *Subnet) toVnc() (*models.IpamSubnetType, error) {
	defaultGateway, err := s.GatewayToVnc()
	if err != nil {
		return nil, err
	}

	subnet, err := s.SubnetTypeToVnc()
	if err != nil {
		return nil, err
	}

	return &models.IpamSubnetType{
		SubnetName: s.Name,
		Created: s.CreatedAt,
		LastModified: s.UpdatedAt,
		EnableDHCP: s.EnableDHCP,
		AddrFromStart: true,
		DHCPOptionList: s.DHCPOptionListToVnc(),
		DefaultGateway: defaultGateway,
		DNSServerAddress: s.DNSServerAddressToVnc(),
		AllocationPools: s.AllocationPoolsToVnc(),
		HostRoutes: s.HostRoutesToVnc(),
		Subnet: subnet,
	}, nil
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

// DHCPOptionListToVnc converts Neutron request to DHCP options list type VNC format.
func (s *Subnet) DHCPOptionListToVnc() *models.DhcpOptionsListType {
	if len(s.DNSNameservers) == 0 {
		return nil
	}

	var optVal []string
	for _, nameserver := range s.DNSNameservers {
		optVal = append(optVal, nameserver.Address)
	}

	return &models.DhcpOptionsListType{
		DHCPOption: []*models.DhcpOptionType{
			{
				DHCPOptionName: "6",
				DHCPOptionValue: strings.Join(optVal, " "),
			},
		},
	}
}

// GatewayToVnc converts Neutron request to Gateway VNC format.
func (s *Subnet) GatewayToVnc() (string, error) {
	if s.GatewayIP != "" {
		return s.GatewayIP, nil
	}

	_, netIP, err := net.ParseCIDR(s.GatewayIP)
	if err != nil {
		return "", err
	}

	firstIP, _ := cidr.AddressRange(netIP)
	return cidr.Inc(firstIP).String(), nil
}

// DNSServerAddressToVnc converts Neutron request to DNS server address VNC format.
func (s *Subnet) DNSServerAddressToVnc() string {
	if strictCompliance {
		return "0.0.0.0"
	}

	if len(s.DNSNameservers) == 0 {
		return ""
	}

	return s.DNSNameservers[0].Address
}

// AllocationPoolsToVnc converts Neutron request to allocation pools VNC format.
func (s *Subnet) AllocationPoolsToVnc() []*models.AllocationPoolType {
	if len(s.AllocationPools) == 0 {
		return nil
	}

	allocationPoolTypes := make([]*models.AllocationPoolType, 0, len(s.AllocationPools))
	for _, allocPool := range s.AllocationPools {
		allocationPoolTypes = append(allocationPoolTypes, &models.AllocationPoolType{
			Start: allocPool.FirstIP,
			End: allocPool.LastIP,
		})
	}
	return allocationPoolTypes
}

// HostRoutesToVnc converts Neutron request to  host routes VNC format.
func (s *Subnet) HostRoutesToVnc() *models.RouteTableType {
	// TODO - not needed for ping by CREATE
	return nil
}

// SubnetTypeToVnc converts Neutron request to subnet type VNC format.
func (s *Subnet) SubnetTypeToVnc() (*models.SubnetType, error) {
	_, netIP, err := net.ParseCIDR(s.Cidr)
	if err != nil {
		return nil, err
	}

	prefixLen, err := strconv.ParseInt(strings.Split(s.Cidr, "/")[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return &models.SubnetType{IPPrefix: netIP.String(), IPPrefixLen: prefixLen}, nil
}

func subnetVncToNeutron(vn *models.VirtualNetwork, ipam *models.IpamSubnetType) *SubnetResponse {
	subnet := &SubnetResponse{
		ID:         ipam.GetSubnetUUID(),
		Name:       ipam.GetSubnetName(),
		TenantID:   contrailUUIDToNeutronID(vn.GetParentUUID()),
		NetworkID:  vn.GetUUID(),
		EnableDHCP: ipam.GetEnableDHCP(),
		Shared:     vn.GetIsShared() || (vn.GetPerms2() != nil && len(vn.GetPerms2().GetShare()) > 0),
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

// CIDRFromVnc converts VNC Subnet Type CIDR to neutron CIDR and IPVersion format.
func (s *SubnetResponse) CIDRFromVnc(ipamType *models.SubnetType) {
	if ipamType == nil {
		s.Cidr = "0.0.0.0/0"
		s.IPVersion = ipV4
	} else {
		s.Cidr = netPrefix(ipamType)
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

func listVirtualNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if len(filters) == 0 {
		return listVNWithoutFilters(ctx, rp)
	}

	if filters.haveKeys(neutronIDKey) {
		kvsResponse, err := rp.UserAgentKV.RetrieveValues(ctx, &services.RetrieveValuesRequest{
			Keys: filters[neutronIDKey],
		})
		if err != nil {
			return nil, err
		}
		return listVNByKeyValues(ctx, rp, kvsResponse.GetValues())
	}

	if filters.haveKeys(neutronSharedKey) || filters.haveKeys(neutronRouterExternalKey) {
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, &listReq{})
	}
	return nil, nil

}

func listVNWithoutFilters(ctx context.Context, rp RequestParameters) ([]*models.VirtualNetwork, error) {
	req := &listReq{}
	if rp.RequestContext.IsAdmin {
		req.ParentID = rp.RequestContext.Tenant
	}

	var vNetworks []*models.VirtualNetwork
	vn, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}
	vNetworks = append(vNetworks, vn...)

	addDBFilter(req, isShared, []string{"true"}, false)
	sharedVNs, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}
	vNetworks = append(vNetworks, sharedVNs...)

	return vNetworks, nil
}

func listVNByKeyValues(ctx context.Context, rp RequestParameters, kvs []string) ([]*models.VirtualNetwork, error) {
	vnIDs := make([]string, 0, len(kvs))
	for _, kv := range kvs {
		vnIDs = append(vnIDs, strings.Split(kv, " ")[0])
	}

	req := &listReq{
		ObjUUIDs: vnIDs,
	}

	vNetworks, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}

	return vNetworks, nil
}

func netPrefix(subnetType *models.SubnetType) string {
	return fmt.Sprintf("%v/%v", subnetType.GetIPPrefix(), subnetType.GetIPPrefixLen())
}
