package models

// ServiceInstanceInterfaceType

import "encoding/json"

type ServiceInstanceInterfaceType struct {
	VirtualNetwork      string               `json:"virtual_network"`
	IPAddress           IpAddressType        `json:"ip_address"`
	AllowedAddressPairs *AllowedAddressPairs `json:"allowed_address_pairs"`
	StaticRoutes        *RouteTableType      `json:"static_routes"`
}

func (model *ServiceInstanceInterfaceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeServiceInstanceInterfaceType() *ServiceInstanceInterfaceType {
	return &ServiceInstanceInterfaceType{
		//TODO(nati): Apply default
		StaticRoutes:        MakeRouteTableType(),
		VirtualNetwork:      "",
		IPAddress:           MakeIpAddressType(),
		AllowedAddressPairs: MakeAllowedAddressPairs(),
	}
}

func InterfaceToServiceInstanceInterfaceType(iData interface{}) *ServiceInstanceInterfaceType {
	data := iData.(map[string]interface{})
	return &ServiceInstanceInterfaceType{
		StaticRoutes: InterfaceToRouteTableType(data["static_routes"]),

		//{"Title":"","Description":"Static routes for this interface (Only V1)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attributes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes"},"next_hop":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string"},"next_hop_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType"},"prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteType","CollectionType":"","Column":"","Item":null,"GoName":"Route","GoType":"RouteType"},"GoName":"Route","GoType":"[]*RouteType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTableType","CollectionType":"","Column":"","Item":null,"GoName":"StaticRoutes","GoType":"RouteTableType"}
		VirtualNetwork: data["virtual_network"].(string),

		//{"Title":"","Description":"Interface belongs to this virtual network.","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}
		IPAddress: InterfaceToIpAddressType(data["ip_address"]),

		//{"Title":"","Description":"Shared ip for this interface (Only V1)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IPAddress","GoType":"IpAddressType"}
		AllowedAddressPairs: InterfaceToAllowedAddressPairs(data["allowed_address_pairs"]),

		//{"Title":"","Description":"Allowed address pairs, list of (IP address, MAC) for this interface","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"allowed_address_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"address_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"","Item":null,"GoName":"AddressMode","GoType":"AddressMode"},"ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"IP","GoType":"SubnetType"},"mac":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Mac","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPair","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPair","GoType":"AllowedAddressPair"},"GoName":"AllowedAddressPair","GoType":"[]*AllowedAddressPair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPairs","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPairs","GoType":"AllowedAddressPairs"}

	}
}

func InterfaceToServiceInstanceInterfaceTypeSlice(data interface{}) []*ServiceInstanceInterfaceType {
	list := data.([]interface{})
	result := MakeServiceInstanceInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInstanceInterfaceType(item))
	}
	return result
}

func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
	return []*ServiceInstanceInterfaceType{}
}
