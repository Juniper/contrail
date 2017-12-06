package models

// CommunityAttribute

type CommunityAttribute []string

// MakeCommunityAttribute makes CommunityAttribute
func MakeCommunityAttribute() CommunityAttribute {
	var data CommunityAttribute
	return data
}

// InterfaceToCommunityAttribute makes CommunityAttribute from interface
func InterfaceToCommunityAttribute(data interface{}) CommunityAttribute {
	return data.(CommunityAttribute)
}

// InterfaceToCommunityAttributeSlice makes a slice of CommunityAttribute from interface
func InterfaceToCommunityAttributeSlice(data interface{}) []CommunityAttribute {
	list := data.([]interface{})
	result := MakeCommunityAttributeSlice()
	for _, item := range list {
		result = append(result, InterfaceToCommunityAttribute(item))
	}
	return result
}

// MakeCommunityAttributeSlice() makes a slice of CommunityAttribute
func MakeCommunityAttributeSlice() []CommunityAttribute {
	return []CommunityAttribute{}
}
