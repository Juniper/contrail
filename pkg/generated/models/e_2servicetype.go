package models

// E2servicetype

type E2servicetype string

// MakeE2servicetype makes E2servicetype
func MakeE2servicetype() E2servicetype {
	var data E2servicetype
	return data
}

// InterfaceToE2servicetype makes E2servicetype from interface
func InterfaceToE2servicetype(data interface{}) E2servicetype {
	return data.(E2servicetype)
}

// InterfaceToE2servicetypeSlice makes a slice of E2servicetype from interface
func InterfaceToE2servicetypeSlice(data interface{}) []E2servicetype {
	list := data.([]interface{})
	result := MakeE2servicetypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToE2servicetype(item))
	}
	return result
}

// MakeE2servicetypeSlice() makes a slice of E2servicetype
func MakeE2servicetypeSlice() []E2servicetype {
	return []E2servicetype{}
}
