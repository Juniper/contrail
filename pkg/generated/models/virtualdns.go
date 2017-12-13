package models

// VirtualDNS

import "encoding/json"

// VirtualDNS
type VirtualDNS struct {
	FQName         []string        `json:"fq_name"`
	ParentType     string          `json:"parent_type"`
	VirtualDNSData *VirtualDnsType `json:"virtual_DNS_data"`
	IDPerms        *IdPermsType    `json:"id_perms"`
	DisplayName    string          `json:"display_name"`
	Annotations    *KeyValuePairs  `json:"annotations"`
	Perms2         *PermType2      `json:"perms2"`
	UUID           string          `json:"uuid"`
	ParentUUID     string          `json:"parent_uuid"`

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
		VirtualDNSData: MakeVirtualDnsType(),
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		UUID:           "",
		ParentUUID:     "",
		FQName:         []string{},
		ParentType:     "",
	}
}

// InterfaceToVirtualDNS makes VirtualDNS from interface
func InterfaceToVirtualDNS(iData interface{}) *VirtualDNS {
	data := iData.(map[string]interface{})
	return &VirtualDNS{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		VirtualDNSData: InterfaceToVirtualDnsType(data["virtual_DNS_data"]),

		//{"description":"Virtual DNS data has configuration for virtual DNS like domain, dynamic records etc.","type":"object","properties":{"default_ttl_seconds":{"type":"integer"},"domain_name":{"type":"string"},"dynamic_records_from_client":{"type":"boolean"},"external_visible":{"type":"boolean"},"floating_ip_record":{"type":"string","enum":["dashed-ip","dashed-ip-tenant-name","vm-name","vm-name-tenant-name"]},"next_virtual_DNS":{"type":"string"},"record_order":{"type":"string","enum":["fixed","random","round-robin"]},"reverse_resolution":{"type":"boolean"}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}

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
