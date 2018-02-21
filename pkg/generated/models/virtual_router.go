package models


// MakeVirtualRouter makes VirtualRouter
func MakeVirtualRouter() *VirtualRouter{
    return &VirtualRouter{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VirtualRouterDPDKEnabled: false,
        VirtualRouterType: "",
        VirtualRouterIPAddress: "",
        
    }
}

// MakeVirtualRouterSlice() makes a slice of VirtualRouter
func MakeVirtualRouterSlice() []*VirtualRouter {
    return []*VirtualRouter{}
}


