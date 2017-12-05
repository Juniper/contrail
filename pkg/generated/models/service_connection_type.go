package models

// ServiceConnectionType

type ServiceConnectionType string

func MakeServiceConnectionType() ServiceConnectionType {
	var data ServiceConnectionType
	return data
}

func InterfaceToServiceConnectionType(data interface{}) ServiceConnectionType {
	return data.(ServiceConnectionType)
}

func InterfaceToServiceConnectionTypeSlice(data interface{}) []ServiceConnectionType {
	list := data.([]interface{})
	result := MakeServiceConnectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceConnectionType(item))
	}
	return result
}

func MakeServiceConnectionTypeSlice() []ServiceConnectionType {
	return []ServiceConnectionType{}
}
