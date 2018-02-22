package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeServiceInterfaceTag makes ServiceInterfaceTag
func MakeServiceInterfaceTag() *ServiceInterfaceTag {
	return &ServiceInterfaceTag{
		//TODO(nati): Apply default
		InterfaceType: "",
	}
}

// MakeServiceInterfaceTag makes ServiceInterfaceTag
func InterfaceToServiceInterfaceTag(i interface{}) *ServiceInterfaceTag {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceInterfaceTag{
		//TODO(nati): Apply default
		InterfaceType: schema.InterfaceToString(m["interface_type"]),
	}
}

// MakeServiceInterfaceTagSlice() makes a slice of ServiceInterfaceTag
func MakeServiceInterfaceTagSlice() []*ServiceInterfaceTag {
	return []*ServiceInterfaceTag{}
}

// InterfaceToServiceInterfaceTagSlice() makes a slice of ServiceInterfaceTag
func InterfaceToServiceInterfaceTagSlice(i interface{}) []*ServiceInterfaceTag {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceInterfaceTag{}
	for _, item := range list {
		result = append(result, InterfaceToServiceInterfaceTag(item))
	}
	return result
}
