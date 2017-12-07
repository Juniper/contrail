package models

// MACLimitExceedActionType

type MACLimitExceedActionType string

// MakeMACLimitExceedActionType makes MACLimitExceedActionType
func MakeMACLimitExceedActionType() MACLimitExceedActionType {
	var data MACLimitExceedActionType
	return data
}

// InterfaceToMACLimitExceedActionType makes MACLimitExceedActionType from interface
func InterfaceToMACLimitExceedActionType(data interface{}) MACLimitExceedActionType {
	return data.(MACLimitExceedActionType)
}

// InterfaceToMACLimitExceedActionTypeSlice makes a slice of MACLimitExceedActionType from interface
func InterfaceToMACLimitExceedActionTypeSlice(data interface{}) []MACLimitExceedActionType {
	list := data.([]interface{})
	result := MakeMACLimitExceedActionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACLimitExceedActionType(item))
	}
	return result
}

// MakeMACLimitExceedActionTypeSlice() makes a slice of MACLimitExceedActionType
func MakeMACLimitExceedActionTypeSlice() []MACLimitExceedActionType {
	return []MACLimitExceedActionType{}
}
