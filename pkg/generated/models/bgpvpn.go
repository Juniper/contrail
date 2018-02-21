package models


// MakeBGPVPN makes BGPVPN
func MakeBGPVPN() *BGPVPN{
    return &BGPVPN{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        RouteTargetList: MakeRouteTargetList(),
        ImportRouteTargetList: MakeRouteTargetList(),
        ExportRouteTargetList: MakeRouteTargetList(),
        BGPVPNType: "",
        
    }
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
    return []*BGPVPN{}
}


