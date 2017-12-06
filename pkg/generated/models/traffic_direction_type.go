package models

// TrafficDirectionType

type TrafficDirectionType string

// MakeTrafficDirectionType makes TrafficDirectionType
func MakeTrafficDirectionType() TrafficDirectionType {
	var data TrafficDirectionType
	return data
}

// InterfaceToTrafficDirectionType makes TrafficDirectionType from interface
func InterfaceToTrafficDirectionType(data interface{}) TrafficDirectionType {
	return data.(TrafficDirectionType)
}

// InterfaceToTrafficDirectionTypeSlice makes a slice of TrafficDirectionType from interface
func InterfaceToTrafficDirectionTypeSlice(data interface{}) []TrafficDirectionType {
	list := data.([]interface{})
	result := MakeTrafficDirectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToTrafficDirectionType(item))
	}
	return result
}

// MakeTrafficDirectionTypeSlice() makes a slice of TrafficDirectionType
func MakeTrafficDirectionTypeSlice() []TrafficDirectionType {
	return []TrafficDirectionType{}
}
