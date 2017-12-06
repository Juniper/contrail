package models

// UuidStringType

type UuidStringType string

// MakeUuidStringType makes UuidStringType
func MakeUuidStringType() UuidStringType {
	var data UuidStringType
	return data
}

// InterfaceToUuidStringType makes UuidStringType from interface
func InterfaceToUuidStringType(data interface{}) UuidStringType {
	return data.(UuidStringType)
}

// InterfaceToUuidStringTypeSlice makes a slice of UuidStringType from interface
func InterfaceToUuidStringTypeSlice(data interface{}) []UuidStringType {
	list := data.([]interface{})
	result := MakeUuidStringTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToUuidStringType(item))
	}
	return result
}

// MakeUuidStringTypeSlice() makes a slice of UuidStringType
func MakeUuidStringTypeSlice() []UuidStringType {
	return []UuidStringType{}
}
