package models
// EncapsulationPrioritiesType



import "encoding/json"

// EncapsulationPrioritiesType 
//proteus:generate
type EncapsulationPrioritiesType struct {

    Encapsulation EncapsulationType `json:"encapsulation,omitempty"`


}



// String returns json representation of the object
func (model *EncapsulationPrioritiesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesType() *EncapsulationPrioritiesType{
    return &EncapsulationPrioritiesType{
    //TODO(nati): Apply default
    
            
                Encapsulation:  MakeEncapsulationType(),
            
        
    }
}



// MakeEncapsulationPrioritiesTypeSlice() makes a slice of EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesTypeSlice() []*EncapsulationPrioritiesType {
    return []*EncapsulationPrioritiesType{}
}
