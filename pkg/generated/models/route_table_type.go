package models

// RouteTableType

import "encoding/json"

// RouteTableType
type RouteTableType struct {
	Route []*RouteType `json:"route"`
}

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

		//{"description":"List of ip routes with following fields.","type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}

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
