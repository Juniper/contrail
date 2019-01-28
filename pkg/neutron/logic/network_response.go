package logic

import "github.com/Juniper/contrail/pkg/models"

func makeNetworkResponse(rp RequestParameters, vn *models.VirtualNetwork) *NetworkResponse {
	parentNeutronUUID := ContrailUUIDToNeutronID(vn.GetParentUUID())
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
		Subnets:                 []string{},
		SubnetIpam:              []*SubnetIpam{},
	}

	if contrailExtensionsEnabled {
		nn.FQName = vn.GetFQName()
	}

	if vn.GetDisplayName() == "" {
		nn.Name = vn.FQName[len(vn.FQName)-1]
	}

	if !nn.Shared && (vn.GetPerms2() != nil && isSharedWithTenant(&rp.RequestContext, vn.GetPerms2().GetShare())) {
		nn.Shared = true
	}

	if vn.GetIDPerms().GetEnable() {
		nn.Status = netStatusActive
	}

	if prop := vn.GetProviderProperties(); prop != nil {
		// TODO: Missing fields provider:physical_network and provider:segmentation_id, have in python
	}

	nn.setSubnets(vn)
	// TODO: Handle field route_table (L1545) - not needed for ping
	return nn
}

func (r *NetworkResponse) setSubnets(vn *models.VirtualNetwork) {
	ipamRefs := vn.GetNetworkIpamRefs()
	for _, ipam := range ipamRefs {
		subnets := ipam.GetAttr().GetIpamSubnets()
		for _, ipamSubnet := range subnets {
			sn := subnetVncToNeutron(vn, ipamSubnet)
			r.Subnets = append(r.Subnets, sn.ID)

			if contrailExtensionsEnabled {
				r.SubnetIpam = append(r.SubnetIpam, &SubnetIpam{SubnetCidr: sn.Cidr, IpamFQName: ipam.GetTo()})

			}
		}
	}
}

func (r *NetworkResponse) setResponseRefs(vn *models.VirtualNetwork) {
	if len(vn.GetNetworkPolicyRefs()) > 0 {
		// TODO: handle policy refs - not needed for ping by CREATE
		// This should be set only for oper READ or LIST => L1535
	}
}
