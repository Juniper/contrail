package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeWidget makes Widget
func MakeWidget() *Widget{
    return &Widget{
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
        ContentConfig: "",
        LayoutConfig: "",
        
    }
}

// MakeWidget makes Widget
func InterfaceToWidget(i interface{}) *Widget{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &Widget{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        ContainerConfig: schema.InterfaceToString(m["container_config"]),
        ContentConfig: schema.InterfaceToString(m["content_config"]),
        LayoutConfig: schema.InterfaceToString(m["layout_config"]),
        
    }
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
    return []*Widget{}
}

// InterfaceToWidgetSlice() makes a slice of Widget
func InterfaceToWidgetSlice(i interface{}) []*Widget {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*Widget{}
    for _, item := range list {
        result = append(result, InterfaceToWidget(item) )
    }
    return result
}



