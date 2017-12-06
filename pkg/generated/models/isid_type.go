package models

// IsidType

type IsidType int

// MakeIsidType makes IsidType
func MakeIsidType() IsidType {
	var data IsidType
	return data
}

// InterfaceToIsidType makes IsidType from interface
func InterfaceToIsidType(data interface{}) IsidType {
	return data.(IsidType)
}

// InterfaceToIsidTypeSlice makes a slice of IsidType from interface
func InterfaceToIsidTypeSlice(data interface{}) []IsidType {
	list := data.([]interface{})
	result := MakeIsidTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIsidType(item))
	}
	return result
}

// MakeIsidTypeSlice() makes a slice of IsidType
func MakeIsidTypeSlice() []IsidType {
	return []IsidType{}
}
