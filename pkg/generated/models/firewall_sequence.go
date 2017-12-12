package models

// FirewallSequence

import "encoding/json"

// FirewallSequence
type FirewallSequence struct {
	Sequence string `json:"sequence"`
}

// String returns json representation of the object
func (model *FirewallSequence) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallSequence makes FirewallSequence
func MakeFirewallSequence() *FirewallSequence {
	return &FirewallSequence{
		//TODO(nati): Apply default
		Sequence: "",
	}
}

// InterfaceToFirewallSequence makes FirewallSequence from interface
func InterfaceToFirewallSequence(iData interface{}) *FirewallSequence {
	data := iData.(map[string]interface{})
	return &FirewallSequence{
		Sequence: data["sequence"].(string),

		//{"type":"string"}

	}
}

// InterfaceToFirewallSequenceSlice makes a slice of FirewallSequence from interface
func InterfaceToFirewallSequenceSlice(data interface{}) []*FirewallSequence {
	list := data.([]interface{})
	result := MakeFirewallSequenceSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallSequence(item))
	}
	return result
}

// MakeFirewallSequenceSlice() makes a slice of FirewallSequence
func MakeFirewallSequenceSlice() []*FirewallSequence {
	return []*FirewallSequence{}
}
