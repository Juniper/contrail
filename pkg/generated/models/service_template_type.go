package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	ImageName                 string                          `json:"image_name,omitempty"`
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
	Version                   int                             `json:"version,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	ServiceScaling            bool                            `json:"service_scaling"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type,omitempty"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
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
		ServiceMode:            MakeServiceModeType(),
		Version:                0,
		ServiceType:            MakeServiceType(),
		Flavor:                 "",
		AvailabilityZoneEnable: false,
		InstanceData:           "",

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:                 "",
		ServiceScaling:            false,
		OrderedInterfaces:         false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		VrouterInstanceType:       MakeVRouterInstanceType(),
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
