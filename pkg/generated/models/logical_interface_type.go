package models

// LogicalInterfaceType

type LogicalInterfaceType string

func MakeLogicalInterfaceType() LogicalInterfaceType {
	var data LogicalInterfaceType
	return data
}

func InterfaceToLogicalInterfaceType(data interface{}) LogicalInterfaceType {
	return data.(LogicalInterfaceType)
}

func InterfaceToLogicalInterfaceTypeSlice(data interface{}) []LogicalInterfaceType {
	list := data.([]interface{})
	result := MakeLogicalInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLogicalInterfaceType(item))
	}
	return result
}

func MakeLogicalInterfaceTypeSlice() []LogicalInterfaceType {
	return []LogicalInterfaceType{}
}
