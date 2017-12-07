package models

// Dot1QTagType

type Dot1QTagType int

// MakeDot1QTagType makes Dot1QTagType
func MakeDot1QTagType() Dot1QTagType {
	var data Dot1QTagType
	return data
}

// InterfaceToDot1QTagType makes Dot1QTagType from interface
func InterfaceToDot1QTagType(data interface{}) Dot1QTagType {
	return data.(Dot1QTagType)
}

// InterfaceToDot1QTagTypeSlice makes a slice of Dot1QTagType from interface
func InterfaceToDot1QTagTypeSlice(data interface{}) []Dot1QTagType {
	list := data.([]interface{})
	result := MakeDot1QTagTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDot1QTagType(item))
	}
	return result
}

// MakeDot1QTagTypeSlice() makes a slice of Dot1QTagType
func MakeDot1QTagTypeSlice() []Dot1QTagType {
	return []Dot1QTagType{}
}
