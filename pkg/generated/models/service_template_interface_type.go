package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceType() *ServiceTemplateInterfaceType {
	return &ServiceTemplateInterfaceType{
		//TODO(nati): Apply default
		StaticRouteEnable:    false,
		SharedIP:             false,
		ServiceInterfaceType: "",
	}
}

// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
func InterfaceToServiceTemplateInterfaceType(i interface{}) *ServiceTemplateInterfaceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceTemplateInterfaceType{
		//TODO(nati): Apply default
		StaticRouteEnable:    schema.InterfaceToBool(m["static_route_enable"]),
		SharedIP:             schema.InterfaceToBool(m["shared_ip"]),
		ServiceInterfaceType: schema.InterfaceToString(m["service_interface_type"]),
	}
}

// MakeServiceTemplateInterfaceTypeSlice() makes a slice of ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceTypeSlice() []*ServiceTemplateInterfaceType {
	return []*ServiceTemplateInterfaceType{}
}

// InterfaceToServiceTemplateInterfaceTypeSlice() makes a slice of ServiceTemplateInterfaceType
func InterfaceToServiceTemplateInterfaceTypeSlice(i interface{}) []*ServiceTemplateInterfaceType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceTemplateInterfaceType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplateInterfaceType(item))
	}
	return result
}
