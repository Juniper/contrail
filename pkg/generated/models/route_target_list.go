package models


// MakeRouteTargetList makes RouteTargetList
func MakeRouteTargetList() *RouteTargetList{
    return &RouteTargetList{
    //TODO(nati): Apply default
    RouteTarget: []string{},
        
    }
}

// MakeRouteTargetListSlice() makes a slice of RouteTargetList
func MakeRouteTargetListSlice() []*RouteTargetList {
    return []*RouteTargetList{}
}


