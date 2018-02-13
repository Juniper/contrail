package models

// CommunityAttributes

import "encoding/json"

// CommunityAttributes
//proteus:generate
type CommunityAttributes struct {
	CommunityAttribute CommunityAttribute `json:"community_attribute,omitempty"`
}

// String returns json representation of the object
func (model *CommunityAttributes) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
