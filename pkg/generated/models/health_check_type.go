package models

// HealthCheckType

type HealthCheckType string

// MakeHealthCheckType makes HealthCheckType
func MakeHealthCheckType() HealthCheckType {
	var data HealthCheckType
	return data
}

// InterfaceToHealthCheckType makes HealthCheckType from interface
func InterfaceToHealthCheckType(data interface{}) HealthCheckType {
	return data.(HealthCheckType)
}

// InterfaceToHealthCheckTypeSlice makes a slice of HealthCheckType from interface
func InterfaceToHealthCheckTypeSlice(data interface{}) []HealthCheckType {
	list := data.([]interface{})
	result := MakeHealthCheckTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToHealthCheckType(item))
	}
	return result
}

// MakeHealthCheckTypeSlice() makes a slice of HealthCheckType
func MakeHealthCheckTypeSlice() []HealthCheckType {
	return []HealthCheckType{}
}
