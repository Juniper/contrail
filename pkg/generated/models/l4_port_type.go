package models

// L4PortType

type L4PortType int

func MakeL4PortType() L4PortType {
	var data L4PortType
	return data
}

func InterfaceToL4PortType(data interface{}) L4PortType {
	return data.(L4PortType)
}

func InterfaceToL4PortTypeSlice(data interface{}) []L4PortType {
	list := data.([]interface{})
	result := MakeL4PortTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToL4PortType(item))
	}
	return result
}

func MakeL4PortTypeSlice() []L4PortType {
	return []L4PortType{}
}
