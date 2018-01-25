package models

// MemberType

// MemberType
//proteus:generate
type MemberType struct {
	Role string `json:"role,omitempty"`
}

// MakeMemberType makes MemberType
func MakeMemberType() *MemberType {
	return &MemberType{
		//TODO(nati): Apply default
		Role: "",
	}
}

// MakeMemberTypeSlice() makes a slice of MemberType
func MakeMemberTypeSlice() []*MemberType {
	return []*MemberType{}
}
