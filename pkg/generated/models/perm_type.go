package models

// PermType

import "encoding/json"

// PermType
type PermType struct {
	Owner       string     `json:"owner,omitempty"`
	OwnerAccess AccessType `json:"owner_access,omitempty"`
	OtherAccess AccessType `json:"other_access,omitempty"`
	Group       string     `json:"group,omitempty"`
	GroupAccess AccessType `json:"group_access,omitempty"`
}

// String returns json representation of the object
func (model *PermType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePermType makes PermType
func MakePermType() *PermType {
	return &PermType{
		//TODO(nati): Apply default
		Owner:       "",
		OwnerAccess: MakeAccessType(),
		OtherAccess: MakeAccessType(),
		Group:       "",
		GroupAccess: MakeAccessType(),
	}
}

// MakePermTypeSlice() makes a slice of PermType
func MakePermTypeSlice() []*PermType {
	return []*PermType{}
}
