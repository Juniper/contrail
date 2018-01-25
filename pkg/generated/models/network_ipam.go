package models

// NetworkIpam

// NetworkIpam
//proteus:generate
type NetworkIpam struct {
	UUID             string           `json:"uuid,omitempty"`
	ParentUUID       string           `json:"parent_uuid,omitempty"`
	ParentType       string           `json:"parent_type,omitempty"`
	FQName           []string         `json:"fq_name,omitempty"`
	IDPerms          *IdPermsType     `json:"id_perms,omitempty"`
	DisplayName      string           `json:"display_name,omitempty"`
	Annotations      *KeyValuePairs   `json:"annotations,omitempty"`
	Perms2           *PermType2       `json:"perms2,omitempty"`
	NetworkIpamMGMT  *IpamType        `json:"network_ipam_mgmt,omitempty"`
	IpamSubnets      *IpamSubnets     `json:"ipam_subnets,omitempty"`
	IpamSubnetMethod SubnetMethodType `json:"ipam_subnet_method,omitempty"`

	VirtualDNSRefs []*NetworkIpamVirtualDNSRef `json:"virtual_DNS_refs,omitempty"`
}

// NetworkIpamVirtualDNSRef references each other
type NetworkIpamVirtualDNSRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// MakeNetworkIpam makes NetworkIpam
func MakeNetworkIpam() *NetworkIpam {
	return &NetworkIpam{
		//TODO(nati): Apply default
		UUID:             "",
		ParentUUID:       "",
		ParentType:       "",
		FQName:           []string{},
		IDPerms:          MakeIdPermsType(),
		DisplayName:      "",
		Annotations:      MakeKeyValuePairs(),
		Perms2:           MakePermType2(),
		NetworkIpamMGMT:  MakeIpamType(),
		IpamSubnets:      MakeIpamSubnets(),
		IpamSubnetMethod: MakeSubnetMethodType(),
	}
}

// MakeNetworkIpamSlice() makes a slice of NetworkIpam
func MakeNetworkIpamSlice() []*NetworkIpam {
	return []*NetworkIpam{}
}
