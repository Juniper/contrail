package models


// MakeOpenStackLink makes OpenStackLink
func MakeOpenStackLink() *OpenStackLink{
    return &OpenStackLink{
    //TODO(nati): Apply default
    Href: "",
        Rel: "",
        Type: "",
        
    }
}

// MakeOpenStackLinkSlice() makes a slice of OpenStackLink
func MakeOpenStackLinkSlice() []*OpenStackLink {
    return []*OpenStackLink{}
}


