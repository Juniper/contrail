package models

// VirtualNetworkIdType

type VirtualNetworkIdType int

// MakeVirtualNetworkIdType makes VirtualNetworkIdType
func MakeVirtualNetworkIdType() VirtualNetworkIdType {
	var data VirtualNetworkIdType
	return data
}

// InterfaceToVirtualNetworkIdType makes VirtualNetworkIdType from interface
func InterfaceToVirtualNetworkIdType(data interface{}) VirtualNetworkIdType {
	return data.(VirtualNetworkIdType)
}

// InterfaceToVirtualNetworkIdTypeSlice makes a slice of VirtualNetworkIdType from interface
func InterfaceToVirtualNetworkIdTypeSlice(data interface{}) []VirtualNetworkIdType {
	list := data.([]interface{})
	result := MakeVirtualNetworkIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkIdType(item))
	}
	return result
}

// MakeVirtualNetworkIdTypeSlice() makes a slice of VirtualNetworkIdType
func MakeVirtualNetworkIdTypeSlice() []VirtualNetworkIdType {
	return []VirtualNetworkIdType{}
}
