package models

// EndOfRibTimeType

type EndOfRibTimeType int

func MakeEndOfRibTimeType() EndOfRibTimeType {
	var data EndOfRibTimeType
	return data
}

func InterfaceToEndOfRibTimeType(data interface{}) EndOfRibTimeType {
	return data.(EndOfRibTimeType)
}

func InterfaceToEndOfRibTimeTypeSlice(data interface{}) []EndOfRibTimeType {
	list := data.([]interface{})
	result := MakeEndOfRibTimeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToEndOfRibTimeType(item))
	}
	return result
}

func MakeEndOfRibTimeTypeSlice() []EndOfRibTimeType {
	return []EndOfRibTimeType{}
}
