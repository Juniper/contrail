package models

// IpamDnsMethodType

type IpamDnsMethodType string

func MakeIpamDnsMethodType() IpamDnsMethodType {
	var data IpamDnsMethodType
	return data
}

func InterfaceToIpamDnsMethodType(data interface{}) IpamDnsMethodType {
	return data.(IpamDnsMethodType)
}

func InterfaceToIpamDnsMethodTypeSlice(data interface{}) []IpamDnsMethodType {
	list := data.([]interface{})
	result := MakeIpamDnsMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamDnsMethodType(item))
	}
	return result
}

func MakeIpamDnsMethodTypeSlice() []IpamDnsMethodType {
	return []IpamDnsMethodType{}
}
