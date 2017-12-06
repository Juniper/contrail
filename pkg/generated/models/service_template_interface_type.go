package models

// ServiceTemplateInterfaceType

import "encoding/json"

// ServiceTemplateInterfaceType
type ServiceTemplateInterfaceType struct {
	StaticRouteEnable    bool                 `json:"static_route_enable"`
	SharedIP             bool                 `json:"shared_ip"`
	ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type"`
}

//  parents relation object

// String returns json representation of the object
func (model *ServiceTemplateInterfaceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceType() *ServiceTemplateInterfaceType {
	return &ServiceTemplateInterfaceType{
		//TODO(nati): Apply default
		SharedIP:             false,
		ServiceInterfaceType: MakeServiceInterfaceType(),
		StaticRouteEnable:    false,
	}
}

// InterfaceToServiceTemplateInterfaceType makes ServiceTemplateInterfaceType from interface
func InterfaceToServiceTemplateInterfaceType(iData interface{}) *ServiceTemplateInterfaceType {
	data := iData.(map[string]interface{})
	return &ServiceTemplateInterfaceType{
		StaticRouteEnable: data["static_route_enable"].(bool),

		//{"Title":"","Description":"Static routes configured required on this interface of service instance (Only V1)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StaticRouteEnable","GoType":"bool","GoPremitive":true}
		SharedIP: data["shared_ip"].(bool),

		//{"Title":"","Description":"Shared ip is required on this interface when service instance is scaled out (Only V1)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SharedIP","GoType":"bool","GoPremitive":true}
		ServiceInterfaceType: InterfaceToServiceInterfaceType(data["service_interface_type"]),

		//{"Title":"","Description":"Type of service interface supported by this template left, right or other.","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInterfaceType","GoType":"ServiceInterfaceType","GoPremitive":false}

	}
}

// InterfaceToServiceTemplateInterfaceTypeSlice makes a slice of ServiceTemplateInterfaceType from interface
func InterfaceToServiceTemplateInterfaceTypeSlice(data interface{}) []*ServiceTemplateInterfaceType {
	list := data.([]interface{})
	result := MakeServiceTemplateInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplateInterfaceType(item))
	}
	return result
}

// MakeServiceTemplateInterfaceTypeSlice() makes a slice of ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceTypeSlice() []*ServiceTemplateInterfaceType {
	return []*ServiceTemplateInterfaceType{}
}
