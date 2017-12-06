package models

// DscpValueType

type DscpValueType int

// MakeDscpValueType makes DscpValueType
func MakeDscpValueType() DscpValueType {
	var data DscpValueType
	return data
}

// InterfaceToDscpValueType makes DscpValueType from interface
func InterfaceToDscpValueType(data interface{}) DscpValueType {
	return data.(DscpValueType)
}

// InterfaceToDscpValueTypeSlice makes a slice of DscpValueType from interface
func InterfaceToDscpValueTypeSlice(data interface{}) []DscpValueType {
	list := data.([]interface{})
	result := MakeDscpValueTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDscpValueType(item))
	}
	return result
}

// MakeDscpValueTypeSlice() makes a slice of DscpValueType
func MakeDscpValueTypeSlice() []DscpValueType {
	return []DscpValueType{}
}
