package models

// RouteTargetList

import "encoding/json"

// RouteTargetList
//proteus:generate
type RouteTargetList struct {
	RouteTarget []string `json:"route_target,omitempty"`
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

// MakeRouteTargetListSlice() makes a slice of RouteTargetList
func MakeRouteTargetListSlice() []*RouteTargetList {
	return []*RouteTargetList{}
}
