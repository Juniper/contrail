package models


// MakeServiceEndpoint makes ServiceEndpoint
func MakeServiceEndpoint() *ServiceEndpoint{
    return &ServiceEndpoint{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
    }
}

// MakeServiceEndpointSlice() makes a slice of ServiceEndpoint
func MakeServiceEndpointSlice() []*ServiceEndpoint {
    return []*ServiceEndpoint{}
}


