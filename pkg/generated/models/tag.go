package models

// Tag

import "encoding/json"

// Tag
type Tag struct {
	ParentType  string         `json:"parent_type,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	TagTypeName string         `json:"tag_type_name,omitempty"`
	TagValue    string         `json:"tag_value,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	TagID       U32BitHexInt   `json:"tag_id,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`

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
		UUID:        "",
		FQName:      []string{},
		TagID:       MakeU32BitHexInt(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		TagTypeName: "",
		TagValue:    "",
	}
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
	return []*Tag{}
}
