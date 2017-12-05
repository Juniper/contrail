package models

// TrafficDirectionType

type TrafficDirectionType string

func MakeTrafficDirectionType() TrafficDirectionType {
	var data TrafficDirectionType
	return data
}

func InterfaceToTrafficDirectionType(data interface{}) TrafficDirectionType {
	return data.(TrafficDirectionType)
}

func InterfaceToTrafficDirectionTypeSlice(data interface{}) []TrafficDirectionType {
	list := data.([]interface{})
	result := MakeTrafficDirectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToTrafficDirectionType(item))
	}
	return result
}

func MakeTrafficDirectionTypeSlice() []TrafficDirectionType {
	return []TrafficDirectionType{}
}
