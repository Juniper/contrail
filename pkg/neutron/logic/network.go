package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	isAdminFieldKey = "is_admin"

	netStatusActive = "ACTIVE"
	netStatusDown   = "DOWN"

	// TODO(pawel.zadrozny) check if this config is still required or can be removed
	contrailExtensionsEnabled = true

	permsRX   = 5
	permsRWX  = 7
	permsNone = 0
)

// UnmarshalJSON unmarshals json into network.
func (n *Network) UnmarshalJSON(data []byte) error {
	type alias Network
	obj := struct {
		*alias
		Policys interface{} `json:"policys"`
	}{alias: (*alias)(n)}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	if policys, ok := obj.Policys.([]interface{}); ok {
		for _, policy := range policys {
			if p, ok := policy.([]string); ok {
				n.Policys = append(n.Policys, p)
			} else {
				n.Policys = [][]string{}
				break
			}
		}
	}
	return nil
}

// ApplyMap applies map onto network.
func (n *Network) ApplyMap(m map[string]interface{}) error {
	_, ok := m[NetworkFieldPolicys].(string)
	if ok {
		delete(m, NetworkFieldPolicys)
	}
	type alias Network
	return format.ApplyMap(m, (*alias)(n))
}

// Create logic
func (n *Network) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	vncNet, err := n.toVncForCreate(ctx, rp)
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

	nn := makeNetworkResponse(rp, vnResp.GetVirtualNetwork(), OperationCreate)
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

	return makeNetworkResponse(rp, vnRes.GetVirtualNetwork(), OperationRead), nil
}

// Delete logic
func (n *Network) Delete(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	if _, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: id}); err != nil {
		if !errutil.IsNotFound(err) {
			return nil, err
		}
		return &NetworkResponse{}, nil
	}

	fippRes, err := rp.ReadService.ListFloatingIPPool(ctx, &services.ListFloatingIPPoolRequest{
		Spec: &baseservices.ListSpec{
			ParentUUIDs: []string{id},
			Detail:      true,
		},
	})
	if err != nil {
		return nil, err
	}

	for _, fipp := range fippRes.GetFloatingIPPools() {
		if err = n.deleteAssociatedFloatingIPsAndPools(ctx, rp, fipp); err != nil {
			return nil, newNeutronError(networkInUse, errorFields{
				"net_id": id,
				"msg":    err.Error(),
			})
		}
	}

	_, err = rp.WriteService.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{
		ID: id,
	})
	if errutil.IsConflict(err) {
		return nil, newNeutronError(networkInUse, errorFields{
			"net_id": id,
			"msg":    err.Error(),
		})
	}
	if err != nil {
		return nil, err
	}

	return &NetworkResponse{}, nil
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

// ReadCount logic
func (n *Network) ReadCount(
	ctx context.Context, rp RequestParameters, filters Filters,
) (Response, error) {
	if vnCount, err := n.getCountByTenantIDs(ctx, rp, filters); err != nil {
		return &NetworkResponse{}, err
	} else if vnCount != nil {
		return *vnCount, nil
	}

	rp.RequestContext = RequestContext{}
	nn, err := n.ReadAll(ctx, rp, filters, Fields{})
	if err != nil {
		return nil, err
	}
	ns, ok := nn.([]*NetworkResponse)
	if !ok {
		return 0, nil
	}

	return len(ns), nil
}

// Update logic
func (n *Network) Update(
	ctx context.Context, rp RequestParameters, id string,
) (Response, error) {
	oldVNRes, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: id,
	})
	if err != nil {
		return nil, newNeutronError(networkNotFound, errorFields{
			"net_id": id,
			"msg":    err.Error(),
		})
	}
	n.ID = id
	oldVN := oldVNRes.GetVirtualNetwork()
	oldShared := oldVN.GetIsShared()
	oldRouterExternal := oldVN.GetRouterExternal()
	var newVN *models.VirtualNetwork
	fm := types.FieldMask{}
	newVN, err = n.toVncForUpdate(ctx, rp, oldVN, &fm)
	if err != nil {
		return nil, err
	}

	if err = n.validateSharedChange(ctx, rp, newVN, oldShared); err != nil {
		return nil, err
	}

	if err = n.validateRouterExternalChange(ctx, rp, newVN, oldRouterExternal, &fm); err != nil {
		return nil, err
	}

	if _, err = rp.WriteService.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
		VirtualNetwork: newVN,
		FieldMask:      fm,
	}); err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "network",
			"msg":      err.Error(),
		})
	}

	return makeNetworkResponse(rp, newVN, OperationRead), nil
}

func (n *Network) getCountByTenantIDs(
	ctx context.Context, rp RequestParameters, filters Filters,
) (*int64, error) {
	// Only if one filter is provided and it is a Tenant ID, we can read virtual networks directly from DB
	// using this filter as parent uuids of vns
	if len(filters) != 1 || !filters.HaveKeys(tenantIDKey) {
		return nil, nil
	}

	var pUUIDs []string
	for _, tenantID := range filters[tenantIDKey] {
		uuid, err := neutronIDToVncUUID(tenantID)
		if err != nil {
			return nil, err
		}
		pUUIDs = append(pUUIDs, uuid)
	}

	vnCount, err := rp.ReadService.ListVirtualNetwork(ctx, &services.ListVirtualNetworkRequest{
		Spec: &baseservices.ListSpec{
			Count:       true,
			ParentUUIDs: pUUIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	return &vnCount.VirtualNetworkCount, nil
}

func (n *Network) validateSharedChange(
	ctx context.Context,
	rp RequestParameters,
	newVN *models.VirtualNetwork,
	oldShared bool,
) error {
	if oldShared && !newVN.GetIsShared() {
		extPortsAssociated, err := n.checkIfExternalPortsAreAssociated(ctx, rp, newVN)
		if err != nil {
			return err
		}
		if extPortsAssociated {
			return newNeutronError(invalidSharedSetting, errorFields{
				"network": newVN.GetDisplayName(),
			})
		}
	}
	return nil
}

func (n *Network) checkIfExternalPortsAreAssociated(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork,
) (bool, error) {
	for _, vmi := range vn.GetVirtualMachineInterfaceBackRefs() {
		vmiRes, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
			ID: vmi.GetUUID(),
		})
		if err != nil {
			return false, err
		}

		if vmiRes.GetVirtualMachineInterface().GetParentType() == models.KindProject &&
			vmiRes.GetVirtualMachineInterface().GetParentUUID() != vn.GetParentUUID() {
			return true, nil
		}
	}
	return false, nil
}

func (n *Network) validateRouterExternalChange(
	ctx context.Context,
	rp RequestParameters,
	newVN *models.VirtualNetwork,
	oldRouterExternal bool,
	fm *types.FieldMask,
) error {
	if newVN.GetRouterExternal() == oldRouterExternal {
		return nil
	}

	if newVN.GetRouterExternal() {
		err := n.createFloatingIPPool(ctx, rp, newVN)
		if err != nil {
			return err
		}
	} else {
		for _, fipp := range newVN.GetFloatingIPPools() {
			if err := n.deleteAssociatedFloatingIPsAndPools(ctx, rp, fipp); err != nil {
				return newNeutronError(networkInUse, errorFields{
					"net_id": newVN.GetUUID(),
					"msg":    err.Error(),
				})
			}
		}
	}
	return nil
}

type listReq struct {
	ParentID string
	Filters  []*baseservices.Filter
	Detail   bool
	Count    bool
	ObjUUIDs []string
}

func (n *Network) toVncForCreate(
	ctx context.Context, rp RequestParameters,
) (vncNet *models.VirtualNetwork, err error) {
	vncNet = models.MakeVirtualNetwork()
	vncNet.Name = n.Name
	vncNet.ParentType = models.KindProject
	vncNet.AddressAllocationMode = models.UserDefinedSubnetOnly //TODO find place where it should be set
	vncNet.RouterExternal = n.RouterExternal
	vncNet.ParentUUID, err = neutronIDToVncUUID(n.TenantID)
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

	vncNet.UUID = n.ID
	if err = n.updateVnc(ctx, rp, vncNet, nil); err != nil {
		return nil, err
	}

	return vncNet, nil
}

func (n *Network) toVncForUpdate(
	ctx context.Context, rp RequestParameters, vncNet *models.VirtualNetwork, fm *types.FieldMask,
) (*models.VirtualNetwork, error) {
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(NetworkFieldRouterExternal)) {
		vncNet.RouterExternal = n.RouterExternal
		basemodels.FieldMaskAppend(fm, models.VirtualNetworkFieldRouterExternal)
	}

	if n.RouterExternal {
		vncNet.Perms2.GlobalAccess = permsRX
	} else {
		vncNet.Perms2.GlobalAccess = permsNone
	}
	basemodels.FieldMaskAppend(fm, models.VirtualNetworkFieldPerms2, models.PermType2FieldGlobalAccess)

	if err := n.updateVnc(ctx, rp, vncNet, fm); err != nil {
		return nil, err
	}
	return vncNet, nil

}

func (n *Network) updateVnc(
	ctx context.Context, rp RequestParameters, vncNet *models.VirtualNetwork, vncFm *types.FieldMask,
) error {
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(nameKey)) {
		vncNet.DisplayName = n.Name
		basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldDisplayName)
	}

	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(sharedKey)) {
		vncNet.IsShared = n.Shared
		basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldIsShared)
	}

	if len(n.ProviderPhysicalNetwork) > 0 || n.ProviderSegmentationID != 0 {
		vncNet.ProviderProperties = &models.ProviderDetails{
			PhysicalNetwork: n.ProviderPhysicalNetwork,
			SegmentationID:  n.ProviderSegmentationID,
		}
		basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldProviderProperties)
	}

	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(NetworkFieldAdminStateUp)) {
		vncNet.IDPerms.Enable = n.AdminStateUp
		basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldIDPerms, models.IdPermsTypeFieldEnable)
	}

	if err := n.setVncRefs(ctx, rp, vncNet, vncFm); err != nil {
		return errors.Wrapf(err, "failed to set refs for virtual network(%v)", vncNet.GetUUID())
	}

	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(NetworkFieldPortSecurityEnabled)) {
		vncNet.PortSecurityEnabled = n.PortSecurityEnabled
		basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldPortSecurityEnabled)
	}

	if len(n.Description) > 0 {
		vncNet.IDPerms.Description = n.Description
		basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldIDPerms, models.IdPermsTypeFieldDescription)
	}

	return nil
}

func (n *Network) setVncRefs(
	ctx context.Context,
	rp RequestParameters,
	vncNet *models.VirtualNetwork,
	vncFm *types.FieldMask,
) error {
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(NetworkFieldPolicys)) {
		n.createNetworkPolicyRef(vncNet, vncFm)
	}
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(NetworkFieldRouteTable)) &&
		len(n.RouteTable) > 0 {
		if err := n.createRouteTableRef(ctx, rp, vncNet, vncFm); err != nil {
			return err
		}
	}
	return nil
}

func (n *Network) deleteAssociatedFloatingIPsAndPools(
	ctx context.Context, rp RequestParameters, fipp *models.FloatingIPPool,
) error {
	fips, err := n.listAssociatedFloatingIPs(ctx, rp, fipp)
	if err != nil {
		return err
	}

	for _, fip := range fips {
		if len(fip.GetVirtualMachineInterfaceRefs()) > 0 {
			return errors.Errorf("floating IP(uuid: %v) is assosiated with a port", fip.GetUUID())
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

func (n *Network) listAssociatedFloatingIPs(
	ctx context.Context, rp RequestParameters, fipp *models.FloatingIPPool,
) ([]*models.FloatingIP, error) {
	uuids := fipp.GetFloatingIPsUUIDs()
	if len(uuids) == 0 {
		return nil, nil
	}

	response, err := rp.ReadService.ListFloatingIP(ctx, &services.ListFloatingIPRequest{
		Spec: &baseservices.ListSpec{
			ObjectUUIDs: uuids,
			Fields: []string{
				models.FloatingIPFieldUUID,
				models.FloatingIPFieldVirtualMachineInterfaceRefs,
			},
		},
	})
	return response.GetFloatingIPs(), err
}

func (n *Network) createFloatingIPPool(
	ctx context.Context, rp RequestParameters, vncNet *models.VirtualNetwork,
) error {
	fipp := models.MakeFloatingIPPool()
	fipp.Name = models.KindFloatingIPPool
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

func (n *Network) createNetworkPolicyRef(
	vncNet *models.VirtualNetwork, vncFm *types.FieldMask,
) {
	vncNet.NetworkPolicyRefs = []*models.VirtualNetworkNetworkPolicyRef{}
	for i, policy := range n.Policys {
		vncNet.AddNetworkPolicyRef(&models.VirtualNetworkNetworkPolicyRef{
			To: policy,
			Attr: &models.VirtualNetworkPolicyType{
				Sequence: &models.SequenceType{
					Major: int64(i),
					Minor: 0,
				},
			},
		})
	}
	basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldNetworkPolicyRefs)
}

func (n *Network) createRouteTableRef(
	ctx context.Context, rp RequestParameters, vncNet *models.VirtualNetwork, vncFm *types.FieldMask,
) error {

	rtUUID, err := rp.FQNameToIDService.FQNameToID(ctx, &services.FQNameToIDRequest{
		FQName: n.RouteTable,
		Type:   models.KindRouteTable,
	})
	//This error is returned by old neutron plugin and marked with TODO to create separate error for this case
	if errutil.IsNotFound(err) {
		return newNeutronError(networkNotFound, errorFields{
			"net_id": vncNet.GetUUID(),
			"msg":    err.Error(),
		})
	}
	if err != nil {
		return err
	}

	vncNet.AddRouteTableRef(&models.VirtualNetworkRouteTableRef{
		UUID: rtUUID.GetUUID(),
		To:   n.RouteTable,
	})
	basemodels.FieldMaskAppend(vncFm, models.VirtualNetworkFieldRouteTableRefs)

	return nil
}

func (n *Network) collectVirtualNetworks(
	ctx context.Context, rp RequestParameters, filters Filters,
) ([]*models.VirtualNetwork, error) {
	if len(filters) > 0 && filters.HaveKeys(idKey) {
		return collectWithoutPrune(ctx, rp, filters[idKey])
	}

	if basemodels.FieldMaskContains(&rp.FieldMask, buildContextPath(isAdminFieldKey)) && !rp.RequestContext.IsAdmin {
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
	if filters.HaveKeys(nameKey) {
		return collectNetworkForTenant(ctx, rp, filters, rp.RequestContext.Tenant)
	}

	var req listReq
	if filters.HaveKeys(sharedKey) || filters.HaveKeys(routerExternalKey) {
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

	addDBFilter(req, models.VirtualNetworkFieldIsShared, []string{"true"}, false)
	sharedVNs, err := listNetworksForProject(ctx, rp, req)
	if err != nil {
		return nil, err
	}
	vns = append(vns, sharedVNs...)

	addDBFilter(req, NetworkFieldRouterExternal, []string{"true"}, true)
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
	if filters.HaveKeys(tenantIDKey) {
		return collectUsingTenantID(ctx, rp, filters)
	}

	req := &listReq{}
	if filters.HaveKeys(nameKey) {
		return listNetworksForProject(ctx, rp, req)
	}

	if filters.HaveKeys(sharedKey) || filters.HaveKeys(routerExternalKey) {
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
		if req.ParentID, err = neutronIDToVncUUID(p); err != nil {
			return nil, err
		}
		vns, err = listNetworksForProject(ctx, rp, req)
		if err != nil {
			return nil, nil
		}
		collectedVNs = append(collectedVNs, vns...)
	}

	if filters.HaveKeys(routerExternalKey) {
		req.ParentID = ""
		addDBFilter(req, NetworkFieldRouterExternal, filters[routerExternalKey], true)
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
	if filters.HaveKeys(sharedKey) {
		addDBFilter(req, models.VirtualNetworkFieldIsShared, filters[sharedKey], false)
	}

	if filters.HaveKeys(routerExternalKey) {
		addDBFilter(req, NetworkFieldRouterExternal, filters[routerExternalKey], false)
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
	if !filters.Match(fqNameKey, vn.GetFQName()...) {
		return false
	}

	if !filters.Match(nameKey, vn.GetName()) &&
		!filters.Match(nameKey, vn.GetDisplayName()) {
		return false
	}

	isShared := vn.GetIsShared()
	if vn.GetPerms2() != nil && isSharedWithTenant(&rp.RequestContext, vn.GetPerms2().GetShare()) {
		isShared = true
	}

	if !filters.Match(sharedKey, strconv.FormatBool(isShared)) {
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

		nn := makeNetworkResponse(rp, vn, OperationReadAll)
		if nn == nil {
			continue
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
