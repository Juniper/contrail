package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	ServiceMode               ServiceModeType                 `json:"service_mode,omitempty"`
	Version                   int                             `json:"version,omitempty"`
	ServiceType               ServiceType                     `json:"service_type,omitempty"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable,omitempty"`
	InstanceData              string                          `json:"instance_data,omitempty"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type,omitempty"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
	ImageName                 string                          `json:"image_name,omitempty"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces,omitempty"`
	Flavor                    string                          `json:"flavor,omitempty"`
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
		AvailabilityZoneEnable:    false,
		InstanceData:              "",
		ServiceVirtualizationType: MakeServiceVirtualizationType(),

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:           "",
		ServiceMode:         MakeServiceModeType(),
		Version:             0,
		ServiceType:         MakeServiceType(),
		OrderedInterfaces:   false,
		Flavor:              "",
		ServiceScaling:      false,
		VrouterInstanceType: MakeVRouterInstanceType(),
	}
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
