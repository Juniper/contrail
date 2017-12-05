package models

// LoadbalancerMethodType

type LoadbalancerMethodType string

func MakeLoadbalancerMethodType() LoadbalancerMethodType {
	var data LoadbalancerMethodType
	return data
}

func InterfaceToLoadbalancerMethodType(data interface{}) LoadbalancerMethodType {
	return data.(LoadbalancerMethodType)
}

func InterfaceToLoadbalancerMethodTypeSlice(data interface{}) []LoadbalancerMethodType {
	list := data.([]interface{})
	result := MakeLoadbalancerMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerMethodType(item))
	}
	return result
}

func MakeLoadbalancerMethodTypeSlice() []LoadbalancerMethodType {
	return []LoadbalancerMethodType{}
}
