package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type"`
	Version                   int                             `json:"version"`
	ServiceType               ServiceType                     `json:"service_type"`
	InstanceData              string                          `json:"instance_data"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type"`
	ImageName                 string                          `json:"image_name"`
	ServiceMode               ServiceModeType                 `json:"service_mode"`
	Flavor                    string                          `json:"flavor"`
	ServiceScaling            bool                            `json:"service_scaling"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
}

// String returns json representation of the object
func (model *ServiceTemplateType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceTemplateType makes ServiceTemplateType
func MakeServiceTemplateType() *ServiceTemplateType {
	return &ServiceTemplateType{
		//TODO(nati): Apply default
		InstanceData:              "",
		OrderedInterfaces:         false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		Version:                   0,
		ServiceType:               MakeServiceType(),
		AvailabilityZoneEnable:    false,

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:           "",
		ServiceMode:         MakeServiceModeType(),
		Flavor:              "",
		ServiceScaling:      false,
		VrouterInstanceType: MakeVRouterInstanceType(),
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
