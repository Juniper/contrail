package models

// EncapsulationType

type EncapsulationType []string

func MakeEncapsulationType() EncapsulationType {
	var data EncapsulationType
	return data
}

func InterfaceToEncapsulationType(data interface{}) EncapsulationType {
	return data.(EncapsulationType)
}

func InterfaceToEncapsulationTypeSlice(data interface{}) []EncapsulationType {
	list := data.([]interface{})
	result := MakeEncapsulationTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToEncapsulationType(item))
	}
	return result
}

func MakeEncapsulationTypeSlice() []EncapsulationType {
	return []EncapsulationType{}
}
