package models

// VirtualDnsRecordType

import "encoding/json"

// VirtualDnsRecordType
type VirtualDnsRecordType struct {
	RecordName         string             `json:"record_name"`
	RecordClass        DnsRecordClassType `json:"record_class"`
	RecordData         string             `json:"record_data"`
	RecordType         DnsRecordTypeType  `json:"record_type"`
	RecordTTLSeconds   int                `json:"record_ttl_seconds"`
	RecordMXPreference int                `json:"record_mx_preference"`
}

// String returns json representation of the object
func (model *VirtualDnsRecordType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDnsRecordType makes VirtualDnsRecordType
func MakeVirtualDnsRecordType() *VirtualDnsRecordType {
	return &VirtualDnsRecordType{
		//TODO(nati): Apply default
		RecordData:         "",
		RecordType:         MakeDnsRecordTypeType(),
		RecordTTLSeconds:   0,
		RecordMXPreference: 0,
		RecordName:         "",
		RecordClass:        MakeDnsRecordClassType(),
	}
}

// InterfaceToVirtualDnsRecordType makes VirtualDnsRecordType from interface
func InterfaceToVirtualDnsRecordType(iData interface{}) *VirtualDnsRecordType {
	data := iData.(map[string]interface{})
	return &VirtualDnsRecordType{
		RecordType: InterfaceToDnsRecordTypeType(data["record_type"]),

		//{"description":"DNS record type can be A, AAAA, CNAME, PTR, NS and MX","type":"string","enum":["A","AAAA","CNAME","PTR","NS","MX"]}
		RecordTTLSeconds: data["record_ttl_seconds"].(int),

		//{"description":"Time To Live for this DNS record","type":"integer"}
		RecordMXPreference: data["record_mx_preference"].(int),

		//{"type":"integer"}
		RecordName: data["record_name"].(string),

		//{"description":"DNS name to be resolved","type":"string"}
		RecordClass: InterfaceToDnsRecordClassType(data["record_class"]),

		//{"description":"DNS record class supported is IN","type":"string","enum":["IN"]}
		RecordData: data["record_data"].(string),

		//{"description":"DNS record data is either ip address or string depending on type","type":"string"}

	}
}

// InterfaceToVirtualDnsRecordTypeSlice makes a slice of VirtualDnsRecordType from interface
func InterfaceToVirtualDnsRecordTypeSlice(data interface{}) []*VirtualDnsRecordType {
	list := data.([]interface{})
	result := MakeVirtualDnsRecordTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualDnsRecordType(item))
	}
	return result
}

// MakeVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
func MakeVirtualDnsRecordTypeSlice() []*VirtualDnsRecordType {
	return []*VirtualDnsRecordType{}
}
