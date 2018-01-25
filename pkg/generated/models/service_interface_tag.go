package models

// ServiceInterfaceTag

// ServiceInterfaceTag
//proteus:generate
type ServiceInterfaceTag struct {
	InterfaceType ServiceInterfaceType `json:"interface_type,omitempty"`
}

// MakeServiceInterfaceTag makes ServiceInterfaceTag
func MakeServiceInterfaceTag() *ServiceInterfaceTag {
	return &ServiceInterfaceTag{
		//TODO(nati): Apply default
		InterfaceType: MakeServiceInterfaceType(),
	}
}

// MakeServiceInterfaceTagSlice() makes a slice of ServiceInterfaceTag
func MakeServiceInterfaceTagSlice() []*ServiceInterfaceTag {
	return []*ServiceInterfaceTag{}
}
