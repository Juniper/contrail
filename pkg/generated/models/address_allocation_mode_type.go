package models

// AddressAllocationModeType

type AddressAllocationModeType string

func MakeAddressAllocationModeType() AddressAllocationModeType {
	var data AddressAllocationModeType
	return data
}

func InterfaceToAddressAllocationModeType(data interface{}) AddressAllocationModeType {
	return data.(AddressAllocationModeType)
}

func InterfaceToAddressAllocationModeTypeSlice(data interface{}) []AddressAllocationModeType {
	list := data.([]interface{})
	result := MakeAddressAllocationModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAddressAllocationModeType(item))
	}
	return result
}

func MakeAddressAllocationModeTypeSlice() []AddressAllocationModeType {
	return []AddressAllocationModeType{}
}
