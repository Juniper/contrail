package models

// Tag

import "encoding/json"

// Tag
type Tag struct {
	DisplayName string         `json:"display_name,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	TagID       U32BitHexInt   `json:"tag_id,omitempty"`
	TagValue    string         `json:"tag_value,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	TagTypeName string         `json:"tag_type_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	UUID        string         `json:"uuid,omitempty"`

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
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		TagID:       MakeU32BitHexInt(),
		TagValue:    "",
		UUID:        "",
		FQName:      []string{},
		TagTypeName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
	return []*Tag{}
}
