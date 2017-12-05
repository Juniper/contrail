package models

// CommunityAttribute

type CommunityAttribute []string

func MakeCommunityAttribute() CommunityAttribute {
	var data CommunityAttribute
	return data
}

func InterfaceToCommunityAttribute(data interface{}) CommunityAttribute {
	return data.(CommunityAttribute)
}

func InterfaceToCommunityAttributeSlice(data interface{}) []CommunityAttribute {
	list := data.([]interface{})
	result := MakeCommunityAttributeSlice()
	for _, item := range list {
		result = append(result, InterfaceToCommunityAttribute(item))
	}
	return result
}

func MakeCommunityAttributeSlice() []CommunityAttribute {
	return []CommunityAttribute{}
}
