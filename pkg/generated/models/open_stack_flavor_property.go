package models

// OpenStackFlavorProperty

// OpenStackFlavorProperty
//proteus:generate
type OpenStackFlavorProperty struct {
	ID    string         `json:"id,omitempty"`
	Links *OpenStackLink `json:"links,omitempty"`
}

// MakeOpenStackFlavorProperty makes OpenStackFlavorProperty
func MakeOpenStackFlavorProperty() *OpenStackFlavorProperty {
	return &OpenStackFlavorProperty{
		//TODO(nati): Apply default
		ID:    "",
		Links: MakeOpenStackLink(),
	}
}

// MakeOpenStackFlavorPropertySlice() makes a slice of OpenStackFlavorProperty
func MakeOpenStackFlavorPropertySlice() []*OpenStackFlavorProperty {
	return []*OpenStackFlavorProperty{}
}
