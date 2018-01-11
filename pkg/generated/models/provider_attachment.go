package models

// ProviderAttachment

import "encoding/json"

// ProviderAttachment
type ProviderAttachment struct {
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`

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
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
	}
}

// MakeProviderAttachmentSlice() makes a slice of ProviderAttachment
func MakeProviderAttachmentSlice() []*ProviderAttachment {
	return []*ProviderAttachment{}
}
