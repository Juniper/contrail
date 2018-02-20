package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLoadbalancerPoolType makes LoadbalancerPoolType
func MakeLoadbalancerPoolType() *LoadbalancerPoolType{
    return &LoadbalancerPoolType{
    //TODO(nati): Apply default
    Status: "",
        Protocol: "",
        SubnetID: "",
        SessionPersistence: "",
        AdminState: false,
        PersistenceCookieName: "",
        StatusDescription: "",
        LoadbalancerMethod: "",
        
    }
}

// MakeLoadbalancerPoolType makes LoadbalancerPoolType
func InterfaceToLoadbalancerPoolType(i interface{}) *LoadbalancerPoolType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &LoadbalancerPoolType{
    //TODO(nati): Apply default
    Status: schema.InterfaceToString(m["status"]),
        Protocol: schema.InterfaceToString(m["protocol"]),
        SubnetID: schema.InterfaceToString(m["subnet_id"]),
        SessionPersistence: schema.InterfaceToString(m["session_persistence"]),
        AdminState: schema.InterfaceToBool(m["admin_state"]),
        PersistenceCookieName: schema.InterfaceToString(m["persistence_cookie_name"]),
        StatusDescription: schema.InterfaceToString(m["status_description"]),
        LoadbalancerMethod: schema.InterfaceToString(m["loadbalancer_method"]),
        
    }
}

// MakeLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
    return []*LoadbalancerPoolType{}
}

// InterfaceToLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
func InterfaceToLoadbalancerPoolTypeSlice(i interface{}) []*LoadbalancerPoolType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*LoadbalancerPoolType{}
    for _, item := range list {
        result = append(result, InterfaceToLoadbalancerPoolType(item) )
    }
    return result
}



