package models

// IpamMethodType

type IpamMethodType string

func MakeIpamMethodType() IpamMethodType {
	var data IpamMethodType
	return data
}

func InterfaceToIpamMethodType(data interface{}) IpamMethodType {
	return data.(IpamMethodType)
}

func InterfaceToIpamMethodTypeSlice(data interface{}) []IpamMethodType {
	list := data.([]interface{})
	result := MakeIpamMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamMethodType(item))
	}
	return result
}

func MakeIpamMethodTypeSlice() []IpamMethodType {
	return []IpamMethodType{}
}
