package logic

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	isShared       = "is_shared"
	routerExternal = "router_external"

	netStatusActive = "ACTIVE"
	netStatusDown   = "DOWN"

	// TODO(pawel.zadrozny) check if this config is still required or can be removed
	contrailExtensionsEnabled = true
)

// Create logic
func (n *Network) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	vncNet, err := n.toVnc()
	if err != nil {
		return nil, err
	}

	vnResp, err := rp.WriteService.CreateVirtualNetwork(ctx,
		&services.CreateVirtualNetworkRequest{VirtualNetwork: vncNet})
	if err != nil {
		return nil, err
	}

	nn := makeNetworkResponse(rp, vnResp.GetVirtualNetwork())
	if contrailExtensionsEnabled {
		nn.setResponseRefs(vnResp.GetVirtualNetwork())
	}

	return nn, nil
}

// Read logic
func (n *Network) Read(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	// TODO: Fix performance
	//       This implementation reads network every time Read is called,
	//       which is a heavy operation. Therefore DB cache needs to be used here.

	vnRes, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: id,
	})
	if errutil.IsNotFound(err) {
		return nil, newNeutronError(networkNotFound, errorFields{
			"net_id": id,
		})
	} else if err != nil {
		return nil, err
	}

	return makeNetworkResponse(rp, vnRes.GetVirtualNetwork()), nil
}

// ReadAll logic
func (n *Network) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO: Fix performance
	//       This implementation reads all networks every time ReadAll is called,
	//       which is a heavy operation. Therefore DB cache needs to be used here.

	vncNets, err := n.collectVirtualNetworks(ctx, rp, filters)
	if err != nil {
		return nil, err
	}

	return convertVNsToNeutronResponse(rp, filters, vncNets), nil
}

type listReq struct {
	ParentID string
	Filters  []*baseservices.Filter
	Detail   bool
	Count    bool
	ObjUUIDs []string
}

func (n *Network) updateVnc(vncNet *models.VirtualNetwork) error {
	var err error

	vncNet.RouterExternal = n.RouterExternal
	if n.RouterExternal {
		vncNet.Perms2 = &models.PermType2{ /* TODO - not needed for ping by CREATE */ }
	}
	// TODO: For Operation == UPDATE do:
	// https://github.com/Juniper/contrail-controller/
	// blob/0b6850b55a63280bfb339113d24bd24c953cf145/src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py#L1432
	vncNet.IsShared = n.Shared
	if vncNet.UUID, err = neutronIDToContrailUUID(n.ID); err != nil {
		return err
	}

	if vncNet.ParentUUID, err = neutronIDToContrailUUID((n.ProjectID)); err != nil {
		return err
	}

	vncNet.DisplayName = n.Name

	// TODO: Handle ProviderProperties L:1441-1445
	if len(n.ProviderPhysicalNetwork) > 0 || len(n.ProviderSegmentationID) > 0 {
		var intSegID int
		if intSegID, err = strconv.Atoi(n.ProviderSegmentationID); err != nil {
			return err
		}

		segID := int64(intSegID)
		//PhysicalNetwork string - not needed for ping by CREATE
		//SegmentationID  int64 - not needed for ping by CREATE
		vncNet.ProviderProperties = &models.ProviderDetails{
			PhysicalNetwork: n.ProviderPhysicalNetwork,
			// TODO: Need to check type of SegmentationID in neutron dumps
			SegmentationID: segID,
		}
	}
	// TODO: This is a bug for operation UPDATE when Admin state up is not set in request.
	vncNet.IDPerms = &models.IdPermsType{Enable: n.AdminStateUp}

	// Handle policys L:1452-1467 - not needed for ping by CREATE
	//TODO: Verify type of 'policys' field with multiple items, currently string but in pytaon array

	// Handle route table L:1469-1478
	if len(n.RouteTable) > 0 {
		/*
			resp := n.FQNameService.FQNameToIDService(services.FQNameToIDRequest{
				FQName: n.RouteTable,
				Type:   models.KindRouteTable,
			})*/
		// TODO: Read route_table by fq_name and set to vncNet - not needed for ping by CREATE
	}

	vncNet.PortSecurityEnabled = n.PortSecurityEnabled

	if len(n.Description) > 0 {
		vncNet.IDPerms.Description = n.Description
	}

	return nil
}

func (n *Network) toVnc() (*models.VirtualNetwork, error) {
	vncNet := models.MakeVirtualNetwork()
	vncNet.Name = n.Name
	vncNet.ParentType = models.KindProject
	vncNet.IDPerms = &models.IdPermsType{Enable: true}
	vncNet.AddressAllocationMode = models.UserDefinedSubnetOnly
	err := n.updateVnc(vncNet)
	if err != nil {
		return nil, err
	}

	return vncNet, nil
}

func (n *Network) collectVirtualNetworks(
	ctx context.Context, rp RequestParameters, filters Filters,
) ([]*models.VirtualNetwork, error) {
	if len(filters) > 0 && filters.haveKeys(idKey) {
		return collectWithoutPrune(ctx, rp, filters[idKey])
	}

	if !rp.RequestContext.IsAdmin {
		return collectNonAdminNetworks(ctx, rp, filters)
	}

	if len(filters) > 0 {
		return collectFilteredAdminNetworks(ctx, rp, filters)
	}

	return collectAllNetworksByProjects(ctx, rp)
}

func collectWithoutPrune(
	ctx context.Context, rp RequestParameters, uuids []string,
) ([]*models.VirtualNetwork, error) {
	var vns []*models.VirtualNetwork
	for _, uuid := range uuids {
		vnResp, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: uuid})
		if errutil.IsNotFound(err) {
			return nil, newNeutronError(networkNotFound, errorFields{
				"net_id": uuid,
			})
		} else if err != nil {
			return nil, err
		}
		vns = append(vns, vnResp.GetVirtualNetwork())
	}

	return vns, nil
}

func collectNonAdminNetworks(
	ctx context.Context, rp RequestParameters, filters Filters,
) ([]*models.VirtualNetwork, error) {
	if filters.haveKeys(nameKey) {
		return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant)
	}

	var req listReq
	if filters.haveKeys(sharedKey) || filters.haveKeys(routerExternalKey) {
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, &req)
	}

	return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant)
}

func collectNetworkForTenant(
	ctx context.Context, rp RequestParameters, filters Filters, tenant string,
) ([]*models.VirtualNetwork, error) {
	var vns []*models.VirtualNetwork
	req := &listReq{
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

func addDBFilter(req *listReq, key string, values []string, clearFirst bool) {
	if clearFirst {
		req.Filters = []*baseservices.Filter{}
	}
	filter := baseservices.Filter{Key: key, Values: values}
	req.Filters = append(req.Filters, &filter)
}

func collectFilteredAdminNetworks(
	ctx context.Context, rp RequestParameters, filters Filters,
) ([]*models.VirtualNetwork, error) {
	if filters.haveKeys(tenantIDKey) {
		return collectUsingTenantID(ctx, rp, filters)
	}

	req := &listReq{}
	if filters.haveKeys(nameKey) {
		return listNetworksForProject(ctx, rp, req)
	}

	if filters.haveKeys(sharedKey) || filters.haveKeys(routerExternalKey) {
		return collectSharedOrRouterExtNetworks(ctx, rp, filters, req)
	}

	return nil, nil
}

func collectUsingTenantID(
	ctx context.Context, rp RequestParameters, filters Filters,
) ([]*models.VirtualNetwork, error) {
	projects := validateProjectByID(rp, filters)

	var collectedVNs, vns []*models.VirtualNetwork
	var err error
	req := &listReq{}

	for _, p := range projects {
		if req.ParentID, err = neutronIDToContrailUUID(p); err != nil {
			return nil, err
		}
		vns, err = listNetworksForProject(ctx, rp, req)
		if err != nil {
			return nil, nil
		}
		collectedVNs = append(collectedVNs, vns...)
	}

	if filters.haveKeys(routerExternalKey) {
		req.ParentID = ""
		addDBFilter(req, routerExternal, filters[routerExternalKey], true)
		vns, err = listNetworksForProject(ctx, rp, req)
		if err != nil {
			return nil, nil
		}
		collectedVNs = append(collectedVNs, vns...)
	}

	return collectedVNs, nil
}

func listNetworksForProject(
	ctx context.Context, rp RequestParameters, req *listReq,
) ([]*models.VirtualNetwork, error) {
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

func collectSharedOrRouterExtNetworks(
	ctx context.Context, rp RequestParameters, filters Filters, req *listReq,
) ([]*models.VirtualNetwork, error) {
	if req == nil {
		req = &listReq{}
	}
	if filters.haveKeys(sharedKey) {
		addDBFilter(req, isShared, filters[sharedKey], false)
	}

	if filters.haveKeys(routerExternalKey) {
		addDBFilter(req, routerExternal, filters[routerExternalKey], false)
	}

	return listNetworksForProject(ctx, rp, req)
}

func collectAllNetworksByProjects(
	ctx context.Context, rp RequestParameters) ([]*models.VirtualNetwork, error,
) {
	req := &listReq{
		Detail: true,
	}

	return listNetworksForProject(ctx, rp, req)
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

	return filters[tenantIDKey]
}

func containsNetworkWithUUID(nns []*NetworkResponse, uuid string) bool {
	for _, nn := range nns {
		if nn.ID == uuid {
			return true
		}
	}
	return false
}

func checkIfVNMatchFilters(rp RequestParameters, filters Filters, vn *models.VirtualNetwork) bool {
	if !filters.checkValue(fqNameKey, vn.GetFQName()...) {
		return false
	}

	if !filters.checkValue(nameKey, vn.GetName()) &&
		!filters.checkValue(nameKey, vn.GetDisplayName()) {
		return false
	}

	isShared := vn.GetIsShared()
	if vn.GetPerms2() != nil && isSharedWithTenant(&rp.RequestContext, vn.GetPerms2().GetShare()) {
		isShared = true
	}

	if !filters.checkValue(sharedKey, strconv.FormatBool(isShared)) {
		return false
	}

	return true
}

func convertVNsToNeutronResponse(
	rp RequestParameters, filters Filters, vns []*models.VirtualNetwork,
) Response {
	nns := []*NetworkResponse{}
	for _, vn := range vns {
		if containsNetworkWithUUID(nns, vn.GetUUID()) {
			continue
		}

		if !checkIfVNMatchFilters(rp, filters, vn) {
			continue
		}

		nn := makeNetworkResponse(rp, vn)
		if nn == nil {
			continue
		}

		if contrailExtensionsEnabled {
			nn.setResponseRefs(vn)
		}

		nns = append(nns, nn)
	}
	return nns
}

func prepareVirtualNetworkListSpec(req *listReq) *baseservices.ListSpec {
	var pUUIDs []string
	if req.ParentID != "" {
		pUUIDs = []string{
			req.ParentID,
		}
	}

	return &baseservices.ListSpec{
		Filters:     req.Filters,
		Detail:      true,
		Count:       req.Count,
		Shared:      false, //TODO: Change to true after JBE-495
		ParentUUIDs: pUUIDs,
		ObjectUUIDs: req.ObjUUIDs,
		Fields: []string{
			models.VirtualNetworkFieldUUID,
			models.VirtualNetworkFieldDisplayName,
			models.VirtualNetworkFieldParentUUID,
			models.VirtualNetworkFieldFQName,
			models.VirtualNetworkFieldIDPerms,
			models.VirtualNetworkFieldPerms2,
			models.VirtualNetworkFieldIsShared,
			models.VirtualNetworkFieldPortSecurityEnabled,
			models.VirtualNetworkFieldProviderProperties,
			models.VirtualNetworkFieldRouterExternal,
			models.VirtualNetworkFieldNetworkIpamRefs,
			models.VirtualNetworkFieldNetworkPolicyRefs,
		},
	}
}
