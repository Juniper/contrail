package models

// VlanIdType

type VlanIdType int

func MakeVlanIdType() VlanIdType {
	var data VlanIdType
	return data
}

func InterfaceToVlanIdType(data interface{}) VlanIdType {
	return data.(VlanIdType)
}

func InterfaceToVlanIdTypeSlice(data interface{}) []VlanIdType {
	list := data.([]interface{})
	result := MakeVlanIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVlanIdType(item))
	}
	return result
}

func MakeVlanIdTypeSlice() []VlanIdType {
	return []VlanIdType{}
}
