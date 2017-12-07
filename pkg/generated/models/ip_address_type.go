package models

// IpAddressType

type IpAddressType string

// MakeIpAddressType makes IpAddressType
func MakeIpAddressType() IpAddressType {
	var data IpAddressType
	return data
}

// InterfaceToIpAddressType makes IpAddressType from interface
func InterfaceToIpAddressType(data interface{}) IpAddressType {
	return data.(IpAddressType)
}

// InterfaceToIpAddressTypeSlice makes a slice of IpAddressType from interface
func InterfaceToIpAddressTypeSlice(data interface{}) []IpAddressType {
	list := data.([]interface{})
	result := MakeIpAddressTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpAddressType(item))
	}
	return result
}

// MakeIpAddressTypeSlice() makes a slice of IpAddressType
func MakeIpAddressTypeSlice() []IpAddressType {
	return []IpAddressType{}
}
