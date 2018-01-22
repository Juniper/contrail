package models

// ServiceScaleOutType

import "encoding/json"

// ServiceScaleOutType
type ServiceScaleOutType struct {
	AutoScale    bool `json:"auto_scale"`
	MaxInstances int  `json:"max_instances,omitempty"`
}

// String returns json representation of the object
func (model *ServiceScaleOutType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceScaleOutType makes ServiceScaleOutType
func MakeServiceScaleOutType() *ServiceScaleOutType {
	return &ServiceScaleOutType{
		//TODO(nati): Apply default
		MaxInstances: 0,
		AutoScale:    false,
	}
}

// MakeServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
	return []*ServiceScaleOutType{}
}
