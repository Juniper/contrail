package models


// MakeNetworkIpam makes NetworkIpam
func MakeNetworkIpam() *NetworkIpam{
    return &NetworkIpam{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        NetworkIpamMGMT: MakeIpamType(),
        IpamSubnets: MakeIpamSubnets(),
        IpamSubnetMethod: "",
        
    }
}

// MakeNetworkIpamSlice() makes a slice of NetworkIpam
func MakeNetworkIpamSlice() []*NetworkIpam {
    return []*NetworkIpam{}
}


