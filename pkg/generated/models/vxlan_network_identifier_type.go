package models

// VxlanNetworkIdentifierType

type VxlanNetworkIdentifierType int

func MakeVxlanNetworkIdentifierType() VxlanNetworkIdentifierType {
	var data VxlanNetworkIdentifierType
	return data
}

func InterfaceToVxlanNetworkIdentifierType(data interface{}) VxlanNetworkIdentifierType {
	return data.(VxlanNetworkIdentifierType)
}

func InterfaceToVxlanNetworkIdentifierTypeSlice(data interface{}) []VxlanNetworkIdentifierType {
	list := data.([]interface{})
	result := MakeVxlanNetworkIdentifierTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVxlanNetworkIdentifierType(item))
	}
	return result
}

func MakeVxlanNetworkIdentifierTypeSlice() []VxlanNetworkIdentifierType {
	return []VxlanNetworkIdentifierType{}
}
