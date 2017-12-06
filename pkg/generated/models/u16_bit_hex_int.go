package models

// U16BitHexInt

type U16BitHexInt string

// MakeU16BitHexInt makes U16BitHexInt
func MakeU16BitHexInt() U16BitHexInt {
	var data U16BitHexInt
	return data
}

// InterfaceToU16BitHexInt makes U16BitHexInt from interface
func InterfaceToU16BitHexInt(data interface{}) U16BitHexInt {
	return data.(U16BitHexInt)
}

// InterfaceToU16BitHexIntSlice makes a slice of U16BitHexInt from interface
func InterfaceToU16BitHexIntSlice(data interface{}) []U16BitHexInt {
	list := data.([]interface{})
	result := MakeU16BitHexIntSlice()
	for _, item := range list {
		result = append(result, InterfaceToU16BitHexInt(item))
	}
	return result
}

// MakeU16BitHexIntSlice() makes a slice of U16BitHexInt
func MakeU16BitHexIntSlice() []U16BitHexInt {
	return []U16BitHexInt{}
}
