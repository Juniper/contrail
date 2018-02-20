package models

// ServiceScaleOutType

// ServiceScaleOutType
//proteus:generate
type ServiceScaleOutType struct {
	AutoScale    bool `json:"auto_scale"`
	MaxInstances int  `json:"max_instances,omitempty"`
}

// MakeServiceScaleOutType makes ServiceScaleOutType
func MakeServiceScaleOutType() *ServiceScaleOutType {
	return &ServiceScaleOutType{
		//TODO(nati): Apply default
		AutoScale:    false,
		MaxInstances: 0,
	}
}

// MakeServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
	return []*ServiceScaleOutType{}
}
