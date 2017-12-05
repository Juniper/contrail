package models

// ServiceType

type ServiceType string

func MakeServiceType() ServiceType {
	var data ServiceType
	return data
}

func InterfaceToServiceType(data interface{}) ServiceType {
	return data.(ServiceType)
}

func InterfaceToServiceTypeSlice(data interface{}) []ServiceType {
	list := data.([]interface{})
	result := MakeServiceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceType(item))
	}
	return result
}

func MakeServiceTypeSlice() []ServiceType {
	return []ServiceType{}
}
