package models

// IpAddressType

type IpAddressType string

func MakeIpAddressType() IpAddressType {
	var data IpAddressType
	return data
}

func InterfaceToIpAddressType(data interface{}) IpAddressType {
	return data.(IpAddressType)
}

func InterfaceToIpAddressTypeSlice(data interface{}) []IpAddressType {
	list := data.([]interface{})
	result := MakeIpAddressTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpAddressType(item))
	}
	return result
}

func MakeIpAddressTypeSlice() []IpAddressType {
	return []IpAddressType{}
}
