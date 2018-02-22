package models


// MakeDashboard makes Dashboard
func MakeDashboard() *Dashboard{
    return &Dashboard{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ContainerConfig: "",
        
    }
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
    return []*Dashboard{}
}


