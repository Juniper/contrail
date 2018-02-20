package models

// IdPermsType

// IdPermsType
//proteus:generate
type IdPermsType struct {
	Enable       bool      `json:"enable"`
	Description  string    `json:"description,omitempty"`
	Created      string    `json:"created,omitempty"`
	Creator      string    `json:"creator,omitempty"`
	UserVisible  bool      `json:"user_visible"`
	LastModified string    `json:"last_modified,omitempty"`
	Permissions  *PermType `json:"permissions,omitempty"`
}

// MakeIdPermsType makes IdPermsType
func MakeIdPermsType() *IdPermsType {
	return &IdPermsType{
		//TODO(nati): Apply default
		Enable:       false,
		Description:  "",
		Created:      "",
		Creator:      "",
		UserVisible:  false,
		LastModified: "",
		Permissions:  MakePermType(),
	}
}

// MakeIdPermsTypeSlice() makes a slice of IdPermsType
func MakeIdPermsTypeSlice() []*IdPermsType {
	return []*IdPermsType{}
}
