package models


// MakeBGPRouter makes BGPRouter
func MakeBGPRouter() *BGPRouter{
    return &BGPRouter{
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

// MakeBGPRouterSlice() makes a slice of BGPRouter
func MakeBGPRouterSlice() []*BGPRouter {
    return []*BGPRouter{}
}


