package models


// MakeInstanceIP makes InstanceIP
func MakeInstanceIP() *InstanceIP{
    return &InstanceIP{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceHealthCheckIP: false,
        SecondaryIPTrackingIP: MakeSubnetType(),
        InstanceIPAddress: "",
        InstanceIPMode: "",
        SubnetUUID: "",
        InstanceIPFamily: "",
        ServiceInstanceIP: false,
        InstanceIPLocalIP: false,
        InstanceIPSecondary: false,
        
    }
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
    return []*InstanceIP{}
}


