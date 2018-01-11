package models

// NetworkIpam

import "encoding/json"

// NetworkIpam
type NetworkIpam struct {
	IpamSubnets      *IpamSubnets     `json:"ipam_subnets"`
	FQName           []string         `json:"fq_name"`
	IDPerms          *IdPermsType     `json:"id_perms"`
	DisplayName      string           `json:"display_name"`
	UUID             string           `json:"uuid"`
	ParentUUID       string           `json:"parent_uuid"`
	NetworkIpamMGMT  *IpamType        `json:"network_ipam_mgmt"`
	IpamSubnetMethod SubnetMethodType `json:"ipam_subnet_method"`
	Annotations      *KeyValuePairs   `json:"annotations"`
	Perms2           *PermType2       `json:"perms2"`
	ParentType       string           `json:"parent_type"`

	VirtualDNSRefs []*NetworkIpamVirtualDNSRef `json:"virtual_DNS_refs"`
}

// NetworkIpamVirtualDNSRef references each other
type NetworkIpamVirtualDNSRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *NetworkIpam) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNetworkIpam makes NetworkIpam
func MakeNetworkIpam() *NetworkIpam {
	return &NetworkIpam{
		//TODO(nati): Apply default
		NetworkIpamMGMT:  MakeIpamType(),
		IpamSubnetMethod: MakeSubnetMethodType(),
		Annotations:      MakeKeyValuePairs(),
		Perms2:           MakePermType2(),
		ParentType:       "",
		IpamSubnets:      MakeIpamSubnets(),
		FQName:           []string{},
		IDPerms:          MakeIdPermsType(),
		DisplayName:      "",
		UUID:             "",
		ParentUUID:       "",
	}
}

// MakeNetworkIpamSlice() makes a slice of NetworkIpam
func MakeNetworkIpamSlice() []*NetworkIpam {
	return []*NetworkIpam{}
}
