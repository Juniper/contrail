package models

// RpfModeType

type RpfModeType string

// MakeRpfModeType makes RpfModeType
func MakeRpfModeType() RpfModeType {
	var data RpfModeType
	return data
}

// InterfaceToRpfModeType makes RpfModeType from interface
func InterfaceToRpfModeType(data interface{}) RpfModeType {
	return data.(RpfModeType)
}

// InterfaceToRpfModeTypeSlice makes a slice of RpfModeType from interface
func InterfaceToRpfModeTypeSlice(data interface{}) []RpfModeType {
	list := data.([]interface{})
	result := MakeRpfModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRpfModeType(item))
	}
	return result
}

// MakeRpfModeTypeSlice() makes a slice of RpfModeType
func MakeRpfModeTypeSlice() []RpfModeType {
	return []RpfModeType{}
}
