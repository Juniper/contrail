package models

// ForwardingModeType

type ForwardingModeType string

// MakeForwardingModeType makes ForwardingModeType
func MakeForwardingModeType() ForwardingModeType {
	var data ForwardingModeType
	return data
}

// InterfaceToForwardingModeType makes ForwardingModeType from interface
func InterfaceToForwardingModeType(data interface{}) ForwardingModeType {
	return data.(ForwardingModeType)
}

// InterfaceToForwardingModeTypeSlice makes a slice of ForwardingModeType from interface
func InterfaceToForwardingModeTypeSlice(data interface{}) []ForwardingModeType {
	list := data.([]interface{})
	result := MakeForwardingModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToForwardingModeType(item))
	}
	return result
}

// MakeForwardingModeTypeSlice() makes a slice of ForwardingModeType
func MakeForwardingModeTypeSlice() []ForwardingModeType {
	return []ForwardingModeType{}
}
