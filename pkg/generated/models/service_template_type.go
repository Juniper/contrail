package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeServiceTemplateType makes ServiceTemplateType
func MakeServiceTemplateType() *ServiceTemplateType{
    return &ServiceTemplateType{
    //TODO(nati): Apply default
    AvailabilityZoneEnable: false,
        InstanceData: "",
        OrderedInterfaces: false,
        ServiceVirtualizationType: "",
        
            
                InterfaceType:  MakeServiceTemplateInterfaceTypeSlice(),
            
        ImageName: "",
        ServiceMode: "",
        Version: 0,
        ServiceType: "",
        Flavor: "",
        ServiceScaling: false,
        VrouterInstanceType: "",
        
    }
}

// MakeServiceTemplateType makes ServiceTemplateType
func InterfaceToServiceTemplateType(i interface{}) *ServiceTemplateType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceTemplateType{
    //TODO(nati): Apply default
    AvailabilityZoneEnable: schema.InterfaceToBool(m["availability_zone_enable"]),
        InstanceData: schema.InterfaceToString(m["instance_data"]),
        OrderedInterfaces: schema.InterfaceToBool(m["ordered_interfaces"]),
        ServiceVirtualizationType: schema.InterfaceToString(m["service_virtualization_type"]),
        
            
                InterfaceType:  InterfaceToServiceTemplateInterfaceTypeSlice(m["interface_type"]),
            
        ImageName: schema.InterfaceToString(m["image_name"]),
        ServiceMode: schema.InterfaceToString(m["service_mode"]),
        Version: schema.InterfaceToInt64(m["version"]),
        ServiceType: schema.InterfaceToString(m["service_type"]),
        Flavor: schema.InterfaceToString(m["flavor"]),
        ServiceScaling: schema.InterfaceToBool(m["service_scaling"]),
        VrouterInstanceType: schema.InterfaceToString(m["vrouter_instance_type"]),
        
    }
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
    return []*ServiceTemplateType{}
}

// InterfaceToServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func InterfaceToServiceTemplateTypeSlice(i interface{}) []*ServiceTemplateType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceTemplateType{}
    for _, item := range list {
        result = append(result, InterfaceToServiceTemplateType(item) )
    }
    return result
}



