package models


// MakeServiceInstanceType makes ServiceInstanceType
func MakeServiceInstanceType() *ServiceInstanceType{
    return &ServiceInstanceType{
    //TODO(nati): Apply default
    RightVirtualNetwork: "",
        RightIPAddress: "",
        AvailabilityZone: "",
        ManagementVirtualNetwork: "",
        ScaleOut: MakeServiceScaleOutType(),
        HaMode: "",
        VirtualRouterID: "",
        
            
                InterfaceList:  MakeServiceInstanceInterfaceTypeSlice(),
            
        LeftIPAddress: "",
        LeftVirtualNetwork: "",
        AutoPolicy: false,
        
    }
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
    return []*ServiceInstanceType{}
}


