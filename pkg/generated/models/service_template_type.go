package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
	ServiceScaling            bool                            `json:"service_scaling"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	ImageName                 string                          `json:"image_name,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type,omitempty"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
	Version                   int                             `json:"version,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
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
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		Version:                   0,
		ServiceType:               MakeServiceType(),
		Flavor:                    "",
		VrouterInstanceType:       MakeVRouterInstanceType(),
		AvailabilityZoneEnable:    false,
		OrderedInterfaces:         false,

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:      "",
		ServiceMode:    MakeServiceModeType(),
		ServiceScaling: false,
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
