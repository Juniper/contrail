package models

// NetworkIpam

import "encoding/json"

// NetworkIpam
type NetworkIpam struct {
	Perms2           *PermType2       `json:"perms2,omitempty"`
	UUID             string           `json:"uuid,omitempty"`
	IpamSubnets      *IpamSubnets     `json:"ipam_subnets,omitempty"`
	ParentType       string           `json:"parent_type,omitempty"`
	DisplayName      string           `json:"display_name,omitempty"`
	IDPerms          *IdPermsType     `json:"id_perms,omitempty"`
	Annotations      *KeyValuePairs   `json:"annotations,omitempty"`
	ParentUUID       string           `json:"parent_uuid,omitempty"`
	NetworkIpamMGMT  *IpamType        `json:"network_ipam_mgmt,omitempty"`
	IpamSubnetMethod SubnetMethodType `json:"ipam_subnet_method,omitempty"`
	FQName           []string         `json:"fq_name,omitempty"`

	VirtualDNSRefs []*NetworkIpamVirtualDNSRef `json:"virtual_DNS_refs,omitempty"`
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
		IpamSubnets:      MakeIpamSubnets(),
		ParentType:       "",
		DisplayName:      "",
		Perms2:           MakePermType2(),
		UUID:             "",
		ParentUUID:       "",
		NetworkIpamMGMT:  MakeIpamType(),
		IpamSubnetMethod: MakeSubnetMethodType(),
		FQName:           []string{},
		IDPerms:          MakeIdPermsType(),
		Annotations:      MakeKeyValuePairs(),
	}
}

// MakeNetworkIpamSlice() makes a slice of NetworkIpam
func MakeNetworkIpamSlice() []*NetworkIpam {
	return []*NetworkIpam{}
}
