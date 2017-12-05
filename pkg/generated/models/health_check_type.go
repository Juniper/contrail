package models

// HealthCheckType

type HealthCheckType string

func MakeHealthCheckType() HealthCheckType {
	var data HealthCheckType
	return data
}

func InterfaceToHealthCheckType(data interface{}) HealthCheckType {
	return data.(HealthCheckType)
}

func InterfaceToHealthCheckTypeSlice(data interface{}) []HealthCheckType {
	list := data.([]interface{})
	result := MakeHealthCheckTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToHealthCheckType(item))
	}
	return result
}

func MakeHealthCheckTypeSlice() []HealthCheckType {
	return []HealthCheckType{}
}
