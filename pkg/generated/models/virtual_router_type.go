package models

// VirtualRouterType

type VirtualRouterType string

func MakeVirtualRouterType() VirtualRouterType {
	var data VirtualRouterType
	return data
}

func InterfaceToVirtualRouterType(data interface{}) VirtualRouterType {
	return data.(VirtualRouterType)
}

func InterfaceToVirtualRouterTypeSlice(data interface{}) []VirtualRouterType {
	list := data.([]interface{})
	result := MakeVirtualRouterTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouterType(item))
	}
	return result
}

func MakeVirtualRouterTypeSlice() []VirtualRouterType {
	return []VirtualRouterType{}
}
