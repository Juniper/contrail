package models


// MakeLogicalRouter makes LogicalRouter
func MakeLogicalRouter() *LogicalRouter{
    return &LogicalRouter{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VxlanNetworkIdentifier: "",
        ConfiguredRouteTargetList: MakeRouteTargetList(),
        
    }
}

// MakeLogicalRouterSlice() makes a slice of LogicalRouter
func MakeLogicalRouterSlice() []*LogicalRouter {
    return []*LogicalRouter{}
}


