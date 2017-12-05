package models

// VRouterInstanceType

type VRouterInstanceType string

func MakeVRouterInstanceType() VRouterInstanceType {
	var data VRouterInstanceType
	return data
}

func InterfaceToVRouterInstanceType(data interface{}) VRouterInstanceType {
	return data.(VRouterInstanceType)
}

func InterfaceToVRouterInstanceTypeSlice(data interface{}) []VRouterInstanceType {
	list := data.([]interface{})
	result := MakeVRouterInstanceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVRouterInstanceType(item))
	}
	return result
}

func MakeVRouterInstanceTypeSlice() []VRouterInstanceType {
	return []VRouterInstanceType{}
}
