package models

// QosConfigType

type QosConfigType string

func MakeQosConfigType() QosConfigType {
	var data QosConfigType
	return data
}

func InterfaceToQosConfigType(data interface{}) QosConfigType {
	return data.(QosConfigType)
}

func InterfaceToQosConfigTypeSlice(data interface{}) []QosConfigType {
	list := data.([]interface{})
	result := MakeQosConfigTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosConfigType(item))
	}
	return result
}

func MakeQosConfigTypeSlice() []QosConfigType {
	return []QosConfigType{}
}
