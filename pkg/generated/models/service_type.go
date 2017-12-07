package models

// ServiceType

type ServiceType string

// MakeServiceType makes ServiceType
func MakeServiceType() ServiceType {
	var data ServiceType
	return data
}

// InterfaceToServiceType makes ServiceType from interface
func InterfaceToServiceType(data interface{}) ServiceType {
	return data.(ServiceType)
}

// InterfaceToServiceTypeSlice makes a slice of ServiceType from interface
func InterfaceToServiceTypeSlice(data interface{}) []ServiceType {
	list := data.([]interface{})
	result := MakeServiceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceType(item))
	}
	return result
}

// MakeServiceTypeSlice() makes a slice of ServiceType
func MakeServiceTypeSlice() []ServiceType {
	return []ServiceType{}
}
