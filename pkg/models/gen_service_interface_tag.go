package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceInterfaceTag makes ServiceInterfaceTag
// nolint
func MakeServiceInterfaceTag() *ServiceInterfaceTag {
	return &ServiceInterfaceTag{
		//TODO(nati): Apply default
		InterfaceType: "",
	}
}

// MakeServiceInterfaceTag makes ServiceInterfaceTag
// nolint
func InterfaceToServiceInterfaceTag(i interface{}) *ServiceInterfaceTag {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceInterfaceTag{
		//TODO(nati): Apply default
		InterfaceType: common.InterfaceToString(m["interface_type"]),
	}
}

// MakeServiceInterfaceTagSlice() makes a slice of ServiceInterfaceTag
// nolint
func MakeServiceInterfaceTagSlice() []*ServiceInterfaceTag {
	return []*ServiceInterfaceTag{}
}

// InterfaceToServiceInterfaceTagSlice() makes a slice of ServiceInterfaceTag
// nolint
func InterfaceToServiceInterfaceTagSlice(i interface{}) []*ServiceInterfaceTag {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceInterfaceTag{}
	for _, item := range list {
		result = append(result, InterfaceToServiceInterfaceTag(item))
	}
	return result
}
