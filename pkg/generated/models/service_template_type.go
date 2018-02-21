package models


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

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
    return []*ServiceTemplateType{}
}


