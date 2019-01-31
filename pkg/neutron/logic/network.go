package logic

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"

	"errors"

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

	permsRX   = 5
	permsRWX  = 7
	permsNone = 0
)

// Create logic
func (n *Network) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	vncNet, err := n.toVnc(ctx, rp)
	if err != nil {
		return nil, err
	}

	vnResp, err := rp.WriteService.CreateVirtualNetwork(ctx,
		&services.CreateVirtualNetworkRequest{VirtualNetwork: vncNet})
	if err != nil {
		return nil, err
	}

	if vncNet.GetRouterExternal() {
		err = n.createFloatingIPPool(ctx, rp, vncNet)
		if err != nil {
			return nil, err
		}
	}

	nn := makeNetworkResponse(rp, vnResp.GetVirtualNetwork(), "CREATE")
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

	return makeNetworkResponse(rp, vnRes.GetVirtualNetwork(), "READ"), nil
}

// Delete logic
func (n *Network) Delete(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	if _, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: id}); err != nil {
		if !errutil.IsNotFound(err) {
			return &NetworkResponse{}, err
		}
		return &NetworkResponse{}, nil
	}
	fippRes, err := rp.ReadService.ListFloatingIPPool(ctx, &services.ListFloatingIPPoolRequest{
		Spec: &baseservices.ListSpec{
			ParentUUIDs: []string{id},
			Fields: []string{
				"floating-ips", // TODO: change to constant
			},
		},
	})
	if err != nil {
		return &NetworkResponse{}, err
	}

	for _, fipp := range fippRes.GetFloatingIPPools() {
		if err = n.processFipp(ctx, rp, fipp); err != nil {
			return &NetworkResponse{}, newNeutronError(networkInUse, errorFields{
				"net_id": id,
			})
		}
	}

	_, err = rp.WriteService.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{
		ID: id,
	})

	return &NetworkResponse{}, err
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

func (n *Network) updateVnc(
	ctx context.Context, rp RequestParameters, vncNet *models.VirtualNetwork,
) error {
	vncNet.IsShared = n.Shared
	vncNet.DisplayName = n.Name

	if len(n.ProviderPhysicalNetwork) > 0 || n.ProviderSegmentationID != 0 {

		vncNet.ProviderProperties = &models.ProviderDetails{
			PhysicalNetwork: n.ProviderPhysicalNetwork,
			SegmentationID:  n.ProviderSegmentationID,
		}
	}
	// TODO: This is a bug for operation UPDATE when Admin state up is not set in request.
	vncNet.IDPerms = &models.IdPermsType{Enable: n.AdminStateUp}

	if len(n.Policys) > 0 {
		//TODO handle policy refs and verify type of 'policys' field with multiple items
	}
	if len(n.RouteTable) > 0 {
		if err := n.handleRouteTableRef(ctx, rp, vncNet); err != nil {
			return err
		}
	}

	operation := rp.RequestContext.Operation
	if operation == "CREATE" || operation == "UPDATE" {
		vncNet.PortSecurityEnabled = n.PortSecurityEnabled
	}

	if len(n.Description) > 0 {
		vncNet.IDPerms.Description = n.Description
	}

	return nil
}

func (n *Network) toVncForCreate(
	ctx context.Context, rp RequestParameters,
) (vncNet *models.VirtualNetwork, err error) {
	vncNet = models.MakeVirtualNetwork()
	vncNet.Name = n.Name
	vncNet.ParentType = models.KindProject
	vncNet.AddressAllocationMode = models.UserDefinedSubnetOnly //TODO find place where it should be set
	vncNet.RouterExternal = n.RouterExternal
	vncNet.ParentUUID, err = neutronIDToContrailUUID((n.TenantID))
	if err != nil {
		return vncNet, err
	}

	vncNet.IDPerms = &models.IdPermsType{Enable: true}
	if n.RouterExternal {
		vncNet.Perms2 = &models.PermType2{
			Owner:        vncNet.ParentUUID,
			OwnerAccess:  permsRWX,
			GlobalAccess: permsRX,
		}
	}

	//TODO consider contrailuuid or neutron
	vncNet.UUID, err = neutronIDToContrailUUID(n.ID)
	return vncNet, err
}

func (n *Network) toVncForUpdate(
	ctx context.Context, rp RequestParameters,
) (vncNet *models.VirtualNetwork, err error) {
	vncNetRes, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: n.ID})
	if err != nil {
		return nil, err
	}
	vncNet = vncNetRes.GetVirtualNetwork()
	vncNet.RouterExternal = n.RouterExternal

	if n.RouterExternal {
		vncNet.Perms2.GlobalAccess = permsRX
	} else {
		vncNet.Perms2.GlobalAccess = permsNone
	}
	return vncNet, nil

}

func (n *Network) toVnc(ctx context.Context, rp RequestParameters) (vncNet *models.VirtualNetwork, err error) {
	operation := rp.RequestContext.Operation
	if operation == "CREATE" {
		vncNet, err = n.toVncForCreate(ctx, rp)

	} else {
		vncNet, err = n.toVncForUpdate(ctx, rp)
	}
	if err != nil {
		return nil, err
	}
	err = n.updateVnc(ctx, rp, vncNet)
	if err != nil {
		return nil, err
	}

	return vncNet, nil
}

func (n *Network) processFipp(
	ctx context.Context, rp RequestParameters, fipp *models.FloatingIPPool,
) error {
	uuids := fipp.GetFloatingIPsUUIDs()
	fips, err := rp.ReadService.ListFloatingIP(ctx, &services.ListFloatingIPRequest{
		Spec: &baseservices.ListSpec{
			ObjectUUIDs: uuids,
			Fields: []string{
				models.FloatingIPFieldVirtualMachineInterfaceRefs,
			},
			Detail: false,
		},
	})
	if err != nil {
		return err
	}

	for _, fip := range fips.GetFloatingIPs() {
		if len(fip.GetVirtualMachineInterfaceRefs()) > 0 {
			return errors.New("floating IP is assosiated with port")
		}
		_, err = rp.WriteService.DeleteFloatingIP(ctx, &services.DeleteFloatingIPRequest{
			ID: fip.GetUUID(),
		})
		if err != nil {
			return err
		}
	}
	_, err = rp.WriteService.DeleteFloatingIPPool(ctx, &services.DeleteFloatingIPPoolRequest{
		ID: fipp.GetUUID(),
	})
	return err
}

func (n *Network) createFloatingIPPool(
	ctx context.Context, rp RequestParameters, vncNet *models.VirtualNetwork,
) error {
	fipp := models.MakeFloatingIPPool()
	fipp.Name = models.KindFloatingIPPool //TODO check correctness of name
	fipp.ParentUUID = vncNet.GetUUID()
	fipp.ParentType = models.KindVirtualNetwork
	fipp.Perms2 = &models.PermType2{
		Owner:        vncNet.GetUUID(),
		OwnerAccess:  permsRWX,
		GlobalAccess: permsRX,
	}

	_, err := rp.WriteService.CreateFloatingIPPool(
		ctx,
		&services.CreateFloatingIPPoolRequest{
			FloatingIPPool: fipp,
		},
	)
	return err
}

func (n *Network) handleRouteTableRef(
	ctx context.Context, rp RequestParameters, vncNet *models.VirtualNetwork,
) error {
	resp, err := rp.FQNameToIDService.FQNameToID(ctx, &services.FQNameToIDRequest{
		FQName: n.RouteTable,
		Type:   models.KindRouteTable,
	})
	if err != nil {
		return err
	}

	_, err = rp.WriteService.CreateVirtualNetworkRouteTableRef(
		ctx,
		&services.CreateVirtualNetworkRouteTableRefRequest{
			ID: vncNet.GetUUID(),
			VirtualNetworkRouteTableRef: &models.VirtualNetworkRouteTableRef{
				UUID: resp.GetUUID(),
			},
		},
	)
	return err
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

		nn := makeNetworkResponse(rp, vn, "LIST")
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
