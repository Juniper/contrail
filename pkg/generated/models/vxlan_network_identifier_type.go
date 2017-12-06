package models

// VxlanNetworkIdentifierType

type VxlanNetworkIdentifierType int

// MakeVxlanNetworkIdentifierType makes VxlanNetworkIdentifierType
func MakeVxlanNetworkIdentifierType() VxlanNetworkIdentifierType {
	var data VxlanNetworkIdentifierType
	return data
}

// InterfaceToVxlanNetworkIdentifierType makes VxlanNetworkIdentifierType from interface
func InterfaceToVxlanNetworkIdentifierType(data interface{}) VxlanNetworkIdentifierType {
	return data.(VxlanNetworkIdentifierType)
}

// InterfaceToVxlanNetworkIdentifierTypeSlice makes a slice of VxlanNetworkIdentifierType from interface
func InterfaceToVxlanNetworkIdentifierTypeSlice(data interface{}) []VxlanNetworkIdentifierType {
	list := data.([]interface{})
	result := MakeVxlanNetworkIdentifierTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVxlanNetworkIdentifierType(item))
	}
	return result
}

// MakeVxlanNetworkIdentifierTypeSlice() makes a slice of VxlanNetworkIdentifierType
func MakeVxlanNetworkIdentifierTypeSlice() []VxlanNetworkIdentifierType {
	return []VxlanNetworkIdentifierType{}
}
