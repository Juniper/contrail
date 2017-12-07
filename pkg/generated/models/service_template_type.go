package models

// ServiceTemplateType

import "encoding/json"

// ServiceTemplateType
type ServiceTemplateType struct {
	InterfaceType             []*ServiceTemplateInterfaceType `json:"interface_type"`
	ImageName                 string                          `json:"image_name"`
	ServiceMode               ServiceModeType                 `json:"service_mode"`
	VrouterInstanceType       VRouterInstanceType             `json:"vrouter_instance_type"`
	InstanceData              string                          `json:"instance_data"`
	OrderedInterfaces         bool                            `json:"ordered_interfaces"`
	ServiceVirtualizationType ServiceVirtualizationType       `json:"service_virtualization_type"`
	Version                   int                             `json:"version"`
	ServiceType               ServiceType                     `json:"service_type"`
	Flavor                    string                          `json:"flavor"`
	ServiceScaling            bool                            `json:"service_scaling"`
	AvailabilityZoneEnable    bool                            `json:"availability_zone_enable"`
}

//  parents relation object

// String returns json representation of the object
func (model *ServiceTemplateType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceTemplateType makes ServiceTemplateType
func MakeServiceTemplateType() *ServiceTemplateType {
	return &ServiceTemplateType{
		//TODO(nati): Apply default
		Version:                   0,
		ServiceType:               MakeServiceType(),
		Flavor:                    "",
		ServiceScaling:            false,
		AvailabilityZoneEnable:    false,
		OrderedInterfaces:         false,
		ServiceVirtualizationType: MakeServiceVirtualizationType(),
		ServiceMode:               MakeServiceModeType(),
		VrouterInstanceType:       MakeVRouterInstanceType(),
		InstanceData:              "",

		InterfaceType: MakeServiceTemplateInterfaceTypeSlice(),

		ImageName: "",
	}
}

// InterfaceToServiceTemplateType makes ServiceTemplateType from interface
func InterfaceToServiceTemplateType(iData interface{}) *ServiceTemplateType {
	data := iData.(map[string]interface{})
	return &ServiceTemplateType{
		ServiceScaling: data["service_scaling"].(bool),

		//{"Title":"","Description":"Enable scaling of service virtual machines, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServiceScaling","GoType":"bool","GoPremitive":true}
		AvailabilityZoneEnable: data["availability_zone_enable"].(bool),

		//{"Title":"","Description":"Enable availability zone for version 1 service instances","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AvailabilityZoneEnable","GoType":"bool","GoPremitive":true}
		OrderedInterfaces: data["ordered_interfaces"].(bool),

		//{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OrderedInterfaces","GoType":"bool","GoPremitive":true}
		ServiceVirtualizationType: InterfaceToServiceVirtualizationType(data["service_virtualization_type"]),

		//{"Title":"","Description":"Service virtualization type decides how individual service instances are instantiated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["virtual-machine","network-namespace","vrouter-instance","physical-device"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceVirtualizationType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceVirtualizationType","GoType":"ServiceVirtualizationType","GoPremitive":false}
		Version: data["version"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Version","GoType":"int","GoPremitive":true}
		ServiceType: InterfaceToServiceType(data["service_type"]),

		//{"Title":"","Description":"Service instance mode decides how routing happens across the service","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["firewall","analyzer","source-nat","loadbalancer"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceType","GoType":"ServiceType","GoPremitive":false}
		Flavor: data["flavor"].(string),

		//{"Title":"","Description":"Nova flavor used for service virtual machines, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Flavor","GoType":"string","GoPremitive":true}
		InstanceData: data["instance_data"].(string),

		//{"Title":"","Description":"Opaque string (typically in json format) used to spawn a vrouter-instance.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"InstanceData","GoType":"string","GoPremitive":true}

		InterfaceType: InterfaceToServiceTemplateInterfaceTypeSlice(data["interface_type"]),

		//{"Title":"","Description":"List of interfaces which decided number of interfaces and type","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"service_interface_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInterfaceType","GoType":"ServiceInterfaceType","GoPremitive":false},"shared_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SharedIP","GoType":"bool","GoPremitive":true},"static_route_enable":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StaticRouteEnable","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceTemplateInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"InterfaceType","GoType":"ServiceTemplateInterfaceType","GoPremitive":false},"GoName":"InterfaceType","GoType":"[]*ServiceTemplateInterfaceType","GoPremitive":true}
		ImageName: data["image_name"].(string),

		//{"Title":"","Description":"Glance image name for the service virtual machine, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ImageName","GoType":"string","GoPremitive":true}
		ServiceMode: InterfaceToServiceModeType(data["service_mode"]),

		//{"Title":"","Description":"Service instance mode decides how packets are forwarded across the service","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["transparent","in-network","in-network-nat"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceModeType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceMode","GoType":"ServiceModeType","GoPremitive":false}
		VrouterInstanceType: InterfaceToVRouterInstanceType(data["vrouter_instance_type"]),

		//{"Title":"","Description":"Mechanism used to spawn service instance, when vrouter is spawning instances.Allowed values libvirt-qemu, docker or netns","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["libvirt-qemu","docker"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/VRouterInstanceType","CollectionType":"","Column":"","Item":null,"GoName":"VrouterInstanceType","GoType":"VRouterInstanceType","GoPremitive":false}

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
