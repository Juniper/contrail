package models

// RouteTableType

import "encoding/json"

// RouteTableType
type RouteTableType struct {
	Route []*RouteType `json:"route"`
}

//  parents relation object

// String returns json representation of the object
func (model *RouteTableType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteTableType makes RouteTableType
func MakeRouteTableType() *RouteTableType {
	return &RouteTableType{
		//TODO(nati): Apply default

		Route: MakeRouteTypeSlice(),
	}
}

// InterfaceToRouteTableType makes RouteTableType from interface
func InterfaceToRouteTableType(iData interface{}) *RouteTableType {
	data := iData.(map[string]interface{})
	return &RouteTableType{

		Route: InterfaceToRouteTypeSlice(data["route"]),

		//{"Title":"","Description":"List of ip routes with following fields.","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attributes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes","GoPremitive":false},"next_hop":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string","GoPremitive":true},"next_hop_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType","GoPremitive":false},"prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteType","CollectionType":"","Column":"","Item":null,"GoName":"Route","GoType":"RouteType","GoPremitive":false},"GoName":"Route","GoType":"[]*RouteType","GoPremitive":true}

	}
}

// InterfaceToRouteTableTypeSlice makes a slice of RouteTableType from interface
func InterfaceToRouteTableTypeSlice(data interface{}) []*RouteTableType {
	list := data.([]interface{})
	result := MakeRouteTableTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteTableType(item))
	}
	return result
}

// MakeRouteTableTypeSlice() makes a slice of RouteTableType
func MakeRouteTableTypeSlice() []*RouteTableType {
	return []*RouteTableType{}
}
