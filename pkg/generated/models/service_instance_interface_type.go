package models

// ServiceInstanceInterfaceType

import "encoding/json"

// ServiceInstanceInterfaceType
type ServiceInstanceInterfaceType struct {
	AllowedAddressPairs *AllowedAddressPairs `json:"allowed_address_pairs"`
	StaticRoutes        *RouteTableType      `json:"static_routes"`
	VirtualNetwork      string               `json:"virtual_network"`
	IPAddress           IpAddressType        `json:"ip_address"`
}

//  parents relation object

// String returns json representation of the object
func (model *ServiceInstanceInterfaceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceInstanceInterfaceType makes ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceType() *ServiceInstanceInterfaceType {
	return &ServiceInstanceInterfaceType{
		//TODO(nati): Apply default
		VirtualNetwork:      "",
		IPAddress:           MakeIpAddressType(),
		AllowedAddressPairs: MakeAllowedAddressPairs(),
		StaticRoutes:        MakeRouteTableType(),
	}
}

// InterfaceToServiceInstanceInterfaceType makes ServiceInstanceInterfaceType from interface
func InterfaceToServiceInstanceInterfaceType(iData interface{}) *ServiceInstanceInterfaceType {
	data := iData.(map[string]interface{})
	return &ServiceInstanceInterfaceType{
		VirtualNetwork: data["virtual_network"].(string),

		//{"Title":"","Description":"Interface belongs to this virtual network.","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string","GoPremitive":true}
		IPAddress: InterfaceToIpAddressType(data["ip_address"]),

		//{"Title":"","Description":"Shared ip for this interface (Only V1)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IPAddress","GoType":"IpAddressType","GoPremitive":false}
		AllowedAddressPairs: InterfaceToAllowedAddressPairs(data["allowed_address_pairs"]),

		//{"Title":"","Description":"Allowed address pairs, list of (IP address, MAC) for this interface","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"allowed_address_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"address_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"","Item":null,"GoName":"AddressMode","GoType":"AddressMode","GoPremitive":false},"ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"IP","GoType":"SubnetType","GoPremitive":false},"mac":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Mac","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPair","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPair","GoType":"AllowedAddressPair","GoPremitive":false},"GoName":"AllowedAddressPair","GoType":"[]*AllowedAddressPair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPairs","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPairs","GoType":"AllowedAddressPairs","GoPremitive":false}
		StaticRoutes: InterfaceToRouteTableType(data["static_routes"]),

		//{"Title":"","Description":"Static routes for this interface (Only V1)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attributes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes","GoPremitive":false},"next_hop":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string","GoPremitive":true},"next_hop_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType","GoPremitive":false},"prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteType","CollectionType":"","Column":"","Item":null,"GoName":"Route","GoType":"RouteType","GoPremitive":false},"GoName":"Route","GoType":"[]*RouteType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTableType","CollectionType":"","Column":"","Item":null,"GoName":"StaticRoutes","GoType":"RouteTableType","GoPremitive":false}

	}
}

// InterfaceToServiceInstanceInterfaceTypeSlice makes a slice of ServiceInstanceInterfaceType from interface
func InterfaceToServiceInstanceInterfaceTypeSlice(data interface{}) []*ServiceInstanceInterfaceType {
	list := data.([]interface{})
	result := MakeServiceInstanceInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInstanceInterfaceType(item))
	}
	return result
}

// MakeServiceInstanceInterfaceTypeSlice() makes a slice of ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
	return []*ServiceInstanceInterfaceType{}
}
