package models

// VirtualNetworkIdType

type VirtualNetworkIdType int

func MakeVirtualNetworkIdType() VirtualNetworkIdType {
	var data VirtualNetworkIdType
	return data
}

func InterfaceToVirtualNetworkIdType(data interface{}) VirtualNetworkIdType {
	return data.(VirtualNetworkIdType)
}

func InterfaceToVirtualNetworkIdTypeSlice(data interface{}) []VirtualNetworkIdType {
	list := data.([]interface{})
	result := MakeVirtualNetworkIdTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkIdType(item))
	}
	return result
}

func MakeVirtualNetworkIdTypeSlice() []VirtualNetworkIdType {
	return []VirtualNetworkIdType{}
}
