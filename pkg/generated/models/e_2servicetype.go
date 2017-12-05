package models

// E2servicetype

type E2servicetype string

func MakeE2servicetype() E2servicetype {
	var data E2servicetype
	return data
}

func InterfaceToE2servicetype(data interface{}) E2servicetype {
	return data.(E2servicetype)
}

func InterfaceToE2servicetypeSlice(data interface{}) []E2servicetype {
	list := data.([]interface{})
	result := MakeE2servicetypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToE2servicetype(item))
	}
	return result
}

func MakeE2servicetypeSlice() []E2servicetype {
	return []E2servicetype{}
}
