package logic

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	netStatusActive = "ACTIVE"
	netStatusDown   = "DOWN"

	contrailExtensionsEnabled = true
)

func (n *Network) updateVnc(vncNet *models.VirtualNetwork) error {
	vncNet.RouterExternal = n.RouterExternal
	if n.RouterExternal {
		vncNet.Perms2 = &models.PermType2{ /* TODO - not needed for ping by CREATE */ }
	}
	// TODO: For Operation == UPDATE do: https://github.com/Juniper/contrail-controller/blob/0b6850b55a63280bfb339113d24bd24c953cf145/src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py#L1432
	vncNet.IsShared = n.Shared
	vncNet.UUID = n.ID

	vncNet.DisplayName = n.Name

	// TODO: Handle ProviderProperties L:1441-1445
	if len(n.ProviderPhysicalNetwork) > 0 || len(n.ProviderSegmentationID) > 0 {
		intSegID, _ := strconv.Atoi(n.ProviderSegmentationID)
		segID := int64(intSegID)
		//PhysicalNetwork string - not needed for ping by CREATE
		//SegmentationID  int64 - not needed for ping by CREATE
		vncNet.ProviderProperties = &models.ProviderDetails{
			PhysicalNetwork: n.ProviderPhysicalNetwork,
			// TODO: Need to check type of SegmentationID in neutron dumps
			SegmentationID: segID,
		}
	}

	vncNet.IDPerms = &models.IdPermsType{Enable: n.AdminStateUp} // TODO: This is a bug for operation UPDATE when Admin state up is not set in request.

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
	err := n.updateVnc(vncNet)
	if err != nil {
		return nil, err
	}

	return vncNet, nil
}

func checkVnSharedWithTennant(vn *models.VirtualNetwork) bool {
	return false
}

func (r *NetworkResponse) setResponseRefs(vn *models.VirtualNetwork, policyRefs bool) {
	if refs := vn.GetNetworkPolicyRefs(); policyRefs && len(refs) > 0 {
		// TODO: handle policy refs - not needed for ping by CREATE
		// This should be set only for oper READ or LIST => L1535
	}
	// TODO: Handle field route_table (L1545) - not needed for ping by CREATE

	// TODO: Handle Subnets (L1552) - not needed for ping by CREATE
	r.Subnets = make([]string, 0)
}

func makeResponseFromVnc(vn *models.VirtualNetwork, policyRefs bool) Response {
	// TODO: handle data processed in extra_dict in python code
	resp := &NetworkResponse{
		ID:             vn.GetUUID(),
		TenantID:       vn.ParentUUID,
		ProjectID:      vn.ParentUUID,
		AdminStateUp:   vn.GetIDPerms().Enable,
		Shared:         false,
		Status:         netStatusDown,
		RouterExternal: vn.GetRouterExternal(),
		CreatedAt:      vn.GetIDPerms().Created,
		UpdatedAt:      vn.GetIDPerms().LastModified,
	}
	if vn.GetDisplayName() != "" {
		resp.Name = vn.GetDisplayName()
	} else {
		fqName := vn.GetFQName()
		resp.Name = fqName[len(fqName)-1]
	}
	if vn.GetIsShared() || vn.GetPerms2() != nil || checkVnSharedWithTennant(vn) {
		resp.Shared = true
	}
	if vn.GetIDPerms().Enable {
		resp.Status = netStatusActive
	}
	if prop := vn.GetProviderProperties(); prop != nil {
		// TODO: Missing fields provider:physical_network and provider:segmentation_id, have in python
	}
	if contrailExtensionsEnabled {
		resp.setResponseRefs(vn, policyRefs)
	}

	if descr := vn.GetIDPerms().Description; len(descr) > 0 {
		resp.Description = descr
	}

	return resp
}

// Create logic
func (n *Network) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	vncNet, err := n.toVnc()
	if err != nil {
		return nil, err
	}

	vn, err := rp.WriteService.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: vncNet})
	if err != nil {
		return nil, err
	}

	return makeResponseFromVnc(vn.VirtualNetwork, false), nil
}
