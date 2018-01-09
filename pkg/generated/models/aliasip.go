package models

// AliasIP

import "encoding/json"

// AliasIP
type AliasIP struct {
	ParentType           string              `json:"parent_type"`
	IDPerms              *IdPermsType        `json:"id_perms"`
	DisplayName          string              `json:"display_name"`
	AliasIPAddress       IpAddressType       `json:"alias_ip_address"`
	AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family"`
	Annotations          *KeyValuePairs      `json:"annotations"`
	Perms2               *PermType2          `json:"perms2"`
	UUID                 string              `json:"uuid"`
	ParentUUID           string              `json:"parent_uuid"`
	FQName               []string            `json:"fq_name"`

	ProjectRefs                 []*AliasIPProjectRef                 `json:"project_refs"`
	VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// AliasIPProjectRef references each other
type AliasIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// AliasIPVirtualMachineInterfaceRef references each other
type AliasIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *AliasIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP {
	return &AliasIP{
		//TODO(nati): Apply default
		ParentType:           "",
		IDPerms:              MakeIdPermsType(),
		AliasIPAddress:       MakeIpAddressType(),
		AliasIPAddressFamily: MakeIpAddressFamilyType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentUUID:           "",
		FQName:               []string{},
		DisplayName:          "",
	}
}

// InterfaceToAliasIP makes AliasIP from interface
func InterfaceToAliasIP(iData interface{}) *AliasIP {
	data := iData.(map[string]interface{})
	return &AliasIP{
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		AliasIPAddress: InterfaceToIpAddressType(data["alias_ip_address"]),

		//{"description":"Alias ip address.","type":"string"}
		AliasIPAddressFamily: InterfaceToIpAddressFamilyType(data["alias_ip_address_family"]),

		//{"description":"Ip address family of the alias ip, IpV4 or IpV6","type":"string","enum":["v4","v6"]}

	}
}

// InterfaceToAliasIPSlice makes a slice of AliasIP from interface
func InterfaceToAliasIPSlice(data interface{}) []*AliasIP {
	list := data.([]interface{})
	result := MakeAliasIPSlice()
	for _, item := range list {
		result = append(result, InterfaceToAliasIP(item))
	}
	return result
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}
