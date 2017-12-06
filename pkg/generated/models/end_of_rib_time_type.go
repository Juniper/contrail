package models

// EndOfRibTimeType

type EndOfRibTimeType int

// MakeEndOfRibTimeType makes EndOfRibTimeType
func MakeEndOfRibTimeType() EndOfRibTimeType {
	var data EndOfRibTimeType
	return data
}

// InterfaceToEndOfRibTimeType makes EndOfRibTimeType from interface
func InterfaceToEndOfRibTimeType(data interface{}) EndOfRibTimeType {
	return data.(EndOfRibTimeType)
}

// InterfaceToEndOfRibTimeTypeSlice makes a slice of EndOfRibTimeType from interface
func InterfaceToEndOfRibTimeTypeSlice(data interface{}) []EndOfRibTimeType {
	list := data.([]interface{})
	result := MakeEndOfRibTimeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToEndOfRibTimeType(item))
	}
	return result
}

// MakeEndOfRibTimeTypeSlice() makes a slice of EndOfRibTimeType
func MakeEndOfRibTimeTypeSlice() []EndOfRibTimeType {
	return []EndOfRibTimeType{}
}
