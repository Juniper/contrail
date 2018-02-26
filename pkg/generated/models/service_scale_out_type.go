package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeServiceScaleOutType makes ServiceScaleOutType
func MakeServiceScaleOutType() *ServiceScaleOutType {
	return &ServiceScaleOutType{
		//TODO(nati): Apply default
		AutoScale:    false,
		MaxInstances: 0,
	}
}

// MakeServiceScaleOutType makes ServiceScaleOutType
func InterfaceToServiceScaleOutType(i interface{}) *ServiceScaleOutType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceScaleOutType{
		//TODO(nati): Apply default
		AutoScale:    schema.InterfaceToBool(m["auto_scale"]),
		MaxInstances: schema.InterfaceToInt64(m["max_instances"]),
	}
}

// MakeServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
	return []*ServiceScaleOutType{}
}

// InterfaceToServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
func InterfaceToServiceScaleOutTypeSlice(i interface{}) []*ServiceScaleOutType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceScaleOutType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceScaleOutType(item))
	}
	return result
}
