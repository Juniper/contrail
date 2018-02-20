package models

// OpenStackImageProperty

// OpenStackImageProperty
//proteus:generate
type OpenStackImageProperty struct {
	ID    string         `json:"id,omitempty"`
	Links *OpenStackLink `json:"links,omitempty"`
}

// MakeOpenStackImageProperty makes OpenStackImageProperty
func MakeOpenStackImageProperty() *OpenStackImageProperty {
	return &OpenStackImageProperty{
		//TODO(nati): Apply default
		ID:    "",
		Links: MakeOpenStackLink(),
	}
}

// MakeOpenStackImagePropertySlice() makes a slice of OpenStackImageProperty
func MakeOpenStackImagePropertySlice() []*OpenStackImageProperty {
	return []*OpenStackImageProperty{}
}
