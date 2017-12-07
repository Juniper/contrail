package models

// QosConfigType

type QosConfigType string

// MakeQosConfigType makes QosConfigType
func MakeQosConfigType() QosConfigType {
	var data QosConfigType
	return data
}

// InterfaceToQosConfigType makes QosConfigType from interface
func InterfaceToQosConfigType(data interface{}) QosConfigType {
	return data.(QosConfigType)
}

// InterfaceToQosConfigTypeSlice makes a slice of QosConfigType from interface
func InterfaceToQosConfigTypeSlice(data interface{}) []QosConfigType {
	list := data.([]interface{})
	result := MakeQosConfigTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosConfigType(item))
	}
	return result
}

// MakeQosConfigTypeSlice() makes a slice of QosConfigType
func MakeQosConfigTypeSlice() []QosConfigType {
	return []QosConfigType{}
}
