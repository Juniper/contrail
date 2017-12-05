package models

// SubnetMethodType

type SubnetMethodType string

func MakeSubnetMethodType() SubnetMethodType {
	var data SubnetMethodType
	return data
}

func InterfaceToSubnetMethodType(data interface{}) SubnetMethodType {
	return data.(SubnetMethodType)
}

func InterfaceToSubnetMethodTypeSlice(data interface{}) []SubnetMethodType {
	list := data.([]interface{})
	result := MakeSubnetMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSubnetMethodType(item))
	}
	return result
}

func MakeSubnetMethodTypeSlice() []SubnetMethodType {
	return []SubnetMethodType{}
}
