package models


// MakeBaremetalNode makes BaremetalNode
func MakeBaremetalNode() *BaremetalNode{
    return &BaremetalNode{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Name: "",
        DriverInfo: MakeDriverInfo(),
        BMProperties: MakeBaremetalProperties(),
        InstanceUUID: "",
        InstanceInfo: MakeInstanceInfo(),
        Maintenance: false,
        MaintenanceReason: "",
        PowerState: "",
        TargetPowerState: "",
        ProvisionState: "",
        TargetProvisionState: "",
        ConsoleEnabled: false,
        CreatedAt: "",
        UpdatedAt: "",
        LastError: "",
        
    }
}

// MakeBaremetalNodeSlice() makes a slice of BaremetalNode
func MakeBaremetalNodeSlice() []*BaremetalNode {
    return []*BaremetalNode{}
}


