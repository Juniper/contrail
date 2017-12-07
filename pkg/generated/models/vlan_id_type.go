package models

// VlanIdType

type VlanIdType int

// MakeVlanIdType makes VlanIdType
func MakeVlanIdType() VlanIdType {
	var data VlanIdType
	return data
}

// InterfaceToVlanIdType makes VlanIdType from interface
func InterfaceToVlanIdType(data interface{}) VlanIdType {
	return data.(VlanIdType)
}

// InterfaceToVlanIdTypeSlice makes a slice of VlanIdType from interface
func InterfaceToVlanIdTypeSlice(data interface{}) []VlanIdType {
	list := data.([]interface{})
	result := MakeVlanIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVlanIdType(item))
	}
	return result
}

// MakeVlanIdTypeSlice() makes a slice of VlanIdType
func MakeVlanIdTypeSlice() []VlanIdType {
	return []VlanIdType{}
}
