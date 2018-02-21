package models


// MakeIdPermsType makes IdPermsType
func MakeIdPermsType() *IdPermsType{
    return &IdPermsType{
    //TODO(nati): Apply default
    Enable: false,
        Description: "",
        Created: "",
        Creator: "",
        UserVisible: false,
        LastModified: "",
        Permissions: MakePermType(),
        
    }
}

// MakeIdPermsTypeSlice() makes a slice of IdPermsType
func MakeIdPermsTypeSlice() []*IdPermsType {
    return []*IdPermsType{}
}


