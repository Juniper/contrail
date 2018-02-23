package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeServiceHealthCheckType makes ServiceHealthCheckType
func MakeServiceHealthCheckType() *ServiceHealthCheckType{
    return &ServiceHealthCheckType{
    //TODO(nati): Apply default
    DelayUsecs: 0,
        TimeoutUsecs: 0,
        Enabled: false,
        Delay: 0,
        ExpectedCodes: "",
        MaxRetries: 0,
        HealthCheckType: "",
        HTTPMethod: "",
        Timeout: 0,
        URLPath: "",
        MonitorType: "",
        
    }
}

// MakeServiceHealthCheckType makes ServiceHealthCheckType
func InterfaceToServiceHealthCheckType(i interface{}) *ServiceHealthCheckType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceHealthCheckType{
    //TODO(nati): Apply default
    DelayUsecs: schema.InterfaceToInt64(m["delayUsecs"]),
        TimeoutUsecs: schema.InterfaceToInt64(m["timeoutUsecs"]),
        Enabled: schema.InterfaceToBool(m["enabled"]),
        Delay: schema.InterfaceToInt64(m["delay"]),
        ExpectedCodes: schema.InterfaceToString(m["expected_codes"]),
        MaxRetries: schema.InterfaceToInt64(m["max_retries"]),
        HealthCheckType: schema.InterfaceToString(m["health_check_type"]),
        HTTPMethod: schema.InterfaceToString(m["http_method"]),
        Timeout: schema.InterfaceToInt64(m["timeout"]),
        URLPath: schema.InterfaceToString(m["url_path"]),
        MonitorType: schema.InterfaceToString(m["monitor_type"]),
        
    }
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
    return []*ServiceHealthCheckType{}
}

// InterfaceToServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func InterfaceToServiceHealthCheckTypeSlice(i interface{}) []*ServiceHealthCheckType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceHealthCheckType{}
    for _, item := range list {
        result = append(result, InterfaceToServiceHealthCheckType(item) )
    }
    return result
}



