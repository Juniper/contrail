package models


// MakeRouteTable makes RouteTable
func MakeRouteTable() *RouteTable{
    return &RouteTable{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Routes: MakeRouteTableType(),
        
    }
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
    return []*RouteTable{}
}


