package models

// VlanPriorityType

type VlanPriorityType int

// MakeVlanPriorityType makes VlanPriorityType
func MakeVlanPriorityType() VlanPriorityType {
	var data VlanPriorityType
	return data
}

// InterfaceToVlanPriorityType makes VlanPriorityType from interface
func InterfaceToVlanPriorityType(data interface{}) VlanPriorityType {
	return data.(VlanPriorityType)
}

// InterfaceToVlanPriorityTypeSlice makes a slice of VlanPriorityType from interface
func InterfaceToVlanPriorityTypeSlice(data interface{}) []VlanPriorityType {
	list := data.([]interface{})
	result := MakeVlanPriorityTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVlanPriorityType(item))
	}
	return result
}

// MakeVlanPriorityTypeSlice() makes a slice of VlanPriorityType
func MakeVlanPriorityTypeSlice() []VlanPriorityType {
	return []VlanPriorityType{}
}
