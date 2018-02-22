package models


// MakeServiceGroup makes ServiceGroup
func MakeServiceGroup() *ServiceGroup{
    return &ServiceGroup{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
        
    }
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
    return []*ServiceGroup{}
}


