package models

// ForwardingClassId

type ForwardingClassId int

func MakeForwardingClassId() ForwardingClassId {
	var data ForwardingClassId
	return data
}

func InterfaceToForwardingClassId(data interface{}) ForwardingClassId {
	return data.(ForwardingClassId)
}

func InterfaceToForwardingClassIdSlice(data interface{}) []ForwardingClassId {
	list := data.([]interface{})
	result := MakeForwardingClassIdSlice()
	for _, item := range list {
		result = append(result, InterfaceToForwardingClassId(item))
	}
	return result
}

func MakeForwardingClassIdSlice() []ForwardingClassId {
	return []ForwardingClassId{}
}
