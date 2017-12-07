package models

// IpamDnsMethodType

type IpamDnsMethodType string

// MakeIpamDnsMethodType makes IpamDnsMethodType
func MakeIpamDnsMethodType() IpamDnsMethodType {
	var data IpamDnsMethodType
	return data
}

// InterfaceToIpamDnsMethodType makes IpamDnsMethodType from interface
func InterfaceToIpamDnsMethodType(data interface{}) IpamDnsMethodType {
	return data.(IpamDnsMethodType)
}

// InterfaceToIpamDnsMethodTypeSlice makes a slice of IpamDnsMethodType from interface
func InterfaceToIpamDnsMethodTypeSlice(data interface{}) []IpamDnsMethodType {
	list := data.([]interface{})
	result := MakeIpamDnsMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamDnsMethodType(item))
	}
	return result
}

// MakeIpamDnsMethodTypeSlice() makes a slice of IpamDnsMethodType
func MakeIpamDnsMethodTypeSlice() []IpamDnsMethodType {
	return []IpamDnsMethodType{}
}
