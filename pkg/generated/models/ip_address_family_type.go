package models

// IpAddressFamilyType

type IpAddressFamilyType string

// MakeIpAddressFamilyType makes IpAddressFamilyType
func MakeIpAddressFamilyType() IpAddressFamilyType {
	var data IpAddressFamilyType
	return data
}

// InterfaceToIpAddressFamilyType makes IpAddressFamilyType from interface
func InterfaceToIpAddressFamilyType(data interface{}) IpAddressFamilyType {
	return data.(IpAddressFamilyType)
}

// InterfaceToIpAddressFamilyTypeSlice makes a slice of IpAddressFamilyType from interface
func InterfaceToIpAddressFamilyTypeSlice(data interface{}) []IpAddressFamilyType {
	list := data.([]interface{})
	result := MakeIpAddressFamilyTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpAddressFamilyType(item))
	}
	return result
}

// MakeIpAddressFamilyTypeSlice() makes a slice of IpAddressFamilyType
func MakeIpAddressFamilyTypeSlice() []IpAddressFamilyType {
	return []IpAddressFamilyType{}
}
