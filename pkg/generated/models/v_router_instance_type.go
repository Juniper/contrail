package models

// VRouterInstanceType

type VRouterInstanceType string

// MakeVRouterInstanceType makes VRouterInstanceType
func MakeVRouterInstanceType() VRouterInstanceType {
	var data VRouterInstanceType
	return data
}

// InterfaceToVRouterInstanceType makes VRouterInstanceType from interface
func InterfaceToVRouterInstanceType(data interface{}) VRouterInstanceType {
	return data.(VRouterInstanceType)
}

// InterfaceToVRouterInstanceTypeSlice makes a slice of VRouterInstanceType from interface
func InterfaceToVRouterInstanceTypeSlice(data interface{}) []VRouterInstanceType {
	list := data.([]interface{})
	result := MakeVRouterInstanceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVRouterInstanceType(item))
	}
	return result
}

// MakeVRouterInstanceTypeSlice() makes a slice of VRouterInstanceType
func MakeVRouterInstanceTypeSlice() []VRouterInstanceType {
	return []VRouterInstanceType{}
}
