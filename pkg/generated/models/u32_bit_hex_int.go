package models

// U32BitHexInt

type U32BitHexInt string

func MakeU32BitHexInt() U32BitHexInt {
	var data U32BitHexInt
	return data
}

func InterfaceToU32BitHexInt(data interface{}) U32BitHexInt {
	return data.(U32BitHexInt)
}

func InterfaceToU32BitHexIntSlice(data interface{}) []U32BitHexInt {
	list := data.([]interface{})
	result := MakeU32BitHexIntSlice()
	for _, item := range list {
		result = append(result, InterfaceToU32BitHexInt(item))
	}
	return result
}

func MakeU32BitHexIntSlice() []U32BitHexInt {
	return []U32BitHexInt{}
}
