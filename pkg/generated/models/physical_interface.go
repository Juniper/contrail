package models


// MakePhysicalInterface makes PhysicalInterface
func MakePhysicalInterface() *PhysicalInterface{
    return &PhysicalInterface{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        EthernetSegmentIdentifier: "",
        
    }
}

// MakePhysicalInterfaceSlice() makes a slice of PhysicalInterface
func MakePhysicalInterfaceSlice() []*PhysicalInterface {
    return []*PhysicalInterface{}
}


