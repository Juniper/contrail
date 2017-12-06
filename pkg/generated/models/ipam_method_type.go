package models

// IpamMethodType

type IpamMethodType string

// MakeIpamMethodType makes IpamMethodType
func MakeIpamMethodType() IpamMethodType {
	var data IpamMethodType
	return data
}

// InterfaceToIpamMethodType makes IpamMethodType from interface
func InterfaceToIpamMethodType(data interface{}) IpamMethodType {
	return data.(IpamMethodType)
}

// InterfaceToIpamMethodTypeSlice makes a slice of IpamMethodType from interface
func InterfaceToIpamMethodTypeSlice(data interface{}) []IpamMethodType {
	list := data.([]interface{})
	result := MakeIpamMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamMethodType(item))
	}
	return result
}

// MakeIpamMethodTypeSlice() makes a slice of IpamMethodType
func MakeIpamMethodTypeSlice() []IpamMethodType {
	return []IpamMethodType{}
}
