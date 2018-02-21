package models


// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceType() *ServiceTemplateInterfaceType{
    return &ServiceTemplateInterfaceType{
    //TODO(nati): Apply default
    StaticRouteEnable: false,
        SharedIP: false,
        ServiceInterfaceType: "",
        
    }
}

// MakeServiceTemplateInterfaceTypeSlice() makes a slice of ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceTypeSlice() []*ServiceTemplateInterfaceType {
    return []*ServiceTemplateInterfaceType{}
}


