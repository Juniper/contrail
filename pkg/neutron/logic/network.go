package logic

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	isShared                 = "is_shared"
	routerExternal           = "router_external"
	neutronIDKey             = "neutronIDKey"
	neutronNameKey           = "neutronNameKey"
	neutronFQNameKey         = "fq_name"
	neutronSharedKey         = "shared"
	neutronRouterExternalKey = "router:external"
	neutronTenantIDKey       = "tenant_id"
	netStatusActive          = "ACTIVE"
	netStatusDown            = "DOWN"
)

type ListReq struct {
	ParentID string
	Filters  []*baseservices.Filter
	Detail   bool
	Count    bool
}

// Create logic
func (n *Network) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}

// ReadAll logic
func (n *Network) ReadAll(ctx context.Context, rp RequestParameters, filters Filters, fields Fields) (Response, error) {
	var resp Response

	vncNets, err := n.collectVirtualNetworks(ctx, rp, filters)
	if err != nil {
		return nil, err
	}

	resp = convertVNsToNeutronResponse(ctx, rp, filters, vncNets)
	return resp, nil
}

func (n *Network) collectVirtualNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if filtersHas(filters, neutronIDKey) {
		return collectWithoutPrune(ctx, rp, filters[neutronIDKey])
	}

	if !rp.RequestContext.IsAdmin {
		return collectNonAdminNetworks(ctx, rp, filters)
	}

	if filters != nil {
		return collectFilteredAdminNetworks(ctx, rp, filters)
	}

	return collectAllNetworksByProjects(ctx, rp, filters)
}

func collectWithoutPrune(ctx context.Context, rp RequestParameters, uuids []string) ([]*models.VirtualNetwork, error) {
	var vns []*models.VirtualNetwork
	for _, uuid := range uuids {
		vnResp, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: uuid})
		if err != nil {
			return nil, err
		}
		vns = append(vns, vnResp.GetVirtualNetwork())
	}

	return vns, nil
}

func collectNonAdminNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if filtersHas(filters, neutronNameKey) {
		return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant)
	}

	var req ListReq
	if filtersHas(filters, neutronSharedKey) || filtersHas(filters, neutronRouterExternalKey) {
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, &req)
	}

	return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant) // filters nil??
}

func collectNetworkForTenant(ctx context.Context, rp RequestParameters, filters Filters, tenant string,
) ([]*models.VirtualNetwork, error) {
	var vns []*models.VirtualNetwork
	req := &ListReq{
		ParentID: tenant,
	}

	tenantVNs, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}
	vns = append(vns, tenantVNs...)

	if len(filters) == 0 {
		return vns, nil
	}

	addDBFilter(req, isShared, []string{"true"}, false)
	sharedVNs, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}
	vns = append(vns, sharedVNs...)

	addDBFilter(req, routerExternal, []string{"true"}, true)
	rtExtVNs, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}
	vns = append(vns, rtExtVNs...)

	return vns, nil
}

func addDBFilter(req *ListReq, key string, values []string, clearFirst bool) {
	if clearFirst {
		req.Filters = []*baseservices.Filter{}
	}
	filter := baseservices.Filter{Key: key, Values: values}
	req.Filters = append(req.Filters, &filter)
}

func collectFilteredAdminNetworks(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	if filtersHas(filters, neutronTenantIDKey) {
		return collectUsingTenantID(ctx, rp, filters)
	}

	req := &ListReq{}
	if filtersHas(filters, neutronNameKey) {
		return listNetworksForProject(ctx, rp, req)
	}

	if filtersHas(filters, neutronSharedKey) || filtersHas(filters, neutronRouterExternalKey) {
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, req)
	}

	return nil, nil
}

func collectUsingTenantID(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	projects := validateProjectByID(rp, filters)

	var collectedVNs, vns []*models.VirtualNetwork
	var err error
	req := &ListReq{}

	for _, p := range projects {
		req.ParentID = p
		vns, err = listNetworksForProject(ctx, rp, req)
		if err != nil {
			return nil, nil
		}
		collectedVNs = append(collectedVNs, vns...)
	}

	if filtersHas(filters, neutronRouterExternalKey) {
		req.ParentID = ""
		addDBFilter(req, routerExternal, filters[neutronRouterExternalKey], true)
		vns, err = listNetworksForProject(ctx, rp, req)
		if err != nil {
			return nil, nil
		}
		collectedVNs = append(collectedVNs, vns...)
	}

	return collectedVNs, nil
}

func listNetworksForProject(ctx context.Context, rp RequestParameters, req *ListReq) ([]*models.VirtualNetwork, error) {
	sp := prepareVirtualNetworkListSpec(req)
	vnResp, err := rp.ReadService.ListVirtualNetwork(
		ctx,
		&services.ListVirtualNetworkRequest{
			Spec: sp,
		},
	)
	if err != nil {
		return nil, err
	}

	return vnResp.GetVirtualNetworks(), nil
}

func collectSharedOrRouterExtNetworks(ctx context.Context, rp RequestParameters, filters Filters, req *ListReq) ([]*models.VirtualNetwork, error) {
	if filtersHas(filters, neutronSharedKey) {
		addDBFilter(req, isShared, filters[neutronSharedKey], false)
	}

	if filtersHas(filters, neutronRouterExternalKey) {
		addDBFilter(req, routerExternal, filters[neutronRouterExternalKey], false)
	}

	return listNetworksForProject(ctx, rp, req)
}

func collectAllNetworksByProjects(ctx context.Context, rp RequestParameters, filters Filters) ([]*models.VirtualNetwork, error) {
	req := &ListReq{
		Detail: true,
	}

	return listNetworksForProject(ctx, rp, req)
}

func convertVNsToNeutronResponse(ctx context.Context, rp RequestParameters, filters Filters, vns []*models.VirtualNetwork) Response {
	var nns []*NetworkResponse
	for _, vn := range vns {
		if containsNetworkWithUUID(nns, vn.GetUUID()) {
			continue
		}

		if !checkFilterValue(filters, neutronFQNameKey, vn.GetFQName()...) {
			continue
		}

		if !checkFilterValue(filters, neutronNameKey, vn.GetName()) && !checkFilterValue(filters, neutronNameKey, vn.GetDisplayName()) {
			continue
		}

		isShared := vn.GetIsShared()
		if vn.GetPerms2() != nil && isSharedWithTenant(&rp.RequestContext, vn.GetPerms2().GetShare()) {
			isShared = true
		}

		if !checkFilterValue(filters, neutronSharedKey, strconv.FormatBool(isShared)) {
			continue
		}

		nn := networkVNCToNeutron(ctx, rp, vn, "LIST")
		if nn == nil {
			continue
		}

		nns = append(nns, nn)
	}
	return nns
}

func networkVNCToNeutron(ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork, operation string) *NetworkResponse {
	parentNeutronUUID := contrailUUIDToNeutronID(vn.GetParentUUID())
	nn := &NetworkResponse{
		ID:                      vn.GetUUID(),
		Name:                    vn.GetDisplayName(),
		TenantID:                parentNeutronUUID,
		ProjectID:               parentNeutronUUID,
		AdminStateUp:            vn.GetIDPerms().GetEnable(),
		Shared:                  vn.GetIsShared(),
		Status:                  netStatusDown,
		RouterExternal:          vn.GetRouterExternal(),
		PortSecurityEnabled:     vn.GetPortSecurityEnabled(),
		Description:             vn.GetIDPerms().GetDescription(),
		CreatedAt:               vn.GetIDPerms().GetCreated(),
		UpdatedAt:               vn.GetIDPerms().GetLastModified(),
		ProviderPhysicalNetwork: vn.GetProviderProperties().GetPhysicalNetwork(),
		ProviderSegmentationID:  vn.GetProviderProperties().GetSegmentationID(),
	}

	//TODO: fqname contrail extension
	nn.FQName = vn.GetFQName()
	if vn.GetDisplayName() == "" {
		nn.Name = vn.FQName[len(vn.FQName)-1]
	}

	if !nn.Shared && (vn.GetPerms2() != nil && isSharedWithTenant(&rp.RequestContext, vn.GetPerms2().GetShare())) {
		nn.Shared = true
	}

	if vn.GetIDPerms().GetEnable() {
		nn.Status = netStatusActive
	}

	if operation == "READ" || operation == "LIST" {
		//TODO: handle network policy refs(contrail extension)
	}

	//TODO: handle route target refs(contrail extension)

	ipamRefs := vn.GetNetworkIpamRefs()
	//TODO: handle subnet_ipam(contrail extension)
	for _, ipam := range ipamRefs {
		subnets := ipam.GetAttr().GetIpamSubnets()
		for _, ipamSubnet := range subnets {
			//TODO: use subnetVNCToNeutron
			nn.Subnets = append(nn.Subnets, ipamSubnet.GetSubnetUUID())
			//TODO: add subnet_ipam to schema and fill this field here
		}
	}

	return nn
}

func isSharedWithTenant(rc *RequestContext, sharedList []*models.ShareType) bool {
	if rc != nil && len(sharedList) > 0 {
		for _, t := range sharedList {
			if rc.Tenant == t.Tenant {
				return true
			}
		}
	}
	return false
}

func validateProjectByID(rp RequestParameters, filters Filters) []string {
	if !rp.RequestContext.IsAdmin {
		return []string{rp.RequestContext.Tenant} //TODO: handle tables in context
	}

	return filters[neutronTenantIDKey]
}

func containsNetworkWithUUID(nns []*NetworkResponse, uuid string) bool {
	for _, nn := range nns {
		if nn.ID == uuid {
			return true
		}
	}
	return false
}

func prepareVirtualNetworkListSpec(req *ListReq) *baseservices.ListSpec {
	return &baseservices.ListSpec{
		Filters:     req.Filters,
		Detail:      true,
		Count:       req.Count,
		Shared:      true,
		ParentUUIDs: []string{req.ParentID},
	}
}
