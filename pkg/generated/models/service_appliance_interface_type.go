package models

// ServiceApplianceInterfaceType

// ServiceApplianceInterfaceType
//proteus:generate
type ServiceApplianceInterfaceType struct {
	InterfaceType ServiceInterfaceType `json:"interface_type,omitempty"`
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
