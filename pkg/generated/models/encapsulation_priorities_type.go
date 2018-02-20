package models

// EncapsulationPrioritiesType

// EncapsulationPrioritiesType
//proteus:generate
type EncapsulationPrioritiesType struct {
	Encapsulation EncapsulationType `json:"encapsulation,omitempty"`
}

// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesType() *EncapsulationPrioritiesType {
	return &EncapsulationPrioritiesType{
		//TODO(nati): Apply default

		Encapsulation: MakeEncapsulationType(),
	}
}

// MakeEncapsulationPrioritiesTypeSlice() makes a slice of EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesTypeSlice() []*EncapsulationPrioritiesType {
	return []*EncapsulationPrioritiesType{}
}
