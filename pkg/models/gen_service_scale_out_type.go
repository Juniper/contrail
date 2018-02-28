package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceScaleOutType makes ServiceScaleOutType
// nolint
func MakeServiceScaleOutType() *ServiceScaleOutType {
	return &ServiceScaleOutType{
		//TODO(nati): Apply default
		AutoScale:    false,
		MaxInstances: 0,
	}
}

// MakeServiceScaleOutType makes ServiceScaleOutType
// nolint
func InterfaceToServiceScaleOutType(i interface{}) *ServiceScaleOutType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceScaleOutType{
		//TODO(nati): Apply default
		AutoScale:    common.InterfaceToBool(m["auto_scale"]),
		MaxInstances: common.InterfaceToInt64(m["max_instances"]),
	}
}

// MakeServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
// nolint
func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
	return []*ServiceScaleOutType{}
}

// InterfaceToServiceScaleOutTypeSlice() makes a slice of ServiceScaleOutType
// nolint
func InterfaceToServiceScaleOutTypeSlice(i interface{}) []*ServiceScaleOutType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceScaleOutType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceScaleOutType(item))
	}
	return result
}
