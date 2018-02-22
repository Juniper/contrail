package models


// MakeServiceTemplate makes ServiceTemplate
func MakeServiceTemplate() *ServiceTemplate{
    return &ServiceTemplate{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceTemplateProperties: MakeServiceTemplateType(),
        
    }
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
    return []*ServiceTemplate{}
}


