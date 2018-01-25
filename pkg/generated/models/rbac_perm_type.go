package models
// RbacPermType



import "encoding/json"

// RbacPermType 
//proteus:generate
type RbacPermType struct {

    RoleCrud string `json:"role_crud,omitempty"`
    RoleName string `json:"role_name,omitempty"`


}



// String returns json representation of the object
func (model *RbacPermType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

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
