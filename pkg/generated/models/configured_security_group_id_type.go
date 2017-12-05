package models

// ConfiguredSecurityGroupIdType

type ConfiguredSecurityGroupIdType int

func MakeConfiguredSecurityGroupIdType() ConfiguredSecurityGroupIdType {
	var data ConfiguredSecurityGroupIdType
	return data
}

func InterfaceToConfiguredSecurityGroupIdType(data interface{}) ConfiguredSecurityGroupIdType {
	return data.(ConfiguredSecurityGroupIdType)
}

func InterfaceToConfiguredSecurityGroupIdTypeSlice(data interface{}) []ConfiguredSecurityGroupIdType {
	list := data.([]interface{})
	result := MakeConfiguredSecurityGroupIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToConfiguredSecurityGroupIdType(item))
	}
	return result
}

func MakeConfiguredSecurityGroupIdTypeSlice() []ConfiguredSecurityGroupIdType {
	return []ConfiguredSecurityGroupIdType{}
}
