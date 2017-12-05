package models

// ServiceVirtualizationType

type ServiceVirtualizationType string

func MakeServiceVirtualizationType() ServiceVirtualizationType {
	var data ServiceVirtualizationType
	return data
}

func InterfaceToServiceVirtualizationType(data interface{}) ServiceVirtualizationType {
	return data.(ServiceVirtualizationType)
}

func InterfaceToServiceVirtualizationTypeSlice(data interface{}) []ServiceVirtualizationType {
	list := data.([]interface{})
	result := MakeServiceVirtualizationTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceVirtualizationType(item))
	}
	return result
}

func MakeServiceVirtualizationTypeSlice() []ServiceVirtualizationType {
	return []ServiceVirtualizationType{}
}
