package models

// VirtualDnsRecordType

import "encoding/json"

// VirtualDnsRecordType
type VirtualDnsRecordType struct {
	RecordMXPreference int                `json:"record_mx_preference"`
	RecordName         string             `json:"record_name"`
	RecordClass        DnsRecordClassType `json:"record_class"`
	RecordData         string             `json:"record_data"`
	RecordType         DnsRecordTypeType  `json:"record_type"`
	RecordTTLSeconds   int                `json:"record_ttl_seconds"`
}

//  parents relation object

// String returns json representation of the object
func (model *VirtualDnsRecordType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDnsRecordType makes VirtualDnsRecordType
func MakeVirtualDnsRecordType() *VirtualDnsRecordType {
	return &VirtualDnsRecordType{
		//TODO(nati): Apply default
		RecordType:         MakeDnsRecordTypeType(),
		RecordTTLSeconds:   0,
		RecordMXPreference: 0,
		RecordName:         "",
		RecordClass:        MakeDnsRecordClassType(),
		RecordData:         "",
	}
}

// InterfaceToVirtualDnsRecordType makes VirtualDnsRecordType from interface
func InterfaceToVirtualDnsRecordType(iData interface{}) *VirtualDnsRecordType {
	data := iData.(map[string]interface{})
	return &VirtualDnsRecordType{
		RecordType: InterfaceToDnsRecordTypeType(data["record_type"]),

		//{"Title":"","Description":"DNS record type can be A, AAAA, CNAME, PTR, NS and MX","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["A","AAAA","CNAME","PTR","NS","MX"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DnsRecordTypeType","CollectionType":"","Column":"","Item":null,"GoName":"RecordType","GoType":"DnsRecordTypeType","GoPremitive":false}
		RecordTTLSeconds: data["record_ttl_seconds"].(int),

		//{"Title":"","Description":"Time To Live for this DNS record","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RecordTTLSeconds","GoType":"int","GoPremitive":true}
		RecordMXPreference: data["record_mx_preference"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RecordMXPreference","GoType":"int","GoPremitive":true}
		RecordName: data["record_name"].(string),

		//{"Title":"","Description":"DNS name to be resolved","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RecordName","GoType":"string","GoPremitive":true}
		RecordClass: InterfaceToDnsRecordClassType(data["record_class"]),

		//{"Title":"","Description":"DNS record class supported is IN","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["IN"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DnsRecordClassType","CollectionType":"","Column":"","Item":null,"GoName":"RecordClass","GoType":"DnsRecordClassType","GoPremitive":false}
		RecordData: data["record_data"].(string),

		//{"Title":"","Description":"DNS record data is either ip address or string depending on type","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RecordData","GoType":"string","GoPremitive":true}

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
