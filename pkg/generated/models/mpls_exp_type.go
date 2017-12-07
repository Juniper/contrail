package models

// MplsExpType

type MplsExpType int

// MakeMplsExpType makes MplsExpType
func MakeMplsExpType() MplsExpType {
	var data MplsExpType
	return data
}

// InterfaceToMplsExpType makes MplsExpType from interface
func InterfaceToMplsExpType(data interface{}) MplsExpType {
	return data.(MplsExpType)
}

// InterfaceToMplsExpTypeSlice makes a slice of MplsExpType from interface
func InterfaceToMplsExpTypeSlice(data interface{}) []MplsExpType {
	list := data.([]interface{})
	result := MakeMplsExpTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMplsExpType(item))
	}
	return result
}

// MakeMplsExpTypeSlice() makes a slice of MplsExpType
func MakeMplsExpTypeSlice() []MplsExpType {
	return []MplsExpType{}
}
