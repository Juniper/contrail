package models


// MakeShareType makes ShareType
func MakeShareType() *ShareType{
    return &ShareType{
    //TODO(nati): Apply default
    TenantAccess: 0,
        Tenant: "",
        
    }
}

// MakeShareTypeSlice() makes a slice of ShareType
func MakeShareTypeSlice() []*ShareType {
    return []*ShareType{}
}


