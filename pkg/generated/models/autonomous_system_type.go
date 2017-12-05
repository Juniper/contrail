package models

// AutonomousSystemType

type AutonomousSystemType int

func MakeAutonomousSystemType() AutonomousSystemType {
	var data AutonomousSystemType
	return data
}

func InterfaceToAutonomousSystemType(data interface{}) AutonomousSystemType {
	return data.(AutonomousSystemType)
}

func InterfaceToAutonomousSystemTypeSlice(data interface{}) []AutonomousSystemType {
	list := data.([]interface{})
	result := MakeAutonomousSystemTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAutonomousSystemType(item))
	}
	return result
}

func MakeAutonomousSystemTypeSlice() []AutonomousSystemType {
	return []AutonomousSystemType{}
}
