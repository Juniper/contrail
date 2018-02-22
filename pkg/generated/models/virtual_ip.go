package models


// MakeVirtualIP makes VirtualIP
func MakeVirtualIP() *VirtualIP{
    return &VirtualIP{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VirtualIPProperties: MakeVirtualIpType(),
        
    }
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
    return []*VirtualIP{}
}


