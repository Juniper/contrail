package models

// ServiceModeType

type ServiceModeType string

// MakeServiceModeType makes ServiceModeType
func MakeServiceModeType() ServiceModeType {
	var data ServiceModeType
	return data
}

// InterfaceToServiceModeType makes ServiceModeType from interface
func InterfaceToServiceModeType(data interface{}) ServiceModeType {
	return data.(ServiceModeType)
}

// InterfaceToServiceModeTypeSlice makes a slice of ServiceModeType from interface
func InterfaceToServiceModeTypeSlice(data interface{}) []ServiceModeType {
	list := data.([]interface{})
	result := MakeServiceModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceModeType(item))
	}
	return result
}

// MakeServiceModeTypeSlice() makes a slice of ServiceModeType
func MakeServiceModeTypeSlice() []ServiceModeType {
	return []ServiceModeType{}
}
