package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	ImageName                 string                          `json:"image_name,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
	Version                   int                             `json:"version,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
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
		OrderedInterfaces:         false,
		ImageName:                 "",
		Flavor:                    "",
		InstanceData:              "",
		ServiceVirtualizationType: MakeServiceVirtualizationType(),

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ServiceMode:            MakeServiceModeType(),
		Version:                0,
		ServiceType:            MakeServiceType(),
		ServiceScaling:         false,
		VrouterInstanceType:    MakeVRouterInstanceType(),
		AvailabilityZoneEnable: false,
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
