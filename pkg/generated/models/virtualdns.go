package models

// VirtualDNS

import "encoding/json"

// VirtualDNS
type VirtualDNS struct {
	ParentUUID     string          `json:"parent_uuid"`
	IDPerms        *IdPermsType    `json:"id_perms"`
	DisplayName    string          `json:"display_name"`
	VirtualDNSData *VirtualDnsType `json:"virtual_DNS_data"`
	Annotations    *KeyValuePairs  `json:"annotations"`
	UUID           string          `json:"uuid"`
	Perms2         *PermType2      `json:"perms2"`
	ParentType     string          `json:"parent_type"`
	FQName         []string        `json:"fq_name"`

	VirtualDNSRecords []*VirtualDNSRecord `json:"virtual_DNS_records"`
}

// String returns json representation of the object
func (model *VirtualDNS) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDNS makes VirtualDNS
func MakeVirtualDNS() *VirtualDNS {
	return &VirtualDNS{
		//TODO(nati): Apply default
		DisplayName:    "",
		VirtualDNSData: MakeVirtualDnsType(),
		Annotations:    MakeKeyValuePairs(),
		UUID:           "",
		ParentUUID:     "",
		IDPerms:        MakeIdPermsType(),
		Perms2:         MakePermType2(),
		ParentType:     "",
		FQName:         []string{},
	}
}

// InterfaceToVirtualDNS makes VirtualDNS from interface
func InterfaceToVirtualDNS(iData interface{}) *VirtualDNS {
	data := iData.(map[string]interface{})
	return &VirtualDNS{
		VirtualDNSData: InterfaceToVirtualDnsType(data["virtual_DNS_data"]),

		//{"description":"Virtual DNS data has configuration for virtual DNS like domain, dynamic records etc.","type":"object","properties":{"default_ttl_seconds":{"type":"integer"},"domain_name":{"type":"string"},"dynamic_records_from_client":{"type":"boolean"},"external_visible":{"type":"boolean"},"floating_ip_record":{"type":"string","enum":["dashed-ip","dashed-ip-tenant-name","vm-name","vm-name-tenant-name"]},"next_virtual_DNS":{"type":"string"},"record_order":{"type":"string","enum":["fixed","random","round-robin"]},"reverse_resolution":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}

	}
}

// InterfaceToVirtualDNSSlice makes a slice of VirtualDNS from interface
func InterfaceToVirtualDNSSlice(data interface{}) []*VirtualDNS {
	list := data.([]interface{})
	result := MakeVirtualDNSSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualDNS(item))
	}
	return result
}

// MakeVirtualDNSSlice() makes a slice of VirtualDNS
func MakeVirtualDNSSlice() []*VirtualDNS {
	return []*VirtualDNS{}
}
