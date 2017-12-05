package models

// MACLimitExceedActionType

type MACLimitExceedActionType string

func MakeMACLimitExceedActionType() MACLimitExceedActionType {
	var data MACLimitExceedActionType
	return data
}

func InterfaceToMACLimitExceedActionType(data interface{}) MACLimitExceedActionType {
	return data.(MACLimitExceedActionType)
}

func InterfaceToMACLimitExceedActionTypeSlice(data interface{}) []MACLimitExceedActionType {
	list := data.([]interface{})
	result := MakeMACLimitExceedActionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACLimitExceedActionType(item))
	}
	return result
}

func MakeMACLimitExceedActionTypeSlice() []MACLimitExceedActionType {
	return []MACLimitExceedActionType{}
}
