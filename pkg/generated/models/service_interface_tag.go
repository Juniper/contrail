package models
// ServiceInterfaceTag



import "encoding/json"

// ServiceInterfaceTag 
//proteus:generate
type ServiceInterfaceTag struct {

    InterfaceType ServiceInterfaceType `json:"interface_type,omitempty"`


}



// String returns json representation of the object
func (model *ServiceInterfaceTag) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceInterfaceTag makes ServiceInterfaceTag
func MakeServiceInterfaceTag() *ServiceInterfaceTag{
    return &ServiceInterfaceTag{
    //TODO(nati): Apply default
    InterfaceType: MakeServiceInterfaceType(),
        
    }
}



// MakeServiceInterfaceTagSlice() makes a slice of ServiceInterfaceTag
func MakeServiceInterfaceTagSlice() []*ServiceInterfaceTag {
    return []*ServiceInterfaceTag{}
}
