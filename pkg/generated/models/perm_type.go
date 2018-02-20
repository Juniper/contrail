package models

// PermType

// PermType
//proteus:generate
type PermType struct {
	Owner       string     `json:"owner,omitempty"`
	OwnerAccess AccessType `json:"owner_access,omitempty"`
	OtherAccess AccessType `json:"other_access,omitempty"`
	Group       string     `json:"group,omitempty"`
	GroupAccess AccessType `json:"group_access,omitempty"`
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
