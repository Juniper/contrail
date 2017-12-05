package models

// VxlanNetworkIdentifierModeType

type VxlanNetworkIdentifierModeType string

func MakeVxlanNetworkIdentifierModeType() VxlanNetworkIdentifierModeType {
	var data VxlanNetworkIdentifierModeType
	return data
}

func InterfaceToVxlanNetworkIdentifierModeType(data interface{}) VxlanNetworkIdentifierModeType {
	return data.(VxlanNetworkIdentifierModeType)
}

func InterfaceToVxlanNetworkIdentifierModeTypeSlice(data interface{}) []VxlanNetworkIdentifierModeType {
	list := data.([]interface{})
	result := MakeVxlanNetworkIdentifierModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVxlanNetworkIdentifierModeType(item))
	}
	return result
}

func MakeVxlanNetworkIdentifierModeTypeSlice() []VxlanNetworkIdentifierModeType {
	return []VxlanNetworkIdentifierModeType{}
}
