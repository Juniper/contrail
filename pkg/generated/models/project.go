package models


// MakeProject makes Project
func MakeProject() *Project{
    return &Project{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VxlanRouting: false,
        AlarmEnable: false,
        Quota: MakeQuotaType(),
        
    }
}

// MakeProjectSlice() makes a slice of Project
func MakeProjectSlice() []*Project {
    return []*Project{}
}


