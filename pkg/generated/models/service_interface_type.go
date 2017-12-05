package models

// ServiceInterfaceType

type ServiceInterfaceType string

func MakeServiceInterfaceType() ServiceInterfaceType {
	var data ServiceInterfaceType
	return data
}

func InterfaceToServiceInterfaceType(data interface{}) ServiceInterfaceType {
	return data.(ServiceInterfaceType)
}

func InterfaceToServiceInterfaceTypeSlice(data interface{}) []ServiceInterfaceType {
	list := data.([]interface{})
	result := MakeServiceInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInterfaceType(item))
	}
	return result
}

func MakeServiceInterfaceTypeSlice() []ServiceInterfaceType {
	return []ServiceInterfaceType{}
}
