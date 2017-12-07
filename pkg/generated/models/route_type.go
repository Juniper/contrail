package models

// RouteType

import "encoding/json"

// RouteType
type RouteType struct {
	Prefix              string               `json:"prefix"`
	NextHop             string               `json:"next_hop"`
	CommunityAttributes *CommunityAttributes `json:"community_attributes"`
	NextHopType         RouteNextHopType     `json:"next_hop_type"`
}

//  parents relation object

// String returns json representation of the object
func (model *RouteType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteType makes RouteType
func MakeRouteType() *RouteType {
	return &RouteType{
		//TODO(nati): Apply default
		NextHop:             "",
		CommunityAttributes: MakeCommunityAttributes(),
		NextHopType:         MakeRouteNextHopType(),
		Prefix:              "",
	}
}

// InterfaceToRouteType makes RouteType from interface
func InterfaceToRouteType(iData interface{}) *RouteType {
	data := iData.(map[string]interface{})
	return &RouteType{
		Prefix: data["prefix"].(string),

		//{"Title":"","Description":"Ip prefix/len format prefix","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string","GoPremitive":true}
		NextHop: data["next_hop"].(string),

		//{"Title":"","Description":"Ip address or service instance name.","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string","GoPremitive":true}
		CommunityAttributes: InterfaceToCommunityAttributes(data["community_attributes"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes","GoPremitive":false}
		NextHopType: InterfaceToRouteNextHopType(data["next_hop_type"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType","GoPremitive":false}

	}
}

// InterfaceToRouteTypeSlice makes a slice of RouteType from interface
func InterfaceToRouteTypeSlice(data interface{}) []*RouteType {
	list := data.([]interface{})
	result := MakeRouteTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteType(item))
	}
	return result
}

// MakeRouteTypeSlice() makes a slice of RouteType
func MakeRouteTypeSlice() []*RouteType {
	return []*RouteType{}
}
