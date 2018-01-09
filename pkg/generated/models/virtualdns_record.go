package models

// VirtualDNSRecord

import "encoding/json"

// VirtualDNSRecord
type VirtualDNSRecord struct {
	ParentUUID           string                `json:"parent_uuid"`
	ParentType           string                `json:"parent_type"`
	Annotations          *KeyValuePairs        `json:"annotations"`
	VirtualDNSRecordData *VirtualDnsRecordType `json:"virtual_DNS_record_data"`
	FQName               []string              `json:"fq_name"`
	IDPerms              *IdPermsType          `json:"id_perms"`
	DisplayName          string                `json:"display_name"`
	Perms2               *PermType2            `json:"perms2"`
	UUID                 string                `json:"uuid"`
}

// String returns json representation of the object
func (model *VirtualDNSRecord) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDNSRecord makes VirtualDNSRecord
func MakeVirtualDNSRecord() *VirtualDNSRecord {
	return &VirtualDNSRecord{
		//TODO(nati): Apply default
		Perms2:               MakePermType2(),
		UUID:                 "",
		VirtualDNSRecordData: MakeVirtualDnsRecordType(),
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		ParentUUID:           "",
		ParentType:           "",
		Annotations:          MakeKeyValuePairs(),
	}
}

// InterfaceToVirtualDNSRecord makes VirtualDNSRecord from interface
func InterfaceToVirtualDNSRecord(iData interface{}) *VirtualDNSRecord {
	data := iData.(map[string]interface{})
	return &VirtualDNSRecord{
		VirtualDNSRecordData: InterfaceToVirtualDnsRecordType(data["virtual_DNS_record_data"]),

		//{"description":"DNS record data has configuration like type, name, ip address, loadbalancing etc.","type":"object","properties":{"record_class":{"type":"string","enum":["IN"]},"record_data":{"type":"string"},"record_mx_preference":{"type":"integer"},"record_name":{"type":"string"},"record_ttl_seconds":{"type":"integer"},"record_type":{"type":"string","enum":["A","AAAA","CNAME","PTR","NS","MX"]}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToVirtualDNSRecordSlice makes a slice of VirtualDNSRecord from interface
func InterfaceToVirtualDNSRecordSlice(data interface{}) []*VirtualDNSRecord {
	list := data.([]interface{})
	result := MakeVirtualDNSRecordSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualDNSRecord(item))
	}
	return result
}

// MakeVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
	return []*VirtualDNSRecord{}
}
