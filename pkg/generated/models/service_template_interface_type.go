package models

// ServiceTemplateInterfaceType

import "encoding/json"

// ServiceTemplateInterfaceType
type ServiceTemplateInterfaceType struct {
	StaticRouteEnable    bool                 `json:"static_route_enable"`
	SharedIP             bool                 `json:"shared_ip"`
	ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type"`
}

// String returns json representation of the object
func (model *ServiceTemplateInterfaceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceType() *ServiceTemplateInterfaceType {
	return &ServiceTemplateInterfaceType{
		//TODO(nati): Apply default
		StaticRouteEnable:    false,
		SharedIP:             false,
		ServiceInterfaceType: MakeServiceInterfaceType(),
	}
}

// InterfaceToServiceTemplateInterfaceType makes ServiceTemplateInterfaceType from interface
func InterfaceToServiceTemplateInterfaceType(iData interface{}) *ServiceTemplateInterfaceType {
	data := iData.(map[string]interface{})
	return &ServiceTemplateInterfaceType{
		StaticRouteEnable: data["static_route_enable"].(bool),

		//{"description":"Static routes configured required on this interface of service instance (Only V1)","type":"boolean"}
		SharedIP: data["shared_ip"].(bool),

		//{"description":"Shared ip is required on this interface when service instance is scaled out (Only V1)","type":"boolean"}
		ServiceInterfaceType: InterfaceToServiceInterfaceType(data["service_interface_type"]),

		//{"description":"Type of service interface supported by this template left, right or other.","type":"string"}

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
