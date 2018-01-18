package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	ServiceScaling            bool                            `json:"service_scaling"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type,omitempty"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	ImageName                 string                          `json:"image_name,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
	Version                   int                             `json:"version,omitempty"`
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
		ServiceType:               MakeServiceType(),
		Flavor:                    "",
		InstanceData:              "",
		OrderedInterfaces:         false,
		ServiceMode:               MakeServiceModeType(),
		Version:                   0,
		ServiceScaling:            false,
		VrouterInstanceType:       MakeVRouterInstanceType(),
		AvailabilityZoneEnable:    false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName: "",
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
