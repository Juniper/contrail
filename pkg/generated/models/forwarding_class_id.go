package models

// ForwardingClassId

type ForwardingClassId int

// MakeForwardingClassId makes ForwardingClassId
func MakeForwardingClassId() ForwardingClassId {
	var data ForwardingClassId
	return data
}

// InterfaceToForwardingClassId makes ForwardingClassId from interface
func InterfaceToForwardingClassId(data interface{}) ForwardingClassId {
	return data.(ForwardingClassId)
}

// InterfaceToForwardingClassIdSlice makes a slice of ForwardingClassId from interface
func InterfaceToForwardingClassIdSlice(data interface{}) []ForwardingClassId {
	list := data.([]interface{})
	result := MakeForwardingClassIdSlice()
	for _, item := range list {
		result = append(result, InterfaceToForwardingClassId(item))
	}
	return result
}

// MakeForwardingClassIdSlice() makes a slice of ForwardingClassId
func MakeForwardingClassIdSlice() []ForwardingClassId {
	return []ForwardingClassId{}
}
