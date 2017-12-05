package models

// ServiceModeType

type ServiceModeType string

func MakeServiceModeType() ServiceModeType {
	var data ServiceModeType
	return data
}

func InterfaceToServiceModeType(data interface{}) ServiceModeType {
	return data.(ServiceModeType)
}

func InterfaceToServiceModeTypeSlice(data interface{}) []ServiceModeType {
	list := data.([]interface{})
	result := MakeServiceModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceModeType(item))
	}
	return result
}

func MakeServiceModeTypeSlice() []ServiceModeType {
	return []ServiceModeType{}
}
