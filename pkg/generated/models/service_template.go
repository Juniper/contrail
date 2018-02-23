package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


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

// MakeServiceTemplate makes ServiceTemplate
func InterfaceToServiceTemplate(i interface{}) *ServiceTemplate{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceTemplate{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        ServiceTemplateProperties: InterfaceToServiceTemplateType(m["service_template_properties"]),
        
    }
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
    return []*ServiceTemplate{}
}

// InterfaceToServiceTemplateSlice() makes a slice of ServiceTemplate
func InterfaceToServiceTemplateSlice(i interface{}) []*ServiceTemplate {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceTemplate{}
    for _, item := range list {
        result = append(result, InterfaceToServiceTemplate(item) )
    }
    return result
}



