package models

// VirtualRouterType

type VirtualRouterType string

// MakeVirtualRouterType makes VirtualRouterType
func MakeVirtualRouterType() VirtualRouterType {
	var data VirtualRouterType
	return data
}

// InterfaceToVirtualRouterType makes VirtualRouterType from interface
func InterfaceToVirtualRouterType(data interface{}) VirtualRouterType {
	return data.(VirtualRouterType)
}

// InterfaceToVirtualRouterTypeSlice makes a slice of VirtualRouterType from interface
func InterfaceToVirtualRouterTypeSlice(data interface{}) []VirtualRouterType {
	list := data.([]interface{})
	result := MakeVirtualRouterTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouterType(item))
	}
	return result
}

// MakeVirtualRouterTypeSlice() makes a slice of VirtualRouterType
func MakeVirtualRouterTypeSlice() []VirtualRouterType {
	return []VirtualRouterType{}
}
