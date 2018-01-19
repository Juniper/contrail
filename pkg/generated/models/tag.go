package models

// Tag

import "encoding/json"

// Tag
type Tag struct {
	TagTypeName string         `json:"tag_type_name,omitempty"`
	TagValue    string         `json:"tag_value,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	TagID       U32BitHexInt   `json:"tag_id,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`

	TagTypeRefs []*TagTagTypeRef `json:"tag_type_refs,omitempty"`
}

// TagTagTypeRef references each other
type TagTagTypeRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *Tag) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTag makes Tag
func MakeTag() *Tag {
	return &Tag{
		//TODO(nati): Apply default
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
		TagID:       MakeU32BitHexInt(),
		UUID:        "",
		ParentUUID:  "",
		DisplayName: "",
		TagTypeName: "",
		TagValue:    "",
		Perms2:      MakePermType2(),
	}
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
	return []*Tag{}
}
