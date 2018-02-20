package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


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

// MakeServiceAppliance makes ServiceAppliance
func InterfaceToServiceAppliance(i interface{}) *ServiceAppliance{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceAppliance{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        ServiceApplianceUserCredentials: InterfaceToUserCredentials(m["service_appliance_user_credentials"]),
        ServiceApplianceIPAddress: schema.InterfaceToString(m["service_appliance_ip_address"]),
        ServiceApplianceProperties: InterfaceToKeyValuePairs(m["service_appliance_properties"]),
        
    }
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
    return []*ServiceAppliance{}
}

// InterfaceToServiceApplianceSlice() makes a slice of ServiceAppliance
func InterfaceToServiceApplianceSlice(i interface{}) []*ServiceAppliance {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceAppliance{}
    for _, item := range list {
        result = append(result, InterfaceToServiceAppliance(item) )
    }
    return result
}



