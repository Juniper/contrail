package models


// MakeRouteTarget makes RouteTarget
func MakeRouteTarget() *RouteTarget{
    return &RouteTarget{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
    }
}

// MakeRouteTargetSlice() makes a slice of RouteTarget
func MakeRouteTargetSlice() []*RouteTarget {
    return []*RouteTarget{}
}


