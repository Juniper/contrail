package models

// IpamType

import "encoding/json"

// IpamType
type IpamType struct {
	IpamMethod     IpamMethodType       `json:"ipam_method"`
	IpamDNSMethod  IpamDnsMethodType    `json:"ipam_dns_method"`
	IpamDNSServer  *IpamDnsAddressType  `json:"ipam_dns_server"`
	DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list"`
	HostRoutes     *RouteTableType      `json:"host_routes"`
	CidrBlock      *SubnetType          `json:"cidr_block"`
}

// String returns json representation of the object
func (model *IpamType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpamType makes IpamType
func MakeIpamType() *IpamType {
	return &IpamType{
		//TODO(nati): Apply default
		IpamDNSServer:  MakeIpamDnsAddressType(),
		DHCPOptionList: MakeDhcpOptionsListType(),
		HostRoutes:     MakeRouteTableType(),
		CidrBlock:      MakeSubnetType(),
		IpamMethod:     MakeIpamMethodType(),
		IpamDNSMethod:  MakeIpamDnsMethodType(),
	}
}

// InterfaceToIpamType makes IpamType from interface
func InterfaceToIpamType(iData interface{}) *IpamType {
	data := iData.(map[string]interface{})
	return &IpamType{
		CidrBlock: InterfaceToSubnetType(data["cidr_block"]),

		//{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}
		IpamMethod: InterfaceToIpamMethodType(data["ipam_method"]),

		//{"type":"string","enum":["dhcp","fixed"]}
		IpamDNSMethod: InterfaceToIpamDnsMethodType(data["ipam_dns_method"]),

		//{"type":"string","enum":["none","default-dns-server","tenant-dns-server","virtual-dns-server"]}
		IpamDNSServer: InterfaceToIpamDnsAddressType(data["ipam_dns_server"]),

		//{"type":"object","properties":{"tenant_dns_server_address":{"type":"object","properties":{"ip_address":{"type":"string"}}},"virtual_dns_server_name":{"type":"string"}}}
		DHCPOptionList: InterfaceToDhcpOptionsListType(data["dhcp_option_list"]),

		//{"type":"object","properties":{"dhcp_option":{"type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}}}
		HostRoutes: InterfaceToRouteTableType(data["host_routes"]),

		//{"type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}}

	}
}

// InterfaceToIpamTypeSlice makes a slice of IpamType from interface
func InterfaceToIpamTypeSlice(data interface{}) []*IpamType {
	list := data.([]interface{})
	result := MakeIpamTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamType(item))
	}
	return result
}

// MakeIpamTypeSlice() makes a slice of IpamType
func MakeIpamTypeSlice() []*IpamType {
	return []*IpamType{}
}
