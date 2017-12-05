package models

// VpnType

type VpnType string

func MakeVpnType() VpnType {
	var data VpnType
	return data
}

func InterfaceToVpnType(data interface{}) VpnType {
	return data.(VpnType)
}

func InterfaceToVpnTypeSlice(data interface{}) []VpnType {
	list := data.([]interface{})
	result := MakeVpnTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVpnType(item))
	}
	return result
}

func MakeVpnTypeSlice() []VpnType {
	return []VpnType{}
}
