package models


// MakePermType makes PermType
func MakePermType() *PermType{
    return &PermType{
    //TODO(nati): Apply default
    Owner: "",
        OwnerAccess: 0,
        OtherAccess: 0,
        Group: "",
        GroupAccess: 0,
        
    }
}

// MakePermTypeSlice() makes a slice of PermType
func MakePermTypeSlice() []*PermType {
    return []*PermType{}
}


