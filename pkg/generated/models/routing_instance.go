package models


// MakeRoutingInstance makes RoutingInstance
func MakeRoutingInstance() *RoutingInstance{
    return &RoutingInstance{
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

// MakeRoutingInstanceSlice() makes a slice of RoutingInstance
func MakeRoutingInstanceSlice() []*RoutingInstance {
    return []*RoutingInstance{}
}


