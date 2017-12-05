package models

// EtherType

type EtherType string

func MakeEtherType() EtherType {
	var data EtherType
	return data
}

func InterfaceToEtherType(data interface{}) EtherType {
	return data.(EtherType)
}

func InterfaceToEtherTypeSlice(data interface{}) []EtherType {
	list := data.([]interface{})
	result := MakeEtherTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToEtherType(item))
	}
	return result
}

func MakeEtherTypeSlice() []EtherType {
	return []EtherType{}
}
