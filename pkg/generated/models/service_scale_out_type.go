package models


// MakeServiceScaleOutType makes ServiceScaleOutType
func MakeServiceScaleOutType() *ServiceScaleOutType{
    return &ServiceScaleOutType{
    //TODO(nati): Apply default
    AutoScale: false,
        MaxInstances: 0,
        
    }
}

// MakeServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
    return []*ServiceScaleOutType{}
}


