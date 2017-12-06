package models

// EtherType

type EtherType string

// MakeEtherType makes EtherType
func MakeEtherType() EtherType {
	var data EtherType
	return data
}

// InterfaceToEtherType makes EtherType from interface
func InterfaceToEtherType(data interface{}) EtherType {
	return data.(EtherType)
}

// InterfaceToEtherTypeSlice makes a slice of EtherType from interface
func InterfaceToEtherTypeSlice(data interface{}) []EtherType {
	list := data.([]interface{})
	result := MakeEtherTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToEtherType(item))
	}
	return result
}

// MakeEtherTypeSlice() makes a slice of EtherType
func MakeEtherTypeSlice() []EtherType {
	return []EtherType{}
}
