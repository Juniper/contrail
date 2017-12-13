package models

// ProviderAttachment

import "encoding/json"

// ProviderAttachment
type ProviderAttachment struct {
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`

	VirtualRouterRefs []*ProviderAttachmentVirtualRouterRef `json:"virtual_router_refs"`
}

// ProviderAttachmentVirtualRouterRef references each other
type ProviderAttachmentVirtualRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ProviderAttachment) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeProviderAttachment makes ProviderAttachment
func MakeProviderAttachment() *ProviderAttachment {
	return &ProviderAttachment{
		//TODO(nati): Apply default
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// InterfaceToProviderAttachment makes ProviderAttachment from interface
func InterfaceToProviderAttachment(iData interface{}) *ProviderAttachment {
	data := iData.(map[string]interface{})
	return &ProviderAttachment{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
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
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToProviderAttachmentSlice makes a slice of ProviderAttachment from interface
func InterfaceToProviderAttachmentSlice(data interface{}) []*ProviderAttachment {
	list := data.([]interface{})
	result := MakeProviderAttachmentSlice()
	for _, item := range list {
		result = append(result, InterfaceToProviderAttachment(item))
	}
	return result
}

// MakeProviderAttachmentSlice() makes a slice of ProviderAttachment
func MakeProviderAttachmentSlice() []*ProviderAttachment {
	return []*ProviderAttachment{}
}
