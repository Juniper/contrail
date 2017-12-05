package models

// SecurityGroupIdType

type SecurityGroupIdType int

func MakeSecurityGroupIdType() SecurityGroupIdType {
	var data SecurityGroupIdType
	return data
}

func InterfaceToSecurityGroupIdType(data interface{}) SecurityGroupIdType {
	return data.(SecurityGroupIdType)
}

func InterfaceToSecurityGroupIdTypeSlice(data interface{}) []SecurityGroupIdType {
	list := data.([]interface{})
	result := MakeSecurityGroupIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityGroupIdType(item))
	}
	return result
}

func MakeSecurityGroupIdTypeSlice() []SecurityGroupIdType {
	return []SecurityGroupIdType{}
}
