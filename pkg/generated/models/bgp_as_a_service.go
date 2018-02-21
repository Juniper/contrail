package models


// MakeBGPAsAService makes BGPAsAService
func MakeBGPAsAService() *BGPAsAService{
    return &BGPAsAService{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        BgpaasShared: false,
        BgpaasSessionAttributes: "",
        BgpaasSuppressRouteAdvertisement: false,
        BgpaasIpv4MappedIpv6Nexthop: false,
        BgpaasIPAddress: "",
        AutonomousSystem: 0,
        
    }
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
    return []*BGPAsAService{}
}


