package models

// UuidStringType

type UuidStringType string

func MakeUuidStringType() UuidStringType {
	var data UuidStringType
	return data
}

func InterfaceToUuidStringType(data interface{}) UuidStringType {
	return data.(UuidStringType)
}

func InterfaceToUuidStringTypeSlice(data interface{}) []UuidStringType {
	list := data.([]interface{})
	result := MakeUuidStringTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToUuidStringType(item))
	}
	return result
}

func MakeUuidStringTypeSlice() []UuidStringType {
	return []UuidStringType{}
}
