package models

// EcmpHashingIncludeFields

import "encoding/json"

// EcmpHashingIncludeFields
type EcmpHashingIncludeFields struct {
	SourcePort        bool `json:"source_port"`
	DestinationPort   bool `json:"destination_port"`
	DestinationIP     bool `json:"destination_ip"`
	IPProtocol        bool `json:"ip_protocol"`
	SourceIP          bool `json:"source_ip"`
	HashingConfigured bool `json:"hashing_configured"`
}

// String returns json representation of the object
func (model *EcmpHashingIncludeFields) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
