package models

// FirewallSequence

// FirewallSequence
//proteus:generate
type FirewallSequence struct {
	Sequence string `json:"sequence,omitempty"`
}

// MakeFirewallSequence makes FirewallSequence
func MakeFirewallSequence() *FirewallSequence {
	return &FirewallSequence{
		//TODO(nati): Apply default
		Sequence: "",
	}
}

// MakeFirewallSequenceSlice() makes a slice of FirewallSequence
func MakeFirewallSequenceSlice() []*FirewallSequence {
	return []*FirewallSequence{}
}
