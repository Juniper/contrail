package models

// NetworkIpam

import "encoding/json"

// NetworkIpam
type NetworkIpam struct {
	IDPerms          *IdPermsType     `json:"id_perms,omitempty"`
	NetworkIpamMGMT  *IpamType        `json:"network_ipam_mgmt,omitempty"`
	IpamSubnets      *IpamSubnets     `json:"ipam_subnets,omitempty"`
	IpamSubnetMethod SubnetMethodType `json:"ipam_subnet_method,omitempty"`
	Annotations      *KeyValuePairs   `json:"annotations,omitempty"`
	ParentUUID       string           `json:"parent_uuid,omitempty"`
	ParentType       string           `json:"parent_type,omitempty"`
	FQName           []string         `json:"fq_name,omitempty"`
	DisplayName      string           `json:"display_name,omitempty"`
	Perms2           *PermType2       `json:"perms2,omitempty"`
	UUID             string           `json:"uuid,omitempty"`

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
		Annotations:      MakeKeyValuePairs(),
		ParentUUID:       "",
		ParentType:       "",
		FQName:           []string{},
		IDPerms:          MakeIdPermsType(),
		NetworkIpamMGMT:  MakeIpamType(),
		IpamSubnets:      MakeIpamSubnets(),
		IpamSubnetMethod: MakeSubnetMethodType(),
		DisplayName:      "",
		Perms2:           MakePermType2(),
		UUID:             "",
	}
}

// MakeNetworkIpamSlice() makes a slice of NetworkIpam
func MakeNetworkIpamSlice() []*NetworkIpam {
	return []*NetworkIpam{}
}
