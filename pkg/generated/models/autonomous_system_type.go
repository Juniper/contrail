package models

// AutonomousSystemType

type AutonomousSystemType int

// MakeAutonomousSystemType makes AutonomousSystemType
func MakeAutonomousSystemType() AutonomousSystemType {
	var data AutonomousSystemType
	return data
}

// InterfaceToAutonomousSystemType makes AutonomousSystemType from interface
func InterfaceToAutonomousSystemType(data interface{}) AutonomousSystemType {
	return data.(AutonomousSystemType)
}

// InterfaceToAutonomousSystemTypeSlice makes a slice of AutonomousSystemType from interface
func InterfaceToAutonomousSystemTypeSlice(data interface{}) []AutonomousSystemType {
	list := data.([]interface{})
	result := MakeAutonomousSystemTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAutonomousSystemType(item))
	}
	return result
}

// MakeAutonomousSystemTypeSlice() makes a slice of AutonomousSystemType
func MakeAutonomousSystemTypeSlice() []AutonomousSystemType {
	return []AutonomousSystemType{}
}
