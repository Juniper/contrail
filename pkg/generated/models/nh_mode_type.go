package models

// NHModeType

type NHModeType string

// MakeNHModeType makes NHModeType
func MakeNHModeType() NHModeType {
	var data NHModeType
	return data
}

// InterfaceToNHModeType makes NHModeType from interface
func InterfaceToNHModeType(data interface{}) NHModeType {
	return data.(NHModeType)
}

// InterfaceToNHModeTypeSlice makes a slice of NHModeType from interface
func InterfaceToNHModeTypeSlice(data interface{}) []NHModeType {
	list := data.([]interface{})
	result := MakeNHModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToNHModeType(item))
	}
	return result
}

// MakeNHModeTypeSlice() makes a slice of NHModeType
func MakeNHModeTypeSlice() []NHModeType {
	return []NHModeType{}
}
