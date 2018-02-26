package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeServiceApplianceInterfaceType makes ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceType() *ServiceApplianceInterfaceType {
	return &ServiceApplianceInterfaceType{
		//TODO(nati): Apply default
		InterfaceType: "",
	}
}

// MakeServiceApplianceInterfaceType makes ServiceApplianceInterfaceType
func InterfaceToServiceApplianceInterfaceType(i interface{}) *ServiceApplianceInterfaceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceApplianceInterfaceType{
		//TODO(nati): Apply default
		InterfaceType: schema.InterfaceToString(m["interface_type"]),
	}
}

// MakeServiceApplianceInterfaceTypeSlice() makes a slice of ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceTypeSlice() []*ServiceApplianceInterfaceType {
	return []*ServiceApplianceInterfaceType{}
}

// InterfaceToServiceApplianceInterfaceTypeSlice() makes a slice of ServiceApplianceInterfaceType
func InterfaceToServiceApplianceInterfaceTypeSlice(i interface{}) []*ServiceApplianceInterfaceType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceApplianceInterfaceType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceApplianceInterfaceType(item))
	}
	return result
}
