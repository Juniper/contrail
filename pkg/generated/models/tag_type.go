package models

// TagType

import "encoding/json"

// TagType
type TagType struct {
	TagTypeID   U16BitHexInt   `json:"tag_type_id,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
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
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		DisplayName: "",
		TagTypeID:   MakeU16BitHexInt(),
		Perms2:      MakePermType2(),
		UUID:        "",
		IDPerms:     MakeIdPermsType(),
	}
}

// MakeTagTypeSlice() makes a slice of TagType
func MakeTagTypeSlice() []*TagType {
	return []*TagType{}
}
