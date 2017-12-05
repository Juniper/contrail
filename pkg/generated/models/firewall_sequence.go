package models

// FirewallSequence

import "encoding/json"

type FirewallSequence struct {
	Sequence string `json:"sequence"`
}

func (model *FirewallSequence) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFirewallSequence() *FirewallSequence {
	return &FirewallSequence{
		//TODO(nati): Apply default
		Sequence: "",
	}
}

func InterfaceToFirewallSequence(iData interface{}) *FirewallSequence {
	data := iData.(map[string]interface{})
	return &FirewallSequence{
		Sequence: data["sequence"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"sequence","Item":null,"GoName":"Sequence","GoType":"string"}

	}
}

func InterfaceToFirewallSequenceSlice(data interface{}) []*FirewallSequence {
	list := data.([]interface{})
	result := MakeFirewallSequenceSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallSequence(item))
	}
	return result
}

func MakeFirewallSequenceSlice() []*FirewallSequence {
	return []*FirewallSequence{}
}
