package models


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

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
    return []*ServiceHealthCheckType{}
}


