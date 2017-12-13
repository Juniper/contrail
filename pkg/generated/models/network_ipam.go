package models

// NetworkIpam

import "encoding/json"

// NetworkIpam
type NetworkIpam struct {
	IpamSubnets      *IpamSubnets     `json:"ipam_subnets"`
	IpamSubnetMethod SubnetMethodType `json:"ipam_subnet_method"`
	Annotations      *KeyValuePairs   `json:"annotations"`
	Perms2           *PermType2       `json:"perms2"`
	DisplayName      string           `json:"display_name"`
	NetworkIpamMGMT  *IpamType        `json:"network_ipam_mgmt"`
	UUID             string           `json:"uuid"`
	ParentUUID       string           `json:"parent_uuid"`
	ParentType       string           `json:"parent_type"`
	FQName           []string         `json:"fq_name"`
	IDPerms          *IdPermsType     `json:"id_perms"`

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
		UUID:             "",
		ParentUUID:       "",
		ParentType:       "",
		FQName:           []string{},
		IDPerms:          MakeIdPermsType(),
		IpamSubnets:      MakeIpamSubnets(),
		IpamSubnetMethod: MakeSubnetMethodType(),
		Annotations:      MakeKeyValuePairs(),
		Perms2:           MakePermType2(),
		DisplayName:      "",
	}
}

// InterfaceToNetworkIpam makes NetworkIpam from interface
func InterfaceToNetworkIpam(iData interface{}) *NetworkIpam {
	data := iData.(map[string]interface{})
	return &NetworkIpam{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		NetworkIpamMGMT: InterfaceToIpamType(data["network_ipam_mgmt"]),

		//{"description":"Network IP Address Management configuration.","type":"object","properties":{"cidr_block":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"dhcp_option_list":{"type":"object","properties":{"dhcp_option":{"type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}}},"host_routes":{"type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}},"ipam_dns_method":{"type":"string","enum":["none","default-dns-server","tenant-dns-server","virtual-dns-server"]},"ipam_dns_server":{"type":"object","properties":{"tenant_dns_server_address":{"type":"object","properties":{"ip_address":{"type":"string"}}},"virtual_dns_server_name":{"type":"string"}}},"ipam_method":{"type":"string","enum":["dhcp","fixed"]}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		IpamSubnets: InterfaceToIpamSubnets(data["ipam_subnets"]),

		//{"description":"List of subnets for this ipam.","type":"object","properties":{"subnets":{"type":"array","item":{"type":"object","properties":{"addr_from_start":{"type":"boolean"},"alloc_unit":{"type":"integer"},"allocation_pools":{"type":"array","item":{"type":"object","properties":{"end":{"type":"string"},"start":{"type":"string"},"vrouter_specific_pool":{"type":"boolean"}}}},"created":{"type":"string"},"default_gateway":{"type":"string"},"dhcp_option_list":{"type":"object","properties":{"dhcp_option":{"type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}}},"dns_nameservers":{"type":"array","item":{"type":"string"}},"dns_server_address":{"type":"string"},"enable_dhcp":{"type":"boolean"},"host_routes":{"type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}},"last_modified":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_name":{"type":"string"},"subnet_uuid":{"type":"string"}}}}}}
		IpamSubnetMethod: InterfaceToSubnetMethodType(data["ipam_subnet_method"]),

		//{"description":"Subnet method configuration for ipam, user can configure user-defined, flat or auto.","type":"string","enum":["user-defined-subnet","flat-subnet","auto-subnet"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}

	}
}

// InterfaceToNetworkIpamSlice makes a slice of NetworkIpam from interface
func InterfaceToNetworkIpamSlice(data interface{}) []*NetworkIpam {
	list := data.([]interface{})
	result := MakeNetworkIpamSlice()
	for _, item := range list {
		result = append(result, InterfaceToNetworkIpam(item))
	}
	return result
}

// MakeNetworkIpamSlice() makes a slice of NetworkIpam
func MakeNetworkIpamSlice() []*NetworkIpam {
	return []*NetworkIpam{}
}
