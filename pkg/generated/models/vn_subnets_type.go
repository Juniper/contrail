package models

// VnSubnetsType

import "encoding/json"

// VnSubnetsType
type VnSubnetsType struct {
	HostRoutes  *RouteTableType   `json:"host_routes"`
	IpamSubnets []*IpamSubnetType `json:"ipam_subnets"`
}

// String returns json representation of the object
func (model *VnSubnetsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVnSubnetsType makes VnSubnetsType
func MakeVnSubnetsType() *VnSubnetsType {
	return &VnSubnetsType{
		//TODO(nati): Apply default

		IpamSubnets: MakeIpamSubnetTypeSlice(),

		HostRoutes: MakeRouteTableType(),
	}
}

// InterfaceToVnSubnetsType makes VnSubnetsType from interface
func InterfaceToVnSubnetsType(iData interface{}) *VnSubnetsType {
	data := iData.(map[string]interface{})
	return &VnSubnetsType{
		HostRoutes: InterfaceToRouteTableType(data["host_routes"]),

		//{"description":"Common host routes to be sent via DHCP for VM(s) in all the subnets, Next hop for these routes is always default gateway","type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}}

		IpamSubnets: InterfaceToIpamSubnetTypeSlice(data["ipam_subnets"]),

		//{"type":"array","item":{"type":"object","properties":{"addr_from_start":{"type":"boolean"},"alloc_unit":{"type":"integer"},"allocation_pools":{"type":"array","item":{"type":"object","properties":{"end":{"type":"string"},"start":{"type":"string"},"vrouter_specific_pool":{"type":"boolean"}}}},"created":{"type":"string"},"default_gateway":{"type":"string"},"dhcp_option_list":{"type":"object","properties":{"dhcp_option":{"type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}}},"dns_nameservers":{"type":"array","item":{"type":"string"}},"dns_server_address":{"type":"string"},"enable_dhcp":{"type":"boolean"},"host_routes":{"type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}},"last_modified":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_name":{"type":"string"},"subnet_uuid":{"type":"string"}}}}

	}
}

// InterfaceToVnSubnetsTypeSlice makes a slice of VnSubnetsType from interface
func InterfaceToVnSubnetsTypeSlice(data interface{}) []*VnSubnetsType {
	list := data.([]interface{})
	result := MakeVnSubnetsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVnSubnetsType(item))
	}
	return result
}

// MakeVnSubnetsTypeSlice() makes a slice of VnSubnetsType
func MakeVnSubnetsTypeSlice() []*VnSubnetsType {
	return []*VnSubnetsType{}
}
