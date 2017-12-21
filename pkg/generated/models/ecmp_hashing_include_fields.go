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

// InterfaceToEcmpHashingIncludeFields makes EcmpHashingIncludeFields from interface
func InterfaceToEcmpHashingIncludeFields(iData interface{}) *EcmpHashingIncludeFields {
	data := iData.(map[string]interface{})
	return &EcmpHashingIncludeFields{
		SourceIP: data["source_ip"].(bool),

		//{"description":"When false, do not use source ip in the ECMP hash calculation","type":"boolean"}
		HashingConfigured: data["hashing_configured"].(bool),

		//{"description":"When True, Packet header fields used in calculating ECMP hash is decided by following flags","type":"boolean"}
		SourcePort: data["source_port"].(bool),

		//{"description":"When false, do not use source port in the ECMP hash calculation","type":"boolean"}
		DestinationPort: data["destination_port"].(bool),

		//{"description":"When false, do not use destination port in the ECMP hash calculation","type":"boolean"}
		DestinationIP: data["destination_ip"].(bool),

		//{"description":"When false, do not use destination ip in the ECMP hash calculation","type":"boolean"}
		IPProtocol: data["ip_protocol"].(bool),

		//{"description":"When false, do not use ip protocol in the ECMP hash calculation","type":"boolean"}

	}
}

// InterfaceToEcmpHashingIncludeFieldsSlice makes a slice of EcmpHashingIncludeFields from interface
func InterfaceToEcmpHashingIncludeFieldsSlice(data interface{}) []*EcmpHashingIncludeFields {
	list := data.([]interface{})
	result := MakeEcmpHashingIncludeFieldsSlice()
	for _, item := range list {
		result = append(result, InterfaceToEcmpHashingIncludeFields(item))
	}
	return result
}

// MakeEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
	return []*EcmpHashingIncludeFields{}
}
