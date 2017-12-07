package models

// SimpleActionType

type SimpleActionType string

// MakeSimpleActionType makes SimpleActionType
func MakeSimpleActionType() SimpleActionType {
	var data SimpleActionType
	return data
}

// InterfaceToSimpleActionType makes SimpleActionType from interface
func InterfaceToSimpleActionType(data interface{}) SimpleActionType {
	return data.(SimpleActionType)
}

// InterfaceToSimpleActionTypeSlice makes a slice of SimpleActionType from interface
func InterfaceToSimpleActionTypeSlice(data interface{}) []SimpleActionType {
	list := data.([]interface{})
	result := MakeSimpleActionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSimpleActionType(item))
	}
	return result
}

// MakeSimpleActionTypeSlice() makes a slice of SimpleActionType
func MakeSimpleActionTypeSlice() []SimpleActionType {
	return []SimpleActionType{}
}
