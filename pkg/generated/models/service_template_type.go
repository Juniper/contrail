package models

// ServiceTemplateType

import "encoding/json"

type ServiceTemplateType struct {
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
	InstanceData              string                          `json:"instance_data"`
	Version                   int                             `json:"version"`
	ServiceType               ServiceType                     `json:"service_type"`
	ServiceMode               ServiceModeType                 `json:"service_mode"`
	Flavor                    string                          `json:"flavor"`
	ServiceScaling            bool                            `json:"service_scaling"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type"`
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type"`
	ImageName                 string                          `json:"image_name"`
}

func (model *ServiceTemplateType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeServiceTemplateType() *ServiceTemplateType {
	return &ServiceTemplateType{
		//TODO(nati): Apply default

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName:                 "",
		ServiceMode:               MakeServiceModeType(),
		Flavor:                    "",
		ServiceScaling:            false,
		VrouterInstanceType:       MakeVRouterInstanceType(),
		OrderedInterfaces:         false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		Version:                   0,
		ServiceType:               MakeServiceType(),
		AvailabilityZoneEnable:    false,
		InstanceData:              "",
	}
}

func InterfaceToServiceTemplateType(iData interface{}) *ServiceTemplateType {
	data := iData.(map[string]interface{})
	return &ServiceTemplateType{
		ImageName: data["image_name"].(string),

		//{"Title":"","Description":"Glance image name for the service virtual machine, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ImageName","GoType":"string"}
		ServiceMode: InterfaceToServiceModeType(data["service_mode"]),

		//{"Title":"","Description":"Service instance mode decides how packets are forwarded across the service","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["transparent","in-network","in-network-nat"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceModeType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceMode","GoType":"ServiceModeType"}
		Flavor: data["flavor"].(string),

		//{"Title":"","Description":"Nova flavor used for service virtual machines, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Flavor","GoType":"string"}
		ServiceScaling: data["service_scaling"].(bool),

		//{"Title":"","Description":"Enable scaling of service virtual machines, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServiceScaling","GoType":"bool"}
		VrouterInstanceType: InterfaceToVRouterInstanceType(data["vrouter_instance_type"]),

		//{"Title":"","Description":"Mechanism used to spawn service instance, when vrouter is spawning instances.Allowed values libvirt-qemu, docker or netns","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["libvirt-qemu","docker"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/VRouterInstanceType","CollectionType":"","Column":"","Item":null,"GoName":"VrouterInstanceType","GoType":"VRouterInstanceType"}
		OrderedInterfaces: data["ordered_interfaces"].(bool),

		//{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OrderedInterfaces","GoType":"bool"}
		ServiceVirtualizationType: InterfaceToServiceVirtualizationType(data["service_virtualization_type"]),

		//{"Title":"","Description":"Service virtualization type decides how individual service instances are instantiated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["virtual-machine","network-namespace","vrouter-instance","physical-device"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceVirtualizationType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceVirtualizationType","GoType":"ServiceVirtualizationType"}

		InterfaceType: InterfaceToServiceTemplateInterfaceTypeSlice(data["interface_type"]),

		//{"Title":"","Description":"List of interfaces which decided number of interfaces and type","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"service_interface_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInterfaceType","GoType":"ServiceInterfaceType"},"shared_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SharedIP","GoType":"bool"},"static_route_enable":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StaticRouteEnable","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceTemplateInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"InterfaceType","GoType":"ServiceTemplateInterfaceType"},"GoName":"InterfaceType","GoType":"[]*ServiceTemplateInterfaceType"}
		ServiceType: InterfaceToServiceType(data["service_type"]),

		//{"Title":"","Description":"Service instance mode decides how routing happens across the service","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["firewall","analyzer","source-nat","loadbalancer"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceType","GoType":"ServiceType"}
		AvailabilityZoneEnable: data["availability_zone_enable"].(bool),

		//{"Title":"","Description":"Enable availability zone for version 1 service instances","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AvailabilityZoneEnable","GoType":"bool"}
		InstanceData: data["instance_data"].(string),

		//{"Title":"","Description":"Opaque string (typically in json format) used to spawn a vrouter-instance.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"InstanceData","GoType":"string"}
		Version: data["version"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Version","GoType":"int"}

	}
}

func InterfaceToServiceTemplateTypeSlice(data interface{}) []*ServiceTemplateType {
	list := data.([]interface{})
	result := MakeServiceTemplateTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplateType(item))
	}
	return result
}

func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
	return []*ServiceTemplateType{}
}
