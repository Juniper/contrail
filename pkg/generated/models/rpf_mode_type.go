package models

// RpfModeType

type RpfModeType string

func MakeRpfModeType() RpfModeType {
	var data RpfModeType
	return data
}

func InterfaceToRpfModeType(data interface{}) RpfModeType {
	return data.(RpfModeType)
}

func InterfaceToRpfModeTypeSlice(data interface{}) []RpfModeType {
	list := data.([]interface{})
	result := MakeRpfModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRpfModeType(item))
	}
	return result
}

func MakeRpfModeTypeSlice() []RpfModeType {
	return []RpfModeType{}
}
