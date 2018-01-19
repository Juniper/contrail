package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
	ServiceScaling            bool                            `json:"service_scaling"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	Version                   int                             `json:"version,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type,omitempty"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	ImageName                 string                          `json:"image_name,omitempty"`
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
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

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:                 "",
		ServiceMode:               MakeServiceModeType(),
		Version:                   0,
		ServiceType:               MakeServiceType(),
		VrouterInstanceType:       MakeVRouterInstanceType(),
		AvailabilityZoneEnable:    false,
		InstanceData:              "",
		OrderedInterfaces:         false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		Flavor:         "",
		ServiceScaling: false,
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
