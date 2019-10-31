package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/format"
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

	interfaceRouteTablePrefix = "NEUTRON_IFACE_RT"

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

// UnmarshalJSON unmarshals json into subnet.
func (s *Subnet) UnmarshalJSON(data []byte) error {
	type alias Subnet
	obj := struct {
		*alias
		IpamFQName interface{} `json:"ipam_fq_name"`
	}{alias: (*alias)(s)}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	if ipamFQName, ok := obj.IpamFQName.([]string); ok {
		s.IpamFQName = ipamFQName
	}
	return nil
}

// ApplyMap applies map onto subnet.
func (s *Subnet) ApplyMap(m map[string]interface{}) error {
	_, ok := m[SubnetFieldIpamFQName].(string)
	if ok {
		delete(m, SubnetFieldIpamFQName)
	}
	type alias Subnet
	return format.ApplyMap(m, (*alias)(s))
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

	addDBFilter(req, models.VirtualNetworkFieldIsShared, []string{"true"}, false)
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
		return nil, newSubnetError(networkNotFound, "failed to fetch network: %v", err)
	}

	networkIpam, err := s.getNetworkIpam(ctx, rp, virtualNetwork)
	if err != nil {
		return nil, newSubnetError(badRequest, "failed to fetch network ipam: %v", err)
	}

	err = s.createOrUpdateVirtualNetworkIpamRefs(ctx, rp, virtualNetwork, networkIpam)
	if err != nil {
		return nil, newSubnetError(badRequest, "failed to update network ipam refs: %v", err)
	}

	// read subnet again to get updated values for gw, etc.
	virtualNetwork, err = getVirtualNetworkByID(ctx, rp, s.NetworkID)
	if err != nil {
		return nil, newSubnetError(networkNotFound, "failed to fetch network: %v", err)
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
	if err := s.checkNotUpdatableFields(rp.FieldMask); err != nil {
		return nil, err
	}

	virtualNetworks, err := collectVNsUsingKV(ctx, rp, []string{id})
	if err != nil {
		return nil, newSubnetError(subnetNotFound, "failed to fetch networks: %v", err)
	}

	vn, err := findVirtualNetworkWithSubnet(id, virtualNetworks)
	if err != nil {
		return nil, err
	}

	subnet := findNetworkIpamRefWithSubnet(id, vn.NetworkIpamRefs).FindSubnetByID(id)

	err = s.applyChangesToVncSubnet(ctx, &rp, subnet, id, vn.GetUUID())
	if err != nil {
		return nil, err
	}
	subnet.LastModified = basemodels.ToVNCTime(time.Now())

	_, err = rp.WriteService.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
		VirtualNetwork: vn,
		FieldMask: types.FieldMask{
			Paths: []string{models.VirtualNetworkFieldNetworkIpamRefs},
		},
	})
	if err != nil {
		return nil, err
	}
	return subnetVncToNeutron(vn, subnet), nil
}

func (s *Subnet) checkNotUpdatableFields(fm types.FieldMask) error {
	if basemodels.FieldMaskContains(&fm, buildDataResourcePath(SubnetFieldGatewayIP)) {
		return newSubnetError(badRequest, "update of gateway is not supported")
	}

	if basemodels.FieldMaskContains(&fm, buildDataResourcePath(SubnetFieldAllocationPools)) {
		return newSubnetError(badRequest, "update of allocation_pools is not allowed")
	}

	return nil
}

func findVirtualNetworkWithSubnet(
	subnetID string, vns []*models.VirtualNetwork,
) (*models.VirtualNetwork, error) {
	for _, vn := range vns {
		if ipam := findNetworkIpamRefWithSubnet(subnetID, vn.GetNetworkIpamRefs()); ipam != nil {
			return vn, nil
		}
	}
	return nil, errors.Errorf("no Virtual Network with Subnet uuid: %v", subnetID)
}

func findNetworkIpamRefWithSubnet(
	subnetID string, refs []*models.VirtualNetworkNetworkIpamRef,
) *models.VirtualNetworkNetworkIpamRef {
	for _, ref := range refs {
		if ref.FindSubnetByID(subnetID) != nil {
			return ref
		}
	}
	return nil
}

func (s *Subnet) applyChangesToVncSubnet(
	ctx context.Context, rp *RequestParameters, vncSubnet *models.IpamSubnetType,
	subnetUUID, vnUUID string,
) error {
	fm := &rp.FieldMask
	if basemodels.FieldMaskContains(fm, buildDataResourcePath(SubnetFieldName)) {
		vncSubnet.SubnetName = s.Name
	}
	if basemodels.FieldMaskContains(fm, buildDataResourcePath(SubnetFieldEnableDHCP)) {
		vncSubnet.EnableDHCP = s.EnableDHCP
	}
	if basemodels.FieldMaskContains(fm, buildDataResourcePath(SubnetFieldDNSNameservers)) {
		vncSubnet.SetDNSNameservers(s.DNSNameservers)
	}
	if basemodels.FieldMaskContains(fm, buildDataResourcePath(SubnetFieldHostRoutes)) {
		if applySubnetHostRoutes() {
			cidr := vncSubnet.GetSubnet().CIDR()
			oldHostRoutes := vncHostRoutesToNeutronHostRoutes(vncSubnet.GetHostRoutes())
			err := portUpdateIfaceRouteTable(ctx, rp, vnUUID, cidr, subnetUUID, s.HostRoutes, oldHostRoutes)
			if err != nil {
				return err
			}
		}

		vncSubnet.HostRoutes = neutronHostRoutesToVncHostRoutes(s.HostRoutes)
	}
	return nil
}

// TODO: implement_apply_subnet_host_routes variable from neutron_plugin_db.py
func applySubnetHostRoutes() bool {
	return true
}

func vncHostRoutesToNeutronHostRoutes(rtt *models.RouteTableType) []*RouteTableType {
	if rtt == nil {
		return nil
	}
	rts := make([]*RouteTableType, len(rtt.Route))
	for i, r := range rtt.Route {
		rts[i] = &RouteTableType{}
		rts[i].Destination = r.Prefix
		rts[i].Nexthop = r.NextHop
	}
	return rts
}

func portUpdateIfaceRouteTable(
	ctx context.Context, rp *RequestParameters, vnUUID, subnetCIDR, subnetID string,
	newHR []*RouteTableType, oldHR []*RouteTableType,
) error {
	hostPrefixesToDelete, hostPrefixesToAdd, err := extractPrefixesRequiredForUpdate(newHR, oldHR, subnetCIDR)
	if err != nil {
		return err
	}

	if hostPrefixesToAdd.isEmpty() && hostPrefixesToDelete.isEmpty() {
		// nothing to do
		return nil
	}

	ipsResponse, err := rp.ReadService.ListInstanceIP(ctx, &services.ListInstanceIPRequest{
		Spec: &baseservices.ListSpec{
			Fields: []string{
				models.InstanceIPFieldInstanceIPAddress,
				models.InstanceIPFieldVirtualMachineInterfaceRefs,
			},
			BackRefUUIDs: []string{vnUUID},
		},
	})
	if err != nil {
		return err
	}

	ips := ipsResponse.GetInstanceIPs()
	if err = handleDeprecatedHostPrefixes(ctx, rp, ips, hostPrefixesToDelete, subnetID); err != nil {
		return err
	}

	return handleNewHostPrefixes(ctx, rp, ips, hostPrefixesToAdd, subnetID)
}

func extractPrefixesRequiredForUpdate(
	newHostRoutes, oldHostRoutes []*RouteTableType, subnetCIDR string,
) (hostPrefixes, hostPrefixes, error) {
	hostPrefixesToDelete, err := getHostPrefixes(oldHostRoutes, subnetCIDR)
	if err != nil {
		return nil, nil, err
	}

	hostPrefixesToAdd, err := getHostPrefixes(newHostRoutes, subnetCIDR)
	if err != nil {
		return nil, nil, err
	}

	for _, nexthop := range hostPrefixesToDelete.getIPAddresses() {
		if hostPrefixesToAdd.hasAnyDestinationsForIP(nexthop) {
			oldDestinations := hostPrefixesToDelete.getDestinationsForIP(nexthop)
			newDestinations := hostPrefixesToAdd.getDestinationsForIP(nexthop)
			if format.StringSetsEqual(oldDestinations, newDestinations) {
				// Nothing changed for this nexthop address so there is no need to change it.
				hostPrefixesToAdd.removeDestinationsForIP(nexthop)
			}
			// There are still destinations for this nexthop address so we don't want to delete it.
			hostPrefixesToDelete.removeDestinationsForIP(nexthop)
		}
	}

	return hostPrefixesToDelete, hostPrefixesToAdd, nil
}

func handleDeprecatedHostPrefixes(
	ctx context.Context, rp *RequestParameters, ips []*models.InstanceIP, prefixes hostPrefixes, subnetID string,
) error {
	for _, ip := range ips {
		addr := ip.GetInstanceIPAddress()
		if prefixes.hasAnyDestinationsForIP(addr) {
			if err := portRemoveIfaceRouteTable(ctx, rp, ip, subnetID); err != nil {
				return err
			}
		}
	}
	return nil
}

func portRemoveIfaceRouteTable(
	ctx context.Context, rp *RequestParameters, ip *models.InstanceIP, subnetUUID string,
) error {
	portRefs := ip.GetVirtualMachineInterfaceRefs()

	for _, port := range portRefs {
		vmiResp, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
			ID: port.GetUUID(),
		})
		if err != nil {
			return err
		}

		vmi := vmiResp.VirtualMachineInterface

		irtName := getInterfaceRouteTableName(subnetUUID, vmi.GetUUID())

		irt := vmi.FindInterfaceRouteTableRef(func(ref *models.VirtualMachineInterfaceInterfaceRouteTableRef) bool {
			return basemodels.FQNameToName(ref.To) == irtName
		})

		if irt == nil {
			continue
		}

		vmi.RemoveInterfaceRouteTableRef(irt)

		_, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
			VirtualMachineInterface: vmi,
			FieldMask: types.FieldMask{
				Paths: []string{models.VirtualMachineInterfaceFieldInterfaceRouteTableRefs},
			},
		})
		if err != nil {
			return err
		}

		_, err = rp.WriteService.DeleteInterfaceRouteTable(ctx, &services.DeleteInterfaceRouteTableRequest{
			ID: irt.UUID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func handleNewHostPrefixes(
	ctx context.Context, rp *RequestParameters, ips []*models.InstanceIP, prefixes hostPrefixes, subnetID string,
) error {
	for _, ip := range ips {
		addr := ip.GetInstanceIPAddress()
		if !prefixes.hasAnyDestinationsForIP(addr) {
			continue
		}
		vmiBackRefs := ip.GetVirtualMachineInterfaceRefs()
		for _, ref := range vmiBackRefs {
			vmi, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
				ID: ref.UUID,
			})
			if err != nil {
				return err
			}
			destinations := prefixes.getDestinationsForIP(addr)
			err = portAddIfaceRouteTable(ctx, rp, vmi.VirtualMachineInterface, subnetID, destinations)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func portAddIfaceRouteTable(
	ctx context.Context, rp *RequestParameters, vmi *models.VirtualMachineInterface,
	subnetUUID string, prefixes []string,
) error {

	irt, err := locateInterfaceRouteTableWithSubnet(ctx, rp, subnetUUID, vmi)
	if err != nil {
		return nil
	}

	irt.SetPrefixes(prefixes)

	_, err = rp.WriteService.UpdateInterfaceRouteTable(ctx, &services.UpdateInterfaceRouteTableRequest{
		InterfaceRouteTable: irt,
	})
	if err != nil {
		return err
	}

	vmi.AddInterfaceRouteTableRef(&models.VirtualMachineInterfaceInterfaceRouteTableRef{
		UUID: irt.UUID,
		To:   irt.FQName,
	})

	_, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})
	return err
}

func locateInterfaceRouteTableWithSubnet(
	ctx context.Context, rp *RequestParameters, subnetUUID string, vmi *models.VirtualMachineInterface,
) (*models.InterfaceRouteTable, error) {
	// TODO: Ensure that vmi object always contain valid fqname
	irtName := getInterfaceRouteTableName(subnetUUID, vmi.GetUUID())
	irtFQName := basemodels.ChildFQName(vmi.GetFQName(), irtName)

	if irt, err := findInterfaceRouteTable(ctx, rp, irtFQName); err != nil {
		return nil, err
	} else if irt != nil {
		return irt, nil
	}

	return createInterfaceRouteTable(ctx, rp, irtName, vmi.ParentUUID)
}

func findInterfaceRouteTable(
	ctx context.Context, rp *RequestParameters, FQName []string,
) (*models.InterfaceRouteTable, error) {
	irts, err := rp.ReadService.ListInterfaceRouteTable(ctx, &services.ListInterfaceRouteTableRequest{
		Spec: &baseservices.ListSpec{
			Filters: []*baseservices.Filter{
				{
					Key:    models.FabricFieldFQName,
					Values: FQName,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(irts.InterfaceRouteTables) > 0 {
		return nil, nil
	}
	return irts.InterfaceRouteTables[0], nil
}

func createInterfaceRouteTable(
	ctx context.Context, rp *RequestParameters, name string, parentUUID string,
) (*models.InterfaceRouteTable, error) {
	irt, err := rp.WriteService.CreateInterfaceRouteTable(ctx, &services.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: &models.InterfaceRouteTable{
			Name:       name,
			ParentType: models.KindProject,
			ParentUUID: parentUUID,
			InterfaceRouteTableRoutes: &models.RouteTableType{
				Route: []*models.RouteType{},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return irt.InterfaceRouteTable, nil
}

func neutronHostRoutesToVncHostRoutes(hr []*RouteTableType) *models.RouteTableType {
	rts := make([]*models.RouteType, len(hr))
	for i, r := range hr {
		rts[i] = &models.RouteType{}
		rts[i].Prefix = r.Destination
		rts[i].NextHop = r.Nexthop
	}
	return &models.RouteTableType{
		Route: rts,
	}
}

func getInterfaceRouteTableName(subnetUUID, vmiUUID string) string {
	return strings.Join([]string{interfaceRouteTablePrefix, subnetUUID, vmiUUID}, "_")
}

// Delete subnet with specified id.
func (s *Subnet) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vns, err := collectVNsUsingKV(ctx, rp, []string{id})
	if err != nil {
		return nil, newSubnetError(internalServerError, "failed to fetch networks: %v", err)
	}

	vn, err := findVirtualNetworkWithSubnet(id, vns)
	if err != nil {
		return nil, newSubnetError(subnetNotFound, "cannot find virtual network: %v", err)
	}

	findNetworkIpamRefWithSubnet(id, vn.NetworkIpamRefs).RemoveSubnet(id)

	_, err = rp.WriteService.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
		VirtualNetwork: vn,
		FieldMask: types.FieldMask{
			Paths: []string{models.VirtualNetworkFieldNetworkIpamRefs},
		},
	})
	if err != nil {
		return nil, newSubnetError(internalServerError, "cannot update virtual network: %v", err)
	}
	return &SubnetResponse{}, nil
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
	if len(s.IpamFQName) > 0 {
		n := models.MakeNetworkIpam()
		n.FQName = s.IpamFQName
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
	return neutronHostRoutesToVncHostRoutes(s.HostRoutes)
}

// subnetTypeToVnc converts Neutron request to subnet type VNC format.
func (s *Subnet) subnetTypeToVnc() (*models.SubnetType, error) {
	_, subnetPrefixIP, prefixLen, err := getIPPrefixAndPrefixLen(s.Cidr)
	if err != nil {
		return nil, err
	}

	return &models.SubnetType{IPPrefix: subnetPrefixIP, IPPrefixLen: prefixLen}, nil
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
		TenantID:   VncUUIDToNeutronID(vn.GetParentUUID()),
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
