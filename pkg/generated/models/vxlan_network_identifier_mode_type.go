package models

// VxlanNetworkIdentifierModeType

type VxlanNetworkIdentifierModeType string

// MakeVxlanNetworkIdentifierModeType makes VxlanNetworkIdentifierModeType
func MakeVxlanNetworkIdentifierModeType() VxlanNetworkIdentifierModeType {
	var data VxlanNetworkIdentifierModeType
	return data
}

// InterfaceToVxlanNetworkIdentifierModeType makes VxlanNetworkIdentifierModeType from interface
func InterfaceToVxlanNetworkIdentifierModeType(data interface{}) VxlanNetworkIdentifierModeType {
	return data.(VxlanNetworkIdentifierModeType)
}

// InterfaceToVxlanNetworkIdentifierModeTypeSlice makes a slice of VxlanNetworkIdentifierModeType from interface
func InterfaceToVxlanNetworkIdentifierModeTypeSlice(data interface{}) []VxlanNetworkIdentifierModeType {
	list := data.([]interface{})
	result := MakeVxlanNetworkIdentifierModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVxlanNetworkIdentifierModeType(item))
	}
	return result
}

// MakeVxlanNetworkIdentifierModeTypeSlice() makes a slice of VxlanNetworkIdentifierModeType
func MakeVxlanNetworkIdentifierModeTypeSlice() []VxlanNetworkIdentifierModeType {
	return []VxlanNetworkIdentifierModeType{}
}
