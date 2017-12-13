package models

// IpamSubnets

import "encoding/json"

// IpamSubnets
type IpamSubnets struct {
	Subnets []*IpamSubnetType `json:"subnets"`
}

// String returns json representation of the object
func (model *IpamSubnets) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpamSubnets makes IpamSubnets
func MakeIpamSubnets() *IpamSubnets {
	return &IpamSubnets{
		//TODO(nati): Apply default

		Subnets: MakeIpamSubnetTypeSlice(),
	}
}

// InterfaceToIpamSubnets makes IpamSubnets from interface
func InterfaceToIpamSubnets(iData interface{}) *IpamSubnets {
	data := iData.(map[string]interface{})
	return &IpamSubnets{

		Subnets: InterfaceToIpamSubnetTypeSlice(data["subnets"]),

		//{"type":"array","item":{"type":"object","properties":{"addr_from_start":{"type":"boolean"},"alloc_unit":{"type":"integer"},"allocation_pools":{"type":"array","item":{"type":"object","properties":{"end":{"type":"string"},"start":{"type":"string"},"vrouter_specific_pool":{"type":"boolean"}}}},"created":{"type":"string"},"default_gateway":{"type":"string"},"dhcp_option_list":{"type":"object","properties":{"dhcp_option":{"type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}}},"dns_nameservers":{"type":"array","item":{"type":"string"}},"dns_server_address":{"type":"string"},"enable_dhcp":{"type":"boolean"},"host_routes":{"type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}},"last_modified":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_name":{"type":"string"},"subnet_uuid":{"type":"string"}}}}

	}
}

// InterfaceToIpamSubnetsSlice makes a slice of IpamSubnets from interface
func InterfaceToIpamSubnetsSlice(data interface{}) []*IpamSubnets {
	list := data.([]interface{})
	result := MakeIpamSubnetsSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamSubnets(item))
	}
	return result
}

// MakeIpamSubnetsSlice() makes a slice of IpamSubnets
func MakeIpamSubnetsSlice() []*IpamSubnets {
	return []*IpamSubnets{}
}
