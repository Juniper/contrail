package models

// EcmpHashingIncludeFields

import "encoding/json"

// EcmpHashingIncludeFields
type EcmpHashingIncludeFields struct {
	DestinationIP     bool `json:"destination_ip,omitempty"`
	IPProtocol        bool `json:"ip_protocol,omitempty"`
	SourceIP          bool `json:"source_ip,omitempty"`
	HashingConfigured bool `json:"hashing_configured,omitempty"`
	SourcePort        bool `json:"source_port,omitempty"`
	DestinationPort   bool `json:"destination_port,omitempty"`
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
		SourceIP:          false,
		HashingConfigured: false,
		SourcePort:        false,
		DestinationPort:   false,
		DestinationIP:     false,
		IPProtocol:        false,
	}
}

// MakeEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
	return []*EcmpHashingIncludeFields{}
}
