package models


// MakeRouteAggregate makes RouteAggregate
func MakeRouteAggregate() *RouteAggregate{
    return &RouteAggregate{
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

// MakeRouteAggregateSlice() makes a slice of RouteAggregate
func MakeRouteAggregateSlice() []*RouteAggregate {
    return []*RouteAggregate{}
}


