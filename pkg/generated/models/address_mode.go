package models

// AddressMode

type AddressMode string

func MakeAddressMode() AddressMode {
	var data AddressMode
	return data
}

func InterfaceToAddressMode(data interface{}) AddressMode {
	return data.(AddressMode)
}

func InterfaceToAddressModeSlice(data interface{}) []AddressMode {
	list := data.([]interface{})
	result := MakeAddressModeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAddressMode(item))
	}
	return result
}

func MakeAddressModeSlice() []AddressMode {
	return []AddressMode{}
}
