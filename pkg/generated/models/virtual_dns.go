package models


// MakeVirtualDNS makes VirtualDNS
func MakeVirtualDNS() *VirtualDNS{
    return &VirtualDNS{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VirtualDNSData: MakeVirtualDnsType(),
        
    }
}

// MakeVirtualDNSSlice() makes a slice of VirtualDNS
func MakeVirtualDNSSlice() []*VirtualDNS {
    return []*VirtualDNS{}
}


