package models

// Tag

import "encoding/json"

// Tag
type Tag struct {
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	TagTypeName string         `json:"tag_type_name,omitempty"`
	TagID       U32BitHexInt   `json:"tag_id,omitempty"`
	TagValue    string         `json:"tag_value,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`

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
		DisplayName: "",
		Perms2:      MakePermType2(),
		TagTypeName: "",
		TagID:       MakeU32BitHexInt(),
		TagValue:    "",
		UUID:        "",
		ParentUUID:  "",
		IDPerms:     MakeIdPermsType(),
		ParentType:  "",
		FQName:      []string{},
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
	return []*Tag{}
}
