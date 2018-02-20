package models

// EcmpHashingIncludeFields

// EcmpHashingIncludeFields
//proteus:generate
type EcmpHashingIncludeFields struct {
	DestinationIP     bool `json:"destination_ip"`
	IPProtocol        bool `json:"ip_protocol"`
	SourceIP          bool `json:"source_ip"`
	HashingConfigured bool `json:"hashing_configured"`
	SourcePort        bool `json:"source_port"`
	DestinationPort   bool `json:"destination_port"`
}

// MakeEcmpHashingIncludeFields makes EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFields() *EcmpHashingIncludeFields {
	return &EcmpHashingIncludeFields{
		//TODO(nati): Apply default
		DestinationIP:     false,
		IPProtocol:        false,
		SourceIP:          false,
		HashingConfigured: false,
		SourcePort:        false,
		DestinationPort:   false,
	}
}

// MakeEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
	return []*EcmpHashingIncludeFields{}
}
