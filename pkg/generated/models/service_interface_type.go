package models

// ServiceInterfaceType

type ServiceInterfaceType string

// MakeServiceInterfaceType makes ServiceInterfaceType
func MakeServiceInterfaceType() ServiceInterfaceType {
	var data ServiceInterfaceType
	return data
}

// InterfaceToServiceInterfaceType makes ServiceInterfaceType from interface
func InterfaceToServiceInterfaceType(data interface{}) ServiceInterfaceType {
	return data.(ServiceInterfaceType)
}

// InterfaceToServiceInterfaceTypeSlice makes a slice of ServiceInterfaceType from interface
func InterfaceToServiceInterfaceTypeSlice(data interface{}) []ServiceInterfaceType {
	list := data.([]interface{})
	result := MakeServiceInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInterfaceType(item))
	}
	return result
}

// MakeServiceInterfaceTypeSlice() makes a slice of ServiceInterfaceType
func MakeServiceInterfaceTypeSlice() []ServiceInterfaceType {
	return []ServiceInterfaceType{}
}
