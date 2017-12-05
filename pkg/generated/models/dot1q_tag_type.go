package models

// Dot1QTagType

type Dot1QTagType int

func MakeDot1QTagType() Dot1QTagType {
	var data Dot1QTagType
	return data
}

func InterfaceToDot1QTagType(data interface{}) Dot1QTagType {
	return data.(Dot1QTagType)
}

func InterfaceToDot1QTagTypeSlice(data interface{}) []Dot1QTagType {
	list := data.([]interface{})
	result := MakeDot1QTagTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDot1QTagType(item))
	}
	return result
}

func MakeDot1QTagTypeSlice() []Dot1QTagType {
	return []Dot1QTagType{}
}
