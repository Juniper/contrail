package models

// CommunityAttributes

// CommunityAttributes
//proteus:generate
type CommunityAttributes struct {
	CommunityAttribute CommunityAttribute `json:"community_attribute,omitempty"`
}

// MakeCommunityAttributes makes CommunityAttributes
func MakeCommunityAttributes() *CommunityAttributes {
	return &CommunityAttributes{
		//TODO(nati): Apply default

		CommunityAttribute: MakeCommunityAttribute(),
	}
}

// MakeCommunityAttributesSlice() makes a slice of CommunityAttributes
func MakeCommunityAttributesSlice() []*CommunityAttributes {
	return []*CommunityAttributes{}
}
