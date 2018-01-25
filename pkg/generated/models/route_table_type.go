package models

// RouteTableType

// RouteTableType
//proteus:generate
type RouteTableType struct {
	Route []*RouteType `json:"route,omitempty"`
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
