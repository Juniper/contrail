package models

// SimpleActionType

type SimpleActionType string

func MakeSimpleActionType() SimpleActionType {
	var data SimpleActionType
	return data
}

func InterfaceToSimpleActionType(data interface{}) SimpleActionType {
	return data.(SimpleActionType)
}

func InterfaceToSimpleActionTypeSlice(data interface{}) []SimpleActionType {
	list := data.([]interface{})
	result := MakeSimpleActionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSimpleActionType(item))
	}
	return result
}

func MakeSimpleActionTypeSlice() []SimpleActionType {
	return []SimpleActionType{}
}
