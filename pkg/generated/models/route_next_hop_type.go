package models

// RouteNextHopType

type RouteNextHopType string

// MakeRouteNextHopType makes RouteNextHopType
func MakeRouteNextHopType() RouteNextHopType {
	var data RouteNextHopType
	return data
}

// InterfaceToRouteNextHopType makes RouteNextHopType from interface
func InterfaceToRouteNextHopType(data interface{}) RouteNextHopType {
	return data.(RouteNextHopType)
}

// InterfaceToRouteNextHopTypeSlice makes a slice of RouteNextHopType from interface
func InterfaceToRouteNextHopTypeSlice(data interface{}) []RouteNextHopType {
	list := data.([]interface{})
	result := MakeRouteNextHopTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteNextHopType(item))
	}
	return result
}

// MakeRouteNextHopTypeSlice() makes a slice of RouteNextHopType
func MakeRouteNextHopTypeSlice() []RouteNextHopType {
	return []RouteNextHopType{}
}
