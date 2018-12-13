package logic

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
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

// TODO: Needs error?
func makeResponseFromVnc(vn *models.VirtualNetwork) Response {
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

	vn, err := n.writeService.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: vncNet})
	if err != nil {
		return nil, err
	}

	return MakeResponseFromVnc(vn)
}

func main() {
	nn := Network{}
	n, e := nn.ToVnc()
}
