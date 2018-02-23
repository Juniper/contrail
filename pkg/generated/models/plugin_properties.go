package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakePluginProperties makes PluginProperties
func MakePluginProperties() *PluginProperties{
    return &PluginProperties{
    //TODO(nati): Apply default
    
            
                PluginProperty:  MakePluginPropertySlice(),
            
        
    }
}

// MakePluginProperties makes PluginProperties
func InterfaceToPluginProperties(i interface{}) *PluginProperties{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &PluginProperties{
    //TODO(nati): Apply default
    
            
                PluginProperty:  InterfaceToPluginPropertySlice(m["plugin_property"]),
            
        
    }
}

// MakePluginPropertiesSlice() makes a slice of PluginProperties
func MakePluginPropertiesSlice() []*PluginProperties {
    return []*PluginProperties{}
}

// InterfaceToPluginPropertiesSlice() makes a slice of PluginProperties
func InterfaceToPluginPropertiesSlice(i interface{}) []*PluginProperties {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*PluginProperties{}
    for _, item := range list {
        result = append(result, InterfaceToPluginProperties(item) )
    }
    return result
}



