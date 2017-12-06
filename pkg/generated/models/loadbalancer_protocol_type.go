package models

// LoadbalancerProtocolType

type LoadbalancerProtocolType string

// MakeLoadbalancerProtocolType makes LoadbalancerProtocolType
func MakeLoadbalancerProtocolType() LoadbalancerProtocolType {
	var data LoadbalancerProtocolType
	return data
}

// InterfaceToLoadbalancerProtocolType makes LoadbalancerProtocolType from interface
func InterfaceToLoadbalancerProtocolType(data interface{}) LoadbalancerProtocolType {
	return data.(LoadbalancerProtocolType)
}

// InterfaceToLoadbalancerProtocolTypeSlice makes a slice of LoadbalancerProtocolType from interface
func InterfaceToLoadbalancerProtocolTypeSlice(data interface{}) []LoadbalancerProtocolType {
	list := data.([]interface{})
	result := MakeLoadbalancerProtocolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerProtocolType(item))
	}
	return result
}

// MakeLoadbalancerProtocolTypeSlice() makes a slice of LoadbalancerProtocolType
func MakeLoadbalancerProtocolTypeSlice() []LoadbalancerProtocolType {
	return []LoadbalancerProtocolType{}
}
