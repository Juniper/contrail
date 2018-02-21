package models


// MakeVirtualMachineInterface makes VirtualMachineInterface
func MakeVirtualMachineInterface() *VirtualMachineInterface{
    return &VirtualMachineInterface{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
        VirtualMachineInterfaceHostRoutes: MakeRouteTableType(),
        VirtualMachineInterfaceMacAddresses: MakeMacAddressesType(),
        VirtualMachineInterfaceDHCPOptionList: MakeDhcpOptionsListType(),
        VirtualMachineInterfaceBindings: MakeKeyValuePairs(),
        VirtualMachineInterfaceDisablePolicy: false,
        VirtualMachineInterfaceAllowedAddressPairs: MakeAllowedAddressPairs(),
        VirtualMachineInterfaceFatFlowProtocols: MakeFatFlowProtocols(),
        VlanTagBasedBridgeDomain: false,
        VirtualMachineInterfaceDeviceOwner: "",
        VRFAssignTable: MakeVrfAssignTableType(),
        PortSecurityEnabled: false,
        VirtualMachineInterfaceProperties: MakeVirtualMachineInterfacePropertiesType(),
        
    }
}

// MakeVirtualMachineInterfaceSlice() makes a slice of VirtualMachineInterface
func MakeVirtualMachineInterfaceSlice() []*VirtualMachineInterface {
    return []*VirtualMachineInterface{}
}


