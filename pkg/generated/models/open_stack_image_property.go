package models


// MakeOpenStackImageProperty makes OpenStackImageProperty
func MakeOpenStackImageProperty() *OpenStackImageProperty{
    return &OpenStackImageProperty{
    //TODO(nati): Apply default
    ID: "",
        Links: MakeOpenStackLink(),
        
    }
}

// MakeOpenStackImagePropertySlice() makes a slice of OpenStackImageProperty
func MakeOpenStackImagePropertySlice() []*OpenStackImageProperty {
    return []*OpenStackImageProperty{}
}


