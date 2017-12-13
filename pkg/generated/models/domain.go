package models

// Domain

import "encoding/json"

// Domain
type Domain struct {
	Perms2       *PermType2        `json:"perms2"`
	ParentType   string            `json:"parent_type"`
	FQName       []string          `json:"fq_name"`
	DisplayName  string            `json:"display_name"`
	Annotations  *KeyValuePairs    `json:"annotations"`
	DomainLimits *DomainLimitsType `json:"domain_limits"`
	ParentUUID   string            `json:"parent_uuid"`
	IDPerms      *IdPermsType      `json:"id_perms"`
	UUID         string            `json:"uuid"`

	APIAccessLists   []*APIAccessList   `json:"api_access_lists"`
	Namespaces       []*Namespace       `json:"namespaces"`
	Projects         []*Project         `json:"projects"`
	ServiceTemplates []*ServiceTemplate `json:"service_templates"`
	VirtualDNSs      []*VirtualDNS      `json:"virtual_DNSs"`
}

// String returns json representation of the object
func (model *Domain) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDomain makes Domain
func MakeDomain() *Domain {
	return &Domain{
		//TODO(nati): Apply default
		DomainLimits: MakeDomainLimitsType(),
		Perms2:       MakePermType2(),
		ParentType:   "",
		FQName:       []string{},
		DisplayName:  "",
		Annotations:  MakeKeyValuePairs(),
		UUID:         "",
		ParentUUID:   "",
		IDPerms:      MakeIdPermsType(),
	}
}

// InterfaceToDomain makes Domain from interface
func InterfaceToDomain(iData interface{}) *Domain {
	data := iData.(map[string]interface{})
	return &Domain{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		DomainLimits: InterfaceToDomainLimitsType(data["domain_limits"]),

		//{"description":"Domain level quota, not currently implemented","type":"object","properties":{"project_limit":{"type":"integer"},"security_group_limit":{"type":"integer"},"virtual_network_limit":{"type":"integer"}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToDomainSlice makes a slice of Domain from interface
func InterfaceToDomainSlice(data interface{}) []*Domain {
	list := data.([]interface{})
	result := MakeDomainSlice()
	for _, item := range list {
		result = append(result, InterfaceToDomain(item))
	}
	return result
}

// MakeDomainSlice() makes a slice of Domain
func MakeDomainSlice() []*Domain {
	return []*Domain{}
}
