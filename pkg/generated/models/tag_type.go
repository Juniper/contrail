package models

// TagType

import "encoding/json"

// TagType
type TagType struct {
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	TagTypeID   U16BitHexInt   `json:"tag_type_id,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
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
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		IDPerms:     MakeIdPermsType(),
		TagTypeID:   MakeU16BitHexInt(),
		DisplayName: "",
		ParentType:  "",
		FQName:      []string{},
	}
}

// MakeTagTypeSlice() makes a slice of TagType
func MakeTagTypeSlice() []*TagType {
	return []*TagType{}
}
