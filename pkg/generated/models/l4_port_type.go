package models

// L4PortType

type L4PortType int

// MakeL4PortType makes L4PortType
func MakeL4PortType() L4PortType {
	var data L4PortType
	return data
}

// InterfaceToL4PortType makes L4PortType from interface
func InterfaceToL4PortType(data interface{}) L4PortType {
	return data.(L4PortType)
}

// InterfaceToL4PortTypeSlice makes a slice of L4PortType from interface
func InterfaceToL4PortTypeSlice(data interface{}) []L4PortType {
	list := data.([]interface{})
	result := MakeL4PortTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToL4PortType(item))
	}
	return result
}

// MakeL4PortTypeSlice() makes a slice of L4PortType
func MakeL4PortTypeSlice() []L4PortType {
	return []L4PortType{}
}
