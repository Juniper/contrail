package logic

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/gogo/protobuf/types"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	neutronNetworkIDKey = "network_id"
	neutronCIDRKey      = "cidr"

	// TODO(pawel.zadrozny) check if this config is still required or can be removed
	strictCompliance = false
)

// ReadAll will fetch all subnets.
func (*Subnet) ReadAll(ctx context.Context, rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	response := make([]*SubnetResponse, 0)

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
	virtualNetwork, err := getVirtualNetworkByID(ctx, rp, s.NetworkID)
	if err != nil {
		return nil, err
	}

	networkIpam, err := s.getNetworkIpam(ctx, rp, virtualNetwork)
	if err != nil {
		return nil, err
	}

	err = s.updateVirtualNetworkIpamRefs(ctx, rp, virtualNetwork, networkIpam)
	if err != nil {
		return nil, err
	}

	virtualNetwork, err = getVirtualNetworkByID(ctx, rp, s.NetworkID)
	if err != nil {
		return nil, err
	}
	for _, ipamRef := range virtualNetwork.GetNetworkIpamRefs() {
		for _, subnet := range ipamRef.GetAttr().GetIpamSubnets() {
			subnetCidr := fmt.Sprintf("%s/%d", subnet.GetSubnet().GetIPPrefix(), subnet.GetSubnet().GetIPPrefixLen())
			if subnetCidr == s.Cidr {
				return subnetVncToNeutron(virtualNetwork, subnet), nil
			}
		}
	}

	// TODO(pawel.zadrozny) New error message if something goes wrong
	return nil, nil
}

// Locate list of subnets to which this subnet has to be appended
func locateNetworkIpamRef(vn *models.VirtualNetwork, ni *models.NetworkIpam) *models.VirtualNetworkNetworkIpamRef {
	var networkIpamRef *models.VirtualNetworkNetworkIpamRef
	for _, ipamRef := range vn.GetNetworkIpamRefs() {
		if strings.Join(ipamRef.GetTo(), "-") == strings.Join(ni.GetFQName(), "-") {
			networkIpamRef = ipamRef
			break
		}
	}
	return networkIpamRef
}

func (s *Subnet) updateVirtualNetworkIpamRefs(
	ctx context.Context,
	rp RequestParameters,
	vn *models.VirtualNetwork,
	ni *models.NetworkIpam,
) (err error) {
	subnetType, err := s.ipamSubnetType()
	if err != nil {
		return err
	}
	networkIpamRef := locateNetworkIpamRef(vn, ni)
	if networkIpamRef == nil {
		// First link from net to this ipam
		var ipamRefRes *services.CreateVirtualNetworkNetworkIpamRefResponse
		ipamRefRes, err = rp.WriteService.CreateVirtualNetworkNetworkIpamRef(
			ctx,
			&services.CreateVirtualNetworkNetworkIpamRefRequest{
				ID: vn.GetUUID(),
				VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
					To: ni.GetFQName(),
					Attr: &models.VnSubnetsType{
						HostRoutes:  subnetType.GetHostRoutes(),
						IpamSubnets: []*models.IpamSubnetType{subnetType},
					},
				},
			},
		)
		if err != nil {
			return err
		}
		vn.AddNetworkIpamRef(ipamRefRes.GetVirtualNetworkNetworkIpamRef())
	} else {
		// virtual-network already linked to this ipam
		// TODO(pawel.zadrozny) check if subnet exists and does not overlap neutron_plugin_db.py:3217
		networkIpamRef.Attr.IpamSubnets = append(networkIpamRef.Attr.IpamSubnets, subnetType)
	}

	_, err = rp.WriteService.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
		VirtualNetwork: vn,
		FieldMask: types.FieldMask{
			Paths: []string{models.VirtualNetworkFieldNetworkIpamRefs},
		},
	})
	return err
}

func (s *Subnet) getNetworkIpam(ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork) (*models.NetworkIpam, error) {
	networkIpam := models.MakeNetworkIpam()

	if s.IpamFQName != "" {
		networkIpam.FQName = strings.Split(s.IpamFQName, "-")
		return networkIpam, nil
	}

	ipamFQName := make([]string, 0, 3)
	ipamFQName = append(ipamFQName, vn.GetFQName()[:len(vn.GetFQName())-1]...)
	ipamFQName = append(ipamFQName, "default-network-ipam")

	networkIpamRes, err := rp.ReadService.ListNetworkIpam(ctx, &services.ListNetworkIpamRequest{
		Spec: &baseservices.ListSpec{
			Filters: []*baseservices.Filter{
				{Key: models.NetworkIpamFieldFQName, Values: ipamFQName},
			},
			Limit: 1,
		},
	})
	if err != nil {
		return nil, err
	}
	if networkIpamRes.Count() == 1 {
		return networkIpamRes.GetNetworkIpams()[0], nil
	}

	networkIpam.FQName = []string{"default-domain", "default-project", "default-network-ipam"}
	networkIpam.Name = "default-network-ipam"
	return networkIpam, nil
}

func (s *Subnet) ipamSubnetType() (*models.IpamSubnetType, error) {
	subnet, err := s.toVnc()
	if err != nil {
		return nil, err
	}

	defaultGateway, err := s.GatewayToVnc()
	if err != nil {
		return nil, err
	}

	return &models.IpamSubnetType{
		SubnetUUID:       subnet.GetUUID(),
		SubnetName:       subnet.GetName(),
		Subnet:           subnet.GetSubnetIPPrefix(),
		EnableDHCP:       s.EnableDHCP,
		AddrFromStart:    true,
		DHCPOptionList:   s.DHCPOptionListToVnc(),
		DefaultGateway:   defaultGateway,
		DNSServerAddress: s.DNSServerAddressToVnc(),
		AllocationPools:  s.AllocationPoolsToVnc(),
		HostRoutes:       s.HostRoutesToVnc(),
		Created:          s.CreatedAt,
		LastModified:     s.UpdatedAt,
	}, nil
}

func (s *Subnet) toVnc() (*models.Subnet, error) {
	subnetIPPrefix, err := s.SubnetTypeToVnc()
	if err != nil {
		return nil, err
	}

	subnet := models.MakeSubnet()
	subnet.Name = s.Name
	subnet.ParentUUID = s.NetworkID
	subnet.SubnetIPPrefix = subnetIPPrefix

	return subnet, nil
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
				DHCPOptionName:  "6",
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

	_, netIP, err := net.ParseCIDR(s.Cidr)
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
		return s.GatewayIP
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
			Start: allocPool.Start,
			End:   allocPool.End,
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
	prefix := strings.Split(netIP.String(), "/")

	prefixIP := prefix[0]
	prefixLen, err := strconv.ParseInt(prefix[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return &models.SubnetType{IPPrefix: prefixIP, IPPrefixLen: prefixLen}, nil
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
		s.Cidr = fmt.Sprintf("%s/%d", ipamType.GetIPPrefix(), ipamType.GetIPPrefixLen())
		ipV, err := getIPVersionFromCIDR(ipamType.GetIPPrefix())
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
			Start: ap.GetStart(),
			End:   ap.GetEnd(),
		})
	}

	if !ipamHasSubnet {
		s.AllocationPools = append(s.AllocationPools, &AllocationPool{
			Start: "0.0.0.0",
			End:   "255.255.255.255",
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
		Start: firstIP.String(),
		End:   lastIP.String(),
	}
}

// DNSNameServersFromVnc converts VNC DHCP Option List Type to Neutron DNS Nameservers format.
func (s *SubnetResponse) DNSNameServersFromVnc(dhcpOptions *models.DhcpOptionsListType) {
	s.DNSNameservers = make([]*DnsNameserver, 0)
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
	s.HostRoutes = make([]*RouteTableType, 0)
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

func getVirtualNetworkByID(ctx context.Context, rp RequestParameters, networkID string) (*models.VirtualNetwork, error) {
	virtualNetworkRequest := &services.GetVirtualNetworkRequest{ID: networkID}
	virtualNetworkResponse, err := rp.ReadService.GetVirtualNetwork(ctx, virtualNetworkRequest)
	if err != nil {
		return nil, err
	}
	return virtualNetworkResponse.GetVirtualNetwork(), nil
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
