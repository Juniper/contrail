package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
// nolint
func MakeServiceTemplateInterfaceType() *ServiceTemplateInterfaceType {
	return &ServiceTemplateInterfaceType{
		//TODO(nati): Apply default
		StaticRouteEnable:    false,
		SharedIP:             false,
		ServiceInterfaceType: "",
	}
}

// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
// nolint
func InterfaceToServiceTemplateInterfaceType(i interface{}) *ServiceTemplateInterfaceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceTemplateInterfaceType{
		//TODO(nati): Apply default
		StaticRouteEnable:    common.InterfaceToBool(m["static_route_enable"]),
		SharedIP:             common.InterfaceToBool(m["shared_ip"]),
		ServiceInterfaceType: common.InterfaceToString(m["service_interface_type"]),
	}
}

// MakeServiceTemplateInterfaceTypeSlice() makes a slice of ServiceTemplateInterfaceType
// nolint
func MakeServiceTemplateInterfaceTypeSlice() []*ServiceTemplateInterfaceType {
	return []*ServiceTemplateInterfaceType{}
}

// InterfaceToServiceTemplateInterfaceTypeSlice() makes a slice of ServiceTemplateInterfaceType
// nolint
func InterfaceToServiceTemplateInterfaceTypeSlice(i interface{}) []*ServiceTemplateInterfaceType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceTemplateInterfaceType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplateInterfaceType(item))
	}
	return result
}
