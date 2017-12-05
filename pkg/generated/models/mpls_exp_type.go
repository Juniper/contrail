package models

// MplsExpType

type MplsExpType int

func MakeMplsExpType() MplsExpType {
	var data MplsExpType
	return data
}

func InterfaceToMplsExpType(data interface{}) MplsExpType {
	return data.(MplsExpType)
}

func InterfaceToMplsExpTypeSlice(data interface{}) []MplsExpType {
	list := data.([]interface{})
	result := MakeMplsExpTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMplsExpType(item))
	}
	return result
}

func MakeMplsExpTypeSlice() []MplsExpType {
	return []MplsExpType{}
}
