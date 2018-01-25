package models

// ServiceTemplateInterfaceType

// ServiceTemplateInterfaceType
//proteus:generate
type ServiceTemplateInterfaceType struct {
	StaticRouteEnable    bool                 `json:"static_route_enable"`
	SharedIP             bool                 `json:"shared_ip"`
	ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type,omitempty"`
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

// MakeServiceTemplateInterfaceTypeSlice() makes a slice of ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceTypeSlice() []*ServiceTemplateInterfaceType {
	return []*ServiceTemplateInterfaceType{}
}
