package models

// ConfiguredSecurityGroupIdType

type ConfiguredSecurityGroupIdType int

// MakeConfiguredSecurityGroupIdType makes ConfiguredSecurityGroupIdType
func MakeConfiguredSecurityGroupIdType() ConfiguredSecurityGroupIdType {
	var data ConfiguredSecurityGroupIdType
	return data
}

// InterfaceToConfiguredSecurityGroupIdType makes ConfiguredSecurityGroupIdType from interface
func InterfaceToConfiguredSecurityGroupIdType(data interface{}) ConfiguredSecurityGroupIdType {
	return data.(ConfiguredSecurityGroupIdType)
}

// InterfaceToConfiguredSecurityGroupIdTypeSlice makes a slice of ConfiguredSecurityGroupIdType from interface
func InterfaceToConfiguredSecurityGroupIdTypeSlice(data interface{}) []ConfiguredSecurityGroupIdType {
	list := data.([]interface{})
	result := MakeConfiguredSecurityGroupIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToConfiguredSecurityGroupIdType(item))
	}
	return result
}

// MakeConfiguredSecurityGroupIdTypeSlice() makes a slice of ConfiguredSecurityGroupIdType
func MakeConfiguredSecurityGroupIdTypeSlice() []ConfiguredSecurityGroupIdType {
	return []ConfiguredSecurityGroupIdType{}
}
