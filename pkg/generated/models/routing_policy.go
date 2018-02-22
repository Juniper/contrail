package models


// MakeRoutingPolicy makes RoutingPolicy
func MakeRoutingPolicy() *RoutingPolicy{
    return &RoutingPolicy{
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

// MakeRoutingPolicySlice() makes a slice of RoutingPolicy
func MakeRoutingPolicySlice() []*RoutingPolicy {
    return []*RoutingPolicy{}
}


