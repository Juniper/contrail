package models

// U16BitHexInt

type U16BitHexInt string

func MakeU16BitHexInt() U16BitHexInt {
	var data U16BitHexInt
	return data
}

func InterfaceToU16BitHexInt(data interface{}) U16BitHexInt {
	return data.(U16BitHexInt)
}

func InterfaceToU16BitHexIntSlice(data interface{}) []U16BitHexInt {
	list := data.([]interface{})
	result := MakeU16BitHexIntSlice()
	for _, item := range list {
		result = append(result, InterfaceToU16BitHexInt(item))
	}
	return result
}

func MakeU16BitHexIntSlice() []U16BitHexInt {
	return []U16BitHexInt{}
}
