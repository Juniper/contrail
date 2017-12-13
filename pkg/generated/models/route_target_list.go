package models

// RouteTargetList

import "encoding/json"

// RouteTargetList
type RouteTargetList struct {
	RouteTarget []string `json:"route_target"`
}

// String returns json representation of the object
func (model *RouteTargetList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteTargetList makes RouteTargetList
func MakeRouteTargetList() *RouteTargetList {
	return &RouteTargetList{
		//TODO(nati): Apply default
		RouteTarget: []string{},
	}
}

// InterfaceToRouteTargetList makes RouteTargetList from interface
func InterfaceToRouteTargetList(iData interface{}) *RouteTargetList {
	data := iData.(map[string]interface{})
	return &RouteTargetList{
		RouteTarget: data["route_target"].([]string),

		//{"type":"array","item":{"type":"string"}}

	}
}

// InterfaceToRouteTargetListSlice makes a slice of RouteTargetList from interface
func InterfaceToRouteTargetListSlice(data interface{}) []*RouteTargetList {
	list := data.([]interface{})
	result := MakeRouteTargetListSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteTargetList(item))
	}
	return result
}

// MakeRouteTargetListSlice() makes a slice of RouteTargetList
func MakeRouteTargetListSlice() []*RouteTargetList {
	return []*RouteTargetList{}
}
