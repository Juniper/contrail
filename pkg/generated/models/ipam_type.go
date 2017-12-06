package models

// IpamType

import "encoding/json"

// IpamType
type IpamType struct {
	IpamDNSServer  *IpamDnsAddressType  `json:"ipam_dns_server"`
	DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list"`
	HostRoutes     *RouteTableType      `json:"host_routes"`
	CidrBlock      *SubnetType          `json:"cidr_block"`
	IpamMethod     IpamMethodType       `json:"ipam_method"`
	IpamDNSMethod  IpamDnsMethodType    `json:"ipam_dns_method"`
}

//  parents relation object

// String returns json representation of the object
func (model *IpamType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpamType makes IpamType
func MakeIpamType() *IpamType {
	return &IpamType{
		//TODO(nati): Apply default
		IpamMethod:     MakeIpamMethodType(),
		IpamDNSMethod:  MakeIpamDnsMethodType(),
		IpamDNSServer:  MakeIpamDnsAddressType(),
		DHCPOptionList: MakeDhcpOptionsListType(),
		HostRoutes:     MakeRouteTableType(),
		CidrBlock:      MakeSubnetType(),
	}
}

// InterfaceToIpamType makes IpamType from interface
func InterfaceToIpamType(iData interface{}) *IpamType {
	data := iData.(map[string]interface{})
	return &IpamType{
		IpamDNSServer: InterfaceToIpamDnsAddressType(data["ipam_dns_server"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"tenant_dns_server_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IPAddress","GoType":"IpAddressType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressesType","CollectionType":"","Column":"","Item":null,"GoName":"TenantDNSServerAddress","GoType":"IpAddressesType","GoPremitive":false},"virtual_dns_server_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualDNSServerName","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpamDnsAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IpamDNSServer","GoType":"IpamDnsAddressType","GoPremitive":false}
		DHCPOptionList: InterfaceToDhcpOptionsListType(data["dhcp_option_list"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"dhcp_option":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dhcp_option_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionName","GoType":"string","GoPremitive":true},"dhcp_option_value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValue","GoType":"string","GoPremitive":true},"dhcp_option_value_bytes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValueBytes","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DhcpOptionType","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOption","GoType":"DhcpOptionType","GoPremitive":false},"GoName":"DHCPOption","GoType":"[]*DhcpOptionType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DhcpOptionsListType","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionList","GoType":"DhcpOptionsListType","GoPremitive":false}
		HostRoutes: InterfaceToRouteTableType(data["host_routes"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attributes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes","GoPremitive":false},"next_hop":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string","GoPremitive":true},"next_hop_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType","GoPremitive":false},"prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteType","CollectionType":"","Column":"","Item":null,"GoName":"Route","GoType":"RouteType","GoPremitive":false},"GoName":"Route","GoType":"[]*RouteType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTableType","CollectionType":"","Column":"","Item":null,"GoName":"HostRoutes","GoType":"RouteTableType","GoPremitive":false}
		CidrBlock: InterfaceToSubnetType(data["cidr_block"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"CidrBlock","GoType":"SubnetType","GoPremitive":false}
		IpamMethod: InterfaceToIpamMethodType(data["ipam_method"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dhcp","fixed"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpamMethodType","CollectionType":"","Column":"","Item":null,"GoName":"IpamMethod","GoType":"IpamMethodType","GoPremitive":false}
		IpamDNSMethod: InterfaceToIpamDnsMethodType(data["ipam_dns_method"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["none","default-dns-server","tenant-dns-server","virtual-dns-server"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpamDnsMethodType","CollectionType":"","Column":"","Item":null,"GoName":"IpamDNSMethod","GoType":"IpamDnsMethodType","GoPremitive":false}

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
