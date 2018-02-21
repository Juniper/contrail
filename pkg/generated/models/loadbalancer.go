package models


// MakeLoadbalancer makes Loadbalancer
func MakeLoadbalancer() *Loadbalancer{
    return &Loadbalancer{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        LoadbalancerProperties: MakeLoadbalancerType(),
        LoadbalancerProvider: "",
        
    }
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
    return []*Loadbalancer{}
}


