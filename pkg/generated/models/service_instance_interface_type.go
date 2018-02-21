package models


// MakeServiceInstanceInterfaceType makes ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceType() *ServiceInstanceInterfaceType{
    return &ServiceInstanceInterfaceType{
    //TODO(nati): Apply default
    VirtualNetwork: "",
        IPAddress: "",
        AllowedAddressPairs: MakeAllowedAddressPairs(),
        StaticRoutes: MakeRouteTableType(),
        
    }
}

// MakeServiceInstanceInterfaceTypeSlice() makes a slice of ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
    return []*ServiceInstanceInterfaceType{}
}


