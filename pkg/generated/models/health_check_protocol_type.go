package models

// HealthCheckProtocolType

type HealthCheckProtocolType string

func MakeHealthCheckProtocolType() HealthCheckProtocolType {
	var data HealthCheckProtocolType
	return data
}

func InterfaceToHealthCheckProtocolType(data interface{}) HealthCheckProtocolType {
	return data.(HealthCheckProtocolType)
}

func InterfaceToHealthCheckProtocolTypeSlice(data interface{}) []HealthCheckProtocolType {
	list := data.([]interface{})
	result := MakeHealthCheckProtocolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToHealthCheckProtocolType(item))
	}
	return result
}

func MakeHealthCheckProtocolTypeSlice() []HealthCheckProtocolType {
	return []HealthCheckProtocolType{}
}
