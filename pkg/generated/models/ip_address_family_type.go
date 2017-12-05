package models

// IpAddressFamilyType

type IpAddressFamilyType string

func MakeIpAddressFamilyType() IpAddressFamilyType {
	var data IpAddressFamilyType
	return data
}

func InterfaceToIpAddressFamilyType(data interface{}) IpAddressFamilyType {
	return data.(IpAddressFamilyType)
}

func InterfaceToIpAddressFamilyTypeSlice(data interface{}) []IpAddressFamilyType {
	list := data.([]interface{})
	result := MakeIpAddressFamilyTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpAddressFamilyType(item))
	}
	return result
}

func MakeIpAddressFamilyTypeSlice() []IpAddressFamilyType {
	return []IpAddressFamilyType{}
}
