package models

// DscpValueType

type DscpValueType int

func MakeDscpValueType() DscpValueType {
	var data DscpValueType
	return data
}

func InterfaceToDscpValueType(data interface{}) DscpValueType {
	return data.(DscpValueType)
}

func InterfaceToDscpValueTypeSlice(data interface{}) []DscpValueType {
	list := data.([]interface{})
	result := MakeDscpValueTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDscpValueType(item))
	}
	return result
}

func MakeDscpValueTypeSlice() []DscpValueType {
	return []DscpValueType{}
}
