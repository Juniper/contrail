package models

// VlanPriorityType

type VlanPriorityType int

func MakeVlanPriorityType() VlanPriorityType {
	var data VlanPriorityType
	return data
}

func InterfaceToVlanPriorityType(data interface{}) VlanPriorityType {
	return data.(VlanPriorityType)
}

func InterfaceToVlanPriorityTypeSlice(data interface{}) []VlanPriorityType {
	list := data.([]interface{})
	result := MakeVlanPriorityTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVlanPriorityType(item))
	}
	return result
}

func MakeVlanPriorityTypeSlice() []VlanPriorityType {
	return []VlanPriorityType{}
}
