package models


// MakeVirtualNetwork makes VirtualNetwork
func MakeVirtualNetwork() *VirtualNetwork{
    return &VirtualNetwork{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VirtualNetworkProperties: MakeVirtualNetworkType(),
        EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
        VirtualNetworkNetworkID: 0,
        AddressAllocationMode: "",
        PBBEvpnEnable: false,
        RouterExternal: false,
        ImportRouteTargetList: MakeRouteTargetList(),
        MacAgingTime: 0,
        ProviderProperties: MakeProviderDetails(),
        RouteTargetList: MakeRouteTargetList(),
        MacLearningEnabled: false,
        ExportRouteTargetList: MakeRouteTargetList(),
        FloodUnknownUnicast: false,
        PBBEtreeEnable: false,
        Layer2ControlWord: false,
        ExternalIpam: false,
        PortSecurityEnabled: false,
        MacMoveControl: MakeMACMoveLimitControlType(),
        MultiPolicyServiceChainsEnabled: false,
        MacLimitControl: MakeMACLimitControlType(),
        IsShared: false,
        
    }
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
    return []*VirtualNetwork{}
}


