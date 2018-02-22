package models


// MakeServiceAppliance makes ServiceAppliance
func MakeServiceAppliance() *ServiceAppliance{
    return &ServiceAppliance{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceApplianceUserCredentials: MakeUserCredentials(),
        ServiceApplianceIPAddress: "",
        ServiceApplianceProperties: MakeKeyValuePairs(),
        
    }
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
    return []*ServiceAppliance{}
}


