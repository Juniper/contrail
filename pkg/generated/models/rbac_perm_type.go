package models

// RbacPermType

// RbacPermType
//proteus:generate
type RbacPermType struct {
	RoleCrud string `json:"role_crud,omitempty"`
	RoleName string `json:"role_name,omitempty"`
}

// MakeRbacPermType makes RbacPermType
func MakeRbacPermType() *RbacPermType {
	return &RbacPermType{
		//TODO(nati): Apply default
		RoleCrud: "",
		RoleName: "",
	}
}

// MakeRbacPermTypeSlice() makes a slice of RbacPermType
func MakeRbacPermTypeSlice() []*RbacPermType {
	return []*RbacPermType{}
}
