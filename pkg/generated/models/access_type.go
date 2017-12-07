package models

// AccessType

type AccessType int

// MakeAccessType makes AccessType
func MakeAccessType() AccessType {
	var data AccessType
	return data
}

// InterfaceToAccessType makes AccessType from interface
func InterfaceToAccessType(data interface{}) AccessType {
	return data.(AccessType)
}

// InterfaceToAccessTypeSlice makes a slice of AccessType from interface
func InterfaceToAccessTypeSlice(data interface{}) []AccessType {
	list := data.([]interface{})
	result := MakeAccessTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAccessType(item))
	}
	return result
}

// MakeAccessTypeSlice() makes a slice of AccessType
func MakeAccessTypeSlice() []AccessType {
	return []AccessType{}
}
