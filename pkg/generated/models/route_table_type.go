package models

// RouteTableType

import "encoding/json"

// RouteTableType
//proteus:generate
type RouteTableType struct {
	Route []*RouteType `json:"route,omitempty"`
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

// MakeRouteTableTypeSlice() makes a slice of RouteTableType
func MakeRouteTableTypeSlice() []*RouteTableType {
	return []*RouteTableType{}
}
