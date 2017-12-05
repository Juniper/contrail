package models

// DirectionType

type DirectionType string

func MakeDirectionType() DirectionType {
	var data DirectionType
	return data
}

func InterfaceToDirectionType(data interface{}) DirectionType {
	return data.(DirectionType)
}

func InterfaceToDirectionTypeSlice(data interface{}) []DirectionType {
	list := data.([]interface{})
	result := MakeDirectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDirectionType(item))
	}
	return result
}

func MakeDirectionTypeSlice() []DirectionType {
	return []DirectionType{}
}
