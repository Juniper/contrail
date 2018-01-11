package models

// TagType

import "encoding/json"

// TagType
type TagType struct {
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Perms2      *PermType2     `json:"perms2"`
	TagTypeID   U16BitHexInt   `json:"tag_type_id"`
	UUID        string         `json:"uuid"`
	Annotations *KeyValuePairs `json:"annotations"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
}

// String returns json representation of the object
func (model *TagType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTagType makes TagType
func MakeTagType() *TagType {
	return &TagType{
		//TODO(nati): Apply default
		TagTypeID:   MakeU16BitHexInt(),
		UUID:        "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeTagTypeSlice() makes a slice of TagType
func MakeTagTypeSlice() []*TagType {
	return []*TagType{}
}
