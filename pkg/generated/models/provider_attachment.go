package models

// ProviderAttachment

import "encoding/json"

// ProviderAttachment
type ProviderAttachment struct {
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`

	VirtualRouterRefs []*ProviderAttachmentVirtualRouterRef `json:"virtual_router_refs,omitempty"`
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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeProviderAttachmentSlice() makes a slice of ProviderAttachment
func MakeProviderAttachmentSlice() []*ProviderAttachment {
	return []*ProviderAttachment{}
}
