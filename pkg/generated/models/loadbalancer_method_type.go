package models

// LoadbalancerMethodType

type LoadbalancerMethodType string

// MakeLoadbalancerMethodType makes LoadbalancerMethodType
func MakeLoadbalancerMethodType() LoadbalancerMethodType {
	var data LoadbalancerMethodType
	return data
}

// InterfaceToLoadbalancerMethodType makes LoadbalancerMethodType from interface
func InterfaceToLoadbalancerMethodType(data interface{}) LoadbalancerMethodType {
	return data.(LoadbalancerMethodType)
}

// InterfaceToLoadbalancerMethodTypeSlice makes a slice of LoadbalancerMethodType from interface
func InterfaceToLoadbalancerMethodTypeSlice(data interface{}) []LoadbalancerMethodType {
	list := data.([]interface{})
	result := MakeLoadbalancerMethodTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerMethodType(item))
	}
	return result
}

// MakeLoadbalancerMethodTypeSlice() makes a slice of LoadbalancerMethodType
func MakeLoadbalancerMethodTypeSlice() []LoadbalancerMethodType {
	return []LoadbalancerMethodType{}
}
