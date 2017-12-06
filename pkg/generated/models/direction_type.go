package models

// DirectionType

type DirectionType string

// MakeDirectionType makes DirectionType
func MakeDirectionType() DirectionType {
	var data DirectionType
	return data
}

// InterfaceToDirectionType makes DirectionType from interface
func InterfaceToDirectionType(data interface{}) DirectionType {
	return data.(DirectionType)
}

// InterfaceToDirectionTypeSlice makes a slice of DirectionType from interface
func InterfaceToDirectionTypeSlice(data interface{}) []DirectionType {
	list := data.([]interface{})
	result := MakeDirectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDirectionType(item))
	}
	return result
}

// MakeDirectionTypeSlice() makes a slice of DirectionType
func MakeDirectionTypeSlice() []DirectionType {
	return []DirectionType{}
}
