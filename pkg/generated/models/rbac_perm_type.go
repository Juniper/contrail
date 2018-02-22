package models


// MakeRbacPermType makes RbacPermType
func MakeRbacPermType() *RbacPermType{
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


