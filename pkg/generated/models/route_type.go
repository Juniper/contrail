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

// String returns json representation of the object
func (model *RouteType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteType makes RouteType
func MakeRouteType() *RouteType {
	return &RouteType{
		//TODO(nati): Apply default
		CommunityAttributes: MakeCommunityAttributes(),
		NextHopType:         MakeRouteNextHopType(),
		Prefix:              "",
		NextHop:             "",
	}
}

// InterfaceToRouteType makes RouteType from interface
func InterfaceToRouteType(iData interface{}) *RouteType {
	data := iData.(map[string]interface{})
	return &RouteType{
		Prefix: data["prefix"].(string),

		//{"description":"Ip prefix/len format prefix","type":"string"}
		NextHop: data["next_hop"].(string),

		//{"description":"Ip address or service instance name.","type":"string"}
		CommunityAttributes: InterfaceToCommunityAttributes(data["community_attributes"]),

		//{"type":"object","properties":{"community_attribute":{"type":"array"}}}
		NextHopType: InterfaceToRouteNextHopType(data["next_hop_type"]),

		//{"type":"string","enum":["service-instance","ip-address"]}

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
