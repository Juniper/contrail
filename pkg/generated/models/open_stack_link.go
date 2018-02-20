package models

// OpenStackLink

// OpenStackLink
//proteus:generate
type OpenStackLink struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
	Type string `json:"type,omitempty"`
}

// MakeOpenStackLink makes OpenStackLink
func MakeOpenStackLink() *OpenStackLink {
	return &OpenStackLink{
		//TODO(nati): Apply default
		Href: "",
		Rel:  "",
		Type: "",
	}
}

// MakeOpenStackLinkSlice() makes a slice of OpenStackLink
func MakeOpenStackLinkSlice() []*OpenStackLink {
	return []*OpenStackLink{}
}
