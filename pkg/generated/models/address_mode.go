package models

// AddressMode

type AddressMode string

// MakeAddressMode makes AddressMode
func MakeAddressMode() AddressMode {
	var data AddressMode
	return data
}

// InterfaceToAddressMode makes AddressMode from interface
func InterfaceToAddressMode(data interface{}) AddressMode {
	return data.(AddressMode)
}

// InterfaceToAddressModeSlice makes a slice of AddressMode from interface
func InterfaceToAddressModeSlice(data interface{}) []AddressMode {
	list := data.([]interface{})
	result := MakeAddressModeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAddressMode(item))
	}
	return result
}

// MakeAddressModeSlice() makes a slice of AddressMode
func MakeAddressModeSlice() []AddressMode {
	return []AddressMode{}
}
