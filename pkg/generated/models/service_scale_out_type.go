package models

// ServiceScaleOutType

import "encoding/json"

// ServiceScaleOutType
type ServiceScaleOutType struct {
	AutoScale    bool `json:"auto_scale"`
	MaxInstances int  `json:"max_instances"`
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
		AutoScale:    false,
		MaxInstances: 0,
	}
}

// InterfaceToServiceScaleOutType makes ServiceScaleOutType from interface
func InterfaceToServiceScaleOutType(iData interface{}) *ServiceScaleOutType {
	data := iData.(map[string]interface{})
	return &ServiceScaleOutType{
		MaxInstances: data["max_instances"].(int),

		//{"description":"Maximum number of scale out factor(virtual machines). can be changed dynamically","type":"integer"}
		AutoScale: data["auto_scale"].(bool),

		//{"description":"Automatically change the number of virtual machines. Not implemented","type":"boolean"}

	}
}

// InterfaceToServiceScaleOutTypeSlice makes a slice of ServiceScaleOutType from interface
func InterfaceToServiceScaleOutTypeSlice(data interface{}) []*ServiceScaleOutType {
	list := data.([]interface{})
	result := MakeServiceScaleOutTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceScaleOutType(item))
	}
	return result
}

// MakeServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
	return []*ServiceScaleOutType{}
}
