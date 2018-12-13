package logic

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	NetStatusActive = "ACTIVE"
	NetStatusDown   = "DOWN"
)

func (n *Network) updateVnc(vncNet *models.VirtualNetwork) error {
	vncNet.RouterExternal = n.RouterExternal
	if n.RouterExternal {
		vncNet.Perms2 = &models.PermType2{ /* TODO */ }
	}
	// TODO: For Operation == UPDATE do: https://github.com/Juniper/contrail-controller/blob/0b6850b55a63280bfb339113d24bd24c953cf145/src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py#L1432
	vncNet.IsShared = n.Shared
	vncNet.UUID = n.ID

	vncNet.DisplayName = n.Name

	// TODO: Handle ProviderProperties L:1441-1445
	if len(n.ProviderPhysicalNetwork) > 0 || len(n.ProviderSegmentationID) > 0 {
		intSegID, _ := strconv.Atoi(n.ProviderSegmentationID)
		segID := int64(intSegID)
		//PhysicalNetwork string
		//SegmentationID  int64
		vncNet.ProviderProperties = &models.ProviderDetails{
			PhysicalNetwork: n.ProviderPhysicalNetwork,
			// TODO: Need to check type of SegmentationID in neutron dumps
			SegmentationID: segID,
		}
	}

	vncNet.IDPerms = models.IDPermsType{Enabled: n.AdminStateUp} // TODO: This is a bug for operation UPDATE when Admin state up is not set in request.

	// Handle policys L:1452-1467
	//TODO: Verify type of 'policys' field with multiple items, currently string but in pytaon array

	// Handle route table L:1469-1478
	if len(n.RouteTable) > 0 {
		resp := n.FQNameService.FQNameToIDService(services.FQNameToIDRequest{
			FQName: n.RouteTable,
			Type:   models.KindRouteTable,
		})
		// TODO: Read route_table by fq_name and set to vncNet
	}

	if len(n.Description) {
		vncNet.IDPerms.Description = n.Description
	}

	return nil
}

func (n *Network) toVnc() (*models.VirtualNetwork, error) {
	vncNet := models.MakeVirtualNetwork()
	vncNet.Name = n.Name
	vncNet.ParentType = models.KindProject
	vncNet.IDPerms = models.IDPermsType{Enabled: true}
	err := n.UpdateVnc(vncNet)

	return vncNet, nil
}

func checkVnSharedWithTennant(vn *models.VierualNetwork) bool {
	return false
}

// TODO: Needs error?
func makeResponseFromVnc(vn *models.VirtualNetwork) Response {
	// TODO: handle data processed in extra_dict in python code
	resp := &NetworkResponse{
		ID: vn.GetUUID(),
		// TODO: field TenantID
		// TODO: field ProjectID
		AdminStateUp:   vn.GetIDPerms().Enabled,
		Shared:         false,
		Status:         NetStatusDown,
		RouterExternal: vn.GetRouterExternal(),
		CreatedAt:      vn.GetIDPerms().Created,
		UpdatedAt:      vn.GetIDPerms().LastModified,
	}
	if nv.GetDisplayName() != "" {
		resp.Name = nv.GetDisplayName()
	} else {
		fqName := nv.GetFQName()
		resp.Name = fqName[len(fqName)-1]
	}
	if vn.GetIsShared() || vn.GetPerms2() != nil || checkVnSharedWithTennant(vn) {
		resp.Shared = true
	}
	if vn.GetIDPerms().Enabled {
		resp.Status = NetStatusActive
	}
	if prop := vn.GetProviderProperties(); prop != nil {
		// TODO: Missing fields provider:physical_network and provider:segmentation_id, have in python
	}
	// TODO: Handle code in L1535 (with extra_dict)

	// TODO: Missong field route_table (L1545)

	// TODO: Subnets handled by extra_dict ....

	// TODO: extra_dict is used only if _contrail_extensions_enabled is true in python

	if descr := vn.GetIDPerms().Description; len(descr) > 0 {
		resp.Description = descr
	}

	// TODO: Remove stub below
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}

// Create logic
func (n *Network) Create(rp RequestParameters) (Response, error) {
	ctx := context.Background()
	vncNet, err := n.ToVnc()
	if err != nil {
		return nil, err
	}

	vn, err := n.WriteService.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: vncNet})
	if err != nil {
		return nil, err
	}

	return MakeResponseFromVnc(vn)
}

func main() {
	nn := Network{}
	n, e := nn.ToVnc()
}
