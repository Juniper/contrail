package models


// MakeServiceInstance makes ServiceInstance
func MakeServiceInstance() *ServiceInstance{
    return &ServiceInstance{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceInstanceBindings: MakeKeyValuePairs(),
        ServiceInstanceProperties: MakeServiceInstanceType(),
        
    }
}

// MakeServiceInstanceSlice() makes a slice of ServiceInstance
func MakeServiceInstanceSlice() []*ServiceInstance {
    return []*ServiceInstance{}
}


