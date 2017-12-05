package models

// AccessType

type AccessType int

func MakeAccessType() AccessType {
	var data AccessType
	return data
}

func InterfaceToAccessType(data interface{}) AccessType {
	return data.(AccessType)
}

func InterfaceToAccessTypeSlice(data interface{}) []AccessType {
	list := data.([]interface{})
	result := MakeAccessTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAccessType(item))
	}
	return result
}

func MakeAccessTypeSlice() []AccessType {
	return []AccessType{}
}
