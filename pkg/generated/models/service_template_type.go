package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	ImageName                 string                          `json:"image_name,omitempty"`
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
	Version                   int                             `json:"version,omitempty"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
	ServiceScaling            bool                            `json:"service_scaling"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type,omitempty"`
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
		ServiceScaling:            false,
		VrouterInstanceType:       MakeVRouterInstanceType(),
		AvailabilityZoneEnable:    false,
		InstanceData:              "",
		OrderedInterfaces:         false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		ServiceType:               MakeServiceType(),
		Flavor:                    "",

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:   "",
		ServiceMode: MakeServiceModeType(),
		Version:     0,
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
