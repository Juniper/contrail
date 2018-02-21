package models


// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesType() *VirtualMachineInterfacePropertiesType{
    return &VirtualMachineInterfacePropertiesType{
    //TODO(nati): Apply default
    SubInterfaceVlanTag: 0,
        LocalPreference: 0,
        InterfaceMirror: MakeInterfaceMirrorType(),
        ServiceInterfaceType: "",
        
    }
}

// MakeVirtualMachineInterfacePropertiesTypeSlice() makes a slice of VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
    return []*VirtualMachineInterfacePropertiesType{}
}


