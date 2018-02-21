package models


// MakeVirtualMachine makes VirtualMachine
func MakeVirtualMachine() *VirtualMachine{
    return &VirtualMachine{
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

// MakeVirtualMachineSlice() makes a slice of VirtualMachine
func MakeVirtualMachineSlice() []*VirtualMachine {
    return []*VirtualMachine{}
}


