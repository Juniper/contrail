package models

// IsidType

type IsidType int

func MakeIsidType() IsidType {
	var data IsidType
	return data
}

func InterfaceToIsidType(data interface{}) IsidType {
	return data.(IsidType)
}

func InterfaceToIsidTypeSlice(data interface{}) []IsidType {
	list := data.([]interface{})
	result := MakeIsidTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIsidType(item))
	}
	return result
}

func MakeIsidTypeSlice() []IsidType {
	return []IsidType{}
}
