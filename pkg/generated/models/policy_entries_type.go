package models
// PolicyEntriesType



import "encoding/json"

// PolicyEntriesType 
//proteus:generate
type PolicyEntriesType struct {

    PolicyRule []*PolicyRuleType `json:"policy_rule,omitempty"`


}



// String returns json representation of the object
func (model *PolicyEntriesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePolicyEntriesType makes PolicyEntriesType
func MakePolicyEntriesType() *PolicyEntriesType{
    return &PolicyEntriesType{
    //TODO(nati): Apply default
    
            
                PolicyRule:  MakePolicyRuleTypeSlice(),
            
        
    }
}



// MakePolicyEntriesTypeSlice() makes a slice of PolicyEntriesType
func MakePolicyEntriesTypeSlice() []*PolicyEntriesType {
    return []*PolicyEntriesType{}
}
