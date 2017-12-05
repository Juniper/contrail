package models

// NHModeType

type NHModeType string

func MakeNHModeType() NHModeType {
	var data NHModeType
	return data
}

func InterfaceToNHModeType(data interface{}) NHModeType {
	return data.(NHModeType)
}

func InterfaceToNHModeTypeSlice(data interface{}) []NHModeType {
	list := data.([]interface{})
	result := MakeNHModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToNHModeType(item))
	}
	return result
}

func MakeNHModeTypeSlice() []NHModeType {
	return []NHModeType{}
}
