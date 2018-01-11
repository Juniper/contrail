package models

// ServiceApplianceInterfaceType

import "encoding/json"

// ServiceApplianceInterfaceType
type ServiceApplianceInterfaceType struct {
	InterfaceType ServiceInterfaceType `json:"interface_type"`
}

// String returns json representation of the object
func (model *ServiceApplianceInterfaceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceApplianceInterfaceType makes ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceType() *ServiceApplianceInterfaceType {
	return &ServiceApplianceInterfaceType{
		//TODO(nati): Apply default
		InterfaceType: MakeServiceInterfaceType(),
	}
}

// MakeServiceApplianceInterfaceTypeSlice() makes a slice of ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceTypeSlice() []*ServiceApplianceInterfaceType {
	return []*ServiceApplianceInterfaceType{}
}
