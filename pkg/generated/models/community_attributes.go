package models

// CommunityAttributes

import "encoding/json"

type CommunityAttributes struct {
	CommunityAttribute CommunityAttribute `json:"community_attribute"`
}

func (model *CommunityAttributes) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeCommunityAttributes() *CommunityAttributes {
	return &CommunityAttributes{
		//TODO(nati): Apply default

		CommunityAttribute: MakeCommunityAttribute(),
	}
}

func InterfaceToCommunityAttributes(iData interface{}) *CommunityAttributes {
	data := iData.(map[string]interface{})
	return &CommunityAttributes{

		CommunityAttribute: InterfaceToCommunityAttribute(data["community_attribute"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute"}

	}
}

func InterfaceToCommunityAttributesSlice(data interface{}) []*CommunityAttributes {
	list := data.([]interface{})
	result := MakeCommunityAttributesSlice()
	for _, item := range list {
		result = append(result, InterfaceToCommunityAttributes(item))
	}
	return result
}

func MakeCommunityAttributesSlice() []*CommunityAttributes {
	return []*CommunityAttributes{}
}
