package models


// MakeBaremetalPort makes BaremetalPort
func MakeBaremetalPort() *BaremetalPort{
    return &BaremetalPort{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        MacAddress: "",
        CreatedAt: "",
        UpdatedAt: "",
        Node: "",
        PxeEnabled: false,
        LocalLinkConnection: MakeLocalLinkConnection(),
        
    }
}

// MakeBaremetalPortSlice() makes a slice of BaremetalPort
func MakeBaremetalPortSlice() []*BaremetalPort {
    return []*BaremetalPort{}
}


