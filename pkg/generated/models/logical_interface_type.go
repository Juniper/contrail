package models

// LogicalInterfaceType

type LogicalInterfaceType string

// MakeLogicalInterfaceType makes LogicalInterfaceType
func MakeLogicalInterfaceType() LogicalInterfaceType {
	var data LogicalInterfaceType
	return data
}

// InterfaceToLogicalInterfaceType makes LogicalInterfaceType from interface
func InterfaceToLogicalInterfaceType(data interface{}) LogicalInterfaceType {
	return data.(LogicalInterfaceType)
}

// InterfaceToLogicalInterfaceTypeSlice makes a slice of LogicalInterfaceType from interface
func InterfaceToLogicalInterfaceTypeSlice(data interface{}) []LogicalInterfaceType {
	list := data.([]interface{})
	result := MakeLogicalInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLogicalInterfaceType(item))
	}
	return result
}

// MakeLogicalInterfaceTypeSlice() makes a slice of LogicalInterfaceType
func MakeLogicalInterfaceTypeSlice() []LogicalInterfaceType {
	return []LogicalInterfaceType{}
}
