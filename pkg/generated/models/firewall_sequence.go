package models
// FirewallSequence



import "encoding/json"

// FirewallSequence 
//proteus:generate
type FirewallSequence struct {

    Sequence string `json:"sequence,omitempty"`


}



// String returns json representation of the object
func (model *FirewallSequence) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallSequence makes FirewallSequence
func MakeFirewallSequence() *FirewallSequence{
    return &FirewallSequence{
    //TODO(nati): Apply default
    Sequence: "",
        
    }
}



// MakeFirewallSequenceSlice() makes a slice of FirewallSequence
func MakeFirewallSequenceSlice() []*FirewallSequence {
    return []*FirewallSequence{}
}
