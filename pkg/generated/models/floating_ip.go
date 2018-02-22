package models


// MakeFloatingIP makes FloatingIP
func MakeFloatingIP() *FloatingIP{
    return &FloatingIP{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        FloatingIPAddressFamily: "",
        FloatingIPPortMappings: MakePortMappings(),
        FloatingIPIsVirtualIP: false,
        FloatingIPAddress: "",
        FloatingIPPortMappingsEnable: false,
        FloatingIPFixedIPAddress: "",
        FloatingIPTrafficDirection: "",
        
    }
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
    return []*FloatingIP{}
}


