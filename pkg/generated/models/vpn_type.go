package models

// VpnType

type VpnType string

// MakeVpnType makes VpnType
func MakeVpnType() VpnType {
	var data VpnType
	return data
}

// InterfaceToVpnType makes VpnType from interface
func InterfaceToVpnType(data interface{}) VpnType {
	return data.(VpnType)
}

// InterfaceToVpnTypeSlice makes a slice of VpnType from interface
func InterfaceToVpnTypeSlice(data interface{}) []VpnType {
	list := data.([]interface{})
	result := MakeVpnTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVpnType(item))
	}
	return result
}

// MakeVpnTypeSlice() makes a slice of VpnType
func MakeVpnTypeSlice() []VpnType {
	return []VpnType{}
}
