package models

// ForwardingModeType

type ForwardingModeType string

func MakeForwardingModeType() ForwardingModeType {
	var data ForwardingModeType
	return data
}

func InterfaceToForwardingModeType(data interface{}) ForwardingModeType {
	return data.(ForwardingModeType)
}

func InterfaceToForwardingModeTypeSlice(data interface{}) []ForwardingModeType {
	list := data.([]interface{})
	result := MakeForwardingModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToForwardingModeType(item))
	}
	return result
}

func MakeForwardingModeTypeSlice() []ForwardingModeType {
	return []ForwardingModeType{}
}
