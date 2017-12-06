package models

// SubnetMethodType

type SubnetMethodType string

// MakeSubnetMethodType makes SubnetMethodType
func MakeSubnetMethodType() SubnetMethodType {
	var data SubnetMethodType
	return data
}

// InterfaceToSubnetMethodType makes SubnetMethodType from interface
func InterfaceToSubnetMethodType(data interface{}) SubnetMethodType {
	return data.(SubnetMethodType)
}

// InterfaceToSubnetMethodTypeSlice makes a slice of SubnetMethodType from interface
func InterfaceToSubnetMethodTypeSlice(data interface{}) []SubnetMethodType {
	list := data.([]interface{})
	result := MakeSubnetMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSubnetMethodType(item))
	}
	return result
}

// MakeSubnetMethodTypeSlice() makes a slice of SubnetMethodType
func MakeSubnetMethodTypeSlice() []SubnetMethodType {
	return []SubnetMethodType{}
}
