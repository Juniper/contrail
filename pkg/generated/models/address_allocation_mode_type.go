package models

// AddressAllocationModeType

type AddressAllocationModeType string

// MakeAddressAllocationModeType makes AddressAllocationModeType
func MakeAddressAllocationModeType() AddressAllocationModeType {
	var data AddressAllocationModeType
	return data
}

// InterfaceToAddressAllocationModeType makes AddressAllocationModeType from interface
func InterfaceToAddressAllocationModeType(data interface{}) AddressAllocationModeType {
	return data.(AddressAllocationModeType)
}

// InterfaceToAddressAllocationModeTypeSlice makes a slice of AddressAllocationModeType from interface
func InterfaceToAddressAllocationModeTypeSlice(data interface{}) []AddressAllocationModeType {
	list := data.([]interface{})
	result := MakeAddressAllocationModeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAddressAllocationModeType(item))
	}
	return result
}

// MakeAddressAllocationModeTypeSlice() makes a slice of AddressAllocationModeType
func MakeAddressAllocationModeTypeSlice() []AddressAllocationModeType {
	return []AddressAllocationModeType{}
}
