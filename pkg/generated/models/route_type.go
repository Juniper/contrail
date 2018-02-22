package models


// MakeRouteType makes RouteType
func MakeRouteType() *RouteType{
    return &RouteType{
    //TODO(nati): Apply default
    Prefix: "",
        NextHop: "",
        CommunityAttributes: MakeCommunityAttributes(),
        NextHopType: "",
        
    }
}

// MakeRouteTypeSlice() makes a slice of RouteType
func MakeRouteTypeSlice() []*RouteType {
    return []*RouteType{}
}


