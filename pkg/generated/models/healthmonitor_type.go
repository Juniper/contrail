package models

// HealthmonitorType

type HealthmonitorType string

// MakeHealthmonitorType makes HealthmonitorType
func MakeHealthmonitorType() HealthmonitorType {
	var data HealthmonitorType
	return data
}

// InterfaceToHealthmonitorType makes HealthmonitorType from interface
func InterfaceToHealthmonitorType(data interface{}) HealthmonitorType {
	return data.(HealthmonitorType)
}

// InterfaceToHealthmonitorTypeSlice makes a slice of HealthmonitorType from interface
func InterfaceToHealthmonitorTypeSlice(data interface{}) []HealthmonitorType {
	list := data.([]interface{})
	result := MakeHealthmonitorTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToHealthmonitorType(item))
	}
	return result
}

// MakeHealthmonitorTypeSlice() makes a slice of HealthmonitorType
func MakeHealthmonitorTypeSlice() []HealthmonitorType {
	return []HealthmonitorType{}
}
