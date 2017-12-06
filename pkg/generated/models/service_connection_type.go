package models

// ServiceConnectionType

type ServiceConnectionType string

// MakeServiceConnectionType makes ServiceConnectionType
func MakeServiceConnectionType() ServiceConnectionType {
	var data ServiceConnectionType
	return data
}

// InterfaceToServiceConnectionType makes ServiceConnectionType from interface
func InterfaceToServiceConnectionType(data interface{}) ServiceConnectionType {
	return data.(ServiceConnectionType)
}

// InterfaceToServiceConnectionTypeSlice makes a slice of ServiceConnectionType from interface
func InterfaceToServiceConnectionTypeSlice(data interface{}) []ServiceConnectionType {
	list := data.([]interface{})
	result := MakeServiceConnectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceConnectionType(item))
	}
	return result
}

// MakeServiceConnectionTypeSlice() makes a slice of ServiceConnectionType
func MakeServiceConnectionTypeSlice() []ServiceConnectionType {
	return []ServiceConnectionType{}
}
