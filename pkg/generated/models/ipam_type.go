package models

// IpamType

import "encoding/json"

type IpamType struct {
	DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list"`
	HostRoutes     *RouteTableType      `json:"host_routes"`
	CidrBlock      *SubnetType          `json:"cidr_block"`
	IpamMethod     IpamMethodType       `json:"ipam_method"`
	IpamDNSMethod  IpamDnsMethodType    `json:"ipam_dns_method"`
	IpamDNSServer  *IpamDnsAddressType  `json:"ipam_dns_server"`
}

func (model *IpamType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

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

func InterfaceToIpamType(iData interface{}) *IpamType {
	data := iData.(map[string]interface{})
	return &IpamType{
		IpamDNSMethod: InterfaceToIpamDnsMethodType(data["ipam_dns_method"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["none","default-dns-server","tenant-dns-server","virtual-dns-server"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpamDnsMethodType","CollectionType":"","Column":"","Item":null,"GoName":"IpamDNSMethod","GoType":"IpamDnsMethodType"}
		IpamDNSServer: InterfaceToIpamDnsAddressType(data["ipam_dns_server"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"tenant_dns_server_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IPAddress","GoType":"IpAddressType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressesType","CollectionType":"","Column":"","Item":null,"GoName":"TenantDNSServerAddress","GoType":"IpAddressesType"},"virtual_dns_server_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualDNSServerName","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpamDnsAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IpamDNSServer","GoType":"IpamDnsAddressType"}
		DHCPOptionList: InterfaceToDhcpOptionsListType(data["dhcp_option_list"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"dhcp_option":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dhcp_option_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionName","GoType":"string"},"dhcp_option_value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValue","GoType":"string"},"dhcp_option_value_bytes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValueBytes","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DhcpOptionType","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOption","GoType":"DhcpOptionType"},"GoName":"DHCPOption","GoType":"[]*DhcpOptionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DhcpOptionsListType","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionList","GoType":"DhcpOptionsListType"}
		HostRoutes: InterfaceToRouteTableType(data["host_routes"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attributes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes"},"next_hop":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string"},"next_hop_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType"},"prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteType","CollectionType":"","Column":"","Item":null,"GoName":"Route","GoType":"RouteType"},"GoName":"Route","GoType":"[]*RouteType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTableType","CollectionType":"","Column":"","Item":null,"GoName":"HostRoutes","GoType":"RouteTableType"}
		CidrBlock: InterfaceToSubnetType(data["cidr_block"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"CidrBlock","GoType":"SubnetType"}
		IpamMethod: InterfaceToIpamMethodType(data["ipam_method"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dhcp","fixed"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpamMethodType","CollectionType":"","Column":"","Item":null,"GoName":"IpamMethod","GoType":"IpamMethodType"}

	}
}

func InterfaceToIpamTypeSlice(data interface{}) []*IpamType {
	list := data.([]interface{})
	result := MakeIpamTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamType(item))
	}
	return result
}

func MakeIpamTypeSlice() []*IpamType {
	return []*IpamType{}
}
