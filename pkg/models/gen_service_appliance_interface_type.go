package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceApplianceInterfaceType makes ServiceApplianceInterfaceType
// nolint
func MakeServiceApplianceInterfaceType() *ServiceApplianceInterfaceType {
	return &ServiceApplianceInterfaceType{
		//TODO(nati): Apply default
		InterfaceType: "",
	}
}

// MakeServiceApplianceInterfaceType makes ServiceApplianceInterfaceType
// nolint
func InterfaceToServiceApplianceInterfaceType(i interface{}) *ServiceApplianceInterfaceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceApplianceInterfaceType{
		//TODO(nati): Apply default
		InterfaceType: common.InterfaceToString(m["interface_type"]),
	}
}

// MakeServiceApplianceInterfaceTypeSlice() makes a slice of ServiceApplianceInterfaceType
// nolint
func MakeServiceApplianceInterfaceTypeSlice() []*ServiceApplianceInterfaceType {
	return []*ServiceApplianceInterfaceType{}
}

// InterfaceToServiceApplianceInterfaceTypeSlice() makes a slice of ServiceApplianceInterfaceType
// nolint
func InterfaceToServiceApplianceInterfaceTypeSlice(i interface{}) []*ServiceApplianceInterfaceType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceApplianceInterfaceType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceApplianceInterfaceType(item))
	}
	return result
}
