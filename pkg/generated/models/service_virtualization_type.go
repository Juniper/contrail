package models

// ServiceVirtualizationType

type ServiceVirtualizationType string

// MakeServiceVirtualizationType makes ServiceVirtualizationType
func MakeServiceVirtualizationType() ServiceVirtualizationType {
	var data ServiceVirtualizationType
	return data
}

// InterfaceToServiceVirtualizationType makes ServiceVirtualizationType from interface
func InterfaceToServiceVirtualizationType(data interface{}) ServiceVirtualizationType {
	return data.(ServiceVirtualizationType)
}

// InterfaceToServiceVirtualizationTypeSlice makes a slice of ServiceVirtualizationType from interface
func InterfaceToServiceVirtualizationTypeSlice(data interface{}) []ServiceVirtualizationType {
	list := data.([]interface{})
	result := MakeServiceVirtualizationTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceVirtualizationType(item))
	}
	return result
}

// MakeServiceVirtualizationTypeSlice() makes a slice of ServiceVirtualizationType
func MakeServiceVirtualizationTypeSlice() []ServiceVirtualizationType {
	return []ServiceVirtualizationType{}
}
