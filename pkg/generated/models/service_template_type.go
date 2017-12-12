package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	InstanceData              string                          `json:"instance_data"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type"`
	Version                   int                             `json:"version"`
	ServiceType               ServiceType                     `json:"service_type"`
	ServiceScaling            bool                            `json:"service_scaling"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type"`
	ImageName                 string                          `json:"image_name"`
	ServiceMode               ServiceModeType                 `json:"service_mode"`
	Flavor                    string                          `json:"flavor"`
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
		ImageName:                 "",
		ServiceMode:               MakeServiceModeType(),
		Flavor:                    "",
		ServiceScaling:            false,
		VrouterInstanceType:       MakeVRouterInstanceType(),
		AvailabilityZoneEnable:    false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		Version:           0,
		ServiceType:       MakeServiceType(),
		InstanceData:      "",
		OrderedInterfaces: false,
	}
}

// InterfaceToServiceTemplateType makes ServiceTemplateType from interface
func InterfaceToServiceTemplateType(iData interface{}) *ServiceTemplateType {
	data := iData.(map[string]interface{})
	return &ServiceTemplateType{
		AvailabilityZoneEnable: data["availability_zone_enable"].(bool),

		//{"description":"Enable availability zone for version 1 service instances","type":"boolean"}
		ServiceVirtualizationType: InterfaceToServiceVirtualizationType(data["service_virtualization_type"]),

		//{"description":"Service virtualization type decides how individual service instances are instantiated","type":"string","enum":["virtual-machine","network-namespace","vrouter-instance","physical-device"]}
		ImageName: data["image_name"].(string),

		//{"description":"Glance image name for the service virtual machine, Version 1 only","type":"string"}
		ServiceMode: InterfaceToServiceModeType(data["service_mode"]),

		//{"description":"Service instance mode decides how packets are forwarded across the service","type":"string","enum":["transparent","in-network","in-network-nat"]}
		Flavor: data["flavor"].(string),

		//{"description":"Nova flavor used for service virtual machines, Version 1 only","type":"string"}
		ServiceScaling: data["service_scaling"].(bool),

		//{"description":"Enable scaling of service virtual machines, Version 1 only","type":"boolean"}
		VrouterInstanceType: InterfaceToVRouterInstanceType(data["vrouter_instance_type"]),

		//{"description":"Mechanism used to spawn service instance, when vrouter is spawning instances.Allowed values libvirt-qemu, docker or netns","type":"string","enum":["libvirt-qemu","docker"]}
		InstanceData: data["instance_data"].(string),

		//{"description":"Opaque string (typically in json format) used to spawn a vrouter-instance.","type":"string"}
		OrderedInterfaces: data["ordered_interfaces"].(bool),

		//{"description":"Deprecated","type":"boolean"}

		InterfaceType: InterfaceToServiceTemplateInterfaceTypeSlice(data["interface_type"]),

		//{"description":"List of interfaces which decided number of interfaces and type","type":"array","item":{"type":"object","properties":{"service_interface_type":{"type":"string"},"shared_ip":{"type":"boolean"},"static_route_enable":{"type":"boolean"}}}}
		Version: data["version"].(int),

		//{"type":"integer"}
		ServiceType: InterfaceToServiceType(data["service_type"]),

		//{"description":"Service instance mode decides how routing happens across the service","type":"string","enum":["firewall","analyzer","source-nat","loadbalancer"]}

	}
}

// InterfaceToServiceTemplateTypeSlice makes a slice of ServiceTemplateType from interface
func InterfaceToServiceTemplateTypeSlice(data interface{}) []*ServiceTemplateType {
	list := data.([]interface{})
	result := MakeServiceTemplateTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplateType(item))
	}
	return result
}

// MakeServiceTemplateTypeSlice() makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
