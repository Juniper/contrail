package models
// SecurityLoggingObjectRuleListType



import "encoding/json"

// SecurityLoggingObjectRuleListType 
//proteus:generate
type SecurityLoggingObjectRuleListType struct {

    Rule []*SecurityLoggingObjectRuleEntryType `json:"rule,omitempty"`


}



// String returns json representation of the object
func (model *SecurityLoggingObjectRuleListType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType{
    return &SecurityLoggingObjectRuleListType{
    //TODO(nati): Apply default
    
            
                Rule:  MakeSecurityLoggingObjectRuleEntryTypeSlice(),
            
        
    }
}



// MakeSecurityLoggingObjectRuleListTypeSlice() makes a slice of SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
    return []*SecurityLoggingObjectRuleListType{}
}
