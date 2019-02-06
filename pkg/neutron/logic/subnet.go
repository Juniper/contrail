package logic

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/gogo/protobuf/types"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	networkIDKey = "network_id"
	cidrKey      = "cidr"

	defaultDomainName      = "default-domain"
	defaultProjectName     = "default-project"
	defaultNetworkIpamName = "default-network-ipam"

	subnetFieldName            = "name"
	subnetFieldGatewayIP       = "gateway_ip"
	subnetFieldEnableDHCP      = "enable_dhcp"
	subnetFieldDNSNameservers  = "dns_nameservers"
	subnetFieldAllocationPools = "allocation_pools"
	subnetFieldHostRoutes      = "host_routes"

	// TODO(pawel.zadrozny) check if this config is still required or can be removed
	strictCompliance = false
)

var defaultNetworkIpamFQName = []string{defaultDomainName, defaultProjectName, defaultNetworkIpamName}

func newSubnetError(name errorType, format string, args ...interface{}) error {
	return newNeutronError(name, errorFields{
		"resource": "subnet",
		"msg":      fmt.Sprintf(format, args...),
	})
}

// ReadAll will fetch all subnets.
func (*Subnet) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (r Response, err error) {
	var vns []*models.VirtualNetwork
	vns, err = listVNsForSubnetReadAll(ctx, rp, filters)
	if err != nil {
		return nil, newSubnetError(
			networkNotFound,
			"failed to fetch networks: %+v",
			err,
		)
	}

	if len(vns) == 0 {
		return []*SubnetResponse{}, nil
	}

	vns = uniqueVNs(vns)

	r = getSubnetsFromVNs(vns, filters)

	return r, nil
}

func listVNsForSubnetReadAll(
	ctx context.Context, rp RequestParameters, filters Filters,
) ([]*models.VirtualNetwork, error) {
	switch {
	case filters.HaveKeys(idKey):
		return collectVNsUsingKV(ctx, rp, filters[idKey])
	case filters.HaveKeys(sharedKey, routerExternalKey):
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, nil)
	default:
		return listVirtualNetworksWithShared(ctx, rp, filters)
	}
}

func collectVNsUsingKV(ctx context.Context, rp RequestParameters, keys []string) ([]*models.VirtualNetwork, error) {
	uuids, err := getVirtualNetworkIDsFromKV(ctx, rp, keys)
	if err != nil {
		return nil, err
	}
	return listNetworksForProject(ctx, rp, &listReq{ObjUUIDs: uuids})
}

func getVirtualNetworkIDsFromKV(
	ctx context.Context, rp RequestParameters, keys []string,
) (ids []string, err error) {
	kvsResponse, err := rp.UserAgentKV.RetrieveValues(
		ctx, &services.RetrieveValuesRequest{Keys: keys},
	)
	if err != nil {
		return nil, err
	}
	for _, kv := range kvsResponse.GetValues() {
		v := strings.Split(kv, " ")
		if len(v) < 1 {
			continue
		}
		ids = append(ids, v[0])
	}

	return ids, nil
}

func listVirtualNetworksWithShared(
	ctx context.Context, rp RequestParameters, filters Filters,
) (vns []*models.VirtualNetwork, err error) {
	req := &listReq{}
	if !rp.RequestContext.IsAdmin {
		req.ParentID = rp.RequestContext.Tenant
	}

	tenantVNs, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}

	addDBFilter(req, isShared, []string{"true"}, false)
	sharedVNs, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}

	vns = make([]*models.VirtualNetwork, 0, len(tenantVNs)+len(sharedVNs))
	vns = append(vns, tenantVNs...)
	vns = append(vns, sharedVNs...)

	return vns, nil
}

func uniqueVNs(vns []*models.VirtualNetwork) []*models.VirtualNetwork {
	idMap := make(map[string]bool, len(vns))
	result := vns[:0]
	for _, vn := range vns {
		if idMap[vn.GetUUID()] {
			continue
		}
		idMap[vn.GetUUID()] = true
		result = append(result, vn)
	}

	return result
}

func getSubnetsFromVNs(vns []*models.VirtualNetwork, f Filters) (s []*SubnetResponse) {
	s = make([]*SubnetResponse, 0)

	for _, vn := range vns {
		for _, subnetVnc := range vn.GetIpamSubnets().GetSubnets() {
			neutronSN := subnetVncToNeutron(vn, subnetVnc)
			if shouldSkipSubnet(f, vn, neutronSN) {
				continue
			}
			s = append(s, neutronSN)
		}
	}
	return s
}

// Read will fetch subnet with specified id.
func (*Subnet) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	virtualNetworks, err := collectVNsUsingKV(ctx, rp, []string{id})
	if err != nil {
		return nil, newSubnetError(
			subnetNotFound, "failed to fetch networks: %+v", err,
		)
	}

	sub, vn := findSubnetInVirtualNetworks(virtualNetworks, id)

	if sub == nil {
		return nil, newSubnetError(
			subnetNotFound, "subnet not found in vn subnets: %+v", id,
		)
	}
	return subnetVncToNeutron(vn, sub), nil
}

func findSubnetInVirtualNetworks(
	vns []*models.VirtualNetwork, subnetUUID string,
) (found *models.IpamSubnetType, inVN *models.VirtualNetwork) {
	for _, vn := range vns {
		if found = vn.GetIpamSubnets().Find(func(ipamS *models.IpamSubnetType) bool {
			return ipamS.GetSubnetUUID() == subnetUUID
		}); found != nil {
			return found, vn
		}
	}
	return nil, nil
}

func shouldSkipSubnet(filters Filters, vn *models.VirtualNetwork, neutronSN *SubnetResponse) bool {
	if len(filters) == 0 {
		return false
	}

	if !vn.GetIsShared() && filters.HaveValues(sharedKey, "true") {
		return true
	}

	if !filters.Match(idKey, neutronSN.ID) {
		return true
	}

	if !filters.Match(tenantIDKey, neutronSN.TenantID) {
		return true
	}

	if !filters.Match(networkIDKey, neutronSN.NetworkID) {
		return true
	}

	if !filters.Match(nameKey, neutronSN.Name) {
		return true
	}

	if !filters.Match(cidrKey, neutronSN.Cidr) {
		return true
	}

	return false
}

// Create new subnet for given network.
func (s *Subnet) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	// TODO(pawel.zadrozny) validate if CIDR version is equal to ip_version neutron_plugin_db.py:1585
	virtualNetwork, err := getVirtualNetworkByID(ctx, rp, s.NetworkID)
	if err != nil {
		return nil, newSubnetError(networkNotFound, "failed to fetch network: %+v", err)
	}

	networkIpam, err := s.getNetworkIpam(ctx, rp, virtualNetwork)
	if err != nil {
		return nil, newSubnetError(badRequest, "failed to fetch network ipam: %+v", err)
	}

	err = s.createOrUpdateVirtualNetworkIpamRefs(ctx, rp, virtualNetwork, networkIpam)
	if err != nil {
		return nil, newSubnetError(badRequest, "failed to update network ipam refs: %+v", err)
	}

	// read subnet again to get updated values for gw, etc.
	virtualNetwork, err = getVirtualNetworkByID(ctx, rp, s.NetworkID)
	if err != nil {
		return nil, newSubnetError(networkNotFound, "failed to fetch network: %+v", err)
	}

	// get subnet data processed by the api
	ipamS := virtualNetwork.GetIpamSubnets().Find(func(sub *models.IpamSubnetType) bool {
		return sub.GetSubnet().CIDR() == s.Cidr
	})
	if ipamS == nil {
		return nil, newSubnetError(
			internalServerError,
			"subnet '%s' create failed for virtual network: '%s'", s.Cidr, virtualNetwork.GetUUID(),
		)
	}

	return subnetVncToNeutron(virtualNetwork, ipamS), nil
}

// Update specific subnet.
func (s *Subnet) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	if err := s.prevalidateBeforeUpdate(rp.FieldMask); err != nil {
		return nil, err
	}

	virtualNetworks, err := collectVNsUsingKV(ctx, rp, []string{id})
	if err != nil {
		return nil, newSubnetError(
			subnetNotFound,
			fmt.Sprintf("failed to fetch networks: %+v", err),
		)
	}

	if len(virtualNetworks) == 0 {
		return nil, newSubnetError(
			subnetNotFound,
			fmt.Sprintf("no Virtual Network with Subnet uuid: %+v", id),
		)
	}

	vn := findVirtualNetworkWithSubnet(id, virtualNetworks)
	if vn == nil {
		return nil, newSubnetError(
			subnetNotFound,
			fmt.Sprintf("no Virtual Network with Subnet uuid: %+v", id),
		)
	}

	ipam := findNetworkIpamRefWithSubnet(id, vn.NetworkIpamRefs)
	subnet := ipam.FindSubnet(id)

	updateSubnet(subnet, s, rp.FieldMask)

	subnet.LastModified = basemodels.ToVNCTime(time.Now().UTC())

	rp.WriteService.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
		VirtualNetwork: vn,
		FieldMask: types.FieldMask{
			Paths: []string{models.VirtualNetworkFieldNetworkIpamRefs},
		},
	})
	return subnetVncToNeutron(vn, subnet), nil
}

func updateSubnet(origin *models.IpamSubnetType, data *Subnet, fm types.FieldMask) {
	if basemodels.FieldMaskContains(&fm, subnetFieldName) {
		origin.SubnetName = data.Name
	}
	// Should we check this situation? Will it ever happen?
	if basemodels.FieldMaskContains(&fm, subnetFieldGatewayIP) {
		origin.DefaultGateway = data.GatewayIP
	}
	if basemodels.FieldMaskContains(&fm, subnetFieldEnableDHCP) {
		origin.EnableDHCP = data.EnableDHCP
	}
	if basemodels.FieldMaskContains(&fm, subnetFieldDNSNameservers) {
		if len(data.DNSNameservers) != 0 {
			origin.DHCPOptionList = &models.DhcpOptionsListType{
				DHCPOption: []*models.DhcpOptionType{
					{
						DHCPOptionName:  "6",
						DHCPOptionValue: strings.Join(data.DNSNameservers, " "),
					},
				},
			}
		} else {
			origin.DHCPOptionList = &models.DhcpOptionsListType{}
		}
	}
	if basemodels.FieldMaskContains(&fm, subnetFieldHostRoutes) {
		hostRoutes := make([]*RouteTableType, 0, len(data.HostRoutes))
		for i, hr := range data.HostRoutes {
			hostRoutes[i].Destination = hr.Destination
			hostRoutes[i].Nexthop = hr.Nexthop
		}
		// TODO: Handle apply subnet host routes situation.

		// Should make them nil if there is no host routes?
		origin.HostRoutes = neutronHostRoutesToModelsHostRoutes(hostRoutes)
	}
}

func neutronHostRoutesToModelsHostRoutes(hr []*RouteTableType) *models.RouteTableType {
	rts := make([]*models.RouteType, 0, len(hr))
	for i, r := range hr {
		rts[i].Prefix = r.Destination
		rts[i].NextHop = r.Nexthop
	}
	return &models.RouteTableType{
		Route: rts,
	}
}

func (s *Subnet) prevalidateBeforeUpdate(fm types.FieldMask) error {
	// should we check if they are not empty/nil like python?
	if basemodels.FieldMaskContains(&fm, subnetFieldGatewayIP) {
		return newSubnetError(badRequest, "update of gateway is not supported")
	}

	if basemodels.FieldMaskContains(&fm, subnetFieldAllocationPools) {
		return newSubnetError(badRequest, "update of allocation_pools is not allowed")
	}

	return nil
}

// Delete subnet with specified id.
func (s *Subnet) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	virtualNetworks, err := collectVNsUsingKV(ctx, rp, []string{id})
	if err != nil {
		return nil, newSubnetError(
			subnetNotFound,
			fmt.Sprintf("failed to fetch networks: %+v", err),
		)
	}

	if len(virtualNetworks) == 0 {
		return nil, newSubnetError(
			subnetNotFound,
			fmt.Sprintf("no Virtual Network with Subnet uuid: %+v", id),
		)
	}

	vn := findVirtualNetworkWithSubnet(id, virtualNetworks)
	if vn == nil {
		return nil, newSubnetError(
			subnetNotFound,
			fmt.Sprintf("no Virtual Network with Subnet uuid: %+v", id),
		)
	}

	ipam := findNetworkIpamRefWithSubnet(id, vn.NetworkIpamRefs)
	ipam.RemoveSubnet(id)

	_, err = rp.WriteService.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
		VirtualNetwork: vn,
		FieldMask: types.FieldMask{
			Paths: []string{models.VirtualNetworkFieldNetworkIpamRefs},
		},
	})
	return nil, err
}

func findVirtualNetworkWithSubnet(subnetID string, vns []*models.VirtualNetwork) *models.VirtualNetwork {
	for _, vn := range vns {
		if ipam := findNetworkIpamRefWithSubnet(subnetID, vn.GetNetworkIpamRefs()); ipam != nil {
			return vn
		}
	}
	return nil
}

func findNetworkIpamRefWithSubnet(
	subnetID string, refs []*models.VirtualNetworkNetworkIpamRef,
) *models.VirtualNetworkNetworkIpamRef {
	for _, ref := range refs {
		s := models.IpamSubnets{
			Subnets: ref.Attr.IpamSubnets,
		}
		found := s.Find(func(i *models.IpamSubnetType) bool {
			return i.SubnetUUID == subnetID
		})
		if found != nil {
			return ref
		}
	}
	return nil
}

func (s *Subnet) createOrUpdateVirtualNetworkIpamRefs(
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
		vn.AddNetworkIpamRef(&models.VirtualNetworkNetworkIpamRef{
			To: ni.GetFQName(),
			Attr: &models.VnSubnetsType{
				HostRoutes:  subnetType.GetHostRoutes(),
				IpamSubnets: []*models.IpamSubnetType{subnetType},
			},
		})
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

func (s *Subnet) getNetworkIpam(
	ctx context.Context,
	rp RequestParameters,
	vn *models.VirtualNetwork,
) (*models.NetworkIpam, error) {
	if s.IpamFQName != "" {
		n := models.MakeNetworkIpam()
		n.FQName = strings.Split(s.IpamFQName, "-")
		return n, nil
	}

	parentFQName := vn.GetFQName()[:len(vn.GetFQName())-1]
	ipamFQName := basemodels.ChildFQName(parentFQName, defaultNetworkIpamName)

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

	networkIpams := networkIpamRes.GetNetworkIpams()
	if len(networkIpams) == 0 {
		n := models.MakeNetworkIpam()
		n.FQName = defaultNetworkIpamFQName
		n.Name = defaultNetworkIpamName
		return n, nil
	}

	return networkIpams[0], nil

}

func (s *Subnet) ipamSubnetType() (*models.IpamSubnetType, error) {
	subnet, err := s.toVnc()
	if err != nil {
		return nil, err
	}

	defaultGateway, err := s.gatewayToVnc()
	if err != nil {
		return nil, err
	}

	return &models.IpamSubnetType{
		SubnetUUID:       subnet.GetUUID(),
		SubnetName:       subnet.GetName(),
		Subnet:           subnet.GetSubnetIPPrefix(),
		EnableDHCP:       s.EnableDHCP,
		AddrFromStart:    true,
		DHCPOptionList:   s.dhcpOptionListToVnc(),
		DefaultGateway:   defaultGateway,
		DNSServerAddress: s.dnsServerAddressToVnc(),
		AllocationPools:  s.allocationPoolType(),
		HostRoutes:       s.routeTableType(),
		Created:          s.CreatedAt,
		LastModified:     s.UpdatedAt,
	}, nil
}

func (s *Subnet) toVnc() (*models.Subnet, error) {
	subnetIPPrefix, err := s.subnetTypeToVnc()
	if err != nil {
		return nil, err
	}

	subnet := models.MakeSubnet()
	subnet.Name = s.Name
	subnet.ParentUUID = s.NetworkID
	subnet.SubnetIPPrefix = subnetIPPrefix

	return subnet, nil
}

// dhcpOptionListToVnc converts Neutron request to DHCP options list type VNC format.
func (s *Subnet) dhcpOptionListToVnc() *models.DhcpOptionsListType {
	if len(s.DNSNameservers) == 0 {
		return nil
	}

	var optVal []string
	for _, nameserver := range s.DNSNameservers {
		optVal = append(optVal, nameserver)
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

// gatewayToVnc converts Neutron request to Gateway VNC format.
func (s *Subnet) gatewayToVnc() (string, error) {
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

// dnsServerAddressToVnc converts Neutron request to DNS server address VNC format.
func (s *Subnet) dnsServerAddressToVnc() string {
	if strictCompliance {
		return "0.0.0.0"
	}

	if len(s.DNSNameservers) == 0 {
		return s.GatewayIP
	}

	return s.DNSNameservers[0]
}

// allocationPoolType converts Neutron request to allocation pools VNC format.
func (s *Subnet) allocationPoolType() []*models.AllocationPoolType {
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

// routeTableType converts Neutron request to  host routes VNC format.
func (s *Subnet) routeTableType() *models.RouteTableType {
	// TODO - not needed for ping by CREATE
	return nil
}

// subnetTypeToVnc converts Neutron request to subnet type VNC format.
func (s *Subnet) subnetTypeToVnc() (*models.SubnetType, error) {
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

func subnetVncToNeutron(vn *models.VirtualNetwork, subnetVnc *models.IpamSubnetType) *SubnetResponse {
	subnet := &SubnetResponse{
		ID:         subnetVnc.GetSubnetUUID(),
		Name:       subnetVnc.GetSubnetName(),
		TenantID:   contrailUUIDToNeutronID(vn.GetParentUUID()),
		NetworkID:  vn.GetUUID(),
		EnableDHCP: subnetVnc.GetEnableDHCP(),
		Shared:     vn.GetIsShared() || (vn.GetPerms2() != nil && len(vn.GetPerms2().GetShare()) > 0),
		CreatedAt:  subnetVnc.GetCreated(),
		UpdatedAt:  subnetVnc.GetLastModified(),
	}

	subnet.CIDRFromVnc(subnetVnc.GetSubnet())
	subnet.GatewayFromVnc(subnetVnc.GetDefaultGateway())
	subnet.HostRoutesFromVnc(subnetVnc.GetHostRoutes())

	subnet.DNSNameServersFromVnc(subnetVnc.GetDHCPOptionList())
	subnet.DNSServerAddressFromVnc(subnetVnc.GetDNSServerAddress())

	ipamHasSubnet := subnetVnc.GetSubnet() != nil
	subnet.AllocationPoolsFromVnc(subnetVnc.GetAllocationPools(), ipamHasSubnet)

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
	s.DNSNameservers = make([]string, 0)
	if dhcpOptions == nil {
		return
	}
	splitter := regexp.MustCompile("[^\\s]+")
	for _, opt := range dhcpOptions.GetDHCPOption() {
		if opt.GetDHCPOptionName() == "6" {
			dnsServers := splitter.FindAllString(opt.GetDHCPOptionValue(), -1)
			for _, dnsServer := range dnsServers {
				s.DNSNameservers = append(s.DNSNameservers, dnsServer)
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

func getVirtualNetworkByID(ctx context.Context, rp RequestParameters, nvID string) (*models.VirtualNetwork, error) {
	virtualNetworkRequest := &services.GetVirtualNetworkRequest{ID: nvID}
	virtualNetworkResponse, err := rp.ReadService.GetVirtualNetwork(ctx, virtualNetworkRequest)
	if err != nil {
		return nil, err
	}
	return virtualNetworkResponse.GetVirtualNetwork(), nil
}
