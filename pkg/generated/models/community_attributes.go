package models

// CommunityAttributes

import "encoding/json"

// CommunityAttributes
type CommunityAttributes struct {
	CommunityAttribute CommunityAttribute `json:"community_attribute"`
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

// InterfaceToCommunityAttributes makes CommunityAttributes from interface
func InterfaceToCommunityAttributes(iData interface{}) *CommunityAttributes {
	data := iData.(map[string]interface{})
	return &CommunityAttributes{

		CommunityAttribute: InterfaceToCommunityAttribute(data["community_attribute"]),

		//{"type":"array"}

	}
}

// InterfaceToCommunityAttributesSlice makes a slice of CommunityAttributes from interface
func InterfaceToCommunityAttributesSlice(data interface{}) []*CommunityAttributes {
	list := data.([]interface{})
	result := MakeCommunityAttributesSlice()
	for _, item := range list {
		result = append(result, InterfaceToCommunityAttributes(item))
	}
	return result
}

// MakeCommunityAttributesSlice() makes a slice of CommunityAttributes
func MakeCommunityAttributesSlice() []*CommunityAttributes {
	return []*CommunityAttributes{}
}
