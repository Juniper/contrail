package models

// SecurityGroupIdType

type SecurityGroupIdType int

// MakeSecurityGroupIdType makes SecurityGroupIdType
func MakeSecurityGroupIdType() SecurityGroupIdType {
	var data SecurityGroupIdType
	return data
}

// InterfaceToSecurityGroupIdType makes SecurityGroupIdType from interface
func InterfaceToSecurityGroupIdType(data interface{}) SecurityGroupIdType {
	return data.(SecurityGroupIdType)
}

// InterfaceToSecurityGroupIdTypeSlice makes a slice of SecurityGroupIdType from interface
func InterfaceToSecurityGroupIdTypeSlice(data interface{}) []SecurityGroupIdType {
	list := data.([]interface{})
	result := MakeSecurityGroupIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityGroupIdType(item))
	}
	return result
}

// MakeSecurityGroupIdTypeSlice() makes a slice of SecurityGroupIdType
func MakeSecurityGroupIdTypeSlice() []SecurityGroupIdType {
	return []SecurityGroupIdType{}
}
