package models

// RouteNextHopType

type RouteNextHopType string

func MakeRouteNextHopType() RouteNextHopType {
	var data RouteNextHopType
	return data
}

func InterfaceToRouteNextHopType(data interface{}) RouteNextHopType {
	return data.(RouteNextHopType)
}

func InterfaceToRouteNextHopTypeSlice(data interface{}) []RouteNextHopType {
	list := data.([]interface{})
	result := MakeRouteNextHopTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteNextHopType(item))
	}
	return result
}

func MakeRouteNextHopTypeSlice() []RouteNextHopType {
	return []RouteNextHopType{}
}
