package models

// LoadbalancerProtocolType

type LoadbalancerProtocolType string

func MakeLoadbalancerProtocolType() LoadbalancerProtocolType {
	var data LoadbalancerProtocolType
	return data
}

func InterfaceToLoadbalancerProtocolType(data interface{}) LoadbalancerProtocolType {
	return data.(LoadbalancerProtocolType)
}

func InterfaceToLoadbalancerProtocolTypeSlice(data interface{}) []LoadbalancerProtocolType {
	list := data.([]interface{})
	result := MakeLoadbalancerProtocolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerProtocolType(item))
	}
	return result
}

func MakeLoadbalancerProtocolTypeSlice() []LoadbalancerProtocolType {
	return []LoadbalancerProtocolType{}
}
