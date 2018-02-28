package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceTemplateType makes ServiceTemplateType
// nolint
func MakeServiceTemplateType() *ServiceTemplateType {
	return &ServiceTemplateType{
		//TODO(nati): Apply default
		AvailabilityZoneEnable:    false,
		InstanceData:              "",
		OrderedInterfaces:         false,
		ServiceVirtualizationType: "",

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:           "",
		ServiceMode:         "",
		Version:             0,
		ServiceType:         "",
		Flavor:              "",
		ServiceScaling:      false,
		VrouterInstanceType: "",
	}
}

// MakeServiceTemplateType makes ServiceTemplateType
// nolint
func InterfaceToServiceTemplateType(i interface{}) *ServiceTemplateType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceTemplateType{
		//TODO(nati): Apply default
		AvailabilityZoneEnable:    common.InterfaceToBool(m["availability_zone_enable"]),
		InstanceData:              common.InterfaceToString(m["instance_data"]),
		OrderedInterfaces:         common.InterfaceToBool(m["ordered_interfaces"]),
		ServiceVirtualizationType: common.InterfaceToString(m["service_virtualization_type"]),

		InterfaceType: InterfaceToServiceTemplateInterfaceTypeSlice(m["interface_type"]),

		ImageName:           common.InterfaceToString(m["image_name"]),
		ServiceMode:         common.InterfaceToString(m["service_mode"]),
		Version:             common.InterfaceToInt64(m["version"]),
		ServiceType:         common.InterfaceToString(m["service_type"]),
		Flavor:              common.InterfaceToString(m["flavor"]),
		ServiceScaling:      common.InterfaceToBool(m["service_scaling"]),
		VrouterInstanceType: common.InterfaceToString(m["vrouter_instance_type"]),
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
// nolint
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}

// InterfaceToServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
// nolint
func InterfaceToServiceTemplateTypeSlice(i interface{}) []*ServiceTemplateType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceTemplateType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplateType(item))
	}
	return result
}
