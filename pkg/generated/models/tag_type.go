package models

// TagType

import "encoding/json"

// TagType
type TagType struct {
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	TagTypeID   U16BitHexInt   `json:"tag_type_id,omitempty"`
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
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeTagTypeSlice() makes a slice of TagType
func MakeTagTypeSlice() []*TagType {
	return []*TagType{}
}
