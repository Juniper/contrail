package models

// HealthmonitorType

type HealthmonitorType string

func MakeHealthmonitorType() HealthmonitorType {
	var data HealthmonitorType
	return data
}

func InterfaceToHealthmonitorType(data interface{}) HealthmonitorType {
	return data.(HealthmonitorType)
}

func InterfaceToHealthmonitorTypeSlice(data interface{}) []HealthmonitorType {
	list := data.([]interface{})
	result := MakeHealthmonitorTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToHealthmonitorType(item))
	}
	return result
}

func MakeHealthmonitorTypeSlice() []HealthmonitorType {
	return []HealthmonitorType{}
}
