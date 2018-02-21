package models


// MakeUser makes User
func MakeUser() *User{
    return &User{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Password: "",
        
    }
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
    return []*User{}
}


