package models


// MakeLoadbalancerPool makes LoadbalancerPool
func MakeLoadbalancerPool() *LoadbalancerPool{
    return &LoadbalancerPool{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        LoadbalancerPoolProperties: MakeLoadbalancerPoolType(),
        LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
        LoadbalancerPoolProvider: "",
        
    }
}

// MakeLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
    return []*LoadbalancerPool{}
}


