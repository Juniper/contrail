package models


// MakeOpenStackFlavorProperty makes OpenStackFlavorProperty
func MakeOpenStackFlavorProperty() *OpenStackFlavorProperty{
    return &OpenStackFlavorProperty{
    //TODO(nati): Apply default
    ID: "",
        Links: MakeOpenStackLink(),
        
    }
}

// MakeOpenStackFlavorPropertySlice() makes a slice of OpenStackFlavorProperty
func MakeOpenStackFlavorPropertySlice() []*OpenStackFlavorProperty {
    return []*OpenStackFlavorProperty{}
}


