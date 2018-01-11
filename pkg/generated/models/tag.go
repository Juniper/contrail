package models

// Tag

import "encoding/json"

// Tag
type Tag struct {
	Perms2      *PermType2     `json:"perms2"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	TagTypeName string         `json:"tag_type_name"`
	TagValue    string         `json:"tag_value"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	TagID       U32BitHexInt   `json:"tag_id"`

	TagTypeRefs []*TagTagTypeRef `json:"tag_type_refs"`
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
		Perms2:      MakePermType2(),
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		TagTypeName: "",
		TagValue:    "",
		UUID:        "",
		ParentUUID:  "",
		TagID:       MakeU32BitHexInt(),
	}
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
	return []*Tag{}
}
