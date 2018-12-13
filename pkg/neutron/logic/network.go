package logic

import (
	"github.com/Juniper/contrail/pkg/models"
)

func (n *Network) UpdateVnc(vncNet *models.VirtualNetwork) error {
	// TODO: id_perms
	vncNet.RouterExternal = n.RouterExternal
	if n.RouterExternal {
		vncNet.Perms2 = &models.PermType2{ /* TODO */ }
	}
	// TODO: For Operation == UPDATE do: https://github.com/Juniper/contrail-controller/blob/0b6850b55a63280bfb339113d24bd24c953cf145/src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py#L1432
	vncNet.IsShared = n.Shared
	vncNet.UUID = n.ID

	vncNet.DisplayName = n.Name

	// TODO: Handle ProviderProperties L:1441-1445

	vncNet.IDPerms = models.IDPermsType{Enabled: n.AdminStateUp} // TODO: This is a bug for operation UPDATE when Admin state up is not set in request.

	// Handle policys L:1452-1467

	// Handle route table L:1469-1478

	if len(n.Description) {
		vncNet.IDPerms.Description = n.Description
	}

	return nil
}

func (n *Network) ToVnc() (*models.VirtualNetwork, error) {
	net := models.MakeVirtualNetwork()
	net.Name = n.Name
	net.ParentType = models.KindProject
	// TODO: id_perms
	err := n.UpdateVnc(net)

	return net, nil
}

// Create logic
func (n *Network) Create(ctx RequestParameters) (Response, error) {
	// TODO: FIX ME!
	vncNet, err := n.ToVnc()

	// TODO: return correct object
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}

func main() {
	nn := Network{}
	n, e := nn.ToVnc()
}
