package models

// IpamSubnetType

import "encoding/json"

type IpamSubnetType struct {
	Created          string                `json:"created"`
	HostRoutes       *RouteTableType       `json:"host_routes"`
	SubnetName       string                `json:"subnet_name"`
	Subnet           *SubnetType           `json:"subnet"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	DefaultGateway   IpAddressType         `json:"default_gateway"`
	AllocUnit        int                   `json:"alloc_unit"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools"`
	LastModified     string                `json:"last_modified"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list"`
	SubnetUUID       string                `json:"subnet_uuid"`
	DNSServerAddress IpAddressType         `json:"dns_server_address"`
	DNSNameservers   []string              `json:"dns_nameservers"`
}

func (model *IpamSubnetType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeIpamSubnetType() *IpamSubnetType {
	return &IpamSubnetType{
		//TODO(nati): Apply default
		DNSNameservers: []string{},
		Created:        "",
		LastModified:   "",
		HostRoutes:     MakeRouteTableType(),
		SubnetName:     "",
		Subnet:         MakeSubnetType(),
		AddrFromStart:  false,
		DefaultGateway: MakeIpAddressType(),
		AllocUnit:      0,

		AllocationPools: MakeAllocationPoolTypeSlice(),

		EnableDHCP:       false,
		DHCPOptionList:   MakeDhcpOptionsListType(),
		SubnetUUID:       "",
		DNSServerAddress: MakeIpAddressType(),
	}
}

func InterfaceToIpamSubnetType(iData interface{}) *IpamSubnetType {
	data := iData.(map[string]interface{})
	return &IpamSubnetType{
		Created: data["created"].(string),

		//{"Title":"","Description":"timestamp when subnet object gets created","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Created","GoType":"string"}
		HostRoutes: InterfaceToRouteTableType(data["host_routes"]),

		//{"Title":"","Description":"Host routes to be sent via DHCP for VM(s) in this subnet, Next hop for these routes is always default gateway","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attributes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes"},"next_hop":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string"},"next_hop_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType"},"prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteType","CollectionType":"","Column":"","Item":null,"GoName":"Route","GoType":"RouteType"},"GoName":"Route","GoType":"[]*RouteType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTableType","CollectionType":"","Column":"","Item":null,"GoName":"HostRoutes","GoType":"RouteTableType"}
		SubnetName: data["subnet_name"].(string),

		//{"Title":"","Description":"User provided name for this subnet","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SubnetName","GoType":"string"}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"Title":"","Description":"ip prefix and length for the subnet","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"}
		AddrFromStart: data["addr_from_start"].(bool),

		//{"Title":"","Description":"Start address allocation from start or from end of address range.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AddrFromStart","GoType":"bool"}
		DefaultGateway: InterfaceToIpAddressType(data["default_gateway"]),

		//{"Title":"","Description":"default-gateway ip address in the subnet, if not provided one is auto generated by the system.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"DefaultGateway","GoType":"IpAddressType"}
		AllocUnit: data["alloc_unit"].(int),

		//{"Title":"","Description":"allocation unit for this subnet to allocate bulk ip addresses","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AllocUnit","GoType":"int"}

		AllocationPools: InterfaceToAllocationPoolTypeSlice(data["allocation_pools"]),

		//{"Title":"","Description":"List of ranges of ip address within the subnet from which to allocate ip address. default is entire prefix","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"End","GoType":"string"},"start":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Start","GoType":"string"},"vrouter_specific_pool":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VrouterSpecificPool","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllocationPoolType","CollectionType":"","Column":"","Item":null,"GoName":"AllocationPools","GoType":"AllocationPoolType"},"GoName":"AllocationPools","GoType":"[]*AllocationPoolType"}
		LastModified: data["last_modified"].(string),

		//{"Title":"","Description":"timestamp when subnet object gets updated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LastModified","GoType":"string"}
		EnableDHCP: data["enable_dhcp"].(bool),

		//{"Title":"","Description":"Enable DHCP for the VM(s) in this subnet","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"EnableDHCP","GoType":"bool"}
		DHCPOptionList: InterfaceToDhcpOptionsListType(data["dhcp_option_list"]),

		//{"Title":"","Description":"DHCP options list to be sent via DHCP for  VM(s) in this subnet","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"dhcp_option":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dhcp_option_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionName","GoType":"string"},"dhcp_option_value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValue","GoType":"string"},"dhcp_option_value_bytes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionValueBytes","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DhcpOptionType","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOption","GoType":"DhcpOptionType"},"GoName":"DHCPOption","GoType":"[]*DhcpOptionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DhcpOptionsListType","CollectionType":"","Column":"","Item":null,"GoName":"DHCPOptionList","GoType":"DhcpOptionsListType"}
		SubnetUUID: data["subnet_uuid"].(string),

		//{"Title":"","Description":"Subnet UUID is auto generated by the system","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SubnetUUID","GoType":"string"}
		DNSServerAddress: InterfaceToIpAddressType(data["dns_server_address"]),

		//{"Title":"","Description":"DNS server ip address in the subnet, if not provided one is auto generated by the system.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"DNSServerAddress","GoType":"IpAddressType"}
		DNSNameservers: data["dns_nameservers"].([]string),

		//{"Title":"","Description":"Tenant DNS servers ip address in tenant DNS method","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DNSNameservers","GoType":"string"},"GoName":"DNSNameservers","GoType":"[]string"}

	}
}

func InterfaceToIpamSubnetTypeSlice(data interface{}) []*IpamSubnetType {
	list := data.([]interface{})
	result := MakeIpamSubnetTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamSubnetType(item))
	}
	return result
}

func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
