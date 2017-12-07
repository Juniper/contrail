package models

// U32BitHexInt

type U32BitHexInt string

// MakeU32BitHexInt makes U32BitHexInt
func MakeU32BitHexInt() U32BitHexInt {
	var data U32BitHexInt
	return data
}

// InterfaceToU32BitHexInt makes U32BitHexInt from interface
func InterfaceToU32BitHexInt(data interface{}) U32BitHexInt {
	return data.(U32BitHexInt)
}

// InterfaceToU32BitHexIntSlice makes a slice of U32BitHexInt from interface
func InterfaceToU32BitHexIntSlice(data interface{}) []U32BitHexInt {
	list := data.([]interface{})
	result := MakeU32BitHexIntSlice()
	for _, item := range list {
		result = append(result, InterfaceToU32BitHexInt(item))
	}
	return result
}

// MakeU32BitHexIntSlice() makes a slice of U32BitHexInt
func MakeU32BitHexIntSlice() []U32BitHexInt {
	return []U32BitHexInt{}
}
