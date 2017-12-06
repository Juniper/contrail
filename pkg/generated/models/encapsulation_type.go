package models

// EncapsulationType

type EncapsulationType []string

// MakeEncapsulationType makes EncapsulationType
func MakeEncapsulationType() EncapsulationType {
	var data EncapsulationType
	return data
}

// InterfaceToEncapsulationType makes EncapsulationType from interface
func InterfaceToEncapsulationType(data interface{}) EncapsulationType {
	return data.(EncapsulationType)
}

// InterfaceToEncapsulationTypeSlice makes a slice of EncapsulationType from interface
func InterfaceToEncapsulationTypeSlice(data interface{}) []EncapsulationType {
	list := data.([]interface{})
	result := MakeEncapsulationTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToEncapsulationType(item))
	}
	return result
}

// MakeEncapsulationTypeSlice() makes a slice of EncapsulationType
func MakeEncapsulationTypeSlice() []EncapsulationType {
	return []EncapsulationType{}
}
