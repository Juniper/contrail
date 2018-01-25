package models

// PolicyEntriesType

// PolicyEntriesType
//proteus:generate
type PolicyEntriesType struct {
	PolicyRule []*PolicyRuleType `json:"policy_rule,omitempty"`
}

// MakePolicyEntriesType makes PolicyEntriesType
func MakePolicyEntriesType() *PolicyEntriesType {
	return &PolicyEntriesType{
		//TODO(nati): Apply default

		PolicyRule: MakePolicyRuleTypeSlice(),
	}
}

// MakePolicyEntriesTypeSlice() makes a slice of PolicyEntriesType
func MakePolicyEntriesTypeSlice() []*PolicyEntriesType {
	return []*PolicyEntriesType{}
}
