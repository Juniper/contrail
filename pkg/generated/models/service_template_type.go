package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable,omitempty"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces,omitempty"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
	ImageName                 string                          `json:"image_name,omitempty"`
	Version                   int                             `json:"version,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
	ServiceScaling            bool                            `json:"service_scaling,omitempty"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type,omitempty"`
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
		ServiceMode: MakeServiceModeType(),
		Flavor:      "",

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		InstanceData:              "",
		OrderedInterfaces:         false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		ImageName:                 "",
		Version:                   0,
		ServiceType:               MakeServiceType(),
		ServiceScaling:            false,
		AvailabilityZoneEnable:    false,
		VrouterInstanceType:       MakeVRouterInstanceType(),
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
