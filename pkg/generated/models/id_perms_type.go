package models

// IdPermsType

import "encoding/json"

// IdPermsType
type IdPermsType struct {
	Creator      string    `json:"creator,omitempty"`
	UserVisible  bool      `json:"user_visible"`
	LastModified string    `json:"last_modified,omitempty"`
	Permissions  *PermType `json:"permissions,omitempty"`
	Enable       bool      `json:"enable"`
	Description  string    `json:"description,omitempty"`
	Created      string    `json:"created,omitempty"`
}

// String returns json representation of the object
func (model *IdPermsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIdPermsType makes IdPermsType
func MakeIdPermsType() *IdPermsType {
	return &IdPermsType{
		//TODO(nati): Apply default
		Creator:      "",
		UserVisible:  false,
		LastModified: "",
		Permissions:  MakePermType(),
		Enable:       false,
		Description:  "",
		Created:      "",
	}
}

// MakeIdPermsTypeSlice() makes a slice of IdPermsType
func MakeIdPermsTypeSlice() []*IdPermsType {
	return []*IdPermsType{}
}
