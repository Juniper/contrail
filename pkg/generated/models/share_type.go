package models
// ShareType



import "encoding/json"

// ShareType 
//proteus:generate
type ShareType struct {

    TenantAccess AccessType `json:"tenant_access,omitempty"`
    Tenant string `json:"tenant,omitempty"`


}



// String returns json representation of the object
func (model *ShareType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeShareType makes ShareType
func MakeShareType() *ShareType{
    return &ShareType{
    //TODO(nati): Apply default
    TenantAccess: MakeAccessType(),
        Tenant: "",
        
    }
}



// MakeShareTypeSlice() makes a slice of ShareType
func MakeShareTypeSlice() []*ShareType {
    return []*ShareType{}
}
